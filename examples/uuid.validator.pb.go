// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: examples/uuid.proto

package validator_examples

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/eugene710899989/go-proto-validators"
	regexp "regexp"
	github_com_eugene710899989_go_proto_validators "github.com/eugene710899989/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

var _regex_UUIDMsg_UserId = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)

func (this *UUIDMsg) Validate() error {
	if !_regex_UUIDMsg_UserId.MatchString(this.UserId) {
		return github_com_eugene710899989_go_proto_validators.FieldError("UserId", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.UserId))
	}
	if this.UserId == "" {
		return github_com_eugene710899989_go_proto_validators.FieldError("UserId", fmt.Errorf(`value '%v' must not be an empty string`, this.UserId))
	}
	return nil
}
