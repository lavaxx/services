# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [entry/db/proto/db.proto](#entry/db/proto/db.proto)
    - [CountRequest](#db.CountRequest)
    - [CountResponse](#db.CountResponse)
    - [CreateRequest](#db.CreateRequest)
    - [CreateResponse](#db.CreateResponse)
    - [DeleteRequest](#db.DeleteRequest)
    - [DeleteResponse](#db.DeleteResponse)
    - [DropTableRequest](#db.DropTableRequest)
    - [DropTableResponse](#db.DropTableResponse)
    - [ListTablesRequest](#db.ListTablesRequest)
    - [ListTablesResponse](#db.ListTablesResponse)
    - [ReadRequest](#db.ReadRequest)
    - [ReadResponse](#db.ReadResponse)
    - [RenameTableRequest](#db.RenameTableRequest)
    - [RenameTableResponse](#db.RenameTableResponse)
    - [TruncateRequest](#db.TruncateRequest)
    - [TruncateResponse](#db.TruncateResponse)
    - [UpdateRequest](#db.UpdateRequest)
    - [UpdateResponse](#db.UpdateResponse)
  
    - [Db](#db.Db)
  
- [Scalar Value Types](#scalar-value-types)



<a name="entry/db/proto/db.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## entry/db/proto/db.proto



<a name="db.CountRequest"></a>

### CountRequest
Count records in a table


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| table | [string](#string) |  | specify the table name |






<a name="db.CountResponse"></a>

### CountResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| count | [int32](#int32) |  | the number of records in the table |






<a name="db.CreateRequest"></a>

### CreateRequest
Create a record in the database. Optionally include an &#34;id&#34; field otherwise it&#39;s set automatically.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| table | [string](#string) |  | Optional table name. Defaults to &#39;default&#39; |
| record | [google.protobuf.Struct](#google.protobuf.Struct) |  | JSON encoded record or records (can be array or object) |






<a name="db.CreateResponse"></a>

### CreateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | The id of the record (either specified or automatically created) |






<a name="db.DeleteRequest"></a>

### DeleteRequest
Delete a record in the database by id.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| table | [string](#string) |  | Optional table name. Defaults to &#39;default&#39; |
| id | [string](#string) |  | id of the record |






<a name="db.DeleteResponse"></a>

### DeleteResponse







<a name="db.DropTableRequest"></a>

### DropTableRequest
Drop a table in the DB


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| table | [string](#string) |  |  |






<a name="db.DropTableResponse"></a>

### DropTableResponse







<a name="db.ListTablesRequest"></a>

### ListTablesRequest
List tables in the DB






<a name="db.ListTablesResponse"></a>

### ListTablesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tables | [string](#string) | repeated | list of tables |






<a name="db.ReadRequest"></a>

### ReadRequest
Read data from a table. Lookup can be by ID or via querying any field in the record.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| table | [string](#string) |  | Optional table name. Defaults to &#39;default&#39; |
| id | [string](#string) |  | Read by id. Equivalent to &#39;id == &#34;your-id&#34;&#39; |
| query | [string](#string) |  | Examples: &#39;age &gt;= 18&#39;, &#39;age &gt;= 18 and verified == true&#39; Comparison operators: &#39;==&#39;, &#39;!=&#39;, &#39;&lt;&#39;, &#39;&gt;&#39;, &#39;&lt;=&#39;, &#39;&gt;=&#39; Logical operator: &#39;and&#39; Dot access is supported, eg: &#39;user.age == 11&#39; Accessing list elements is not supported yet. |
| offset | [int32](#int32) |  |  |
| limit | [int32](#int32) |  | Maximum number of records to return. Default limit is 25. Maximum limit is 1000. Anything higher will return an error. |
| orderBy | [string](#string) |  | field name to order by |
| order | [string](#string) |  | &#39;asc&#39; (default), &#39;desc&#39; |






<a name="db.ReadResponse"></a>

### ReadResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| records | [google.protobuf.Struct](#google.protobuf.Struct) | repeated | JSON encoded records |






<a name="db.RenameTableRequest"></a>

### RenameTableRequest
Rename a table


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| from | [string](#string) |  | current table name |
| to | [string](#string) |  | new table name |






<a name="db.RenameTableResponse"></a>

### RenameTableResponse







<a name="db.TruncateRequest"></a>

### TruncateRequest
Truncate the records in a table


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| table | [string](#string) |  |  |






<a name="db.TruncateResponse"></a>

### TruncateResponse







<a name="db.UpdateRequest"></a>

### UpdateRequest
Update a record in the database. Include an &#34;id&#34; in the record to update.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| table | [string](#string) |  | Optional table name. Defaults to &#39;default&#39; |
| id | [string](#string) |  | The id of the record. If not specified it is inferred from the &#39;id&#39; field of the record |
| record | [google.protobuf.Struct](#google.protobuf.Struct) |  | record, JSON object |






<a name="db.UpdateResponse"></a>

### UpdateResponse






 

 

 


<a name="db.Db"></a>

### Db


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Create | [CreateRequest](#db.CreateRequest) | [CreateResponse](#db.CreateResponse) |  |
| Read | [ReadRequest](#db.ReadRequest) | [ReadResponse](#db.ReadResponse) |  |
| Update | [UpdateRequest](#db.UpdateRequest) | [UpdateResponse](#db.UpdateResponse) |  |
| Delete | [DeleteRequest](#db.DeleteRequest) | [DeleteResponse](#db.DeleteResponse) |  |
| Truncate | [TruncateRequest](#db.TruncateRequest) | [TruncateResponse](#db.TruncateResponse) |  |
| Count | [CountRequest](#db.CountRequest) | [CountResponse](#db.CountResponse) |  |
| RenameTable | [RenameTableRequest](#db.RenameTableRequest) | [RenameTableResponse](#db.RenameTableResponse) |  |
| ListTables | [ListTablesRequest](#db.ListTablesRequest) | [ListTablesResponse](#db.ListTablesResponse) |  |
| DropTable | [DropTableRequest](#db.DropTableRequest) | [DropTableResponse](#db.DropTableResponse) |  |

 



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

