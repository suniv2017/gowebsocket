// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/services/remarketing_action_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/wrappers"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v2/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Request message for [RemarketingActionService.GetRemarketingAction][google.ads.googleads.v2.services.RemarketingActionService.GetRemarketingAction].
type GetRemarketingActionRequest struct {
	// The resource name of the remarketing action to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRemarketingActionRequest) Reset()         { *m = GetRemarketingActionRequest{} }
func (m *GetRemarketingActionRequest) String() string { return proto.CompactTextString(m) }
func (*GetRemarketingActionRequest) ProtoMessage()    {}
func (*GetRemarketingActionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_df39c8bb012d2552, []int{0}
}

func (m *GetRemarketingActionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRemarketingActionRequest.Unmarshal(m, b)
}
func (m *GetRemarketingActionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRemarketingActionRequest.Marshal(b, m, deterministic)
}
func (m *GetRemarketingActionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRemarketingActionRequest.Merge(m, src)
}
func (m *GetRemarketingActionRequest) XXX_Size() int {
	return xxx_messageInfo_GetRemarketingActionRequest.Size(m)
}
func (m *GetRemarketingActionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRemarketingActionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRemarketingActionRequest proto.InternalMessageInfo

func (m *GetRemarketingActionRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for [RemarketingActionService.MutateRemarketingActions][google.ads.googleads.v2.services.RemarketingActionService.MutateRemarketingActions].
type MutateRemarketingActionsRequest struct {
	// The ID of the customer whose remarketing actions are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The list of operations to perform on individual remarketing actions.
	Operations []*RemarketingActionOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	// If true, successful operations will be carried out and invalid
	// operations will return errors. If false, all operations will be carried
	// out in one transaction if and only if they are all valid.
	// Default is false.
	PartialFailure bool `protobuf:"varint,3,opt,name=partial_failure,json=partialFailure,proto3" json:"partial_failure,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly         bool     `protobuf:"varint,4,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateRemarketingActionsRequest) Reset()         { *m = MutateRemarketingActionsRequest{} }
func (m *MutateRemarketingActionsRequest) String() string { return proto.CompactTextString(m) }
func (*MutateRemarketingActionsRequest) ProtoMessage()    {}
func (*MutateRemarketingActionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_df39c8bb012d2552, []int{1}
}

func (m *MutateRemarketingActionsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateRemarketingActionsRequest.Unmarshal(m, b)
}
func (m *MutateRemarketingActionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateRemarketingActionsRequest.Marshal(b, m, deterministic)
}
func (m *MutateRemarketingActionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateRemarketingActionsRequest.Merge(m, src)
}
func (m *MutateRemarketingActionsRequest) XXX_Size() int {
	return xxx_messageInfo_MutateRemarketingActionsRequest.Size(m)
}
func (m *MutateRemarketingActionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateRemarketingActionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateRemarketingActionsRequest proto.InternalMessageInfo

func (m *MutateRemarketingActionsRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateRemarketingActionsRequest) GetOperations() []*RemarketingActionOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *MutateRemarketingActionsRequest) GetPartialFailure() bool {
	if m != nil {
		return m.PartialFailure
	}
	return false
}

func (m *MutateRemarketingActionsRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

// A single operation (create, update) on a remarketing action.
type RemarketingActionOperation struct {
	// FieldMask that determines which resource fields are modified in an update.
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*RemarketingActionOperation_Create
	//	*RemarketingActionOperation_Update
	Operation            isRemarketingActionOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                               `json:"-"`
	XXX_unrecognized     []byte                                 `json:"-"`
	XXX_sizecache        int32                                  `json:"-"`
}

func (m *RemarketingActionOperation) Reset()         { *m = RemarketingActionOperation{} }
func (m *RemarketingActionOperation) String() string { return proto.CompactTextString(m) }
func (*RemarketingActionOperation) ProtoMessage()    {}
func (*RemarketingActionOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_df39c8bb012d2552, []int{2}
}

func (m *RemarketingActionOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemarketingActionOperation.Unmarshal(m, b)
}
func (m *RemarketingActionOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemarketingActionOperation.Marshal(b, m, deterministic)
}
func (m *RemarketingActionOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemarketingActionOperation.Merge(m, src)
}
func (m *RemarketingActionOperation) XXX_Size() int {
	return xxx_messageInfo_RemarketingActionOperation.Size(m)
}
func (m *RemarketingActionOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_RemarketingActionOperation.DiscardUnknown(m)
}

var xxx_messageInfo_RemarketingActionOperation proto.InternalMessageInfo

func (m *RemarketingActionOperation) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

type isRemarketingActionOperation_Operation interface {
	isRemarketingActionOperation_Operation()
}

type RemarketingActionOperation_Create struct {
	Create *resources.RemarketingAction `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type RemarketingActionOperation_Update struct {
	Update *resources.RemarketingAction `protobuf:"bytes,2,opt,name=update,proto3,oneof"`
}

func (*RemarketingActionOperation_Create) isRemarketingActionOperation_Operation() {}

func (*RemarketingActionOperation_Update) isRemarketingActionOperation_Operation() {}

func (m *RemarketingActionOperation) GetOperation() isRemarketingActionOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *RemarketingActionOperation) GetCreate() *resources.RemarketingAction {
	if x, ok := m.GetOperation().(*RemarketingActionOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *RemarketingActionOperation) GetUpdate() *resources.RemarketingAction {
	if x, ok := m.GetOperation().(*RemarketingActionOperation_Update); ok {
		return x.Update
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*RemarketingActionOperation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*RemarketingActionOperation_Create)(nil),
		(*RemarketingActionOperation_Update)(nil),
	}
}

// Response message for remarketing action mutate.
type MutateRemarketingActionsResponse struct {
	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (e.g. auth errors),
	// we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results              []*MutateRemarketingActionResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *MutateRemarketingActionsResponse) Reset()         { *m = MutateRemarketingActionsResponse{} }
func (m *MutateRemarketingActionsResponse) String() string { return proto.CompactTextString(m) }
func (*MutateRemarketingActionsResponse) ProtoMessage()    {}
func (*MutateRemarketingActionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_df39c8bb012d2552, []int{3}
}

func (m *MutateRemarketingActionsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateRemarketingActionsResponse.Unmarshal(m, b)
}
func (m *MutateRemarketingActionsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateRemarketingActionsResponse.Marshal(b, m, deterministic)
}
func (m *MutateRemarketingActionsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateRemarketingActionsResponse.Merge(m, src)
}
func (m *MutateRemarketingActionsResponse) XXX_Size() int {
	return xxx_messageInfo_MutateRemarketingActionsResponse.Size(m)
}
func (m *MutateRemarketingActionsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateRemarketingActionsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateRemarketingActionsResponse proto.InternalMessageInfo

func (m *MutateRemarketingActionsResponse) GetPartialFailureError() *status.Status {
	if m != nil {
		return m.PartialFailureError
	}
	return nil
}

func (m *MutateRemarketingActionsResponse) GetResults() []*MutateRemarketingActionResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for the remarketing action mutate.
type MutateRemarketingActionResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateRemarketingActionResult) Reset()         { *m = MutateRemarketingActionResult{} }
func (m *MutateRemarketingActionResult) String() string { return proto.CompactTextString(m) }
func (*MutateRemarketingActionResult) ProtoMessage()    {}
func (*MutateRemarketingActionResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_df39c8bb012d2552, []int{4}
}

func (m *MutateRemarketingActionResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateRemarketingActionResult.Unmarshal(m, b)
}
func (m *MutateRemarketingActionResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateRemarketingActionResult.Marshal(b, m, deterministic)
}
func (m *MutateRemarketingActionResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateRemarketingActionResult.Merge(m, src)
}
func (m *MutateRemarketingActionResult) XXX_Size() int {
	return xxx_messageInfo_MutateRemarketingActionResult.Size(m)
}
func (m *MutateRemarketingActionResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateRemarketingActionResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateRemarketingActionResult proto.InternalMessageInfo

func (m *MutateRemarketingActionResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetRemarketingActionRequest)(nil), "google.ads.googleads.v2.services.GetRemarketingActionRequest")
	proto.RegisterType((*MutateRemarketingActionsRequest)(nil), "google.ads.googleads.v2.services.MutateRemarketingActionsRequest")
	proto.RegisterType((*RemarketingActionOperation)(nil), "google.ads.googleads.v2.services.RemarketingActionOperation")
	proto.RegisterType((*MutateRemarketingActionsResponse)(nil), "google.ads.googleads.v2.services.MutateRemarketingActionsResponse")
	proto.RegisterType((*MutateRemarketingActionResult)(nil), "google.ads.googleads.v2.services.MutateRemarketingActionResult")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/services/remarketing_action_service.proto", fileDescriptor_df39c8bb012d2552)
}

var fileDescriptor_df39c8bb012d2552 = []byte{
	// 721 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x95, 0x4f, 0x6b, 0xd4, 0x40,
	0x14, 0xc0, 0x4d, 0x56, 0xaa, 0x9d, 0xad, 0x0a, 0xa3, 0x62, 0xd8, 0x5a, 0xbb, 0xc4, 0x82, 0x65,
	0x0f, 0x09, 0xc4, 0xa2, 0x90, 0xb6, 0x48, 0x16, 0x6d, 0xeb, 0xa1, 0x7f, 0x48, 0xa1, 0xa0, 0x2c,
	0x84, 0x69, 0x32, 0x0d, 0xa1, 0x49, 0x26, 0xce, 0x4c, 0x56, 0x4a, 0xe9, 0x45, 0xc4, 0x2f, 0xe0,
	0x37, 0xf0, 0xe8, 0xf7, 0xf0, 0x60, 0xaf, 0x7e, 0x05, 0xbd, 0x28, 0xf8, 0x19, 0x24, 0x99, 0xcc,
	0x6e, 0xdb, 0x6d, 0xba, 0xd2, 0xde, 0xde, 0xbe, 0xf7, 0xf2, 0x7b, 0x7f, 0xe7, 0x2d, 0x70, 0x42,
	0x42, 0xc2, 0x18, 0x9b, 0x28, 0x60, 0xa6, 0x10, 0x0b, 0xa9, 0x6f, 0x99, 0x0c, 0xd3, 0x7e, 0xe4,
	0x63, 0x66, 0x52, 0x9c, 0x20, 0xba, 0x8f, 0x79, 0x94, 0x86, 0x1e, 0xf2, 0x79, 0x44, 0x52, 0xaf,
	0xb2, 0x19, 0x19, 0x25, 0x9c, 0xc0, 0xb6, 0xf8, 0xce, 0x40, 0x01, 0x33, 0x06, 0x08, 0xa3, 0x6f,
	0x19, 0x12, 0xd1, 0xb2, 0xeb, 0x82, 0x50, 0xcc, 0x48, 0x4e, 0xcf, 0x8f, 0x22, 0xe8, 0xad, 0x87,
	0xf2, 0xdb, 0x2c, 0x32, 0x51, 0x9a, 0x12, 0x8e, 0x0a, 0x23, 0xab, 0xac, 0x55, 0x6c, 0xb3, 0xfc,
	0xb5, 0x9b, 0xef, 0x99, 0x7b, 0x11, 0x8e, 0x03, 0x2f, 0x41, 0x6c, 0xbf, 0xf2, 0x78, 0x74, 0xd6,
	0xe3, 0x3d, 0x45, 0x59, 0x86, 0xa9, 0x24, 0x3c, 0xa8, 0xec, 0x34, 0xf3, 0x4d, 0xc6, 0x11, 0xcf,
	0xcf, 0x1a, 0x8a, 0xc0, 0x7e, 0x1c, 0xe1, 0x94, 0x0b, 0x83, 0xde, 0x05, 0xd3, 0xab, 0x98, 0xbb,
	0xc3, 0x84, 0x9d, 0x32, 0x5f, 0x17, 0xbf, 0xcb, 0x31, 0xe3, 0xf0, 0x31, 0xb8, 0x25, 0xcb, 0xf2,
	0x52, 0x94, 0x60, 0x4d, 0x69, 0x2b, 0xf3, 0x93, 0xee, 0x94, 0x54, 0x6e, 0xa0, 0x04, 0xeb, 0x7f,
	0x15, 0x30, 0xbb, 0x9e, 0x73, 0xc4, 0xf1, 0x08, 0x87, 0x49, 0xd0, 0x2c, 0x68, 0xfa, 0x39, 0xe3,
	0x24, 0xc1, 0xd4, 0x8b, 0x82, 0x0a, 0x03, 0xa4, 0xea, 0x75, 0x00, 0x7b, 0x00, 0x90, 0x0c, 0x53,
	0xd1, 0x10, 0x4d, 0x6d, 0x37, 0xe6, 0x9b, 0xd6, 0x92, 0x31, 0x6e, 0x1a, 0xc6, 0x48, 0xc4, 0x4d,
	0x09, 0x71, 0x4f, 0xf0, 0xe0, 0x13, 0x70, 0x27, 0x43, 0x94, 0x47, 0x28, 0xf6, 0xf6, 0x50, 0x14,
	0xe7, 0x14, 0x6b, 0x8d, 0xb6, 0x32, 0x7f, 0xd3, 0xbd, 0x5d, 0xa9, 0x57, 0x84, 0xb6, 0x28, 0xb8,
	0x8f, 0xe2, 0x28, 0x40, 0x1c, 0x7b, 0x24, 0x8d, 0x0f, 0xb4, 0xeb, 0xa5, 0xdb, 0x94, 0x54, 0x6e,
	0xa6, 0xf1, 0x81, 0xfe, 0x49, 0x05, 0xad, 0xfa, 0xc0, 0x70, 0x11, 0x34, 0xf3, 0xac, 0x24, 0x14,
	0xa3, 0x2b, 0x09, 0x4d, 0xab, 0x25, 0x6b, 0x91, 0xb3, 0x33, 0x56, 0x8a, 0xe9, 0xae, 0x23, 0xb6,
	0xef, 0x02, 0xe1, 0x5e, 0xc8, 0x70, 0x03, 0x4c, 0xf8, 0x14, 0x23, 0x2e, 0x5a, 0xdd, 0xb4, 0x16,
	0x6a, 0x7b, 0x30, 0xd8, 0xb7, 0xd1, 0x26, 0xac, 0x5d, 0x73, 0x2b, 0x4a, 0xc1, 0x13, 0x74, 0x4d,
	0xbd, 0x1a, 0x4f, 0x50, 0xba, 0x4d, 0x30, 0x39, 0xe8, 0xab, 0xfe, 0x4d, 0x01, 0xed, 0xfa, 0xc9,
	0xb3, 0x8c, 0xa4, 0x0c, 0xc3, 0x15, 0x70, 0xff, 0x4c, 0xef, 0x3d, 0x4c, 0x29, 0xa1, 0xe5, 0x04,
	0x9a, 0x16, 0x94, 0x09, 0xd1, 0xcc, 0x37, 0xb6, 0xcb, 0xa5, 0x75, 0xef, 0x9e, 0x9e, 0xca, 0xab,
	0xc2, 0x1d, 0xbe, 0x01, 0x37, 0x28, 0x66, 0x79, 0xcc, 0xe5, 0x7a, 0xbc, 0x18, 0xbf, 0x1e, 0x35,
	0xc9, 0xb9, 0x25, 0xc7, 0x95, 0x3c, 0xfd, 0x25, 0x98, 0xb9, 0xd0, 0xf3, 0xbf, 0xde, 0x81, 0xf5,
	0xa7, 0x01, 0xb4, 0x11, 0xc0, 0xb6, 0x48, 0x05, 0x7e, 0x57, 0xc0, 0xbd, 0xf3, 0x5e, 0x1a, 0x5c,
	0x1e, 0x5f, 0xc5, 0x05, 0x2f, 0xb4, 0x75, 0xa9, 0x79, 0xea, 0x4b, 0x1f, 0x7e, 0xfc, 0xfc, 0xac,
	0x3e, 0x83, 0x0b, 0xc5, 0xe1, 0x3a, 0x3c, 0x55, 0xda, 0xb2, 0x7c, 0x94, 0xcc, 0xec, 0x9c, 0xbc,
	0x64, 0xd5, 0x58, 0xcd, 0xce, 0x11, 0xfc, 0xa5, 0x00, 0xad, 0x6e, 0xec, 0xd0, 0xb9, 0xf4, 0x54,
	0xe4, 0xb1, 0x68, 0x75, 0xaf, 0x82, 0x10, 0x5b, 0xa7, 0x77, 0xcb, 0x0a, 0x97, 0xf4, 0xe7, 0x45,
	0x85, 0xc3, 0x92, 0x0e, 0x4f, 0x5c, 0xa1, 0xe5, 0xce, 0xd1, 0x39, 0x05, 0xda, 0x49, 0x89, 0xb6,
	0x95, 0x4e, 0x6b, 0xfa, 0xd8, 0xd1, 0x86, 0xe1, 0x2b, 0x29, 0x8b, 0x98, 0xe1, 0x93, 0xa4, 0xfb,
	0x51, 0x05, 0x73, 0x3e, 0x49, 0xc6, 0xa6, 0xda, 0x9d, 0xa9, 0xdb, 0x89, 0xad, 0xe2, 0x12, 0x6c,
	0x29, 0x6f, 0xd7, 0x2a, 0x44, 0x48, 0x62, 0x94, 0x86, 0x06, 0xa1, 0xa1, 0x19, 0xe2, 0xb4, 0xbc,
	0x13, 0xe6, 0x30, 0x68, 0xfd, 0xbf, 0xda, 0xa2, 0x14, 0xbe, 0xa8, 0x8d, 0x55, 0xc7, 0xf9, 0xaa,
	0xb6, 0x57, 0x05, 0xd0, 0x09, 0x98, 0x21, 0xc4, 0x42, 0xda, 0xb1, 0x8c, 0x2a, 0x30, 0x3b, 0x96,
	0x2e, 0x3d, 0x27, 0x60, 0xbd, 0x81, 0x4b, 0x6f, 0xc7, 0xea, 0x49, 0x97, 0xdf, 0xea, 0x9c, 0xd0,
	0xdb, 0xb6, 0x13, 0x30, 0xdb, 0x1e, 0x38, 0xd9, 0xf6, 0x8e, 0x65, 0xdb, 0xd2, 0x6d, 0x77, 0xa2,
	0xcc, 0xf3, 0xe9, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd2, 0x0a, 0x70, 0x8d, 0x7c, 0x07, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RemarketingActionServiceClient is the client API for RemarketingActionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RemarketingActionServiceClient interface {
	// Returns the requested remarketing action in full detail.
	GetRemarketingAction(ctx context.Context, in *GetRemarketingActionRequest, opts ...grpc.CallOption) (*resources.RemarketingAction, error)
	// Creates or updates remarketing actions. Operation statuses are returned.
	MutateRemarketingActions(ctx context.Context, in *MutateRemarketingActionsRequest, opts ...grpc.CallOption) (*MutateRemarketingActionsResponse, error)
}

type remarketingActionServiceClient struct {
	cc *grpc.ClientConn
}

func NewRemarketingActionServiceClient(cc *grpc.ClientConn) RemarketingActionServiceClient {
	return &remarketingActionServiceClient{cc}
}

func (c *remarketingActionServiceClient) GetRemarketingAction(ctx context.Context, in *GetRemarketingActionRequest, opts ...grpc.CallOption) (*resources.RemarketingAction, error) {
	out := new(resources.RemarketingAction)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.RemarketingActionService/GetRemarketingAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remarketingActionServiceClient) MutateRemarketingActions(ctx context.Context, in *MutateRemarketingActionsRequest, opts ...grpc.CallOption) (*MutateRemarketingActionsResponse, error) {
	out := new(MutateRemarketingActionsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.RemarketingActionService/MutateRemarketingActions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RemarketingActionServiceServer is the server API for RemarketingActionService service.
type RemarketingActionServiceServer interface {
	// Returns the requested remarketing action in full detail.
	GetRemarketingAction(context.Context, *GetRemarketingActionRequest) (*resources.RemarketingAction, error)
	// Creates or updates remarketing actions. Operation statuses are returned.
	MutateRemarketingActions(context.Context, *MutateRemarketingActionsRequest) (*MutateRemarketingActionsResponse, error)
}

func RegisterRemarketingActionServiceServer(s *grpc.Server, srv RemarketingActionServiceServer) {
	s.RegisterService(&_RemarketingActionService_serviceDesc, srv)
}

func _RemarketingActionService_GetRemarketingAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRemarketingActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemarketingActionServiceServer).GetRemarketingAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.RemarketingActionService/GetRemarketingAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemarketingActionServiceServer).GetRemarketingAction(ctx, req.(*GetRemarketingActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemarketingActionService_MutateRemarketingActions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateRemarketingActionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemarketingActionServiceServer).MutateRemarketingActions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.RemarketingActionService/MutateRemarketingActions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemarketingActionServiceServer).MutateRemarketingActions(ctx, req.(*MutateRemarketingActionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RemarketingActionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v2.services.RemarketingActionService",
	HandlerType: (*RemarketingActionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRemarketingAction",
			Handler:    _RemarketingActionService_GetRemarketingAction_Handler,
		},
		{
			MethodName: "MutateRemarketingActions",
			Handler:    _RemarketingActionService_MutateRemarketingActions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v2/services/remarketing_action_service.proto",
}
