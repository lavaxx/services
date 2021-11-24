package handler

import (
	"context"
	"time"

	"github.com/micro/services/pkg/cache"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"github.com/pubgo/lava/errors"
	"github.com/pubgo/lava/logger"

	"github.com/lavaxx/services/entry/otp/otp_pb"
)

var _ otp_pb.OtpServer = (*Otp)(nil)

type Otp struct{}

func (e *Otp) Init() {
}

type otpKey struct {
	Secret string
	Expiry uint
}

func (e *Otp) Generate(ctx context.Context, req *otp_pb.GenerateRequest) (*otp_pb.GenerateResponse, error) {
	var log = logger.GetLog(ctx)

	if len(req.Id) == 0 {
		return nil, errors.BadRequest("otp.generate", "missing id")
	}

	// check if a key exists for the user
	okey := new(otpKey)

	if req.Expiry <= 0 {
		req.Expiry = 60
	}

	if req.Size <= 0 {
		req.Size = 6
	}

	if err := cache.Context(ctx).Get("otp:"+req.Id, &okey); err != nil || okey == nil {
		// generate a key
		key, err := totp.Generate(totp.GenerateOpts{
			Issuer:      "Micro",
			AccountName: req.Id,
			Period:      300,
			Algorithm:   otp.AlgorithmSHA1,
		})

		if err != nil {
			log.WithErr(err).Error("Failed to generate secret")
			return nil, errors.InternalServerError("otp.generate", "failed to generate code")
		}

		okey = &otpKey{
			Secret: key.Secret(),
			Expiry: uint(req.Expiry),
		}

		if err := cache.Context(ctx).Set("otp:"+req.Id, okey, time.Now().Add(time.Minute*5)); err != nil {
			log.WithErr(err).Error("Failed to store secret")
			return nil, errors.InternalServerError("otp.generate", "failed to generate code")
		}
	}

	log.Info("generating the code: ", okey.Secret, " ", okey.Expiry)

	// generate a new code
	code, err := totp.GenerateCodeCustom(okey.Secret, time.Now(), totp.ValidateOpts{
		Period:    uint(req.Expiry),
		Skew:      1,
		Digits:    otp.Digits(req.Size),
		Algorithm: otp.AlgorithmSHA1,
	})

	if err != nil {
		return nil, errors.InternalServerError("otp.generate", "failed to generate code: %v", err)
	}

	// we have to replaced the cached value if the expiry is different
	if v := uint(req.Expiry); v != okey.Expiry {
		okey.Expiry = v

		if err := cache.Context(ctx).Set("otp:"+req.Id, okey, time.Now().Add(time.Minute*5)); err != nil {
			log.Errorf("Failed to store secret: %v", err)
			return nil, errors.InternalServerError("otp.generate", "failed to generate code")
		}
	}

	return &otp_pb.GenerateResponse{Code: code}, nil
}

func (e *Otp) Validate(ctx context.Context, req *otp_pb.ValidateRequest) (*otp_pb.ValidateResponse, error) {
	var log = logger.GetLog(ctx)

	if len(req.Id) == 0 {
		return nil, errors.BadRequest("otp.generate", "missing id")
	}
	if len(req.Code) == 0 {
		return nil, errors.BadRequest("otp.generate", "missing code")
	}

	key := new(otpKey)

	if err := cache.Context(ctx).Get("otp:"+req.Id, &key); err != nil {
		log.Errorf("Failed to get secret from store: %v", err)
		return nil, errors.InternalServerError("otp.generate", "failed to validate code")
	}

	log.Info("validating the code: ", key.Secret, " ", key.Expiry)
	ok, err := totp.ValidateCustom(req.Code, key.Secret, time.Now(), totp.ValidateOpts{
		Period:    key.Expiry,
		Skew:      1,
		Digits:    otp.Digits(len(req.Code)),
		Algorithm: otp.AlgorithmSHA1,
	})
	if err != nil {
		return nil, errors.InternalServerError("otp.generate", "failed to validate code")
	}

	return &otp_pb.ValidateResponse{Success: ok}, nil
}
