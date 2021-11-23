// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: entry/db/db_pb/db.proto

package db_pb

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/structpb"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *ReadRequest) Validate() error {
	return nil
}
func (this *ReadResponse) Validate() error {
	for _, item := range this.Records {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Records", err)
			}
		}
	}
	return nil
}
func (this *CreateRequest) Validate() error {
	if this.Record != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Record); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Record", err)
		}
	}
	return nil
}
func (this *CreateResponse) Validate() error {
	return nil
}
func (this *UpdateRequest) Validate() error {
	if this.Record != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Record); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Record", err)
		}
	}
	return nil
}
func (this *UpdateResponse) Validate() error {
	return nil
}
func (this *DeleteRequest) Validate() error {
	return nil
}
func (this *DeleteResponse) Validate() error {
	return nil
}
func (this *TruncateRequest) Validate() error {
	return nil
}
func (this *TruncateResponse) Validate() error {
	return nil
}
func (this *CountRequest) Validate() error {
	return nil
}
func (this *CountResponse) Validate() error {
	return nil
}
func (this *RenameTableRequest) Validate() error {
	return nil
}
func (this *RenameTableResponse) Validate() error {
	return nil
}
func (this *ListTablesRequest) Validate() error {
	return nil
}
func (this *ListTablesResponse) Validate() error {
	return nil
}
func (this *DropTableRequest) Validate() error {
	return nil
}
func (this *DropTableResponse) Validate() error {
	return nil
}