// Code generated by protoc-gen-go.
// source: proto/google_security_e2ekeys_v2/e2ekeys.proto
// DO NOT EDIT!

/*
Package google_security_e2ekeys_v2 is a generated protocol buffer package.

It is generated from these files:
	proto/google_security_e2ekeys_v2/e2ekeys.proto

It has these top-level messages:
	GetEntryResponse
	Profile
	PublicKey
	GetEntryRequest
	ListEntryHistoryRequest
	ListEntryHistoryResponse
	UpdateEntryRequest
	UpdateEntryResponse
	ListSEHRequest
	ListSEHResponse
	ListUpdateRequest
	ListUpdateResponse
	ListStepsRequest
	ListStepsResponse
*/
package google_security_e2ekeys_v2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import security_ctmap "github.com/google/e2e-key-server/proto/security_ctmap"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type GetEntryResponse struct {
	// UserId supports sending this as a complete proof to a third party.
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	// IndexProof is the proof for VUF(user_id).
	IndexProof []byte `protobuf:"bytes,2,opt,name=index_proof,json=indexProof,proto3" json:"index_proof,omitempty"`
	// Index is the output of VUF on (user_id).
	// If index != entry.index,  but they share a common prefix of
	// len(merkle_tree_neighbors), this is a proof of absence for user_id.
	Index []byte `protobuf:"bytes,3,opt,name=index,proto3" json:"index,omitempty"`
	// The signed epoch head.  A server may choose to return multiple signatures
	// of the same head for the same epoch. A client MUST NOT consider the
	// presence of surplus or invalid epoch heads in this field an error.
	SignedEpochHeads []*security_ctmap.SignedEpochHead `protobuf:"bytes,4,rep,name=signed_epoch_heads,json=signedEpochHeads" json:"signed_epoch_heads,omitempty"`
	// MerkleTreeNeighbors is a list of all the adjacent nodes along the path
	// from the bottommost node to the head.
	MerkleTreeNeighbors [][]byte `protobuf:"bytes,5,rep,name=merkle_tree_neighbors,json=merkleTreeNeighbors,proto3" json:"merkle_tree_neighbors,omitempty"`
	// Entry contains the public portion of the user's data, and a comitment to
	// profile.
	Entry *security_ctmap.Entry `protobuf:"bytes,6,opt,name=entry" json:"entry,omitempty"`
	// Profile contains the user's keys.
	Profile []byte `protobuf:"bytes,7,opt,name=profile,proto3" json:"profile,omitempty"`
	// CommitmentKey is 16 random bytes.
	CommitmentKey []byte `protobuf:"bytes,8,opt,name=commitment_key,json=commitmentKey,proto3" json:"commitment_key,omitempty"`
}

func (m *GetEntryResponse) Reset()                    { *m = GetEntryResponse{} }
func (m *GetEntryResponse) String() string            { return proto.CompactTextString(m) }
func (*GetEntryResponse) ProtoMessage()               {}
func (*GetEntryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetEntryResponse) GetSignedEpochHeads() []*security_ctmap.SignedEpochHead {
	if m != nil {
		return m.SignedEpochHeads
	}
	return nil
}

func (m *GetEntryResponse) GetEntry() *security_ctmap.Entry {
	if m != nil {
		return m.Entry
	}
	return nil
}

// Profile contains data hidden behind the crypto comitment.
type Profile struct {
	// Keys is a map of appIds to keys.
	Keys map[string][]byte `protobuf:"bytes,1,rep,name=keys" json:"keys,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *Profile) Reset()                    { *m = Profile{} }
func (m *Profile) String() string            { return proto.CompactTextString(m) }
func (*Profile) ProtoMessage()               {}
func (*Profile) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Profile) GetKeys() map[string][]byte {
	if m != nil {
		return m.Keys
	}
	return nil
}

// PublicKey defines a key this domain uses to sign EpochHeads with.
type PublicKey struct {
	// KeyFormats from Keyczar.
	//
	// Types that are valid to be assigned to KeyType:
	//	*PublicKey_Ed25519
	//	*PublicKey_RsaVerifyingSha256_2048
	//	*PublicKey_EcdsaVerifyingP256
	KeyType isPublicKey_KeyType `protobuf_oneof:"key_type"`
}

func (m *PublicKey) Reset()                    { *m = PublicKey{} }
func (m *PublicKey) String() string            { return proto.CompactTextString(m) }
func (*PublicKey) ProtoMessage()               {}
func (*PublicKey) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type isPublicKey_KeyType interface {
	isPublicKey_KeyType()
}

type PublicKey_Ed25519 struct {
	Ed25519 []byte `protobuf:"bytes,1,opt,name=ed25519,proto3,oneof"`
}
type PublicKey_RsaVerifyingSha256_2048 struct {
	RsaVerifyingSha256_2048 []byte `protobuf:"bytes,2,opt,name=rsa_verifying_sha256_2048,json=rsaVerifyingSha2562048,proto3,oneof"`
}
type PublicKey_EcdsaVerifyingP256 struct {
	EcdsaVerifyingP256 []byte `protobuf:"bytes,3,opt,name=ecdsa_verifying_p256,json=ecdsaVerifyingP256,proto3,oneof"`
}

func (*PublicKey_Ed25519) isPublicKey_KeyType()                 {}
func (*PublicKey_RsaVerifyingSha256_2048) isPublicKey_KeyType() {}
func (*PublicKey_EcdsaVerifyingP256) isPublicKey_KeyType()      {}

func (m *PublicKey) GetKeyType() isPublicKey_KeyType {
	if m != nil {
		return m.KeyType
	}
	return nil
}

func (m *PublicKey) GetEd25519() []byte {
	if x, ok := m.GetKeyType().(*PublicKey_Ed25519); ok {
		return x.Ed25519
	}
	return nil
}

func (m *PublicKey) GetRsaVerifyingSha256_2048() []byte {
	if x, ok := m.GetKeyType().(*PublicKey_RsaVerifyingSha256_2048); ok {
		return x.RsaVerifyingSha256_2048
	}
	return nil
}

func (m *PublicKey) GetEcdsaVerifyingP256() []byte {
	if x, ok := m.GetKeyType().(*PublicKey_EcdsaVerifyingP256); ok {
		return x.EcdsaVerifyingP256
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*PublicKey) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _PublicKey_OneofMarshaler, _PublicKey_OneofUnmarshaler, _PublicKey_OneofSizer, []interface{}{
		(*PublicKey_Ed25519)(nil),
		(*PublicKey_RsaVerifyingSha256_2048)(nil),
		(*PublicKey_EcdsaVerifyingP256)(nil),
	}
}

func _PublicKey_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*PublicKey)
	// key_type
	switch x := m.KeyType.(type) {
	case *PublicKey_Ed25519:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.Ed25519)
	case *PublicKey_RsaVerifyingSha256_2048:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.RsaVerifyingSha256_2048)
	case *PublicKey_EcdsaVerifyingP256:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.EcdsaVerifyingP256)
	case nil:
	default:
		return fmt.Errorf("PublicKey.KeyType has unexpected type %T", x)
	}
	return nil
}

func _PublicKey_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*PublicKey)
	switch tag {
	case 1: // key_type.ed25519
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.KeyType = &PublicKey_Ed25519{x}
		return true, err
	case 2: // key_type.rsa_verifying_sha256_2048
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.KeyType = &PublicKey_RsaVerifyingSha256_2048{x}
		return true, err
	case 3: // key_type.ecdsa_verifying_p256
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.KeyType = &PublicKey_EcdsaVerifyingP256{x}
		return true, err
	default:
		return false, nil
	}
}

func _PublicKey_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*PublicKey)
	// key_type
	switch x := m.KeyType.(type) {
	case *PublicKey_Ed25519:
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Ed25519)))
		n += len(x.Ed25519)
	case *PublicKey_RsaVerifyingSha256_2048:
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.RsaVerifyingSha256_2048)))
		n += len(x.RsaVerifyingSha256_2048)
	case *PublicKey_EcdsaVerifyingP256:
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.EcdsaVerifyingP256)))
		n += len(x.EcdsaVerifyingP256)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Get request for a user object.
type GetEntryRequest struct {
	// Absence of the time field indicates a request for the current value.
	Epoch int64 `protobuf:"varint,1,opt,name=epoch" json:"epoch,omitempty"`
	// User identifier. Most commonly an email address.
	UserId string `protobuf:"bytes,2,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	// TODO(cesarghali): implement app_id filtering.
	// Only return the keys belonging to this app.
	AppId string `protobuf:"bytes,3,opt,name=app_id,json=appId" json:"app_id,omitempty"`
	// TODO: implement key_hash filtering.
	// Only return SEH's that are signed by these keys.
	TrustedKeys []uint64 `protobuf:"fixed64,4,rep,name=trusted_keys,json=trustedKeys" json:"trusted_keys,omitempty"`
}

func (m *GetEntryRequest) Reset()                    { *m = GetEntryRequest{} }
func (m *GetEntryRequest) String() string            { return proto.CompactTextString(m) }
func (*GetEntryRequest) ProtoMessage()               {}
func (*GetEntryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

// Get a list of historical values for a user.
type ListEntryHistoryRequest struct {
	// The user identifier.
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	// from_epoch is the starting epcoh.
	StartEpoch int64 `protobuf:"varint,2,opt,name=start_epoch,json=startEpoch" json:"start_epoch,omitempty"`
	// The maximum number of entries to return.
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
}

func (m *ListEntryHistoryRequest) Reset()                    { *m = ListEntryHistoryRequest{} }
func (m *ListEntryHistoryRequest) String() string            { return proto.CompactTextString(m) }
func (*ListEntryHistoryRequest) ProtoMessage()               {}
func (*ListEntryHistoryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

// A paginated history of values for a user.
type ListEntryHistoryResponse struct {
	// The list of values this user_id has contained over time.
	Values []*GetEntryResponse `protobuf:"bytes,1,rep,name=values" json:"values,omitempty"`
	// The next time to query for pagination.
	NextEpoch int64 `protobuf:"varint,2,opt,name=next_epoch,json=nextEpoch" json:"next_epoch,omitempty"`
}

func (m *ListEntryHistoryResponse) Reset()                    { *m = ListEntryHistoryResponse{} }
func (m *ListEntryHistoryResponse) String() string            { return proto.CompactTextString(m) }
func (*ListEntryHistoryResponse) ProtoMessage()               {}
func (*ListEntryHistoryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ListEntryHistoryResponse) GetValues() []*GetEntryResponse {
	if m != nil {
		return m.Values
	}
	return nil
}

// Update a user's profile.
type UpdateEntryRequest struct {
	// UserID specifies the id for the new account to be registered.
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	// SignedEntryUpdate authorizes the change to profile.
	SignedEntryUpdate *security_ctmap.SignedEntryUpdate `protobuf:"bytes,2,opt,name=signed_entry_update,json=signedEntryUpdate" json:"signed_entry_update,omitempty"`
	// Profile is the serialized protobuf Profile.
	// Profile is private and must not be released to verifiers.
	Profile []byte `protobuf:"bytes,3,opt,name=profile,proto3" json:"profile,omitempty"`
	// CommitmentKey is 16 random bytes.
	CommitmentKey []byte `protobuf:"bytes,4,opt,name=commitment_key,json=commitmentKey,proto3" json:"commitment_key,omitempty"`
	// TODO: Provide a way for clients to specify a quorum of signatures to wait
	// on before returning.
	NotUsed []uint64 `protobuf:"fixed64,5,rep,name=not_used,json=notUsed" json:"not_used,omitempty"`
	// DkimProof is used to vouch for the validity of a new registration.
	// Used when OAuth is not used.
	DkimProof []byte `protobuf:"bytes,1001,opt,name=dkim_proof,json=dkimProof,proto3" json:"dkim_proof,omitempty"`
}

func (m *UpdateEntryRequest) Reset()                    { *m = UpdateEntryRequest{} }
func (m *UpdateEntryRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateEntryRequest) ProtoMessage()               {}
func (*UpdateEntryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *UpdateEntryRequest) GetSignedEntryUpdate() *security_ctmap.SignedEntryUpdate {
	if m != nil {
		return m.SignedEntryUpdate
	}
	return nil
}

// UpdateEntryResponse contains a proof once the update has been included in
// the Merkel Tree.
type UpdateEntryResponse struct {
	Proof *GetEntryResponse `protobuf:"bytes,1,opt,name=proof" json:"proof,omitempty"`
}

func (m *UpdateEntryResponse) Reset()                    { *m = UpdateEntryResponse{} }
func (m *UpdateEntryResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateEntryResponse) ProtoMessage()               {}
func (*UpdateEntryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *UpdateEntryResponse) GetProof() *GetEntryResponse {
	if m != nil {
		return m.Proof
	}
	return nil
}

type ListSEHRequest struct {
	// from_epoch is the starting epcoh.
	StartEpoch int64 `protobuf:"varint,1,opt,name=start_epoch,json=startEpoch" json:"start_epoch,omitempty"`
	// The maximum number of entries to return.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
}

func (m *ListSEHRequest) Reset()                    { *m = ListSEHRequest{} }
func (m *ListSEHRequest) String() string            { return proto.CompactTextString(m) }
func (*ListSEHRequest) ProtoMessage()               {}
func (*ListSEHRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type ListSEHResponse struct {
	Heads []*security_ctmap.SignedEpochHead `protobuf:"bytes,1,rep,name=heads" json:"heads,omitempty"`
}

func (m *ListSEHResponse) Reset()                    { *m = ListSEHResponse{} }
func (m *ListSEHResponse) String() string            { return proto.CompactTextString(m) }
func (*ListSEHResponse) ProtoMessage()               {}
func (*ListSEHResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *ListSEHResponse) GetHeads() []*security_ctmap.SignedEpochHead {
	if m != nil {
		return m.Heads
	}
	return nil
}

type ListUpdateRequest struct {
	// start_commitment_timestamp is the starting commitment timestamp.
	StartCommitmentTimestamp int64 `protobuf:"varint,1,opt,name=start_commitment_timestamp,json=startCommitmentTimestamp" json:"start_commitment_timestamp,omitempty"`
	// The maximum number of entries to return.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
}

func (m *ListUpdateRequest) Reset()                    { *m = ListUpdateRequest{} }
func (m *ListUpdateRequest) String() string            { return proto.CompactTextString(m) }
func (*ListUpdateRequest) ProtoMessage()               {}
func (*ListUpdateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

type ListUpdateResponse struct {
	// updates contains a serialized SignedEntryUpdate.
	Updates [][]byte `protobuf:"bytes,1,rep,name=updates,proto3" json:"updates,omitempty"`
}

func (m *ListUpdateResponse) Reset()                    { *m = ListUpdateResponse{} }
func (m *ListUpdateResponse) String() string            { return proto.CompactTextString(m) }
func (*ListUpdateResponse) ProtoMessage()               {}
func (*ListUpdateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

type ListStepsRequest struct {
	// start_commitment_timestamp is the starting commitment timestamp.
	StartCommitmentTimestamp int64 `protobuf:"varint,1,opt,name=start_commitment_timestamp,json=startCommitmentTimestamp" json:"start_commitment_timestamp,omitempty"`
	// The maximum number of entries to return.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
}

func (m *ListStepsRequest) Reset()                    { *m = ListStepsRequest{} }
func (m *ListStepsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListStepsRequest) ProtoMessage()               {}
func (*ListStepsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

type ListStepsResponse struct {
	Steps []*security_ctmap.Step `protobuf:"bytes,1,rep,name=steps" json:"steps,omitempty"`
}

func (m *ListStepsResponse) Reset()                    { *m = ListStepsResponse{} }
func (m *ListStepsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListStepsResponse) ProtoMessage()               {}
func (*ListStepsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *ListStepsResponse) GetSteps() []*security_ctmap.Step {
	if m != nil {
		return m.Steps
	}
	return nil
}

func init() {
	proto.RegisterType((*GetEntryResponse)(nil), "google.security.e2ekeys.v2.GetEntryResponse")
	proto.RegisterType((*Profile)(nil), "google.security.e2ekeys.v2.Profile")
	proto.RegisterType((*PublicKey)(nil), "google.security.e2ekeys.v2.PublicKey")
	proto.RegisterType((*GetEntryRequest)(nil), "google.security.e2ekeys.v2.GetEntryRequest")
	proto.RegisterType((*ListEntryHistoryRequest)(nil), "google.security.e2ekeys.v2.ListEntryHistoryRequest")
	proto.RegisterType((*ListEntryHistoryResponse)(nil), "google.security.e2ekeys.v2.ListEntryHistoryResponse")
	proto.RegisterType((*UpdateEntryRequest)(nil), "google.security.e2ekeys.v2.UpdateEntryRequest")
	proto.RegisterType((*UpdateEntryResponse)(nil), "google.security.e2ekeys.v2.UpdateEntryResponse")
	proto.RegisterType((*ListSEHRequest)(nil), "google.security.e2ekeys.v2.ListSEHRequest")
	proto.RegisterType((*ListSEHResponse)(nil), "google.security.e2ekeys.v2.ListSEHResponse")
	proto.RegisterType((*ListUpdateRequest)(nil), "google.security.e2ekeys.v2.ListUpdateRequest")
	proto.RegisterType((*ListUpdateResponse)(nil), "google.security.e2ekeys.v2.ListUpdateResponse")
	proto.RegisterType((*ListStepsRequest)(nil), "google.security.e2ekeys.v2.ListStepsRequest")
	proto.RegisterType((*ListStepsResponse)(nil), "google.security.e2ekeys.v2.ListStepsResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for E2EKeyService service

type E2EKeyServiceClient interface {
	// GetEntry returns a user's entry in the Merkle Tree. Entries contain
	// signed commitments to a profile, which is also returned.
	GetEntry(ctx context.Context, in *GetEntryRequest, opts ...grpc.CallOption) (*GetEntryResponse, error)
	// ListEntryHistory returns a list of GetEntryRespons covering several epochs.
	ListEntryHistory(ctx context.Context, in *ListEntryHistoryRequest, opts ...grpc.CallOption) (*ListEntryHistoryResponse, error)
	// blocking or polling?
	// UpdateEntry submits a SignedEntryUpdate.  Returns empty until this update
	// has been included in an epoch.  Clients must retry until this function
	// returns a proof.
	UpdateEntry(ctx context.Context, in *UpdateEntryRequest, opts ...grpc.CallOption) (*UpdateEntryResponse, error)
	// List the Signed Epoch Heads, from epoch to epoch.
	ListSEH(ctx context.Context, in *ListSEHRequest, opts ...grpc.CallOption) (*ListSEHResponse, error)
	// List the EntryUpdates by update number.
	ListUpdate(ctx context.Context, in *ListUpdateRequest, opts ...grpc.CallOption) (*ListUpdateResponse, error)
	// ListSteps combines SEH and EntryUpdates into single list.
	ListSteps(ctx context.Context, in *ListStepsRequest, opts ...grpc.CallOption) (*ListStepsResponse, error)
}

type e2EKeyServiceClient struct {
	cc *grpc.ClientConn
}

func NewE2EKeyServiceClient(cc *grpc.ClientConn) E2EKeyServiceClient {
	return &e2EKeyServiceClient{cc}
}

func (c *e2EKeyServiceClient) GetEntry(ctx context.Context, in *GetEntryRequest, opts ...grpc.CallOption) (*GetEntryResponse, error) {
	out := new(GetEntryResponse)
	err := grpc.Invoke(ctx, "/google.security.e2ekeys.v2.E2EKeyService/GetEntry", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *e2EKeyServiceClient) ListEntryHistory(ctx context.Context, in *ListEntryHistoryRequest, opts ...grpc.CallOption) (*ListEntryHistoryResponse, error) {
	out := new(ListEntryHistoryResponse)
	err := grpc.Invoke(ctx, "/google.security.e2ekeys.v2.E2EKeyService/ListEntryHistory", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *e2EKeyServiceClient) UpdateEntry(ctx context.Context, in *UpdateEntryRequest, opts ...grpc.CallOption) (*UpdateEntryResponse, error) {
	out := new(UpdateEntryResponse)
	err := grpc.Invoke(ctx, "/google.security.e2ekeys.v2.E2EKeyService/UpdateEntry", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *e2EKeyServiceClient) ListSEH(ctx context.Context, in *ListSEHRequest, opts ...grpc.CallOption) (*ListSEHResponse, error) {
	out := new(ListSEHResponse)
	err := grpc.Invoke(ctx, "/google.security.e2ekeys.v2.E2EKeyService/ListSEH", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *e2EKeyServiceClient) ListUpdate(ctx context.Context, in *ListUpdateRequest, opts ...grpc.CallOption) (*ListUpdateResponse, error) {
	out := new(ListUpdateResponse)
	err := grpc.Invoke(ctx, "/google.security.e2ekeys.v2.E2EKeyService/ListUpdate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *e2EKeyServiceClient) ListSteps(ctx context.Context, in *ListStepsRequest, opts ...grpc.CallOption) (*ListStepsResponse, error) {
	out := new(ListStepsResponse)
	err := grpc.Invoke(ctx, "/google.security.e2ekeys.v2.E2EKeyService/ListSteps", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for E2EKeyService service

type E2EKeyServiceServer interface {
	// GetEntry returns a user's entry in the Merkle Tree. Entries contain
	// signed commitments to a profile, which is also returned.
	GetEntry(context.Context, *GetEntryRequest) (*GetEntryResponse, error)
	// ListEntryHistory returns a list of GetEntryRespons covering several epochs.
	ListEntryHistory(context.Context, *ListEntryHistoryRequest) (*ListEntryHistoryResponse, error)
	// blocking or polling?
	// UpdateEntry submits a SignedEntryUpdate.  Returns empty until this update
	// has been included in an epoch.  Clients must retry until this function
	// returns a proof.
	UpdateEntry(context.Context, *UpdateEntryRequest) (*UpdateEntryResponse, error)
	// List the Signed Epoch Heads, from epoch to epoch.
	ListSEH(context.Context, *ListSEHRequest) (*ListSEHResponse, error)
	// List the EntryUpdates by update number.
	ListUpdate(context.Context, *ListUpdateRequest) (*ListUpdateResponse, error)
	// ListSteps combines SEH and EntryUpdates into single list.
	ListSteps(context.Context, *ListStepsRequest) (*ListStepsResponse, error)
}

func RegisterE2EKeyServiceServer(s *grpc.Server, srv E2EKeyServiceServer) {
	s.RegisterService(&_E2EKeyService_serviceDesc, srv)
}

func _E2EKeyService_GetEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(GetEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(E2EKeyServiceServer).GetEntry(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _E2EKeyService_ListEntryHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ListEntryHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(E2EKeyServiceServer).ListEntryHistory(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _E2EKeyService_UpdateEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(UpdateEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(E2EKeyServiceServer).UpdateEntry(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _E2EKeyService_ListSEH_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ListSEHRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(E2EKeyServiceServer).ListSEH(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _E2EKeyService_ListUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ListUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(E2EKeyServiceServer).ListUpdate(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _E2EKeyService_ListSteps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ListStepsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(E2EKeyServiceServer).ListSteps(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _E2EKeyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.security.e2ekeys.v2.E2EKeyService",
	HandlerType: (*E2EKeyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEntry",
			Handler:    _E2EKeyService_GetEntry_Handler,
		},
		{
			MethodName: "ListEntryHistory",
			Handler:    _E2EKeyService_ListEntryHistory_Handler,
		},
		{
			MethodName: "UpdateEntry",
			Handler:    _E2EKeyService_UpdateEntry_Handler,
		},
		{
			MethodName: "ListSEH",
			Handler:    _E2EKeyService_ListSEH_Handler,
		},
		{
			MethodName: "ListUpdate",
			Handler:    _E2EKeyService_ListUpdate_Handler,
		},
		{
			MethodName: "ListSteps",
			Handler:    _E2EKeyService_ListSteps_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 982 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xbc, 0x56, 0x5b, 0x6f, 0x1b, 0x55,
	0x10, 0x66, 0xed, 0xac, 0x2f, 0xe3, 0xb4, 0x4d, 0x4f, 0x12, 0xba, 0x5d, 0x04, 0x6d, 0x57, 0x42,
	0xaa, 0x9a, 0x7a, 0x0d, 0xdb, 0x06, 0xca, 0x45, 0xaa, 0x28, 0x58, 0x18, 0x01, 0x55, 0x58, 0xb7,
	0x48, 0x3c, 0xad, 0xd6, 0xde, 0x13, 0x7b, 0x95, 0xec, 0x85, 0x3d, 0xc7, 0x51, 0x1d, 0x09, 0xf1,
	0xc4, 0xff, 0xe0, 0x81, 0x7f, 0xc4, 0xaf, 0xe0, 0x3f, 0xf0, 0xc0, 0x9c, 0xcb, 0xfa, 0x16, 0xe2,
	0x38, 0x2f, 0x7d, 0x88, 0xb4, 0x33, 0xf3, 0xcd, 0x99, 0xef, 0xcc, 0x7c, 0x73, 0x62, 0x70, 0xf3,
	0x22, 0xe3, 0x59, 0x67, 0x94, 0x65, 0xa3, 0x53, 0x1a, 0x30, 0x3a, 0x9c, 0x14, 0x31, 0x9f, 0x06,
	0xd4, 0xa3, 0x27, 0x74, 0xca, 0x82, 0x33, 0xaf, 0xa3, 0x3f, 0x15, 0x90, 0xd8, 0x0a, 0xe9, 0x96,
	0x48, 0xb7, 0x0c, 0x9f, 0x79, 0xf6, 0xf3, 0x51, 0xcc, 0xc7, 0x93, 0x81, 0x3b, 0xcc, 0x12, 0x7d,
	0xa0, 0x48, 0x6e, 0x63, 0xb8, 0xcd, 0x68, 0x71, 0x46, 0x8b, 0x8e, 0x2a, 0x36, 0xab, 0x32, 0xe4,
	0x49, 0x98, 0x77, 0xf0, 0x4f, 0x1d, 0xee, 0xfc, 0x5d, 0x81, 0x9d, 0x6f, 0x29, 0xef, 0xa6, 0xbc,
	0x98, 0xfa, 0x94, 0xe5, 0x59, 0xca, 0x28, 0xb9, 0x03, 0xf5, 0x09, 0xa6, 0x07, 0x71, 0x64, 0x19,
	0xf7, 0x8d, 0x87, 0x4d, 0xbf, 0x26, 0xcc, 0xef, 0x22, 0x72, 0x0f, 0x5a, 0x71, 0x1a, 0xd1, 0x37,
	0x01, 0x26, 0x67, 0xc7, 0x56, 0x05, 0x83, 0xdb, 0x3e, 0x48, 0xd7, 0x91, 0xf0, 0x90, 0x3d, 0x30,
	0xa5, 0x65, 0x55, 0x65, 0x48, 0x19, 0xe4, 0x47, 0x20, 0x2c, 0x1e, 0xa5, 0x34, 0x0a, 0x68, 0x9e,
	0x0d, 0xc7, 0xc1, 0x98, 0x86, 0x11, 0xb3, 0xb6, 0xee, 0x57, 0x1f, 0xb6, 0xbc, 0x7b, 0xee, 0x32,
	0x37, 0xb7, 0x2f, 0x91, 0x5d, 0x01, 0xec, 0x21, 0xce, 0xdf, 0x61, 0xcb, 0x0e, 0x46, 0x3c, 0xd8,
	0x4f, 0x68, 0x71, 0x82, 0xcd, 0xe3, 0x05, 0xa5, 0x41, 0x4a, 0xe3, 0xd1, 0x78, 0x90, 0x15, 0xcc,
	0x32, 0xf1, 0xc4, 0x6d, 0x7f, 0x57, 0x05, 0x5f, 0x61, 0xec, 0x65, 0x19, 0x22, 0x07, 0x60, 0x52,
	0x71, 0x47, 0xab, 0x86, 0xc4, 0x5a, 0xde, 0xfe, 0x6a, 0x55, 0xd5, 0x00, 0x85, 0x21, 0x16, 0xd4,
	0xf1, 0x82, 0xc7, 0xf1, 0x29, 0xb5, 0xea, 0xf2, 0x1e, 0xa5, 0x49, 0x3e, 0x84, 0x9b, 0xd8, 0xea,
	0x24, 0xe6, 0x09, 0x22, 0x03, 0x6c, 0xb3, 0xd5, 0x90, 0x80, 0x1b, 0x73, 0xef, 0xf7, 0x74, 0xea,
	0xfc, 0x61, 0x40, 0xfd, 0x48, 0xa7, 0x7c, 0x05, 0x5b, 0x62, 0x5a, 0xd8, 0x49, 0x71, 0xdd, 0xb6,
	0x7b, 0xf9, 0x34, 0x5d, 0x9d, 0xe2, 0xe2, 0x11, 0x4c, 0x11, 0x92, 0xa9, 0xf6, 0xa7, 0xd0, 0x9c,
	0xb9, 0xc8, 0x0e, 0x54, 0x45, 0x5d, 0x35, 0x18, 0xf1, 0x29, 0x9a, 0x7e, 0x16, 0x9e, 0x4e, 0xa8,
	0x9e, 0x87, 0x32, 0x3e, 0xaf, 0x3c, 0x33, 0x9c, 0xbf, 0x0c, 0x68, 0x1e, 0x4d, 0x06, 0xa7, 0xf1,
	0x10, 0xf3, 0x89, 0x0d, 0x75, 0x1a, 0x79, 0x87, 0x87, 0x1f, 0x7f, 0x26, 0xb3, 0xb7, 0x7b, 0xef,
	0xf8, 0xa5, 0x83, 0x7c, 0x01, 0x77, 0x0b, 0x16, 0x06, 0xa8, 0x98, 0xf8, 0x78, 0x1a, 0xa7, 0xa3,
	0x80, 0x8d, 0x43, 0xef, 0xf0, 0x93, 0xc0, 0xfb, 0xe8, 0xe9, 0x33, 0x75, 0x2e, 0xa2, 0xdf, 0x45,
	0xc8, 0xcf, 0x25, 0xa2, 0x2f, 0x01, 0x22, 0x8e, 0x03, 0xd9, 0xa3, 0xc3, 0x68, 0x29, 0x3d, 0xc7,
	0x98, 0x12, 0x01, 0xe6, 0x11, 0x19, 0x9d, 0x65, 0x1e, 0x61, 0xec, 0x05, 0x40, 0x03, 0xb9, 0x07,
	0x7c, 0x9a, 0x53, 0xe7, 0x1c, 0x6e, 0xcd, 0x35, 0xf8, 0xeb, 0x84, 0x32, 0x2e, 0xee, 0x24, 0xb5,
	0x22, 0x99, 0x56, 0x7d, 0x65, 0x2c, 0x0a, 0xb3, 0xb2, 0x24, 0xcc, 0x7d, 0xa8, 0x85, 0x79, 0x2e,
	0xfc, 0x55, 0xe9, 0x37, 0xd1, 0x42, 0xf7, 0x03, 0xd8, 0xe6, 0xc5, 0x84, 0x71, 0x54, 0x9e, 0x9c,
	0x81, 0x90, 0x5c, 0xcd, 0x6f, 0x69, 0x9f, 0xe8, 0xa9, 0x93, 0xc3, 0x9d, 0x1f, 0x62, 0xa6, 0x8a,
	0xf7, 0xf0, 0x23, 0x9b, 0x73, 0x58, 0xb7, 0x06, 0x8c, 0x87, 0x05, 0x57, 0x72, 0x96, 0x54, 0xaa,
	0x3e, 0x48, 0x97, 0x94, 0x29, 0x79, 0x0f, 0x9a, 0x79, 0x38, 0xc2, 0xe5, 0x8e, 0xcf, 0xa9, 0x64,
	0x64, 0xfa, 0x0d, 0xe1, 0xe8, 0xa3, 0xed, 0xfc, 0x0e, 0xd6, 0xc5, 0x8a, 0x7a, 0xf3, 0xbe, 0x81,
	0x9a, 0x9c, 0x5e, 0x29, 0x97, 0xc7, 0xeb, 0xe4, 0xb2, 0xba, 0xb7, 0xbe, 0xce, 0x25, 0xef, 0x03,
	0xa4, 0xf4, 0xcd, 0x32, 0xbd, 0xa6, 0xf0, 0x48, 0x76, 0xce, 0xbf, 0x06, 0x90, 0xd7, 0x79, 0x14,
	0x72, 0xba, 0xd4, 0xf2, 0x4b, 0xaf, 0xfb, 0x13, 0xec, 0x96, 0xeb, 0x2b, 0xf0, 0xc1, 0x44, 0xe6,
	0xca, 0x73, 0x5b, 0xde, 0x83, 0x4b, 0xf6, 0x57, 0x20, 0x55, 0x11, 0xff, 0x36, 0x5b, 0x75, 0x2d,
	0x6e, 0x58, 0xf5, 0xaa, 0x0d, 0xdb, 0xfa, 0x9f, 0x0d, 0x23, 0x77, 0xa1, 0x91, 0x66, 0x3c, 0x40,
	0x86, 0x91, 0x5c, 0xfb, 0x9a, 0x5f, 0x47, 0xfb, 0x35, 0x9a, 0xe4, 0x03, 0x80, 0xe8, 0x24, 0x4e,
	0xf4, 0x1b, 0xf5, 0x8f, 0xda, 0xe0, 0xa6, 0x70, 0xc9, 0x37, 0xca, 0xf9, 0x05, 0x76, 0x97, 0x6e,
	0xaf, 0x5b, 0xff, 0x02, 0x4c, 0x95, 0x61, 0xc8, 0x7b, 0x5d, 0xaf, 0xf3, 0x2a, 0xd5, 0x79, 0x09,
	0x37, 0xc5, 0x68, 0xfb, 0xdd, 0x5e, 0xd9, 0xd4, 0x15, 0xa9, 0x18, 0xeb, 0xa5, 0x52, 0x59, 0x91,
	0x4a, 0x0f, 0x6e, 0xcd, 0xce, 0xd3, 0x34, 0x0f, 0xc1, 0x54, 0xcf, 0xa7, 0xb1, 0xd9, 0xf3, 0xa9,
	0xd0, 0x4e, 0x0a, 0xb7, 0xc5, 0x49, 0x7a, 0x22, 0x9a, 0xdc, 0x97, 0x60, 0x2b, 0x72, 0x0b, 0x1d,
	0xe7, 0x71, 0x82, 0xa1, 0x30, 0xc9, 0x35, 0x57, 0x4b, 0x22, 0xbe, 0x9e, 0x01, 0x5e, 0x95, 0xf1,
	0xf5, 0xcc, 0x5d, 0x20, 0x8b, 0xf5, 0x34, 0x79, 0x1c, 0xbb, 0x12, 0x8f, 0xa2, 0x8f, 0x63, 0xd7,
	0xa6, 0x93, 0xc0, 0x8e, 0xbc, 0x29, 0xa7, 0x39, 0x7b, 0x0b, 0xf4, 0x9e, 0xab, 0x76, 0xe8, 0x72,
	0x9a, 0xdd, 0x23, 0x30, 0x99, 0x70, 0xe8, 0xd6, 0xee, 0x5d, 0x68, 0x2d, 0x06, 0x7d, 0x05, 0xf1,
	0xfe, 0x34, 0xe1, 0x46, 0xd7, 0xeb, 0xa2, 0x14, 0xfb, 0xf8, 0x7f, 0x36, 0x1e, 0x52, 0x42, 0xa1,
	0x51, 0xca, 0x82, 0x1c, 0x6c, 0x26, 0x1e, 0x79, 0x4d, 0xfb, 0x5a, 0x4a, 0x23, 0xbf, 0xa9, 0x46,
	0x2d, 0xbe, 0x1e, 0xe4, 0xc9, 0xba, 0x13, 0x2e, 0x79, 0xdd, 0xec, 0xa7, 0xd7, 0x4b, 0xd2, 0xe5,
	0x53, 0x68, 0x2d, 0x2c, 0x0f, 0x71, 0xd7, 0x1d, 0x72, 0xf1, 0x8d, 0xb1, 0x3b, 0x1b, 0xe3, 0x75,
	0xbd, 0x01, 0xd4, 0xf5, 0x06, 0x90, 0x47, 0x57, 0x11, 0x9e, 0xaf, 0x9d, 0x7d, 0xb0, 0x11, 0x56,
	0xd7, 0x38, 0x01, 0x98, 0x6b, 0x95, 0xb4, 0xaf, 0x4a, 0x5d, 0xda, 0x21, 0xdb, 0xdd, 0x14, 0xae,
	0x8b, 0x8d, 0xa1, 0x39, 0x53, 0x1e, 0x79, 0x7c, 0x25, 0xcd, 0x85, 0x7d, 0xb0, 0xdb, 0x1b, 0xa2,
	0x55, 0xa5, 0x41, 0x4d, 0xfe, 0xc2, 0x7b, 0xf2, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x08, 0xe3,
	0xdf, 0xa3, 0x70, 0x0a, 0x00, 0x00,
}