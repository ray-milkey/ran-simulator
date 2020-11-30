// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: atomix/membership/membership.proto

package membership

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Membership group identifier
type GroupId struct {
	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (m *GroupId) Reset()         { *m = GroupId{} }
func (m *GroupId) String() string { return proto.CompactTextString(m) }
func (*GroupId) ProtoMessage()    {}
func (*GroupId) Descriptor() ([]byte, []int) {
	return fileDescriptor_907f8caf41e45230, []int{0}
}
func (m *GroupId) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GroupId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GroupId.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GroupId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupId.Merge(m, src)
}
func (m *GroupId) XXX_Size() int {
	return m.Size()
}
func (m *GroupId) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupId.DiscardUnknown(m)
}

var xxx_messageInfo_GroupId proto.InternalMessageInfo

func (m *GroupId) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GroupId) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

// Member identifier
type MemberId struct {
	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (m *MemberId) Reset()         { *m = MemberId{} }
func (m *MemberId) String() string { return proto.CompactTextString(m) }
func (*MemberId) ProtoMessage()    {}
func (*MemberId) Descriptor() ([]byte, []int) {
	return fileDescriptor_907f8caf41e45230, []int{1}
}
func (m *MemberId) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MemberId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MemberId.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MemberId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemberId.Merge(m, src)
}
func (m *MemberId) XXX_Size() int {
	return m.Size()
}
func (m *MemberId) XXX_DiscardUnknown() {
	xxx_messageInfo_MemberId.DiscardUnknown(m)
}

var xxx_messageInfo_MemberId proto.InternalMessageInfo

func (m *MemberId) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MemberId) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

// Member is a membership member
type Member struct {
	ID   MemberId `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	Host string   `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	Port int32    `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
}

func (m *Member) Reset()         { *m = Member{} }
func (m *Member) String() string { return proto.CompactTextString(m) }
func (*Member) ProtoMessage()    {}
func (*Member) Descriptor() ([]byte, []int) {
	return fileDescriptor_907f8caf41e45230, []int{2}
}
func (m *Member) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Member) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Member.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Member) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Member.Merge(m, src)
}
func (m *Member) XXX_Size() int {
	return m.Size()
}
func (m *Member) XXX_DiscardUnknown() {
	xxx_messageInfo_Member.DiscardUnknown(m)
}

var xxx_messageInfo_Member proto.InternalMessageInfo

func (m *Member) GetID() MemberId {
	if m != nil {
		return m.ID
	}
	return MemberId{}
}

func (m *Member) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *Member) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

// JoinGroupRequest is a request to join a membership group
type JoinGroupRequest struct {
	Member  *Member `protobuf:"bytes,1,opt,name=member,proto3" json:"member,omitempty"`
	GroupID GroupId `protobuf:"bytes,2,opt,name=group_id,json=groupId,proto3" json:"group_id"`
}

func (m *JoinGroupRequest) Reset()         { *m = JoinGroupRequest{} }
func (m *JoinGroupRequest) String() string { return proto.CompactTextString(m) }
func (*JoinGroupRequest) ProtoMessage()    {}
func (*JoinGroupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_907f8caf41e45230, []int{3}
}
func (m *JoinGroupRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *JoinGroupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_JoinGroupRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *JoinGroupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinGroupRequest.Merge(m, src)
}
func (m *JoinGroupRequest) XXX_Size() int {
	return m.Size()
}
func (m *JoinGroupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinGroupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_JoinGroupRequest proto.InternalMessageInfo

func (m *JoinGroupRequest) GetMember() *Member {
	if m != nil {
		return m.Member
	}
	return nil
}

func (m *JoinGroupRequest) GetGroupID() GroupId {
	if m != nil {
		return m.GroupID
	}
	return GroupId{}
}

// JoinGroupResponse is a response to joining a membership group
type JoinGroupResponse struct {
	GroupID GroupId  `protobuf:"bytes,1,opt,name=group_id,json=groupId,proto3" json:"group_id"`
	Members []Member `protobuf:"bytes,2,rep,name=members,proto3" json:"members"`
}

func (m *JoinGroupResponse) Reset()         { *m = JoinGroupResponse{} }
func (m *JoinGroupResponse) String() string { return proto.CompactTextString(m) }
func (*JoinGroupResponse) ProtoMessage()    {}
func (*JoinGroupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_907f8caf41e45230, []int{4}
}
func (m *JoinGroupResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *JoinGroupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_JoinGroupResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *JoinGroupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinGroupResponse.Merge(m, src)
}
func (m *JoinGroupResponse) XXX_Size() int {
	return m.Size()
}
func (m *JoinGroupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinGroupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_JoinGroupResponse proto.InternalMessageInfo

func (m *JoinGroupResponse) GetGroupID() GroupId {
	if m != nil {
		return m.GroupID
	}
	return GroupId{}
}

func (m *JoinGroupResponse) GetMembers() []Member {
	if m != nil {
		return m.Members
	}
	return nil
}

func init() {
	proto.RegisterType((*GroupId)(nil), "atomix.membership.GroupId")
	proto.RegisterType((*MemberId)(nil), "atomix.membership.MemberId")
	proto.RegisterType((*Member)(nil), "atomix.membership.Member")
	proto.RegisterType((*JoinGroupRequest)(nil), "atomix.membership.JoinGroupRequest")
	proto.RegisterType((*JoinGroupResponse)(nil), "atomix.membership.JoinGroupResponse")
}

func init() { proto.RegisterFile("atomix/membership/membership.proto", fileDescriptor_907f8caf41e45230) }

var fileDescriptor_907f8caf41e45230 = []byte{
	// 357 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0x3d, 0x4f, 0xfa, 0x40,
	0x18, 0xef, 0x15, 0xfe, 0xbc, 0x3c, 0x0c, 0x7f, 0xb9, 0x38, 0x54, 0x34, 0x85, 0x54, 0x07, 0xa6,
	0xa2, 0x38, 0x18, 0xa3, 0x53, 0x43, 0x34, 0x98, 0xb0, 0x9c, 0x8b, 0x9b, 0xe1, 0xe5, 0x52, 0x6e,
	0x28, 0x57, 0x7b, 0xc5, 0xf8, 0x25, 0x4c, 0x9c, 0xfc, 0x4c, 0x8c, 0x8c, 0x4e, 0xc4, 0x94, 0x2f,
	0x62, 0x7a, 0x77, 0xbc, 0xa8, 0xc8, 0xc0, 0xd4, 0x5f, 0xee, 0xf9, 0xbd, 0x3d, 0x4f, 0x0a, 0x4e,
	0x37, 0xe6, 0x01, 0x7b, 0x69, 0x04, 0x34, 0xe8, 0xd1, 0x48, 0x0c, 0x59, 0xb8, 0x06, 0xdd, 0x30,
	0xe2, 0x31, 0xc7, 0x65, 0xc5, 0x71, 0x57, 0x83, 0xca, 0xbe, 0xcf, 0x7d, 0x2e, 0xa7, 0x8d, 0x14,
	0x29, 0xa2, 0x73, 0x05, 0xf9, 0xdb, 0x88, 0x8f, 0xc3, 0xf6, 0x00, 0x63, 0xc8, 0x8e, 0xba, 0x01,
	0xb5, 0x50, 0x0d, 0xd5, 0x8b, 0x44, 0x62, 0x7c, 0x04, 0xc5, 0xf4, 0x2b, 0xc2, 0x6e, 0x9f, 0x5a,
	0xa6, 0x1c, 0xac, 0x1e, 0x9c, 0x6b, 0x28, 0x74, 0x64, 0xc0, 0x4e, 0x6a, 0x06, 0x39, 0xa5, 0xc6,
	0x17, 0x60, 0xb2, 0x81, 0x54, 0x96, 0x9a, 0x87, 0xee, 0xaf, 0xea, 0xee, 0x22, 0xc4, 0x83, 0xc9,
	0xac, 0x6a, 0x24, 0xb3, 0xaa, 0xd9, 0x6e, 0x11, 0x93, 0xc9, 0xd0, 0x21, 0x17, 0xb1, 0xf6, 0x96,
	0x38, 0x7d, 0x0b, 0x79, 0x14, 0x5b, 0x99, 0x1a, 0xaa, 0xff, 0x23, 0x12, 0x3b, 0xaf, 0x08, 0xf6,
	0xee, 0x38, 0x1b, 0xc9, 0x55, 0x09, 0x7d, 0x1a, 0x53, 0x11, 0xe3, 0x33, 0xc8, 0xa9, 0x0c, 0x9d,
	0x7c, 0xf0, 0x67, 0x32, 0xd1, 0x44, 0x7c, 0x03, 0x05, 0x3f, 0xb5, 0x78, 0x64, 0x03, 0x99, 0x59,
	0x6a, 0x56, 0x36, 0x88, 0xf4, 0x41, 0xbd, 0xff, 0xba, 0xad, 0xbe, 0x70, 0x8b, 0xe4, 0x7d, 0x35,
	0x71, 0xde, 0x11, 0x94, 0xd7, 0xfa, 0x88, 0x90, 0x8f, 0x04, 0xfd, 0xe6, 0x8e, 0x76, 0x77, 0xc7,
	0x97, 0x90, 0xd7, 0x7c, 0xcb, 0xac, 0x65, 0xb6, 0x6e, 0xe6, 0x65, 0x53, 0x17, 0xb2, 0xe0, 0x37,
	0x03, 0x28, 0x77, 0x96, 0x9c, 0x7b, 0x1a, 0x3d, 0xb3, 0x3e, 0xc5, 0x0f, 0x50, 0x5c, 0x96, 0xc5,
	0xc7, 0x1b, 0xbc, 0x7e, 0x9e, 0xb6, 0x72, 0xb2, 0x9d, 0xa4, 0xf6, 0x3d, 0x45, 0x9e, 0x35, 0x49,
	0x6c, 0x34, 0x4d, 0x6c, 0xf4, 0x99, 0xd8, 0xe8, 0x6d, 0x6e, 0x1b, 0xd3, 0xb9, 0x6d, 0x7c, 0xcc,
	0x6d, 0xa3, 0x97, 0x93, 0xbf, 0xe7, 0xf9, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x96, 0x4a, 0x19,
	0xbf, 0xed, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MembershipServiceClient is the client API for MembershipService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MembershipServiceClient interface {
	// Joins a member to a cluster
	JoinGroup(ctx context.Context, in *JoinGroupRequest, opts ...grpc.CallOption) (MembershipService_JoinGroupClient, error)
}

type membershipServiceClient struct {
	cc *grpc.ClientConn
}

func NewMembershipServiceClient(cc *grpc.ClientConn) MembershipServiceClient {
	return &membershipServiceClient{cc}
}

func (c *membershipServiceClient) JoinGroup(ctx context.Context, in *JoinGroupRequest, opts ...grpc.CallOption) (MembershipService_JoinGroupClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MembershipService_serviceDesc.Streams[0], "/atomix.membership.MembershipService/JoinGroup", opts...)
	if err != nil {
		return nil, err
	}
	x := &membershipServiceJoinGroupClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MembershipService_JoinGroupClient interface {
	Recv() (*JoinGroupResponse, error)
	grpc.ClientStream
}

type membershipServiceJoinGroupClient struct {
	grpc.ClientStream
}

func (x *membershipServiceJoinGroupClient) Recv() (*JoinGroupResponse, error) {
	m := new(JoinGroupResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MembershipServiceServer is the server API for MembershipService service.
type MembershipServiceServer interface {
	// Joins a member to a cluster
	JoinGroup(*JoinGroupRequest, MembershipService_JoinGroupServer) error
}

// UnimplementedMembershipServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMembershipServiceServer struct {
}

func (*UnimplementedMembershipServiceServer) JoinGroup(req *JoinGroupRequest, srv MembershipService_JoinGroupServer) error {
	return status.Errorf(codes.Unimplemented, "method JoinGroup not implemented")
}

func RegisterMembershipServiceServer(s *grpc.Server, srv MembershipServiceServer) {
	s.RegisterService(&_MembershipService_serviceDesc, srv)
}

func _MembershipService_JoinGroup_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(JoinGroupRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MembershipServiceServer).JoinGroup(m, &membershipServiceJoinGroupServer{stream})
}

type MembershipService_JoinGroupServer interface {
	Send(*JoinGroupResponse) error
	grpc.ServerStream
}

type membershipServiceJoinGroupServer struct {
	grpc.ServerStream
}

func (x *membershipServiceJoinGroupServer) Send(m *JoinGroupResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _MembershipService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "atomix.membership.MembershipService",
	HandlerType: (*MembershipServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "JoinGroup",
			Handler:       _MembershipService_JoinGroup_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "atomix/membership/membership.proto",
}

func (m *GroupId) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GroupId) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GroupId) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Namespace) > 0 {
		i -= len(m.Namespace)
		copy(dAtA[i:], m.Namespace)
		i = encodeVarintMembership(dAtA, i, uint64(len(m.Namespace)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintMembership(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MemberId) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MemberId) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MemberId) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Namespace) > 0 {
		i -= len(m.Namespace)
		copy(dAtA[i:], m.Namespace)
		i = encodeVarintMembership(dAtA, i, uint64(len(m.Namespace)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintMembership(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Member) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Member) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Member) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Port != 0 {
		i = encodeVarintMembership(dAtA, i, uint64(m.Port))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Host) > 0 {
		i -= len(m.Host)
		copy(dAtA[i:], m.Host)
		i = encodeVarintMembership(dAtA, i, uint64(len(m.Host)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.ID.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMembership(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *JoinGroupRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *JoinGroupRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *JoinGroupRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.GroupID.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMembership(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.Member != nil {
		{
			size, err := m.Member.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMembership(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *JoinGroupResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *JoinGroupResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *JoinGroupResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Members) > 0 {
		for iNdEx := len(m.Members) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Members[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMembership(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.GroupID.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMembership(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintMembership(dAtA []byte, offset int, v uint64) int {
	offset -= sovMembership(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GroupId) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovMembership(uint64(l))
	}
	l = len(m.Namespace)
	if l > 0 {
		n += 1 + l + sovMembership(uint64(l))
	}
	return n
}

func (m *MemberId) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovMembership(uint64(l))
	}
	l = len(m.Namespace)
	if l > 0 {
		n += 1 + l + sovMembership(uint64(l))
	}
	return n
}

func (m *Member) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ID.Size()
	n += 1 + l + sovMembership(uint64(l))
	l = len(m.Host)
	if l > 0 {
		n += 1 + l + sovMembership(uint64(l))
	}
	if m.Port != 0 {
		n += 1 + sovMembership(uint64(m.Port))
	}
	return n
}

func (m *JoinGroupRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Member != nil {
		l = m.Member.Size()
		n += 1 + l + sovMembership(uint64(l))
	}
	l = m.GroupID.Size()
	n += 1 + l + sovMembership(uint64(l))
	return n
}

func (m *JoinGroupResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.GroupID.Size()
	n += 1 + l + sovMembership(uint64(l))
	if len(m.Members) > 0 {
		for _, e := range m.Members {
			l = e.Size()
			n += 1 + l + sovMembership(uint64(l))
		}
	}
	return n
}

func sovMembership(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMembership(x uint64) (n int) {
	return sovMembership(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GroupId) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMembership
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GroupId: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GroupId: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMembership
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMembership
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMembership
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMembership
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMembership
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMembership
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Namespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMembership(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMembership
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMembership
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MemberId) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMembership
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MemberId: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MemberId: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMembership
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMembership
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMembership
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMembership
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMembership
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMembership
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Namespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMembership(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMembership
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMembership
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Member) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMembership
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Member: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Member: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMembership
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMembership
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMembership
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Host", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMembership
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMembership
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMembership
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Host = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Port", wireType)
			}
			m.Port = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMembership
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Port |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMembership(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMembership
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMembership
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *JoinGroupRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMembership
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: JoinGroupRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: JoinGroupRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Member", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMembership
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMembership
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMembership
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Member == nil {
				m.Member = &Member{}
			}
			if err := m.Member.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMembership
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMembership
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMembership
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.GroupID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMembership(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMembership
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMembership
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *JoinGroupResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMembership
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: JoinGroupResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: JoinGroupResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMembership
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMembership
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMembership
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.GroupID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Members", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMembership
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMembership
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMembership
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Members = append(m.Members, Member{})
			if err := m.Members[len(m.Members)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMembership(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMembership
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMembership
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMembership(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMembership
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMembership
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMembership
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthMembership
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMembership
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMembership
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMembership        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMembership          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMembership = fmt.Errorf("proto: unexpected end of group")
)
