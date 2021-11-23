# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [entry/user/proto/user.proto](#entry/user/proto/user.proto)
    - [Account](#user.Account)
    - [Account.ProfileEntry](#user.Account.ProfileEntry)
    - [CreateRequest](#user.CreateRequest)
    - [CreateRequest.ProfileEntry](#user.CreateRequest.ProfileEntry)
    - [CreateResponse](#user.CreateResponse)
    - [DeleteRequest](#user.DeleteRequest)
    - [DeleteResponse](#user.DeleteResponse)
    - [ListRequest](#user.ListRequest)
    - [ListResponse](#user.ListResponse)
    - [LoginRequest](#user.LoginRequest)
    - [LoginResponse](#user.LoginResponse)
    - [LogoutRequest](#user.LogoutRequest)
    - [LogoutResponse](#user.LogoutResponse)
    - [ReadRequest](#user.ReadRequest)
    - [ReadResponse](#user.ReadResponse)
    - [ReadSessionRequest](#user.ReadSessionRequest)
    - [ReadSessionResponse](#user.ReadSessionResponse)
    - [ResetPasswordRequest](#user.ResetPasswordRequest)
    - [ResetPasswordResponse](#user.ResetPasswordResponse)
    - [SendPasswordResetEmailRequest](#user.SendPasswordResetEmailRequest)
    - [SendPasswordResetEmailResponse](#user.SendPasswordResetEmailResponse)
    - [SendVerificationEmailRequest](#user.SendVerificationEmailRequest)
    - [SendVerificationEmailResponse](#user.SendVerificationEmailResponse)
    - [Session](#user.Session)
    - [UpdatePasswordRequest](#user.UpdatePasswordRequest)
    - [UpdatePasswordResponse](#user.UpdatePasswordResponse)
    - [UpdateRequest](#user.UpdateRequest)
    - [UpdateRequest.ProfileEntry](#user.UpdateRequest.ProfileEntry)
    - [UpdateResponse](#user.UpdateResponse)
    - [VerifyEmailRequest](#user.VerifyEmailRequest)
    - [VerifyEmailResponse](#user.VerifyEmailResponse)
  
    - [User](#user.User)
  
- [Scalar Value Types](#scalar-value-types)



<a name="entry/user/proto/user.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## entry/user/proto/user.proto



<a name="user.Account"></a>

### Account



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | unique account id |
| username | [string](#string) |  | alphanumeric username |
| email | [string](#string) |  | an email address |
| created | [int64](#int64) |  | unix timestamp |
| updated | [int64](#int64) |  | unix timestamp |
| verified | [bool](#bool) |  | if the account is verified |
| verificationDate | [int64](#int64) |  | date of verification |
| profile | [Account.ProfileEntry](#user.Account.ProfileEntry) | repeated | Store any custom data you want about your users in this fields. |






<a name="user.Account.ProfileEntry"></a>

### Account.ProfileEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="user.CreateRequest"></a>

### CreateRequest
Create a new user account. The email address and username for the account must be unique.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | optional account id |
| username | [string](#string) |  | the username |
| email | [string](#string) |  | the email address |
| password | [string](#string) |  | the user password |
| profile | [CreateRequest.ProfileEntry](#user.CreateRequest.ProfileEntry) | repeated | optional user profile as map&lt;string,string&gt; |






<a name="user.CreateRequest.ProfileEntry"></a>

### CreateRequest.ProfileEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="user.CreateResponse"></a>

### CreateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| account | [Account](#user.Account) |  |  |






<a name="user.DeleteRequest"></a>

### DeleteRequest
Delete an account by id


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | the account id |






<a name="user.DeleteResponse"></a>

### DeleteResponse







<a name="user.ListRequest"></a>

### ListRequest
List all users. Returns a paged list of results


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| offset | [int32](#int32) |  |  |
| limit | [int32](#int32) |  | Maximum number of records to return. Default limit is 25. Maximum limit is 1000. Anything higher will return an error. |






<a name="user.ListResponse"></a>

### ListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| users | [Account](#user.Account) | repeated |  |






<a name="user.LoginRequest"></a>

### LoginRequest
Login using username or email. The response will return a new session for successful login, 
401 in the case of login failure and 500 for any other error


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| username | [string](#string) |  | The username of the user |
| email | [string](#string) |  | The email address of the user |
| password | [string](#string) |  | The password of the user |






<a name="user.LoginResponse"></a>

### LoginResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| session | [Session](#user.Session) |  | The session of the logged in user |






<a name="user.LogoutRequest"></a>

### LogoutRequest
Logout a user account


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sessionId | [string](#string) |  | the session id for the user to logout |






<a name="user.LogoutResponse"></a>

### LogoutResponse







<a name="user.ReadRequest"></a>

### ReadRequest
Read an account by id, username or email. Only one need to be specified.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | the account id |
| username | [string](#string) |  | the account username |
| email | [string](#string) |  | the account email |






<a name="user.ReadResponse"></a>

### ReadResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| account | [Account](#user.Account) |  |  |






<a name="user.ReadSessionRequest"></a>

### ReadSessionRequest
Read a session by the session id. In the event it has expired or is not found and error is returned.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sessionId | [string](#string) |  | The unique session id |






<a name="user.ReadSessionResponse"></a>

### ReadSessionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| session | [Session](#user.Session) |  | the session for the user |






<a name="user.ResetPasswordRequest"></a>

### ResetPasswordRequest
Reset password with the code sent by the &#34;SendPasswordResetEmail&#34; endoint.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| email | [string](#string) |  | the email to reset the password for |
| code | [string](#string) |  | The code from the verification email |
| newPassword | [string](#string) |  | the new password |
| confirmPassword | [string](#string) |  | confirm new password |






<a name="user.ResetPasswordResponse"></a>

### ResetPasswordResponse







<a name="user.SendPasswordResetEmailRequest"></a>

### SendPasswordResetEmailRequest
Send an email with a verification code to reset password.
Call &#34;ResetPassword&#34; endpoint once user provides the code.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| email | [string](#string) |  | email address to send reset for |
| subject | [string](#string) |  | subject of the email |
| textContent | [string](#string) |  | Text content of the email. Don&#39;t forget to include the string &#39;$code&#39; which will be replaced by the real verification link HTML emails are not available currently. |
| fromName | [string](#string) |  | Display name of the sender for the email. Note: the email address will still be &#39;noreply@email.m3ocontent.com&#39; |






<a name="user.SendPasswordResetEmailResponse"></a>

### SendPasswordResetEmailResponse







<a name="user.SendVerificationEmailRequest"></a>

### SendVerificationEmailRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| email | [string](#string) |  | email address to send the verification code |
| subject | [string](#string) |  | subject of the email |
| textContent | [string](#string) |  | Text content of the email. Don&#39;t forget to include the string &#39;$micro_verification_link&#39; which will be replaced by the real verification link HTML emails are not available currently. |
| redirectUrl | [string](#string) |  |  |
| failureRedirectUrl | [string](#string) |  |  |
| fromName | [string](#string) |  | Display name of the sender for the email. Note: the email address will still be &#39;noreply@email.m3ocontent.com&#39; |






<a name="user.SendVerificationEmailResponse"></a>

### SendVerificationEmailResponse







<a name="user.Session"></a>

### Session



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | the session id |
| userId | [string](#string) |  | the associated user id |
| created | [int64](#int64) |  | unix timestamp |
| expires | [int64](#int64) |  | unix timestamp |






<a name="user.UpdatePasswordRequest"></a>

### UpdatePasswordRequest
Update the account password


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| userId | [string](#string) |  | the account id |
| oldPassword | [string](#string) |  | the old password |
| newPassword | [string](#string) |  | the new password |
| confirm_password | [string](#string) |  | confirm new password |






<a name="user.UpdatePasswordResponse"></a>

### UpdatePasswordResponse







<a name="user.UpdateRequest"></a>

### UpdateRequest
Update the account username or email


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | the account id |
| username | [string](#string) |  | the new username |
| email | [string](#string) |  | the new email address |
| profile | [UpdateRequest.ProfileEntry](#user.UpdateRequest.ProfileEntry) | repeated | the user profile as map&lt;string,string&gt; |






<a name="user.UpdateRequest.ProfileEntry"></a>

### UpdateRequest.ProfileEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="user.UpdateResponse"></a>

### UpdateResponse







<a name="user.VerifyEmailRequest"></a>

### VerifyEmailRequest
Verify the email address of an account from a token sent in an email to the user.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| email | [string](#string) |  | the email address to verify |
| token | [string](#string) |  | The token from the verification email |






<a name="user.VerifyEmailResponse"></a>

### VerifyEmailResponse






 

 

 


<a name="user.User"></a>

### User


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Create | [CreateRequest](#user.CreateRequest) | [CreateResponse](#user.CreateResponse) |  |
| Read | [ReadRequest](#user.ReadRequest) | [ReadResponse](#user.ReadResponse) |  |
| Update | [UpdateRequest](#user.UpdateRequest) | [UpdateResponse](#user.UpdateResponse) |  |
| Delete | [DeleteRequest](#user.DeleteRequest) | [DeleteResponse](#user.DeleteResponse) |  |
| UpdatePassword | [UpdatePasswordRequest](#user.UpdatePasswordRequest) | [UpdatePasswordResponse](#user.UpdatePasswordResponse) |  |
| Login | [LoginRequest](#user.LoginRequest) | [LoginResponse](#user.LoginResponse) |  |
| Logout | [LogoutRequest](#user.LogoutRequest) | [LogoutResponse](#user.LogoutResponse) |  |
| ReadSession | [ReadSessionRequest](#user.ReadSessionRequest) | [ReadSessionResponse](#user.ReadSessionResponse) |  |
| VerifyEmail | [VerifyEmailRequest](#user.VerifyEmailRequest) | [VerifyEmailResponse](#user.VerifyEmailResponse) |  |
| SendVerificationEmail | [SendVerificationEmailRequest](#user.SendVerificationEmailRequest) | [SendVerificationEmailResponse](#user.SendVerificationEmailResponse) |  |
| SendPasswordResetEmail | [SendPasswordResetEmailRequest](#user.SendPasswordResetEmailRequest) | [SendPasswordResetEmailResponse](#user.SendPasswordResetEmailResponse) |  |
| ResetPassword | [ResetPasswordRequest](#user.ResetPasswordRequest) | [ResetPasswordResponse](#user.ResetPasswordResponse) |  |
| List | [ListRequest](#user.ListRequest) | [ListResponse](#user.ListResponse) |  |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

