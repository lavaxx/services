package handler

import (
	goctx "context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/pubgo/lava/errors"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"

	"github.com/lavaxx/services/entry/db"
	"github.com/lavaxx/services/entry/db/db_pb"
	"github.com/lavaxx/services/entry/otp"
	"github.com/lavaxx/services/entry/otp/otp_pb"
	"github.com/lavaxx/services/entry/user/domain"
	"github.com/lavaxx/services/entry/user/user_pb"
)

const (
	x = "cruft123"
)

var (
	alphanum    = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	emailFormat = regexp.MustCompile("^[\\w-\\.\\+]+@([\\w-]+\\.)+[\\w-]{2,4}$")
)

func random(i int) string {
	bytes := make([]byte, i)
	for {
		rand.Read(bytes)
		for i, b := range bytes {
			bytes[i] = alphanum[b%byte(len(alphanum))]
		}
		return string(bytes)
	}
	return "ughwhy?!!!"
}

type User struct {
	domain *domain.Domain
	dbCli  db_pb.DbClient
	otpCli otp_pb.OtpClient
}

func (s *User) Init() {
	s.dbCli = db_pb.GetDbClient(db.Name)
	s.domain = domain.New(s.dbCli)
	s.otpCli = otp_pb.GetOtpClient(otp.Name)
}

func (s *User) Create(ctx context.Context, req *user_pb.CreateRequest, rsp *user_pb.CreateResponse) error {
	if !emailFormat.MatchString(req.Email) {
		return errors.BadRequest("create.email-format-check", "email has wrong format")
	}
	if len(req.Password) < 8 {
		return errors.InternalServerError("user.Create.Check", "Password is less than 8 characters")
	}
	req.Username = strings.ToLower(req.Username)
	req.Email = strings.ToLower(req.Email)
	usernames, err := s.domain.Search(ctx, req.Username, "")
	if err != nil && err.Error() != "not found" {
		return err
	}
	if len(usernames) > 0 {
		return errors.BadRequest("create.username-check", "username already exists")
	}

	// TODO: don't error out here
	emails, err := s.domain.Search(ctx, "", req.Email)
	if err != nil && err.Error() != "not found" {
		return err
	}
	if len(emails) > 0 {
		return errors.BadRequest("create.email-check", "email already exists")
	}

	salt := random(16)
	h, err := bcrypt.GenerateFromPassword([]byte(x+salt+req.Password), 10)
	if err != nil {
		return errors.InternalServerError("user.Create", err.Error())
	}
	pp := base64.StdEncoding.EncodeToString(h)
	if req.Id == "" {
		req.Id = uuid.New().String()
	}

	acc := &user_pb.Account{
		Id:       req.Id,
		Username: req.Username,
		Email:    req.Email,
		Profile:  req.Profile,
	}

	err = s.domain.Create(ctx, acc, salt, pp)
	if err != nil {
		return err
	}

	// return the account
	rsp.Account = acc

	return nil
}

func (s *User) Read(ctx context.Context, req *user_pb.ReadRequest, rsp *user_pb.ReadResponse) error {
	switch {
	case req.Id != "":
		account, err := s.domain.Read(ctx, req.Id)
		if err != nil {
			return err
		}
		rsp.Account = account
		return nil
	case req.Username != "" || req.Email != "":
		accounts, err := s.domain.Search(ctx, req.Username, req.Email)
		if err != nil {
			return err
		}
		rsp.Account = accounts[0]
		return nil
	}
	return nil
}

func (s *User) Update(ctx context.Context, req *user_pb.UpdateRequest, rsp *user_pb.UpdateResponse) error {
	return s.domain.Update(ctx, &user_pb.Account{
		Id:       req.Id,
		Username: strings.ToLower(req.Username),
		Email:    strings.ToLower(req.Email),
		Profile:  req.Profile,
	})
}

func (s *User) Delete(ctx context.Context, req *user_pb.DeleteRequest, rsp *user_pb.DeleteResponse) error {
	return s.domain.Delete(ctx, req.Id)
}

func (s *User) UpdatePassword(ctx context.Context, req *user_pb.UpdatePasswordRequest, rsp *user_pb.UpdatePasswordResponse) error {
	usr, err := s.domain.Read(ctx, req.UserId)
	if err != nil {
		return errors.InternalServerError("user.updatepassword", err.Error())
	}
	if req.NewPassword != req.ConfirmPassword {
		return errors.InternalServerError("user.updatepassword", "Passwords don't match")
	}

	salt, hashed, err := s.domain.SaltAndPassword(ctx, usr.Id)
	if err != nil {
		return errors.InternalServerError("user.updatepassword", err.Error())
	}

	hh, err := base64.StdEncoding.DecodeString(hashed)
	if err != nil {
		return errors.InternalServerError("user.updatepassword", err.Error())
	}

	if err := bcrypt.CompareHashAndPassword(hh, []byte(x+salt+req.OldPassword)); err != nil {
		return errors.Unauthorized("user.updatepassword", err.Error())
	}

	salt = random(16)
	h, err := bcrypt.GenerateFromPassword([]byte(x+salt+req.NewPassword), 10)
	if err != nil {
		return errors.InternalServerError("user.updatepassword", err.Error())
	}
	pp := base64.StdEncoding.EncodeToString(h)

	if err := s.domain.UpdatePassword(ctx, req.UserId, salt, pp); err != nil {
		return errors.InternalServerError("user.updatepassword", err.Error())
	}
	return nil
}

func (s *User) Login(ctx context.Context, req *user_pb.LoginRequest, rsp *user_pb.LoginResponse) error {
	username := strings.ToLower(req.Username)
	email := strings.ToLower(req.Email)

	accounts, err := s.domain.Search(ctx, username, email)
	if err != nil {
		return err
	}
	if len(accounts) == 0 {
		return fmt.Errorf("account not found")
	}
	salt, hashed, err := s.domain.SaltAndPassword(ctx, accounts[0].Id)
	if err != nil {
		return err
	}

	hh, err := base64.StdEncoding.DecodeString(hashed)
	if err != nil {
		return errors.InternalServerError("user.Login", err.Error())
	}

	if err := bcrypt.CompareHashAndPassword(hh, []byte(x+salt+req.Password)); err != nil {
		return errors.Unauthorized("user.login", err.Error())
	}
	// save session
	sess := &user_pb.Session{
		Id:      random(128),
		Created: time.Now().Unix(),
		Expires: time.Now().Add(time.Hour * 24 * 7).Unix(),
		UserId:  accounts[0].Id,
	}

	if err := s.domain.CreateSession(ctx, sess); err != nil {
		return errors.InternalServerError("user.Login", err.Error())
	}
	rsp.Session = sess
	return nil
}

func (s *User) Logout(ctx context.Context, req *user_pb.LogoutRequest, rsp *user_pb.LogoutResponse) error {
	return s.domain.DeleteSession(ctx, req.SessionId)
}

func (s *User) ReadSession(ctx context.Context, req *user_pb.ReadSessionRequest, rsp *user_pb.ReadSessionResponse) error {
	sess, err := s.domain.ReadSession(ctx, req.SessionId)
	if err != nil {
		return err
	}
	rsp.Session = sess
	return nil
}

func (s *User) VerifyEmail(ctx context.Context, req *user_pb.VerifyEmailRequest, rsp *user_pb.VerifyEmailResponse) error {
	if len(req.Email) == 0 {
		return errors.BadRequest("user.verifyemail", "missing email")
	}
	if len(req.Token) == 0 {
		return errors.BadRequest("user.verifyemail", "missing token")
	}

	// check the token exists
	userId, err := s.domain.ReadToken(ctx, req.Email, req.Token)
	if err != nil {
		return err
	}

	// validate the code, e.g its an OTP token and hasn't expired
	resp, err := s.otpCli.Validate(ctx, &otp_pb.ValidateRequest{
		Id:   req.Email,
		Code: req.Token,
	})
	if err != nil {
		return err
	}

	// check if the code is actually valid
	if !resp.Success {
		return errors.BadRequest("user.resetpassword", "invalid code")
	}

	// mark user as verified
	user, err := s.domain.Read(ctx, userId)
	user.Verified = true

	// update the user
	return s.domain.Update(ctx, user)
}

func (s *User) SendVerificationEmail(ctx context.Context, req *user_pb.SendVerificationEmailRequest, rsp *user_pb.SendVerificationEmailResponse) error {
	if len(req.Email) == 0 {
		return errors.BadRequest("user.sendverificationemail", "missing email")
	}

	// search for the user
	users, err := s.domain.Search(ctx, "", req.Email)
	if err != nil {
		return err
	}

	// generate a new OTP code
	resp, err := s.otpCli.Generate(ctx, &otp_pb.GenerateRequest{
		Expiry: 300,
		Id:     req.Email,
	})

	if err != nil {
		return err
	}

	// generate/save a token for verification
	token, err := s.domain.CreateToken(ctx, req.Email, resp.Code)
	if err != nil {
		return err
	}

	return s.domain.SendEmail(req.FromName, req.Email, users[0].Username, req.Subject, req.TextContent, token, req.RedirectUrl, req.FailureRedirectUrl)
}

func (s *User) SendPasswordResetEmail(ctx context.Context, req *user_pb.SendPasswordResetEmailRequest, rsp *user_pb.SendPasswordResetEmailResponse) error {
	if len(req.Email) == 0 {
		return errors.BadRequest("user.sendpasswordresetemail", "missing email")
	}

	// look for an existing user
	users, err := s.domain.Search(ctx, "", req.Email)
	if err != nil {
		return err
	}

	// generate a new OTP code
	resp, err := s.otpCli.Generate(ctx, &otp_pb.GenerateRequest{
		Expiry: 300,
		Id:     req.Email,
	})

	if err != nil {
		return err
	}

	// save the code in the database and then send via email
	return s.domain.SendPasswordResetEmail(ctx, users[0].Id, resp.Code, req.FromName, req.Email, users[0].Username, req.Subject, req.TextContent)
}

func (s *User) ResetPassword(ctx context.Context, req *user_pb.ResetPasswordRequest, rsp *user_pb.ResetPasswordResponse) error {
	if len(req.Email) == 0 {
		return errors.BadRequest("user.resetpassword", "missing email")
	}
	if len(req.Code) == 0 {
		return errors.BadRequest("user.resetpassword", "missing code")
	}
	if len(req.ConfirmPassword) == 0 {
		return errors.BadRequest("user.resetpassword", "missing confirm password")
	}
	if len(req.NewPassword) == 0 {
		return errors.BadRequest("user.resetpassword", "missing new password")
	}
	if req.ConfirmPassword != req.NewPassword {
		return errors.BadRequest("user.resetpassword", "passwords do not match")
	}

	// check if a request was made to reset the password, we should have saved it
	code, err := s.domain.ReadPasswordResetCode(ctx, req.Email, req.Code)
	if err != nil {
		return err
	}

	// validate the code, e.g its an OTP token and hasn't expired
	resp, err := s.otpCli.Validate(ctx, &otp_pb.ValidateRequest{
		Id:   req.Email,
		Code: req.Code,
	})
	if err != nil {
		return err
	}

	// check if the code is actually valid
	if !resp.Success {
		return errors.BadRequest("user.resetpassword", "invalid code")
	}

	// no error means it exists and not expired
	salt := random(16)
	h, err := bcrypt.GenerateFromPassword([]byte(x+salt+req.NewPassword), 10)
	if err != nil {
		return errors.InternalServerError("user.ResetPassword", err.Error())
	}
	pp := base64.StdEncoding.EncodeToString(h)

	// update the user password
	if err := s.domain.UpdatePassword(ctx, code.UserID, salt, pp); err != nil {
		return errors.InternalServerError("user.resetpassword", err.Error())
	}

	// delete our saved code
	s.domain.DeletePasswordRestCode(ctx, req.Email, req.Code)

	return nil
}

func (s *User) List(ctx goctx.Context, request *user_pb.ListRequest, response *user_pb.ListResponse) error {
	accs, err := s.domain.List(ctx, request.Offset, request.Limit)
	if err != nil && err != domain.ErrNotFound {
		return errors.InternalServerError("user.List", "Error retrieving user list")
	}
	response.Users = make([]*user_pb.Account, len(accs))
	for i, v := range accs {
		response.Users[i] = v
	}
	return nil
}
