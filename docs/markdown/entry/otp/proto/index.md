# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [entry/otp/proto/otp.proto](#entry/otp/proto/otp.proto)
    - [GenerateRequest](#otp.GenerateRequest)
    - [GenerateResponse](#otp.GenerateResponse)
    - [ValidateRequest](#otp.ValidateRequest)
    - [ValidateResponse](#otp.ValidateResponse)
  
    - [Otp](#otp.Otp)
  
- [Scalar Value Types](#scalar-value-types)



<a name="entry/otp/proto/otp.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## entry/otp/proto/otp.proto



<a name="otp.GenerateRequest"></a>

### GenerateRequest
Generate an OTP (one time pass) code


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | unique id, email or user to generate an OTP for |
| size | [int64](#int64) |  | number of characters (default: 6) |
| expiry | [int64](#int64) |  | expiration in seconds (default: 60) |






<a name="otp.GenerateResponse"></a>

### GenerateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| code | [string](#string) |  | one time pass code |






<a name="otp.ValidateRequest"></a>

### ValidateRequest
Validate the OTP code


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | unique id, email or user for which the code was generated |
| code | [string](#string) |  | one time pass code to validate |






<a name="otp.ValidateResponse"></a>

### ValidateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| success | [bool](#bool) |  | returns true if the code is valid for the ID |





 

 

 


<a name="otp.Otp"></a>

### Otp


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Generate | [GenerateRequest](#otp.GenerateRequest) | [GenerateResponse](#otp.GenerateResponse) |  |
| Validate | [ValidateRequest](#otp.ValidateRequest) | [ValidateResponse](#otp.ValidateResponse) |  |

 



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

