package groupv1beta1

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	_ "github.com/gogo/protobuf/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	io "io"
	reflect "reflect"
	sync "sync"
)

var _ protoreflect.List = (*_MsgCreateGroup_2_list)(nil)

type _MsgCreateGroup_2_list struct {
	list *[]*Member
}

func (x *_MsgCreateGroup_2_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_MsgCreateGroup_2_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_MsgCreateGroup_2_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*Member)
	(*x.list)[i] = concreteValue
}

func (x *_MsgCreateGroup_2_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*Member)
	*x.list = append(*x.list, concreteValue)
}

func (x *_MsgCreateGroup_2_list) AppendMutable() protoreflect.Value {
	v := new(Member)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_MsgCreateGroup_2_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_MsgCreateGroup_2_list) NewElement() protoreflect.Value {
	v := new(Member)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_MsgCreateGroup_2_list) IsValid() bool {
	return x.list != nil
}

var (
	md_MsgCreateGroup          protoreflect.MessageDescriptor
	fd_MsgCreateGroup_admin    protoreflect.FieldDescriptor
	fd_MsgCreateGroup_members  protoreflect.FieldDescriptor
	fd_MsgCreateGroup_metadata protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_group_v1beta1_tx_proto_init()
	md_MsgCreateGroup = File_cosmos_group_v1beta1_tx_proto.Messages().ByName("MsgCreateGroup")
	fd_MsgCreateGroup_admin = md_MsgCreateGroup.Fields().ByName("admin")
	fd_MsgCreateGroup_members = md_MsgCreateGroup.Fields().ByName("members")
	fd_MsgCreateGroup_metadata = md_MsgCreateGroup.Fields().ByName("metadata")
}

var _ protoreflect.Message = (*fastReflection_MsgCreateGroup)(nil)

type fastReflection_MsgCreateGroup MsgCreateGroup

func (x *MsgCreateGroup) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgCreateGroup)(x)
}

func (x *MsgCreateGroup) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_group_v1beta1_tx_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgCreateGroup_messageType fastReflection_MsgCreateGroup_messageType
var _ protoreflect.MessageType = fastReflection_MsgCreateGroup_messageType{}

type fastReflection_MsgCreateGroup_messageType struct{}

func (x fastReflection_MsgCreateGroup_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgCreateGroup)(nil)
}
func (x fastReflection_MsgCreateGroup_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgCreateGroup)
}
func (x fastReflection_MsgCreateGroup_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgCreateGroup
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgCreateGroup) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgCreateGroup
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgCreateGroup) Type() protoreflect.MessageType {
	return _fastReflection_MsgCreateGroup_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgCreateGroup) New() protoreflect.Message {
	return new(fastReflection_MsgCreateGroup)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgCreateGroup) Interface() protoreflect.ProtoMessage {
	return (*MsgCreateGroup)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgCreateGroup) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Admin != "" {
		value := protoreflect.ValueOfString(x.Admin)
		if !f(fd_MsgCreateGroup_admin, value) {
			return
		}
	}
	if len(x.Members) != 0 {
		value := protoreflect.ValueOfList(&_MsgCreateGroup_2_list{list: &x.Members})
		if !f(fd_MsgCreateGroup_members, value) {
			return
		}
	}
	if len(x.Metadata) != 0 {
		value := protoreflect.ValueOfBytes(x.Metadata)
		if !f(fd_MsgCreateGroup_metadata, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_MsgCreateGroup) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.group.v1beta1.MsgCreateGroup.admin":
		return x.Admin != ""
	case "cosmos.group.v1beta1.MsgCreateGroup.members":
		return len(x.Members) != 0
	case "cosmos.group.v1beta1.MsgCreateGroup.metadata":
		return len(x.Metadata) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgCreateGroup"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgCreateGroup does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateGroup) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.group.v1beta1.MsgCreateGroup.admin":
		x.Admin = ""
	case "cosmos.group.v1beta1.MsgCreateGroup.members":
		x.Members = nil
	case "cosmos.group.v1beta1.MsgCreateGroup.metadata":
		x.Metadata = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgCreateGroup"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgCreateGroup does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgCreateGroup) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.group.v1beta1.MsgCreateGroup.admin":
		value := x.Admin
		return protoreflect.ValueOfString(value)
	case "cosmos.group.v1beta1.MsgCreateGroup.members":
		if len(x.Members) == 0 {
			return protoreflect.ValueOfList(&_MsgCreateGroup_2_list{})
		}
		listValue := &_MsgCreateGroup_2_list{list: &x.Members}
		return protoreflect.ValueOfList(listValue)
	case "cosmos.group.v1beta1.MsgCreateGroup.metadata":
		value := x.Metadata
		return protoreflect.ValueOfBytes(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgCreateGroup"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgCreateGroup does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateGroup) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.group.v1beta1.MsgCreateGroup.admin":
		x.Admin = value.Interface().(string)
	case "cosmos.group.v1beta1.MsgCreateGroup.members":
		lv := value.List()
		clv := lv.(*_MsgCreateGroup_2_list)
		x.Members = *clv.list
	case "cosmos.group.v1beta1.MsgCreateGroup.metadata":
		x.Metadata = value.Bytes()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgCreateGroup"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgCreateGroup does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateGroup) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.group.v1beta1.MsgCreateGroup.members":
		if x.Members == nil {
			x.Members = []*Member{}
		}
		value := &_MsgCreateGroup_2_list{list: &x.Members}
		return protoreflect.ValueOfList(value)
	case "cosmos.group.v1beta1.MsgCreateGroup.admin":
		panic(fmt.Errorf("field admin of message cosmos.group.v1beta1.MsgCreateGroup is not mutable"))
	case "cosmos.group.v1beta1.MsgCreateGroup.metadata":
		panic(fmt.Errorf("field metadata of message cosmos.group.v1beta1.MsgCreateGroup is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgCreateGroup"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgCreateGroup does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgCreateGroup) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.group.v1beta1.MsgCreateGroup.admin":
		return protoreflect.ValueOfString("")
	case "cosmos.group.v1beta1.MsgCreateGroup.members":
		list := []*Member{}
		return protoreflect.ValueOfList(&_MsgCreateGroup_2_list{list: &list})
	case "cosmos.group.v1beta1.MsgCreateGroup.metadata":
		return protoreflect.ValueOfBytes(nil)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgCreateGroup"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgCreateGroup does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgCreateGroup) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.group.v1beta1.MsgCreateGroup", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgCreateGroup) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateGroup) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_MsgCreateGroup) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgCreateGroup) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgCreateGroup)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.Admin)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if len(x.Members) > 0 {
			for _, e := range x.Members {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		l = len(x.Metadata)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*MsgCreateGroup)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.Metadata) > 0 {
			i -= len(x.Metadata)
			copy(dAtA[i:], x.Metadata)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Metadata)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.Members) > 0 {
			for iNdEx := len(x.Members) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Members[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0x12
			}
		}
		if len(x.Admin) > 0 {
			i -= len(x.Admin)
			copy(dAtA[i:], x.Admin)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Admin)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*MsgCreateGroup)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgCreateGroup: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgCreateGroup: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Admin", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Admin = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Members", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Members = append(x.Members, &Member{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Members[len(x.Members)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Metadata = append(x.Metadata[:0], dAtA[iNdEx:postIndex]...)
				if x.Metadata == nil {
					x.Metadata = []byte{}
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_MsgCreateGroupResponse          protoreflect.MessageDescriptor
	fd_MsgCreateGroupResponse_group_id protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_group_v1beta1_tx_proto_init()
	md_MsgCreateGroupResponse = File_cosmos_group_v1beta1_tx_proto.Messages().ByName("MsgCreateGroupResponse")
	fd_MsgCreateGroupResponse_group_id = md_MsgCreateGroupResponse.Fields().ByName("group_id")
}

var _ protoreflect.Message = (*fastReflection_MsgCreateGroupResponse)(nil)

type fastReflection_MsgCreateGroupResponse MsgCreateGroupResponse

func (x *MsgCreateGroupResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgCreateGroupResponse)(x)
}

func (x *MsgCreateGroupResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_group_v1beta1_tx_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgCreateGroupResponse_messageType fastReflection_MsgCreateGroupResponse_messageType
var _ protoreflect.MessageType = fastReflection_MsgCreateGroupResponse_messageType{}

type fastReflection_MsgCreateGroupResponse_messageType struct{}

func (x fastReflection_MsgCreateGroupResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgCreateGroupResponse)(nil)
}
func (x fastReflection_MsgCreateGroupResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgCreateGroupResponse)
}
func (x fastReflection_MsgCreateGroupResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgCreateGroupResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgCreateGroupResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgCreateGroupResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgCreateGroupResponse) Type() protoreflect.MessageType {
	return _fastReflection_MsgCreateGroupResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgCreateGroupResponse) New() protoreflect.Message {
	return new(fastReflection_MsgCreateGroupResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgCreateGroupResponse) Interface() protoreflect.ProtoMessage {
	return (*MsgCreateGroupResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgCreateGroupResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.GroupId != uint64(0) {
		value := protoreflect.ValueOfUint64(x.GroupId)
		if !f(fd_MsgCreateGroupResponse_group_id, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_MsgCreateGroupResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.group.v1beta1.MsgCreateGroupResponse.group_id":
		return x.GroupId != uint64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgCreateGroupResponse"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgCreateGroupResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateGroupResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.group.v1beta1.MsgCreateGroupResponse.group_id":
		x.GroupId = uint64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgCreateGroupResponse"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgCreateGroupResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgCreateGroupResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.group.v1beta1.MsgCreateGroupResponse.group_id":
		value := x.GroupId
		return protoreflect.ValueOfUint64(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgCreateGroupResponse"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgCreateGroupResponse does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateGroupResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.group.v1beta1.MsgCreateGroupResponse.group_id":
		x.GroupId = value.Uint()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgCreateGroupResponse"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgCreateGroupResponse does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateGroupResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.group.v1beta1.MsgCreateGroupResponse.group_id":
		panic(fmt.Errorf("field group_id of message cosmos.group.v1beta1.MsgCreateGroupResponse is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgCreateGroupResponse"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgCreateGroupResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgCreateGroupResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.group.v1beta1.MsgCreateGroupResponse.group_id":
		return protoreflect.ValueOfUint64(uint64(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgCreateGroupResponse"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgCreateGroupResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgCreateGroupResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.group.v1beta1.MsgCreateGroupResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgCreateGroupResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateGroupResponse) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_MsgCreateGroupResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgCreateGroupResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgCreateGroupResponse)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.GroupId != 0 {
			n += 1 + runtime.Sov(uint64(x.GroupId))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*MsgCreateGroupResponse)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.GroupId != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.GroupId))
			i--
			dAtA[i] = 0x8
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*MsgCreateGroupResponse)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgCreateGroupResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgCreateGroupResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field GroupId", wireType)
				}
				x.GroupId = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.GroupId |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var _ protoreflect.List = (*_MsgUpdateGroupMembers_3_list)(nil)

type _MsgUpdateGroupMembers_3_list struct {
	list *[]*Member
}

func (x *_MsgUpdateGroupMembers_3_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_MsgUpdateGroupMembers_3_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_MsgUpdateGroupMembers_3_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*Member)
	(*x.list)[i] = concreteValue
}

func (x *_MsgUpdateGroupMembers_3_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*Member)
	*x.list = append(*x.list, concreteValue)
}

func (x *_MsgUpdateGroupMembers_3_list) AppendMutable() protoreflect.Value {
	v := new(Member)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_MsgUpdateGroupMembers_3_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_MsgUpdateGroupMembers_3_list) NewElement() protoreflect.Value {
	v := new(Member)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_MsgUpdateGroupMembers_3_list) IsValid() bool {
	return x.list != nil
}

var (
	md_MsgUpdateGroupMembers                protoreflect.MessageDescriptor
	fd_MsgUpdateGroupMembers_admin          protoreflect.FieldDescriptor
	fd_MsgUpdateGroupMembers_group_id       protoreflect.FieldDescriptor
	fd_MsgUpdateGroupMembers_member_updates protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_group_v1beta1_tx_proto_init()
	md_MsgUpdateGroupMembers = File_cosmos_group_v1beta1_tx_proto.Messages().ByName("MsgUpdateGroupMembers")
	fd_MsgUpdateGroupMembers_admin = md_MsgUpdateGroupMembers.Fields().ByName("admin")
	fd_MsgUpdateGroupMembers_group_id = md_MsgUpdateGroupMembers.Fields().ByName("group_id")
	fd_MsgUpdateGroupMembers_member_updates = md_MsgUpdateGroupMembers.Fields().ByName("member_updates")
}

var _ protoreflect.Message = (*fastReflection_MsgUpdateGroupMembers)(nil)

type fastReflection_MsgUpdateGroupMembers MsgUpdateGroupMembers

func (x *MsgUpdateGroupMembers) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgUpdateGroupMembers)(x)
}

func (x *MsgUpdateGroupMembers) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_group_v1beta1_tx_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgUpdateGroupMembers_messageType fastReflection_MsgUpdateGroupMembers_messageType
var _ protoreflect.MessageType = fastReflection_MsgUpdateGroupMembers_messageType{}

type fastReflection_MsgUpdateGroupMembers_messageType struct{}

func (x fastReflection_MsgUpdateGroupMembers_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgUpdateGroupMembers)(nil)
}
func (x fastReflection_MsgUpdateGroupMembers_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgUpdateGroupMembers)
}
func (x fastReflection_MsgUpdateGroupMembers_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgUpdateGroupMembers
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgUpdateGroupMembers) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgUpdateGroupMembers
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgUpdateGroupMembers) Type() protoreflect.MessageType {
	return _fastReflection_MsgUpdateGroupMembers_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgUpdateGroupMembers) New() protoreflect.Message {
	return new(fastReflection_MsgUpdateGroupMembers)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgUpdateGroupMembers) Interface() protoreflect.ProtoMessage {
	return (*MsgUpdateGroupMembers)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgUpdateGroupMembers) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Admin != "" {
		value := protoreflect.ValueOfString(x.Admin)
		if !f(fd_MsgUpdateGroupMembers_admin, value) {
			return
		}
	}
	if x.GroupId != uint64(0) {
		value := protoreflect.ValueOfUint64(x.GroupId)
		if !f(fd_MsgUpdateGroupMembers_group_id, value) {
			return
		}
	}
	if len(x.MemberUpdates) != 0 {
		value := protoreflect.ValueOfList(&_MsgUpdateGroupMembers_3_list{list: &x.MemberUpdates})
		if !f(fd_MsgUpdateGroupMembers_member_updates, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_MsgUpdateGroupMembers) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.group.v1beta1.MsgUpdateGroupMembers.admin":
		return x.Admin != ""
	case "cosmos.group.v1beta1.MsgUpdateGroupMembers.group_id":
		return x.GroupId != uint64(0)
	case "cosmos.group.v1beta1.MsgUpdateGroupMembers.member_updates":
		return len(x.MemberUpdates) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgUpdateGroupMembers"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgUpdateGroupMembers does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupMembers) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.group.v1beta1.MsgUpdateGroupMembers.admin":
		x.Admin = ""
	case "cosmos.group.v1beta1.MsgUpdateGroupMembers.group_id":
		x.GroupId = uint64(0)
	case "cosmos.group.v1beta1.MsgUpdateGroupMembers.member_updates":
		x.MemberUpdates = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgUpdateGroupMembers"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgUpdateGroupMembers does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgUpdateGroupMembers) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.group.v1beta1.MsgUpdateGroupMembers.admin":
		value := x.Admin
		return protoreflect.ValueOfString(value)
	case "cosmos.group.v1beta1.MsgUpdateGroupMembers.group_id":
		value := x.GroupId
		return protoreflect.ValueOfUint64(value)
	case "cosmos.group.v1beta1.MsgUpdateGroupMembers.member_updates":
		if len(x.MemberUpdates) == 0 {
			return protoreflect.ValueOfList(&_MsgUpdateGroupMembers_3_list{})
		}
		listValue := &_MsgUpdateGroupMembers_3_list{list: &x.MemberUpdates}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgUpdateGroupMembers"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgUpdateGroupMembers does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupMembers) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.group.v1beta1.MsgUpdateGroupMembers.admin":
		x.Admin = value.Interface().(string)
	case "cosmos.group.v1beta1.MsgUpdateGroupMembers.group_id":
		x.GroupId = value.Uint()
	case "cosmos.group.v1beta1.MsgUpdateGroupMembers.member_updates":
		lv := value.List()
		clv := lv.(*_MsgUpdateGroupMembers_3_list)
		x.MemberUpdates = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgUpdateGroupMembers"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgUpdateGroupMembers does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupMembers) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.group.v1beta1.MsgUpdateGroupMembers.member_updates":
		if x.MemberUpdates == nil {
			x.MemberUpdates = []*Member{}
		}
		value := &_MsgUpdateGroupMembers_3_list{list: &x.MemberUpdates}
		return protoreflect.ValueOfList(value)
	case "cosmos.group.v1beta1.MsgUpdateGroupMembers.admin":
		panic(fmt.Errorf("field admin of message cosmos.group.v1beta1.MsgUpdateGroupMembers is not mutable"))
	case "cosmos.group.v1beta1.MsgUpdateGroupMembers.group_id":
		panic(fmt.Errorf("field group_id of message cosmos.group.v1beta1.MsgUpdateGroupMembers is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgUpdateGroupMembers"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgUpdateGroupMembers does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgUpdateGroupMembers) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.group.v1beta1.MsgUpdateGroupMembers.admin":
		return protoreflect.ValueOfString("")
	case "cosmos.group.v1beta1.MsgUpdateGroupMembers.group_id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "cosmos.group.v1beta1.MsgUpdateGroupMembers.member_updates":
		list := []*Member{}
		return protoreflect.ValueOfList(&_MsgUpdateGroupMembers_3_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgUpdateGroupMembers"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgUpdateGroupMembers does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgUpdateGroupMembers) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.group.v1beta1.MsgUpdateGroupMembers", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgUpdateGroupMembers) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupMembers) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_MsgUpdateGroupMembers) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgUpdateGroupMembers) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgUpdateGroupMembers)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.Admin)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.GroupId != 0 {
			n += 1 + runtime.Sov(uint64(x.GroupId))
		}
		if len(x.MemberUpdates) > 0 {
			for _, e := range x.MemberUpdates {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*MsgUpdateGroupMembers)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.MemberUpdates) > 0 {
			for iNdEx := len(x.MemberUpdates) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.MemberUpdates[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0x1a
			}
		}
		if x.GroupId != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.GroupId))
			i--
			dAtA[i] = 0x10
		}
		if len(x.Admin) > 0 {
			i -= len(x.Admin)
			copy(dAtA[i:], x.Admin)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Admin)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*MsgUpdateGroupMembers)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgUpdateGroupMembers: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgUpdateGroupMembers: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Admin", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Admin = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field GroupId", wireType)
				}
				x.GroupId = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.GroupId |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field MemberUpdates", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.MemberUpdates = append(x.MemberUpdates, &Member{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.MemberUpdates[len(x.MemberUpdates)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_MsgUpdateGroupMembersResponse protoreflect.MessageDescriptor
)

func init() {
	file_cosmos_group_v1beta1_tx_proto_init()
	md_MsgUpdateGroupMembersResponse = File_cosmos_group_v1beta1_tx_proto.Messages().ByName("MsgUpdateGroupMembersResponse")
}

var _ protoreflect.Message = (*fastReflection_MsgUpdateGroupMembersResponse)(nil)

type fastReflection_MsgUpdateGroupMembersResponse MsgUpdateGroupMembersResponse

func (x *MsgUpdateGroupMembersResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgUpdateGroupMembersResponse)(x)
}

func (x *MsgUpdateGroupMembersResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_group_v1beta1_tx_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgUpdateGroupMembersResponse_messageType fastReflection_MsgUpdateGroupMembersResponse_messageType
var _ protoreflect.MessageType = fastReflection_MsgUpdateGroupMembersResponse_messageType{}

type fastReflection_MsgUpdateGroupMembersResponse_messageType struct{}

func (x fastReflection_MsgUpdateGroupMembersResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgUpdateGroupMembersResponse)(nil)
}
func (x fastReflection_MsgUpdateGroupMembersResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgUpdateGroupMembersResponse)
}
func (x fastReflection_MsgUpdateGroupMembersResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgUpdateGroupMembersResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgUpdateGroupMembersResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgUpdateGroupMembersResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgUpdateGroupMembersResponse) Type() protoreflect.MessageType {
	return _fastReflection_MsgUpdateGroupMembersResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgUpdateGroupMembersResponse) New() protoreflect.Message {
	return new(fastReflection_MsgUpdateGroupMembersResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgUpdateGroupMembersResponse) Interface() protoreflect.ProtoMessage {
	return (*MsgUpdateGroupMembersResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgUpdateGroupMembersResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_MsgUpdateGroupMembersResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgUpdateGroupMembersResponse"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgUpdateGroupMembersResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupMembersResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgUpdateGroupMembersResponse"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgUpdateGroupMembersResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgUpdateGroupMembersResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgUpdateGroupMembersResponse"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgUpdateGroupMembersResponse does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupMembersResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgUpdateGroupMembersResponse"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgUpdateGroupMembersResponse does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupMembersResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgUpdateGroupMembersResponse"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgUpdateGroupMembersResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgUpdateGroupMembersResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgUpdateGroupMembersResponse"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgUpdateGroupMembersResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgUpdateGroupMembersResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.group.v1beta1.MsgUpdateGroupMembersResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgUpdateGroupMembersResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupMembersResponse) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_MsgUpdateGroupMembersResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgUpdateGroupMembersResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgUpdateGroupMembersResponse)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*MsgUpdateGroupMembersResponse)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*MsgUpdateGroupMembersResponse)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgUpdateGroupMembersResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgUpdateGroupMembersResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_MsgUpdateGroupAdmin           protoreflect.MessageDescriptor
	fd_MsgUpdateGroupAdmin_admin     protoreflect.FieldDescriptor
	fd_MsgUpdateGroupAdmin_group_id  protoreflect.FieldDescriptor
	fd_MsgUpdateGroupAdmin_new_admin protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_group_v1beta1_tx_proto_init()
	md_MsgUpdateGroupAdmin = File_cosmos_group_v1beta1_tx_proto.Messages().ByName("MsgUpdateGroupAdmin")
	fd_MsgUpdateGroupAdmin_admin = md_MsgUpdateGroupAdmin.Fields().ByName("admin")
	fd_MsgUpdateGroupAdmin_group_id = md_MsgUpdateGroupAdmin.Fields().ByName("group_id")
	fd_MsgUpdateGroupAdmin_new_admin = md_MsgUpdateGroupAdmin.Fields().ByName("new_admin")
}

var _ protoreflect.Message = (*fastReflection_MsgUpdateGroupAdmin)(nil)

type fastReflection_MsgUpdateGroupAdmin MsgUpdateGroupAdmin

func (x *MsgUpdateGroupAdmin) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgUpdateGroupAdmin)(x)
}

func (x *MsgUpdateGroupAdmin) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_group_v1beta1_tx_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgUpdateGroupAdmin_messageType fastReflection_MsgUpdateGroupAdmin_messageType
var _ protoreflect.MessageType = fastReflection_MsgUpdateGroupAdmin_messageType{}

type fastReflection_MsgUpdateGroupAdmin_messageType struct{}

func (x fastReflection_MsgUpdateGroupAdmin_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgUpdateGroupAdmin)(nil)
}
func (x fastReflection_MsgUpdateGroupAdmin_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgUpdateGroupAdmin)
}
func (x fastReflection_MsgUpdateGroupAdmin_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgUpdateGroupAdmin
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgUpdateGroupAdmin) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgUpdateGroupAdmin
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgUpdateGroupAdmin) Type() protoreflect.MessageType {
	return _fastReflection_MsgUpdateGroupAdmin_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgUpdateGroupAdmin) New() protoreflect.Message {
	return new(fastReflection_MsgUpdateGroupAdmin)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgUpdateGroupAdmin) Interface() protoreflect.ProtoMessage {
	return (*MsgUpdateGroupAdmin)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgUpdateGroupAdmin) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Admin != "" {
		value := protoreflect.ValueOfString(x.Admin)
		if !f(fd_MsgUpdateGroupAdmin_admin, value) {
			return
		}
	}
	if x.GroupId != uint64(0) {
		value := protoreflect.ValueOfUint64(x.GroupId)
		if !f(fd_MsgUpdateGroupAdmin_group_id, value) {
			return
		}
	}
	if x.NewAdmin != "" {
		value := protoreflect.ValueOfString(x.NewAdmin)
		if !f(fd_MsgUpdateGroupAdmin_new_admin, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_MsgUpdateGroupAdmin) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.group.v1beta1.MsgUpdateGroupAdmin.admin":
		return x.Admin != ""
	case "cosmos.group.v1beta1.MsgUpdateGroupAdmin.group_id":
		return x.GroupId != uint64(0)
	case "cosmos.group.v1beta1.MsgUpdateGroupAdmin.new_admin":
		return x.NewAdmin != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgUpdateGroupAdmin"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgUpdateGroupAdmin does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupAdmin) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.group.v1beta1.MsgUpdateGroupAdmin.admin":
		x.Admin = ""
	case "cosmos.group.v1beta1.MsgUpdateGroupAdmin.group_id":
		x.GroupId = uint64(0)
	case "cosmos.group.v1beta1.MsgUpdateGroupAdmin.new_admin":
		x.NewAdmin = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgUpdateGroupAdmin"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgUpdateGroupAdmin does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgUpdateGroupAdmin) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.group.v1beta1.MsgUpdateGroupAdmin.admin":
		value := x.Admin
		return protoreflect.ValueOfString(value)
	case "cosmos.group.v1beta1.MsgUpdateGroupAdmin.group_id":
		value := x.GroupId
		return protoreflect.ValueOfUint64(value)
	case "cosmos.group.v1beta1.MsgUpdateGroupAdmin.new_admin":
		value := x.NewAdmin
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgUpdateGroupAdmin"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgUpdateGroupAdmin does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupAdmin) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.group.v1beta1.MsgUpdateGroupAdmin.admin":
		x.Admin = value.Interface().(string)
	case "cosmos.group.v1beta1.MsgUpdateGroupAdmin.group_id":
		x.GroupId = value.Uint()
	case "cosmos.group.v1beta1.MsgUpdateGroupAdmin.new_admin":
		x.NewAdmin = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgUpdateGroupAdmin"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgUpdateGroupAdmin does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupAdmin) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.group.v1beta1.MsgUpdateGroupAdmin.admin":
		panic(fmt.Errorf("field admin of message cosmos.group.v1beta1.MsgUpdateGroupAdmin is not mutable"))
	case "cosmos.group.v1beta1.MsgUpdateGroupAdmin.group_id":
		panic(fmt.Errorf("field group_id of message cosmos.group.v1beta1.MsgUpdateGroupAdmin is not mutable"))
	case "cosmos.group.v1beta1.MsgUpdateGroupAdmin.new_admin":
		panic(fmt.Errorf("field new_admin of message cosmos.group.v1beta1.MsgUpdateGroupAdmin is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgUpdateGroupAdmin"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgUpdateGroupAdmin does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgUpdateGroupAdmin) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.group.v1beta1.MsgUpdateGroupAdmin.admin":
		return protoreflect.ValueOfString("")
	case "cosmos.group.v1beta1.MsgUpdateGroupAdmin.group_id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "cosmos.group.v1beta1.MsgUpdateGroupAdmin.new_admin":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgUpdateGroupAdmin"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgUpdateGroupAdmin does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgUpdateGroupAdmin) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.group.v1beta1.MsgUpdateGroupAdmin", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgUpdateGroupAdmin) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupAdmin) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_MsgUpdateGroupAdmin) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgUpdateGroupAdmin) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgUpdateGroupAdmin)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.Admin)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.GroupId != 0 {
			n += 1 + runtime.Sov(uint64(x.GroupId))
		}
		l = len(x.NewAdmin)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*MsgUpdateGroupAdmin)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.NewAdmin) > 0 {
			i -= len(x.NewAdmin)
			copy(dAtA[i:], x.NewAdmin)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.NewAdmin)))
			i--
			dAtA[i] = 0x1a
		}
		if x.GroupId != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.GroupId))
			i--
			dAtA[i] = 0x10
		}
		if len(x.Admin) > 0 {
			i -= len(x.Admin)
			copy(dAtA[i:], x.Admin)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Admin)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*MsgUpdateGroupAdmin)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgUpdateGroupAdmin: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgUpdateGroupAdmin: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Admin", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Admin = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field GroupId", wireType)
				}
				x.GroupId = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.GroupId |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field NewAdmin", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.NewAdmin = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_MsgUpdateGroupAdminResponse protoreflect.MessageDescriptor
)

func init() {
	file_cosmos_group_v1beta1_tx_proto_init()
	md_MsgUpdateGroupAdminResponse = File_cosmos_group_v1beta1_tx_proto.Messages().ByName("MsgUpdateGroupAdminResponse")
}

var _ protoreflect.Message = (*fastReflection_MsgUpdateGroupAdminResponse)(nil)

type fastReflection_MsgUpdateGroupAdminResponse MsgUpdateGroupAdminResponse

func (x *MsgUpdateGroupAdminResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgUpdateGroupAdminResponse)(x)
}

func (x *MsgUpdateGroupAdminResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_group_v1beta1_tx_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgUpdateGroupAdminResponse_messageType fastReflection_MsgUpdateGroupAdminResponse_messageType
var _ protoreflect.MessageType = fastReflection_MsgUpdateGroupAdminResponse_messageType{}

type fastReflection_MsgUpdateGroupAdminResponse_messageType struct{}

func (x fastReflection_MsgUpdateGroupAdminResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgUpdateGroupAdminResponse)(nil)
}
func (x fastReflection_MsgUpdateGroupAdminResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgUpdateGroupAdminResponse)
}
func (x fastReflection_MsgUpdateGroupAdminResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgUpdateGroupAdminResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgUpdateGroupAdminResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgUpdateGroupAdminResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgUpdateGroupAdminResponse) Type() protoreflect.MessageType {
	return _fastReflection_MsgUpdateGroupAdminResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgUpdateGroupAdminResponse) New() protoreflect.Message {
	return new(fastReflection_MsgUpdateGroupAdminResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgUpdateGroupAdminResponse) Interface() protoreflect.ProtoMessage {
	return (*MsgUpdateGroupAdminResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgUpdateGroupAdminResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_MsgUpdateGroupAdminResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgUpdateGroupAdminResponse"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgUpdateGroupAdminResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupAdminResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgUpdateGroupAdminResponse"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgUpdateGroupAdminResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgUpdateGroupAdminResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgUpdateGroupAdminResponse"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgUpdateGroupAdminResponse does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupAdminResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgUpdateGroupAdminResponse"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgUpdateGroupAdminResponse does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupAdminResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgUpdateGroupAdminResponse"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgUpdateGroupAdminResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgUpdateGroupAdminResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgUpdateGroupAdminResponse"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgUpdateGroupAdminResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgUpdateGroupAdminResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.group.v1beta1.MsgUpdateGroupAdminResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgUpdateGroupAdminResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupAdminResponse) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_MsgUpdateGroupAdminResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgUpdateGroupAdminResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgUpdateGroupAdminResponse)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*MsgUpdateGroupAdminResponse)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*MsgUpdateGroupAdminResponse)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgUpdateGroupAdminResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgUpdateGroupAdminResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_MsgUpdateGroupMetadata          protoreflect.MessageDescriptor
	fd_MsgUpdateGroupMetadata_admin    protoreflect.FieldDescriptor
	fd_MsgUpdateGroupMetadata_group_id protoreflect.FieldDescriptor
	fd_MsgUpdateGroupMetadata_metadata protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_group_v1beta1_tx_proto_init()
	md_MsgUpdateGroupMetadata = File_cosmos_group_v1beta1_tx_proto.Messages().ByName("MsgUpdateGroupMetadata")
	fd_MsgUpdateGroupMetadata_admin = md_MsgUpdateGroupMetadata.Fields().ByName("admin")
	fd_MsgUpdateGroupMetadata_group_id = md_MsgUpdateGroupMetadata.Fields().ByName("group_id")
	fd_MsgUpdateGroupMetadata_metadata = md_MsgUpdateGroupMetadata.Fields().ByName("metadata")
}

var _ protoreflect.Message = (*fastReflection_MsgUpdateGroupMetadata)(nil)

type fastReflection_MsgUpdateGroupMetadata MsgUpdateGroupMetadata

func (x *MsgUpdateGroupMetadata) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgUpdateGroupMetadata)(x)
}

func (x *MsgUpdateGroupMetadata) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_group_v1beta1_tx_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgUpdateGroupMetadata_messageType fastReflection_MsgUpdateGroupMetadata_messageType
var _ protoreflect.MessageType = fastReflection_MsgUpdateGroupMetadata_messageType{}

type fastReflection_MsgUpdateGroupMetadata_messageType struct{}

func (x fastReflection_MsgUpdateGroupMetadata_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgUpdateGroupMetadata)(nil)
}
func (x fastReflection_MsgUpdateGroupMetadata_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgUpdateGroupMetadata)
}
func (x fastReflection_MsgUpdateGroupMetadata_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgUpdateGroupMetadata
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgUpdateGroupMetadata) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgUpdateGroupMetadata
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgUpdateGroupMetadata) Type() protoreflect.MessageType {
	return _fastReflection_MsgUpdateGroupMetadata_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgUpdateGroupMetadata) New() protoreflect.Message {
	return new(fastReflection_MsgUpdateGroupMetadata)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgUpdateGroupMetadata) Interface() protoreflect.ProtoMessage {
	return (*MsgUpdateGroupMetadata)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgUpdateGroupMetadata) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Admin != "" {
		value := protoreflect.ValueOfString(x.Admin)
		if !f(fd_MsgUpdateGroupMetadata_admin, value) {
			return
		}
	}
	if x.GroupId != uint64(0) {
		value := protoreflect.ValueOfUint64(x.GroupId)
		if !f(fd_MsgUpdateGroupMetadata_group_id, value) {
			return
		}
	}
	if len(x.Metadata) != 0 {
		value := protoreflect.ValueOfBytes(x.Metadata)
		if !f(fd_MsgUpdateGroupMetadata_metadata, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_MsgUpdateGroupMetadata) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.group.v1beta1.MsgUpdateGroupMetadata.admin":
		return x.Admin != ""
	case "cosmos.group.v1beta1.MsgUpdateGroupMetadata.group_id":
		return x.GroupId != uint64(0)
	case "cosmos.group.v1beta1.MsgUpdateGroupMetadata.metadata":
		return len(x.Metadata) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgUpdateGroupMetadata"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgUpdateGroupMetadata does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupMetadata) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.group.v1beta1.MsgUpdateGroupMetadata.admin":
		x.Admin = ""
	case "cosmos.group.v1beta1.MsgUpdateGroupMetadata.group_id":
		x.GroupId = uint64(0)
	case "cosmos.group.v1beta1.MsgUpdateGroupMetadata.metadata":
		x.Metadata = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgUpdateGroupMetadata"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgUpdateGroupMetadata does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgUpdateGroupMetadata) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.group.v1beta1.MsgUpdateGroupMetadata.admin":
		value := x.Admin
		return protoreflect.ValueOfString(value)
	case "cosmos.group.v1beta1.MsgUpdateGroupMetadata.group_id":
		value := x.GroupId
		return protoreflect.ValueOfUint64(value)
	case "cosmos.group.v1beta1.MsgUpdateGroupMetadata.metadata":
		value := x.Metadata
		return protoreflect.ValueOfBytes(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgUpdateGroupMetadata"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgUpdateGroupMetadata does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupMetadata) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.group.v1beta1.MsgUpdateGroupMetadata.admin":
		x.Admin = value.Interface().(string)
	case "cosmos.group.v1beta1.MsgUpdateGroupMetadata.group_id":
		x.GroupId = value.Uint()
	case "cosmos.group.v1beta1.MsgUpdateGroupMetadata.metadata":
		x.Metadata = value.Bytes()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgUpdateGroupMetadata"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgUpdateGroupMetadata does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupMetadata) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.group.v1beta1.MsgUpdateGroupMetadata.admin":
		panic(fmt.Errorf("field admin of message cosmos.group.v1beta1.MsgUpdateGroupMetadata is not mutable"))
	case "cosmos.group.v1beta1.MsgUpdateGroupMetadata.group_id":
		panic(fmt.Errorf("field group_id of message cosmos.group.v1beta1.MsgUpdateGroupMetadata is not mutable"))
	case "cosmos.group.v1beta1.MsgUpdateGroupMetadata.metadata":
		panic(fmt.Errorf("field metadata of message cosmos.group.v1beta1.MsgUpdateGroupMetadata is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgUpdateGroupMetadata"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgUpdateGroupMetadata does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgUpdateGroupMetadata) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.group.v1beta1.MsgUpdateGroupMetadata.admin":
		return protoreflect.ValueOfString("")
	case "cosmos.group.v1beta1.MsgUpdateGroupMetadata.group_id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "cosmos.group.v1beta1.MsgUpdateGroupMetadata.metadata":
		return protoreflect.ValueOfBytes(nil)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgUpdateGroupMetadata"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgUpdateGroupMetadata does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgUpdateGroupMetadata) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.group.v1beta1.MsgUpdateGroupMetadata", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgUpdateGroupMetadata) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupMetadata) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_MsgUpdateGroupMetadata) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgUpdateGroupMetadata) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgUpdateGroupMetadata)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.Admin)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.GroupId != 0 {
			n += 1 + runtime.Sov(uint64(x.GroupId))
		}
		l = len(x.Metadata)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*MsgUpdateGroupMetadata)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.Metadata) > 0 {
			i -= len(x.Metadata)
			copy(dAtA[i:], x.Metadata)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Metadata)))
			i--
			dAtA[i] = 0x1a
		}
		if x.GroupId != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.GroupId))
			i--
			dAtA[i] = 0x10
		}
		if len(x.Admin) > 0 {
			i -= len(x.Admin)
			copy(dAtA[i:], x.Admin)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Admin)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*MsgUpdateGroupMetadata)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgUpdateGroupMetadata: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgUpdateGroupMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Admin", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Admin = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field GroupId", wireType)
				}
				x.GroupId = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.GroupId |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Metadata = append(x.Metadata[:0], dAtA[iNdEx:postIndex]...)
				if x.Metadata == nil {
					x.Metadata = []byte{}
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_MsgUpdateGroupMetadataResponse protoreflect.MessageDescriptor
)

func init() {
	file_cosmos_group_v1beta1_tx_proto_init()
	md_MsgUpdateGroupMetadataResponse = File_cosmos_group_v1beta1_tx_proto.Messages().ByName("MsgUpdateGroupMetadataResponse")
}

var _ protoreflect.Message = (*fastReflection_MsgUpdateGroupMetadataResponse)(nil)

type fastReflection_MsgUpdateGroupMetadataResponse MsgUpdateGroupMetadataResponse

func (x *MsgUpdateGroupMetadataResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgUpdateGroupMetadataResponse)(x)
}

func (x *MsgUpdateGroupMetadataResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_group_v1beta1_tx_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgUpdateGroupMetadataResponse_messageType fastReflection_MsgUpdateGroupMetadataResponse_messageType
var _ protoreflect.MessageType = fastReflection_MsgUpdateGroupMetadataResponse_messageType{}

type fastReflection_MsgUpdateGroupMetadataResponse_messageType struct{}

func (x fastReflection_MsgUpdateGroupMetadataResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgUpdateGroupMetadataResponse)(nil)
}
func (x fastReflection_MsgUpdateGroupMetadataResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgUpdateGroupMetadataResponse)
}
func (x fastReflection_MsgUpdateGroupMetadataResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgUpdateGroupMetadataResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgUpdateGroupMetadataResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgUpdateGroupMetadataResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgUpdateGroupMetadataResponse) Type() protoreflect.MessageType {
	return _fastReflection_MsgUpdateGroupMetadataResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgUpdateGroupMetadataResponse) New() protoreflect.Message {
	return new(fastReflection_MsgUpdateGroupMetadataResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgUpdateGroupMetadataResponse) Interface() protoreflect.ProtoMessage {
	return (*MsgUpdateGroupMetadataResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgUpdateGroupMetadataResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_MsgUpdateGroupMetadataResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgUpdateGroupMetadataResponse"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgUpdateGroupMetadataResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupMetadataResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgUpdateGroupMetadataResponse"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgUpdateGroupMetadataResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgUpdateGroupMetadataResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgUpdateGroupMetadataResponse"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgUpdateGroupMetadataResponse does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupMetadataResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgUpdateGroupMetadataResponse"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgUpdateGroupMetadataResponse does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupMetadataResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgUpdateGroupMetadataResponse"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgUpdateGroupMetadataResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgUpdateGroupMetadataResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgUpdateGroupMetadataResponse"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgUpdateGroupMetadataResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgUpdateGroupMetadataResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.group.v1beta1.MsgUpdateGroupMetadataResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgUpdateGroupMetadataResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupMetadataResponse) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_MsgUpdateGroupMetadataResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgUpdateGroupMetadataResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgUpdateGroupMetadataResponse)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*MsgUpdateGroupMetadataResponse)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*MsgUpdateGroupMetadataResponse)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgUpdateGroupMetadataResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgUpdateGroupMetadataResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_MsgCreateGroupPolicy                 protoreflect.MessageDescriptor
	fd_MsgCreateGroupPolicy_admin           protoreflect.FieldDescriptor
	fd_MsgCreateGroupPolicy_group_id        protoreflect.FieldDescriptor
	fd_MsgCreateGroupPolicy_metadata        protoreflect.FieldDescriptor
	fd_MsgCreateGroupPolicy_decision_policy protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_group_v1beta1_tx_proto_init()
	md_MsgCreateGroupPolicy = File_cosmos_group_v1beta1_tx_proto.Messages().ByName("MsgCreateGroupPolicy")
	fd_MsgCreateGroupPolicy_admin = md_MsgCreateGroupPolicy.Fields().ByName("admin")
	fd_MsgCreateGroupPolicy_group_id = md_MsgCreateGroupPolicy.Fields().ByName("group_id")
	fd_MsgCreateGroupPolicy_metadata = md_MsgCreateGroupPolicy.Fields().ByName("metadata")
	fd_MsgCreateGroupPolicy_decision_policy = md_MsgCreateGroupPolicy.Fields().ByName("decision_policy")
}

var _ protoreflect.Message = (*fastReflection_MsgCreateGroupPolicy)(nil)

type fastReflection_MsgCreateGroupPolicy MsgCreateGroupPolicy

func (x *MsgCreateGroupPolicy) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgCreateGroupPolicy)(x)
}

func (x *MsgCreateGroupPolicy) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_group_v1beta1_tx_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgCreateGroupPolicy_messageType fastReflection_MsgCreateGroupPolicy_messageType
var _ protoreflect.MessageType = fastReflection_MsgCreateGroupPolicy_messageType{}

type fastReflection_MsgCreateGroupPolicy_messageType struct{}

func (x fastReflection_MsgCreateGroupPolicy_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgCreateGroupPolicy)(nil)
}
func (x fastReflection_MsgCreateGroupPolicy_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgCreateGroupPolicy)
}
func (x fastReflection_MsgCreateGroupPolicy_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgCreateGroupPolicy
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgCreateGroupPolicy) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgCreateGroupPolicy
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgCreateGroupPolicy) Type() protoreflect.MessageType {
	return _fastReflection_MsgCreateGroupPolicy_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgCreateGroupPolicy) New() protoreflect.Message {
	return new(fastReflection_MsgCreateGroupPolicy)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgCreateGroupPolicy) Interface() protoreflect.ProtoMessage {
	return (*MsgCreateGroupPolicy)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgCreateGroupPolicy) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Admin != "" {
		value := protoreflect.ValueOfString(x.Admin)
		if !f(fd_MsgCreateGroupPolicy_admin, value) {
			return
		}
	}
	if x.GroupId != uint64(0) {
		value := protoreflect.ValueOfUint64(x.GroupId)
		if !f(fd_MsgCreateGroupPolicy_group_id, value) {
			return
		}
	}
	if len(x.Metadata) != 0 {
		value := protoreflect.ValueOfBytes(x.Metadata)
		if !f(fd_MsgCreateGroupPolicy_metadata, value) {
			return
		}
	}
	if x.DecisionPolicy != nil {
		value := protoreflect.ValueOfMessage(x.DecisionPolicy.ProtoReflect())
		if !f(fd_MsgCreateGroupPolicy_decision_policy, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_MsgCreateGroupPolicy) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.group.v1beta1.MsgCreateGroupPolicy.admin":
		return x.Admin != ""
	case "cosmos.group.v1beta1.MsgCreateGroupPolicy.group_id":
		return x.GroupId != uint64(0)
	case "cosmos.group.v1beta1.MsgCreateGroupPolicy.metadata":
		return len(x.Metadata) != 0
	case "cosmos.group.v1beta1.MsgCreateGroupPolicy.decision_policy":
		return x.DecisionPolicy != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgCreateGroupPolicy"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgCreateGroupPolicy does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateGroupPolicy) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.group.v1beta1.MsgCreateGroupPolicy.admin":
		x.Admin = ""
	case "cosmos.group.v1beta1.MsgCreateGroupPolicy.group_id":
		x.GroupId = uint64(0)
	case "cosmos.group.v1beta1.MsgCreateGroupPolicy.metadata":
		x.Metadata = nil
	case "cosmos.group.v1beta1.MsgCreateGroupPolicy.decision_policy":
		x.DecisionPolicy = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgCreateGroupPolicy"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgCreateGroupPolicy does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgCreateGroupPolicy) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.group.v1beta1.MsgCreateGroupPolicy.admin":
		value := x.Admin
		return protoreflect.ValueOfString(value)
	case "cosmos.group.v1beta1.MsgCreateGroupPolicy.group_id":
		value := x.GroupId
		return protoreflect.ValueOfUint64(value)
	case "cosmos.group.v1beta1.MsgCreateGroupPolicy.metadata":
		value := x.Metadata
		return protoreflect.ValueOfBytes(value)
	case "cosmos.group.v1beta1.MsgCreateGroupPolicy.decision_policy":
		value := x.DecisionPolicy
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgCreateGroupPolicy"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgCreateGroupPolicy does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateGroupPolicy) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.group.v1beta1.MsgCreateGroupPolicy.admin":
		x.Admin = value.Interface().(string)
	case "cosmos.group.v1beta1.MsgCreateGroupPolicy.group_id":
		x.GroupId = value.Uint()
	case "cosmos.group.v1beta1.MsgCreateGroupPolicy.metadata":
		x.Metadata = value.Bytes()
	case "cosmos.group.v1beta1.MsgCreateGroupPolicy.decision_policy":
		x.DecisionPolicy = value.Message().Interface().(*anypb.Any)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgCreateGroupPolicy"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgCreateGroupPolicy does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateGroupPolicy) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.group.v1beta1.MsgCreateGroupPolicy.decision_policy":
		if x.DecisionPolicy == nil {
			x.DecisionPolicy = new(anypb.Any)
		}
		return protoreflect.ValueOfMessage(x.DecisionPolicy.ProtoReflect())
	case "cosmos.group.v1beta1.MsgCreateGroupPolicy.admin":
		panic(fmt.Errorf("field admin of message cosmos.group.v1beta1.MsgCreateGroupPolicy is not mutable"))
	case "cosmos.group.v1beta1.MsgCreateGroupPolicy.group_id":
		panic(fmt.Errorf("field group_id of message cosmos.group.v1beta1.MsgCreateGroupPolicy is not mutable"))
	case "cosmos.group.v1beta1.MsgCreateGroupPolicy.metadata":
		panic(fmt.Errorf("field metadata of message cosmos.group.v1beta1.MsgCreateGroupPolicy is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgCreateGroupPolicy"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgCreateGroupPolicy does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgCreateGroupPolicy) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.group.v1beta1.MsgCreateGroupPolicy.admin":
		return protoreflect.ValueOfString("")
	case "cosmos.group.v1beta1.MsgCreateGroupPolicy.group_id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "cosmos.group.v1beta1.MsgCreateGroupPolicy.metadata":
		return protoreflect.ValueOfBytes(nil)
	case "cosmos.group.v1beta1.MsgCreateGroupPolicy.decision_policy":
		m := new(anypb.Any)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgCreateGroupPolicy"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgCreateGroupPolicy does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgCreateGroupPolicy) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.group.v1beta1.MsgCreateGroupPolicy", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgCreateGroupPolicy) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateGroupPolicy) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_MsgCreateGroupPolicy) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgCreateGroupPolicy) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgCreateGroupPolicy)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.Admin)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.GroupId != 0 {
			n += 1 + runtime.Sov(uint64(x.GroupId))
		}
		l = len(x.Metadata)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.DecisionPolicy != nil {
			l = options.Size(x.DecisionPolicy)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*MsgCreateGroupPolicy)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.DecisionPolicy != nil {
			encoded, err := options.Marshal(x.DecisionPolicy)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x22
		}
		if len(x.Metadata) > 0 {
			i -= len(x.Metadata)
			copy(dAtA[i:], x.Metadata)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Metadata)))
			i--
			dAtA[i] = 0x1a
		}
		if x.GroupId != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.GroupId))
			i--
			dAtA[i] = 0x10
		}
		if len(x.Admin) > 0 {
			i -= len(x.Admin)
			copy(dAtA[i:], x.Admin)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Admin)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*MsgCreateGroupPolicy)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgCreateGroupPolicy: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgCreateGroupPolicy: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Admin", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Admin = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field GroupId", wireType)
				}
				x.GroupId = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.GroupId |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Metadata = append(x.Metadata[:0], dAtA[iNdEx:postIndex]...)
				if x.Metadata == nil {
					x.Metadata = []byte{}
				}
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field DecisionPolicy", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.DecisionPolicy == nil {
					x.DecisionPolicy = &anypb.Any{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.DecisionPolicy); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_MsgCreateGroupPolicyResponse         protoreflect.MessageDescriptor
	fd_MsgCreateGroupPolicyResponse_address protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_group_v1beta1_tx_proto_init()
	md_MsgCreateGroupPolicyResponse = File_cosmos_group_v1beta1_tx_proto.Messages().ByName("MsgCreateGroupPolicyResponse")
	fd_MsgCreateGroupPolicyResponse_address = md_MsgCreateGroupPolicyResponse.Fields().ByName("address")
}

var _ protoreflect.Message = (*fastReflection_MsgCreateGroupPolicyResponse)(nil)

type fastReflection_MsgCreateGroupPolicyResponse MsgCreateGroupPolicyResponse

func (x *MsgCreateGroupPolicyResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgCreateGroupPolicyResponse)(x)
}

func (x *MsgCreateGroupPolicyResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_group_v1beta1_tx_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgCreateGroupPolicyResponse_messageType fastReflection_MsgCreateGroupPolicyResponse_messageType
var _ protoreflect.MessageType = fastReflection_MsgCreateGroupPolicyResponse_messageType{}

type fastReflection_MsgCreateGroupPolicyResponse_messageType struct{}

func (x fastReflection_MsgCreateGroupPolicyResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgCreateGroupPolicyResponse)(nil)
}
func (x fastReflection_MsgCreateGroupPolicyResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgCreateGroupPolicyResponse)
}
func (x fastReflection_MsgCreateGroupPolicyResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgCreateGroupPolicyResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgCreateGroupPolicyResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgCreateGroupPolicyResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgCreateGroupPolicyResponse) Type() protoreflect.MessageType {
	return _fastReflection_MsgCreateGroupPolicyResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgCreateGroupPolicyResponse) New() protoreflect.Message {
	return new(fastReflection_MsgCreateGroupPolicyResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgCreateGroupPolicyResponse) Interface() protoreflect.ProtoMessage {
	return (*MsgCreateGroupPolicyResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgCreateGroupPolicyResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Address != "" {
		value := protoreflect.ValueOfString(x.Address)
		if !f(fd_MsgCreateGroupPolicyResponse_address, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_MsgCreateGroupPolicyResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.group.v1beta1.MsgCreateGroupPolicyResponse.address":
		return x.Address != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgCreateGroupPolicyResponse"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgCreateGroupPolicyResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateGroupPolicyResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.group.v1beta1.MsgCreateGroupPolicyResponse.address":
		x.Address = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgCreateGroupPolicyResponse"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgCreateGroupPolicyResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgCreateGroupPolicyResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.group.v1beta1.MsgCreateGroupPolicyResponse.address":
		value := x.Address
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgCreateGroupPolicyResponse"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgCreateGroupPolicyResponse does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateGroupPolicyResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.group.v1beta1.MsgCreateGroupPolicyResponse.address":
		x.Address = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgCreateGroupPolicyResponse"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgCreateGroupPolicyResponse does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateGroupPolicyResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.group.v1beta1.MsgCreateGroupPolicyResponse.address":
		panic(fmt.Errorf("field address of message cosmos.group.v1beta1.MsgCreateGroupPolicyResponse is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgCreateGroupPolicyResponse"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgCreateGroupPolicyResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgCreateGroupPolicyResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.group.v1beta1.MsgCreateGroupPolicyResponse.address":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgCreateGroupPolicyResponse"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgCreateGroupPolicyResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgCreateGroupPolicyResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.group.v1beta1.MsgCreateGroupPolicyResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgCreateGroupPolicyResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateGroupPolicyResponse) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_MsgCreateGroupPolicyResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgCreateGroupPolicyResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgCreateGroupPolicyResponse)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.Address)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*MsgCreateGroupPolicyResponse)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.Address) > 0 {
			i -= len(x.Address)
			copy(dAtA[i:], x.Address)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Address)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*MsgCreateGroupPolicyResponse)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgCreateGroupPolicyResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgCreateGroupPolicyResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Address = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_MsgUpdateGroupPolicyAdmin           protoreflect.MessageDescriptor
	fd_MsgUpdateGroupPolicyAdmin_admin     protoreflect.FieldDescriptor
	fd_MsgUpdateGroupPolicyAdmin_address   protoreflect.FieldDescriptor
	fd_MsgUpdateGroupPolicyAdmin_new_admin protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_group_v1beta1_tx_proto_init()
	md_MsgUpdateGroupPolicyAdmin = File_cosmos_group_v1beta1_tx_proto.Messages().ByName("MsgUpdateGroupPolicyAdmin")
	fd_MsgUpdateGroupPolicyAdmin_admin = md_MsgUpdateGroupPolicyAdmin.Fields().ByName("admin")
	fd_MsgUpdateGroupPolicyAdmin_address = md_MsgUpdateGroupPolicyAdmin.Fields().ByName("address")
	fd_MsgUpdateGroupPolicyAdmin_new_admin = md_MsgUpdateGroupPolicyAdmin.Fields().ByName("new_admin")
}

var _ protoreflect.Message = (*fastReflection_MsgUpdateGroupPolicyAdmin)(nil)

type fastReflection_MsgUpdateGroupPolicyAdmin MsgUpdateGroupPolicyAdmin

func (x *MsgUpdateGroupPolicyAdmin) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgUpdateGroupPolicyAdmin)(x)
}

func (x *MsgUpdateGroupPolicyAdmin) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_group_v1beta1_tx_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgUpdateGroupPolicyAdmin_messageType fastReflection_MsgUpdateGroupPolicyAdmin_messageType
var _ protoreflect.MessageType = fastReflection_MsgUpdateGroupPolicyAdmin_messageType{}

type fastReflection_MsgUpdateGroupPolicyAdmin_messageType struct{}

func (x fastReflection_MsgUpdateGroupPolicyAdmin_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgUpdateGroupPolicyAdmin)(nil)
}
func (x fastReflection_MsgUpdateGroupPolicyAdmin_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgUpdateGroupPolicyAdmin)
}
func (x fastReflection_MsgUpdateGroupPolicyAdmin_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgUpdateGroupPolicyAdmin
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgUpdateGroupPolicyAdmin) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgUpdateGroupPolicyAdmin
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgUpdateGroupPolicyAdmin) Type() protoreflect.MessageType {
	return _fastReflection_MsgUpdateGroupPolicyAdmin_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgUpdateGroupPolicyAdmin) New() protoreflect.Message {
	return new(fastReflection_MsgUpdateGroupPolicyAdmin)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgUpdateGroupPolicyAdmin) Interface() protoreflect.ProtoMessage {
	return (*MsgUpdateGroupPolicyAdmin)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgUpdateGroupPolicyAdmin) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Admin != "" {
		value := protoreflect.ValueOfString(x.Admin)
		if !f(fd_MsgUpdateGroupPolicyAdmin_admin, value) {
			return
		}
	}
	if x.Address != "" {
		value := protoreflect.ValueOfString(x.Address)
		if !f(fd_MsgUpdateGroupPolicyAdmin_address, value) {
			return
		}
	}
	if x.NewAdmin != "" {
		value := protoreflect.ValueOfString(x.NewAdmin)
		if !f(fd_MsgUpdateGroupPolicyAdmin_new_admin, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_MsgUpdateGroupPolicyAdmin) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.group.v1beta1.MsgUpdateGroupPolicyAdmin.admin":
		return x.Admin != ""
	case "cosmos.group.v1beta1.MsgUpdateGroupPolicyAdmin.address":
		return x.Address != ""
	case "cosmos.group.v1beta1.MsgUpdateGroupPolicyAdmin.new_admin":
		return x.NewAdmin != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgUpdateGroupPolicyAdmin"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgUpdateGroupPolicyAdmin does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupPolicyAdmin) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.group.v1beta1.MsgUpdateGroupPolicyAdmin.admin":
		x.Admin = ""
	case "cosmos.group.v1beta1.MsgUpdateGroupPolicyAdmin.address":
		x.Address = ""
	case "cosmos.group.v1beta1.MsgUpdateGroupPolicyAdmin.new_admin":
		x.NewAdmin = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgUpdateGroupPolicyAdmin"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgUpdateGroupPolicyAdmin does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgUpdateGroupPolicyAdmin) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.group.v1beta1.MsgUpdateGroupPolicyAdmin.admin":
		value := x.Admin
		return protoreflect.ValueOfString(value)
	case "cosmos.group.v1beta1.MsgUpdateGroupPolicyAdmin.address":
		value := x.Address
		return protoreflect.ValueOfString(value)
	case "cosmos.group.v1beta1.MsgUpdateGroupPolicyAdmin.new_admin":
		value := x.NewAdmin
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgUpdateGroupPolicyAdmin"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgUpdateGroupPolicyAdmin does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupPolicyAdmin) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.group.v1beta1.MsgUpdateGroupPolicyAdmin.admin":
		x.Admin = value.Interface().(string)
	case "cosmos.group.v1beta1.MsgUpdateGroupPolicyAdmin.address":
		x.Address = value.Interface().(string)
	case "cosmos.group.v1beta1.MsgUpdateGroupPolicyAdmin.new_admin":
		x.NewAdmin = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgUpdateGroupPolicyAdmin"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgUpdateGroupPolicyAdmin does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupPolicyAdmin) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.group.v1beta1.MsgUpdateGroupPolicyAdmin.admin":
		panic(fmt.Errorf("field admin of message cosmos.group.v1beta1.MsgUpdateGroupPolicyAdmin is not mutable"))
	case "cosmos.group.v1beta1.MsgUpdateGroupPolicyAdmin.address":
		panic(fmt.Errorf("field address of message cosmos.group.v1beta1.MsgUpdateGroupPolicyAdmin is not mutable"))
	case "cosmos.group.v1beta1.MsgUpdateGroupPolicyAdmin.new_admin":
		panic(fmt.Errorf("field new_admin of message cosmos.group.v1beta1.MsgUpdateGroupPolicyAdmin is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgUpdateGroupPolicyAdmin"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgUpdateGroupPolicyAdmin does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgUpdateGroupPolicyAdmin) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.group.v1beta1.MsgUpdateGroupPolicyAdmin.admin":
		return protoreflect.ValueOfString("")
	case "cosmos.group.v1beta1.MsgUpdateGroupPolicyAdmin.address":
		return protoreflect.ValueOfString("")
	case "cosmos.group.v1beta1.MsgUpdateGroupPolicyAdmin.new_admin":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgUpdateGroupPolicyAdmin"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgUpdateGroupPolicyAdmin does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgUpdateGroupPolicyAdmin) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.group.v1beta1.MsgUpdateGroupPolicyAdmin", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgUpdateGroupPolicyAdmin) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupPolicyAdmin) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_MsgUpdateGroupPolicyAdmin) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgUpdateGroupPolicyAdmin) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgUpdateGroupPolicyAdmin)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.Admin)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Address)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.NewAdmin)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*MsgUpdateGroupPolicyAdmin)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.NewAdmin) > 0 {
			i -= len(x.NewAdmin)
			copy(dAtA[i:], x.NewAdmin)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.NewAdmin)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.Address) > 0 {
			i -= len(x.Address)
			copy(dAtA[i:], x.Address)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Address)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Admin) > 0 {
			i -= len(x.Admin)
			copy(dAtA[i:], x.Admin)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Admin)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*MsgUpdateGroupPolicyAdmin)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgUpdateGroupPolicyAdmin: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgUpdateGroupPolicyAdmin: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Admin", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Admin = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Address = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field NewAdmin", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.NewAdmin = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_MsgUpdateGroupPolicyAdminResponse protoreflect.MessageDescriptor
)

func init() {
	file_cosmos_group_v1beta1_tx_proto_init()
	md_MsgUpdateGroupPolicyAdminResponse = File_cosmos_group_v1beta1_tx_proto.Messages().ByName("MsgUpdateGroupPolicyAdminResponse")
}

var _ protoreflect.Message = (*fastReflection_MsgUpdateGroupPolicyAdminResponse)(nil)

type fastReflection_MsgUpdateGroupPolicyAdminResponse MsgUpdateGroupPolicyAdminResponse

func (x *MsgUpdateGroupPolicyAdminResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgUpdateGroupPolicyAdminResponse)(x)
}

func (x *MsgUpdateGroupPolicyAdminResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_group_v1beta1_tx_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgUpdateGroupPolicyAdminResponse_messageType fastReflection_MsgUpdateGroupPolicyAdminResponse_messageType
var _ protoreflect.MessageType = fastReflection_MsgUpdateGroupPolicyAdminResponse_messageType{}

type fastReflection_MsgUpdateGroupPolicyAdminResponse_messageType struct{}

func (x fastReflection_MsgUpdateGroupPolicyAdminResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgUpdateGroupPolicyAdminResponse)(nil)
}
func (x fastReflection_MsgUpdateGroupPolicyAdminResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgUpdateGroupPolicyAdminResponse)
}
func (x fastReflection_MsgUpdateGroupPolicyAdminResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgUpdateGroupPolicyAdminResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgUpdateGroupPolicyAdminResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgUpdateGroupPolicyAdminResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgUpdateGroupPolicyAdminResponse) Type() protoreflect.MessageType {
	return _fastReflection_MsgUpdateGroupPolicyAdminResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgUpdateGroupPolicyAdminResponse) New() protoreflect.Message {
	return new(fastReflection_MsgUpdateGroupPolicyAdminResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgUpdateGroupPolicyAdminResponse) Interface() protoreflect.ProtoMessage {
	return (*MsgUpdateGroupPolicyAdminResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgUpdateGroupPolicyAdminResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_MsgUpdateGroupPolicyAdminResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgUpdateGroupPolicyAdminResponse"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgUpdateGroupPolicyAdminResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupPolicyAdminResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgUpdateGroupPolicyAdminResponse"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgUpdateGroupPolicyAdminResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgUpdateGroupPolicyAdminResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgUpdateGroupPolicyAdminResponse"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgUpdateGroupPolicyAdminResponse does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupPolicyAdminResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgUpdateGroupPolicyAdminResponse"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgUpdateGroupPolicyAdminResponse does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupPolicyAdminResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgUpdateGroupPolicyAdminResponse"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgUpdateGroupPolicyAdminResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgUpdateGroupPolicyAdminResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgUpdateGroupPolicyAdminResponse"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgUpdateGroupPolicyAdminResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgUpdateGroupPolicyAdminResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.group.v1beta1.MsgUpdateGroupPolicyAdminResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgUpdateGroupPolicyAdminResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupPolicyAdminResponse) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_MsgUpdateGroupPolicyAdminResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgUpdateGroupPolicyAdminResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgUpdateGroupPolicyAdminResponse)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*MsgUpdateGroupPolicyAdminResponse)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*MsgUpdateGroupPolicyAdminResponse)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgUpdateGroupPolicyAdminResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgUpdateGroupPolicyAdminResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_MsgUpdateGroupPolicyDecisionPolicy                 protoreflect.MessageDescriptor
	fd_MsgUpdateGroupPolicyDecisionPolicy_admin           protoreflect.FieldDescriptor
	fd_MsgUpdateGroupPolicyDecisionPolicy_address         protoreflect.FieldDescriptor
	fd_MsgUpdateGroupPolicyDecisionPolicy_decision_policy protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_group_v1beta1_tx_proto_init()
	md_MsgUpdateGroupPolicyDecisionPolicy = File_cosmos_group_v1beta1_tx_proto.Messages().ByName("MsgUpdateGroupPolicyDecisionPolicy")
	fd_MsgUpdateGroupPolicyDecisionPolicy_admin = md_MsgUpdateGroupPolicyDecisionPolicy.Fields().ByName("admin")
	fd_MsgUpdateGroupPolicyDecisionPolicy_address = md_MsgUpdateGroupPolicyDecisionPolicy.Fields().ByName("address")
	fd_MsgUpdateGroupPolicyDecisionPolicy_decision_policy = md_MsgUpdateGroupPolicyDecisionPolicy.Fields().ByName("decision_policy")
}

var _ protoreflect.Message = (*fastReflection_MsgUpdateGroupPolicyDecisionPolicy)(nil)

type fastReflection_MsgUpdateGroupPolicyDecisionPolicy MsgUpdateGroupPolicyDecisionPolicy

func (x *MsgUpdateGroupPolicyDecisionPolicy) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgUpdateGroupPolicyDecisionPolicy)(x)
}

func (x *MsgUpdateGroupPolicyDecisionPolicy) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_group_v1beta1_tx_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgUpdateGroupPolicyDecisionPolicy_messageType fastReflection_MsgUpdateGroupPolicyDecisionPolicy_messageType
var _ protoreflect.MessageType = fastReflection_MsgUpdateGroupPolicyDecisionPolicy_messageType{}

type fastReflection_MsgUpdateGroupPolicyDecisionPolicy_messageType struct{}

func (x fastReflection_MsgUpdateGroupPolicyDecisionPolicy_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgUpdateGroupPolicyDecisionPolicy)(nil)
}
func (x fastReflection_MsgUpdateGroupPolicyDecisionPolicy_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgUpdateGroupPolicyDecisionPolicy)
}
func (x fastReflection_MsgUpdateGroupPolicyDecisionPolicy_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgUpdateGroupPolicyDecisionPolicy
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgUpdateGroupPolicyDecisionPolicy) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgUpdateGroupPolicyDecisionPolicy
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgUpdateGroupPolicyDecisionPolicy) Type() protoreflect.MessageType {
	return _fastReflection_MsgUpdateGroupPolicyDecisionPolicy_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgUpdateGroupPolicyDecisionPolicy) New() protoreflect.Message {
	return new(fastReflection_MsgUpdateGroupPolicyDecisionPolicy)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgUpdateGroupPolicyDecisionPolicy) Interface() protoreflect.ProtoMessage {
	return (*MsgUpdateGroupPolicyDecisionPolicy)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgUpdateGroupPolicyDecisionPolicy) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Admin != "" {
		value := protoreflect.ValueOfString(x.Admin)
		if !f(fd_MsgUpdateGroupPolicyDecisionPolicy_admin, value) {
			return
		}
	}
	if x.Address != "" {
		value := protoreflect.ValueOfString(x.Address)
		if !f(fd_MsgUpdateGroupPolicyDecisionPolicy_address, value) {
			return
		}
	}
	if x.DecisionPolicy != nil {
		value := protoreflect.ValueOfMessage(x.DecisionPolicy.ProtoReflect())
		if !f(fd_MsgUpdateGroupPolicyDecisionPolicy_decision_policy, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_MsgUpdateGroupPolicyDecisionPolicy) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.group.v1beta1.MsgUpdateGroupPolicyDecisionPolicy.admin":
		return x.Admin != ""
	case "cosmos.group.v1beta1.MsgUpdateGroupPolicyDecisionPolicy.address":
		return x.Address != ""
	case "cosmos.group.v1beta1.MsgUpdateGroupPolicyDecisionPolicy.decision_policy":
		return x.DecisionPolicy != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgUpdateGroupPolicyDecisionPolicy"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgUpdateGroupPolicyDecisionPolicy does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupPolicyDecisionPolicy) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.group.v1beta1.MsgUpdateGroupPolicyDecisionPolicy.admin":
		x.Admin = ""
	case "cosmos.group.v1beta1.MsgUpdateGroupPolicyDecisionPolicy.address":
		x.Address = ""
	case "cosmos.group.v1beta1.MsgUpdateGroupPolicyDecisionPolicy.decision_policy":
		x.DecisionPolicy = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgUpdateGroupPolicyDecisionPolicy"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgUpdateGroupPolicyDecisionPolicy does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgUpdateGroupPolicyDecisionPolicy) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.group.v1beta1.MsgUpdateGroupPolicyDecisionPolicy.admin":
		value := x.Admin
		return protoreflect.ValueOfString(value)
	case "cosmos.group.v1beta1.MsgUpdateGroupPolicyDecisionPolicy.address":
		value := x.Address
		return protoreflect.ValueOfString(value)
	case "cosmos.group.v1beta1.MsgUpdateGroupPolicyDecisionPolicy.decision_policy":
		value := x.DecisionPolicy
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgUpdateGroupPolicyDecisionPolicy"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgUpdateGroupPolicyDecisionPolicy does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupPolicyDecisionPolicy) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.group.v1beta1.MsgUpdateGroupPolicyDecisionPolicy.admin":
		x.Admin = value.Interface().(string)
	case "cosmos.group.v1beta1.MsgUpdateGroupPolicyDecisionPolicy.address":
		x.Address = value.Interface().(string)
	case "cosmos.group.v1beta1.MsgUpdateGroupPolicyDecisionPolicy.decision_policy":
		x.DecisionPolicy = value.Message().Interface().(*anypb.Any)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgUpdateGroupPolicyDecisionPolicy"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgUpdateGroupPolicyDecisionPolicy does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupPolicyDecisionPolicy) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.group.v1beta1.MsgUpdateGroupPolicyDecisionPolicy.decision_policy":
		if x.DecisionPolicy == nil {
			x.DecisionPolicy = new(anypb.Any)
		}
		return protoreflect.ValueOfMessage(x.DecisionPolicy.ProtoReflect())
	case "cosmos.group.v1beta1.MsgUpdateGroupPolicyDecisionPolicy.admin":
		panic(fmt.Errorf("field admin of message cosmos.group.v1beta1.MsgUpdateGroupPolicyDecisionPolicy is not mutable"))
	case "cosmos.group.v1beta1.MsgUpdateGroupPolicyDecisionPolicy.address":
		panic(fmt.Errorf("field address of message cosmos.group.v1beta1.MsgUpdateGroupPolicyDecisionPolicy is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgUpdateGroupPolicyDecisionPolicy"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgUpdateGroupPolicyDecisionPolicy does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgUpdateGroupPolicyDecisionPolicy) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.group.v1beta1.MsgUpdateGroupPolicyDecisionPolicy.admin":
		return protoreflect.ValueOfString("")
	case "cosmos.group.v1beta1.MsgUpdateGroupPolicyDecisionPolicy.address":
		return protoreflect.ValueOfString("")
	case "cosmos.group.v1beta1.MsgUpdateGroupPolicyDecisionPolicy.decision_policy":
		m := new(anypb.Any)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgUpdateGroupPolicyDecisionPolicy"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgUpdateGroupPolicyDecisionPolicy does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgUpdateGroupPolicyDecisionPolicy) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.group.v1beta1.MsgUpdateGroupPolicyDecisionPolicy", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgUpdateGroupPolicyDecisionPolicy) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupPolicyDecisionPolicy) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_MsgUpdateGroupPolicyDecisionPolicy) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgUpdateGroupPolicyDecisionPolicy) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgUpdateGroupPolicyDecisionPolicy)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.Admin)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Address)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.DecisionPolicy != nil {
			l = options.Size(x.DecisionPolicy)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*MsgUpdateGroupPolicyDecisionPolicy)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.DecisionPolicy != nil {
			encoded, err := options.Marshal(x.DecisionPolicy)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.Address) > 0 {
			i -= len(x.Address)
			copy(dAtA[i:], x.Address)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Address)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Admin) > 0 {
			i -= len(x.Admin)
			copy(dAtA[i:], x.Admin)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Admin)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*MsgUpdateGroupPolicyDecisionPolicy)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgUpdateGroupPolicyDecisionPolicy: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgUpdateGroupPolicyDecisionPolicy: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Admin", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Admin = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Address = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field DecisionPolicy", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.DecisionPolicy == nil {
					x.DecisionPolicy = &anypb.Any{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.DecisionPolicy); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_MsgUpdateGroupPolicyDecisionPolicyResponse protoreflect.MessageDescriptor
)

func init() {
	file_cosmos_group_v1beta1_tx_proto_init()
	md_MsgUpdateGroupPolicyDecisionPolicyResponse = File_cosmos_group_v1beta1_tx_proto.Messages().ByName("MsgUpdateGroupPolicyDecisionPolicyResponse")
}

var _ protoreflect.Message = (*fastReflection_MsgUpdateGroupPolicyDecisionPolicyResponse)(nil)

type fastReflection_MsgUpdateGroupPolicyDecisionPolicyResponse MsgUpdateGroupPolicyDecisionPolicyResponse

func (x *MsgUpdateGroupPolicyDecisionPolicyResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgUpdateGroupPolicyDecisionPolicyResponse)(x)
}

func (x *MsgUpdateGroupPolicyDecisionPolicyResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_group_v1beta1_tx_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgUpdateGroupPolicyDecisionPolicyResponse_messageType fastReflection_MsgUpdateGroupPolicyDecisionPolicyResponse_messageType
var _ protoreflect.MessageType = fastReflection_MsgUpdateGroupPolicyDecisionPolicyResponse_messageType{}

type fastReflection_MsgUpdateGroupPolicyDecisionPolicyResponse_messageType struct{}

func (x fastReflection_MsgUpdateGroupPolicyDecisionPolicyResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgUpdateGroupPolicyDecisionPolicyResponse)(nil)
}
func (x fastReflection_MsgUpdateGroupPolicyDecisionPolicyResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgUpdateGroupPolicyDecisionPolicyResponse)
}
func (x fastReflection_MsgUpdateGroupPolicyDecisionPolicyResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgUpdateGroupPolicyDecisionPolicyResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgUpdateGroupPolicyDecisionPolicyResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgUpdateGroupPolicyDecisionPolicyResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgUpdateGroupPolicyDecisionPolicyResponse) Type() protoreflect.MessageType {
	return _fastReflection_MsgUpdateGroupPolicyDecisionPolicyResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgUpdateGroupPolicyDecisionPolicyResponse) New() protoreflect.Message {
	return new(fastReflection_MsgUpdateGroupPolicyDecisionPolicyResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgUpdateGroupPolicyDecisionPolicyResponse) Interface() protoreflect.ProtoMessage {
	return (*MsgUpdateGroupPolicyDecisionPolicyResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgUpdateGroupPolicyDecisionPolicyResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_MsgUpdateGroupPolicyDecisionPolicyResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgUpdateGroupPolicyDecisionPolicyResponse"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgUpdateGroupPolicyDecisionPolicyResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupPolicyDecisionPolicyResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgUpdateGroupPolicyDecisionPolicyResponse"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgUpdateGroupPolicyDecisionPolicyResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgUpdateGroupPolicyDecisionPolicyResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgUpdateGroupPolicyDecisionPolicyResponse"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgUpdateGroupPolicyDecisionPolicyResponse does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupPolicyDecisionPolicyResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgUpdateGroupPolicyDecisionPolicyResponse"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgUpdateGroupPolicyDecisionPolicyResponse does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupPolicyDecisionPolicyResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgUpdateGroupPolicyDecisionPolicyResponse"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgUpdateGroupPolicyDecisionPolicyResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgUpdateGroupPolicyDecisionPolicyResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgUpdateGroupPolicyDecisionPolicyResponse"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgUpdateGroupPolicyDecisionPolicyResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgUpdateGroupPolicyDecisionPolicyResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.group.v1beta1.MsgUpdateGroupPolicyDecisionPolicyResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgUpdateGroupPolicyDecisionPolicyResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupPolicyDecisionPolicyResponse) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_MsgUpdateGroupPolicyDecisionPolicyResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgUpdateGroupPolicyDecisionPolicyResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgUpdateGroupPolicyDecisionPolicyResponse)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*MsgUpdateGroupPolicyDecisionPolicyResponse)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*MsgUpdateGroupPolicyDecisionPolicyResponse)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgUpdateGroupPolicyDecisionPolicyResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgUpdateGroupPolicyDecisionPolicyResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_MsgUpdateGroupPolicyMetadata          protoreflect.MessageDescriptor
	fd_MsgUpdateGroupPolicyMetadata_admin    protoreflect.FieldDescriptor
	fd_MsgUpdateGroupPolicyMetadata_address  protoreflect.FieldDescriptor
	fd_MsgUpdateGroupPolicyMetadata_metadata protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_group_v1beta1_tx_proto_init()
	md_MsgUpdateGroupPolicyMetadata = File_cosmos_group_v1beta1_tx_proto.Messages().ByName("MsgUpdateGroupPolicyMetadata")
	fd_MsgUpdateGroupPolicyMetadata_admin = md_MsgUpdateGroupPolicyMetadata.Fields().ByName("admin")
	fd_MsgUpdateGroupPolicyMetadata_address = md_MsgUpdateGroupPolicyMetadata.Fields().ByName("address")
	fd_MsgUpdateGroupPolicyMetadata_metadata = md_MsgUpdateGroupPolicyMetadata.Fields().ByName("metadata")
}

var _ protoreflect.Message = (*fastReflection_MsgUpdateGroupPolicyMetadata)(nil)

type fastReflection_MsgUpdateGroupPolicyMetadata MsgUpdateGroupPolicyMetadata

func (x *MsgUpdateGroupPolicyMetadata) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgUpdateGroupPolicyMetadata)(x)
}

func (x *MsgUpdateGroupPolicyMetadata) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_group_v1beta1_tx_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgUpdateGroupPolicyMetadata_messageType fastReflection_MsgUpdateGroupPolicyMetadata_messageType
var _ protoreflect.MessageType = fastReflection_MsgUpdateGroupPolicyMetadata_messageType{}

type fastReflection_MsgUpdateGroupPolicyMetadata_messageType struct{}

func (x fastReflection_MsgUpdateGroupPolicyMetadata_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgUpdateGroupPolicyMetadata)(nil)
}
func (x fastReflection_MsgUpdateGroupPolicyMetadata_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgUpdateGroupPolicyMetadata)
}
func (x fastReflection_MsgUpdateGroupPolicyMetadata_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgUpdateGroupPolicyMetadata
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgUpdateGroupPolicyMetadata) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgUpdateGroupPolicyMetadata
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgUpdateGroupPolicyMetadata) Type() protoreflect.MessageType {
	return _fastReflection_MsgUpdateGroupPolicyMetadata_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgUpdateGroupPolicyMetadata) New() protoreflect.Message {
	return new(fastReflection_MsgUpdateGroupPolicyMetadata)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgUpdateGroupPolicyMetadata) Interface() protoreflect.ProtoMessage {
	return (*MsgUpdateGroupPolicyMetadata)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgUpdateGroupPolicyMetadata) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Admin != "" {
		value := protoreflect.ValueOfString(x.Admin)
		if !f(fd_MsgUpdateGroupPolicyMetadata_admin, value) {
			return
		}
	}
	if x.Address != "" {
		value := protoreflect.ValueOfString(x.Address)
		if !f(fd_MsgUpdateGroupPolicyMetadata_address, value) {
			return
		}
	}
	if len(x.Metadata) != 0 {
		value := protoreflect.ValueOfBytes(x.Metadata)
		if !f(fd_MsgUpdateGroupPolicyMetadata_metadata, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_MsgUpdateGroupPolicyMetadata) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.group.v1beta1.MsgUpdateGroupPolicyMetadata.admin":
		return x.Admin != ""
	case "cosmos.group.v1beta1.MsgUpdateGroupPolicyMetadata.address":
		return x.Address != ""
	case "cosmos.group.v1beta1.MsgUpdateGroupPolicyMetadata.metadata":
		return len(x.Metadata) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgUpdateGroupPolicyMetadata"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgUpdateGroupPolicyMetadata does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupPolicyMetadata) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.group.v1beta1.MsgUpdateGroupPolicyMetadata.admin":
		x.Admin = ""
	case "cosmos.group.v1beta1.MsgUpdateGroupPolicyMetadata.address":
		x.Address = ""
	case "cosmos.group.v1beta1.MsgUpdateGroupPolicyMetadata.metadata":
		x.Metadata = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgUpdateGroupPolicyMetadata"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgUpdateGroupPolicyMetadata does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgUpdateGroupPolicyMetadata) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.group.v1beta1.MsgUpdateGroupPolicyMetadata.admin":
		value := x.Admin
		return protoreflect.ValueOfString(value)
	case "cosmos.group.v1beta1.MsgUpdateGroupPolicyMetadata.address":
		value := x.Address
		return protoreflect.ValueOfString(value)
	case "cosmos.group.v1beta1.MsgUpdateGroupPolicyMetadata.metadata":
		value := x.Metadata
		return protoreflect.ValueOfBytes(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgUpdateGroupPolicyMetadata"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgUpdateGroupPolicyMetadata does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupPolicyMetadata) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.group.v1beta1.MsgUpdateGroupPolicyMetadata.admin":
		x.Admin = value.Interface().(string)
	case "cosmos.group.v1beta1.MsgUpdateGroupPolicyMetadata.address":
		x.Address = value.Interface().(string)
	case "cosmos.group.v1beta1.MsgUpdateGroupPolicyMetadata.metadata":
		x.Metadata = value.Bytes()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgUpdateGroupPolicyMetadata"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgUpdateGroupPolicyMetadata does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupPolicyMetadata) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.group.v1beta1.MsgUpdateGroupPolicyMetadata.admin":
		panic(fmt.Errorf("field admin of message cosmos.group.v1beta1.MsgUpdateGroupPolicyMetadata is not mutable"))
	case "cosmos.group.v1beta1.MsgUpdateGroupPolicyMetadata.address":
		panic(fmt.Errorf("field address of message cosmos.group.v1beta1.MsgUpdateGroupPolicyMetadata is not mutable"))
	case "cosmos.group.v1beta1.MsgUpdateGroupPolicyMetadata.metadata":
		panic(fmt.Errorf("field metadata of message cosmos.group.v1beta1.MsgUpdateGroupPolicyMetadata is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgUpdateGroupPolicyMetadata"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgUpdateGroupPolicyMetadata does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgUpdateGroupPolicyMetadata) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.group.v1beta1.MsgUpdateGroupPolicyMetadata.admin":
		return protoreflect.ValueOfString("")
	case "cosmos.group.v1beta1.MsgUpdateGroupPolicyMetadata.address":
		return protoreflect.ValueOfString("")
	case "cosmos.group.v1beta1.MsgUpdateGroupPolicyMetadata.metadata":
		return protoreflect.ValueOfBytes(nil)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgUpdateGroupPolicyMetadata"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgUpdateGroupPolicyMetadata does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgUpdateGroupPolicyMetadata) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.group.v1beta1.MsgUpdateGroupPolicyMetadata", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgUpdateGroupPolicyMetadata) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupPolicyMetadata) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_MsgUpdateGroupPolicyMetadata) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgUpdateGroupPolicyMetadata) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgUpdateGroupPolicyMetadata)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.Admin)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Address)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Metadata)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*MsgUpdateGroupPolicyMetadata)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.Metadata) > 0 {
			i -= len(x.Metadata)
			copy(dAtA[i:], x.Metadata)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Metadata)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.Address) > 0 {
			i -= len(x.Address)
			copy(dAtA[i:], x.Address)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Address)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Admin) > 0 {
			i -= len(x.Admin)
			copy(dAtA[i:], x.Admin)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Admin)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*MsgUpdateGroupPolicyMetadata)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgUpdateGroupPolicyMetadata: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgUpdateGroupPolicyMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Admin", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Admin = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Address = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Metadata = append(x.Metadata[:0], dAtA[iNdEx:postIndex]...)
				if x.Metadata == nil {
					x.Metadata = []byte{}
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_MsgUpdateGroupPolicyMetadataResponse protoreflect.MessageDescriptor
)

func init() {
	file_cosmos_group_v1beta1_tx_proto_init()
	md_MsgUpdateGroupPolicyMetadataResponse = File_cosmos_group_v1beta1_tx_proto.Messages().ByName("MsgUpdateGroupPolicyMetadataResponse")
}

var _ protoreflect.Message = (*fastReflection_MsgUpdateGroupPolicyMetadataResponse)(nil)

type fastReflection_MsgUpdateGroupPolicyMetadataResponse MsgUpdateGroupPolicyMetadataResponse

func (x *MsgUpdateGroupPolicyMetadataResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgUpdateGroupPolicyMetadataResponse)(x)
}

func (x *MsgUpdateGroupPolicyMetadataResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_group_v1beta1_tx_proto_msgTypes[15]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgUpdateGroupPolicyMetadataResponse_messageType fastReflection_MsgUpdateGroupPolicyMetadataResponse_messageType
var _ protoreflect.MessageType = fastReflection_MsgUpdateGroupPolicyMetadataResponse_messageType{}

type fastReflection_MsgUpdateGroupPolicyMetadataResponse_messageType struct{}

func (x fastReflection_MsgUpdateGroupPolicyMetadataResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgUpdateGroupPolicyMetadataResponse)(nil)
}
func (x fastReflection_MsgUpdateGroupPolicyMetadataResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgUpdateGroupPolicyMetadataResponse)
}
func (x fastReflection_MsgUpdateGroupPolicyMetadataResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgUpdateGroupPolicyMetadataResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgUpdateGroupPolicyMetadataResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgUpdateGroupPolicyMetadataResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgUpdateGroupPolicyMetadataResponse) Type() protoreflect.MessageType {
	return _fastReflection_MsgUpdateGroupPolicyMetadataResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgUpdateGroupPolicyMetadataResponse) New() protoreflect.Message {
	return new(fastReflection_MsgUpdateGroupPolicyMetadataResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgUpdateGroupPolicyMetadataResponse) Interface() protoreflect.ProtoMessage {
	return (*MsgUpdateGroupPolicyMetadataResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgUpdateGroupPolicyMetadataResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_MsgUpdateGroupPolicyMetadataResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgUpdateGroupPolicyMetadataResponse"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgUpdateGroupPolicyMetadataResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupPolicyMetadataResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgUpdateGroupPolicyMetadataResponse"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgUpdateGroupPolicyMetadataResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgUpdateGroupPolicyMetadataResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgUpdateGroupPolicyMetadataResponse"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgUpdateGroupPolicyMetadataResponse does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupPolicyMetadataResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgUpdateGroupPolicyMetadataResponse"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgUpdateGroupPolicyMetadataResponse does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupPolicyMetadataResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgUpdateGroupPolicyMetadataResponse"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgUpdateGroupPolicyMetadataResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgUpdateGroupPolicyMetadataResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgUpdateGroupPolicyMetadataResponse"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgUpdateGroupPolicyMetadataResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgUpdateGroupPolicyMetadataResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.group.v1beta1.MsgUpdateGroupPolicyMetadataResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgUpdateGroupPolicyMetadataResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupPolicyMetadataResponse) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_MsgUpdateGroupPolicyMetadataResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgUpdateGroupPolicyMetadataResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgUpdateGroupPolicyMetadataResponse)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*MsgUpdateGroupPolicyMetadataResponse)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*MsgUpdateGroupPolicyMetadataResponse)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgUpdateGroupPolicyMetadataResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgUpdateGroupPolicyMetadataResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var _ protoreflect.List = (*_MsgCreateProposal_2_list)(nil)

type _MsgCreateProposal_2_list struct {
	list *[]string
}

func (x *_MsgCreateProposal_2_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_MsgCreateProposal_2_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfString((*x.list)[i])
}

func (x *_MsgCreateProposal_2_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.String()
	concreteValue := valueUnwrapped
	(*x.list)[i] = concreteValue
}

func (x *_MsgCreateProposal_2_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.String()
	concreteValue := valueUnwrapped
	*x.list = append(*x.list, concreteValue)
}

func (x *_MsgCreateProposal_2_list) AppendMutable() protoreflect.Value {
	panic(fmt.Errorf("AppendMutable can not be called on message MsgCreateProposal at list field Proposers as it is not of Message kind"))
}

func (x *_MsgCreateProposal_2_list) Truncate(n int) {
	*x.list = (*x.list)[:n]
}

func (x *_MsgCreateProposal_2_list) NewElement() protoreflect.Value {
	v := ""
	return protoreflect.ValueOfString(v)
}

func (x *_MsgCreateProposal_2_list) IsValid() bool {
	return x.list != nil
}

var _ protoreflect.List = (*_MsgCreateProposal_4_list)(nil)

type _MsgCreateProposal_4_list struct {
	list *[]*anypb.Any
}

func (x *_MsgCreateProposal_4_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_MsgCreateProposal_4_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_MsgCreateProposal_4_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*anypb.Any)
	(*x.list)[i] = concreteValue
}

func (x *_MsgCreateProposal_4_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*anypb.Any)
	*x.list = append(*x.list, concreteValue)
}

func (x *_MsgCreateProposal_4_list) AppendMutable() protoreflect.Value {
	v := new(anypb.Any)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_MsgCreateProposal_4_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_MsgCreateProposal_4_list) NewElement() protoreflect.Value {
	v := new(anypb.Any)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_MsgCreateProposal_4_list) IsValid() bool {
	return x.list != nil
}

var (
	md_MsgCreateProposal           protoreflect.MessageDescriptor
	fd_MsgCreateProposal_address   protoreflect.FieldDescriptor
	fd_MsgCreateProposal_proposers protoreflect.FieldDescriptor
	fd_MsgCreateProposal_metadata  protoreflect.FieldDescriptor
	fd_MsgCreateProposal_msgs      protoreflect.FieldDescriptor
	fd_MsgCreateProposal_exec      protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_group_v1beta1_tx_proto_init()
	md_MsgCreateProposal = File_cosmos_group_v1beta1_tx_proto.Messages().ByName("MsgCreateProposal")
	fd_MsgCreateProposal_address = md_MsgCreateProposal.Fields().ByName("address")
	fd_MsgCreateProposal_proposers = md_MsgCreateProposal.Fields().ByName("proposers")
	fd_MsgCreateProposal_metadata = md_MsgCreateProposal.Fields().ByName("metadata")
	fd_MsgCreateProposal_msgs = md_MsgCreateProposal.Fields().ByName("msgs")
	fd_MsgCreateProposal_exec = md_MsgCreateProposal.Fields().ByName("exec")
}

var _ protoreflect.Message = (*fastReflection_MsgCreateProposal)(nil)

type fastReflection_MsgCreateProposal MsgCreateProposal

func (x *MsgCreateProposal) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgCreateProposal)(x)
}

func (x *MsgCreateProposal) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_group_v1beta1_tx_proto_msgTypes[16]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgCreateProposal_messageType fastReflection_MsgCreateProposal_messageType
var _ protoreflect.MessageType = fastReflection_MsgCreateProposal_messageType{}

type fastReflection_MsgCreateProposal_messageType struct{}

func (x fastReflection_MsgCreateProposal_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgCreateProposal)(nil)
}
func (x fastReflection_MsgCreateProposal_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgCreateProposal)
}
func (x fastReflection_MsgCreateProposal_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgCreateProposal
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgCreateProposal) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgCreateProposal
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgCreateProposal) Type() protoreflect.MessageType {
	return _fastReflection_MsgCreateProposal_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgCreateProposal) New() protoreflect.Message {
	return new(fastReflection_MsgCreateProposal)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgCreateProposal) Interface() protoreflect.ProtoMessage {
	return (*MsgCreateProposal)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgCreateProposal) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Address != "" {
		value := protoreflect.ValueOfString(x.Address)
		if !f(fd_MsgCreateProposal_address, value) {
			return
		}
	}
	if len(x.Proposers) != 0 {
		value := protoreflect.ValueOfList(&_MsgCreateProposal_2_list{list: &x.Proposers})
		if !f(fd_MsgCreateProposal_proposers, value) {
			return
		}
	}
	if len(x.Metadata) != 0 {
		value := protoreflect.ValueOfBytes(x.Metadata)
		if !f(fd_MsgCreateProposal_metadata, value) {
			return
		}
	}
	if len(x.Msgs) != 0 {
		value := protoreflect.ValueOfList(&_MsgCreateProposal_4_list{list: &x.Msgs})
		if !f(fd_MsgCreateProposal_msgs, value) {
			return
		}
	}
	if x.Exec != 0 {
		value := protoreflect.ValueOfEnum((protoreflect.EnumNumber)(x.Exec))
		if !f(fd_MsgCreateProposal_exec, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_MsgCreateProposal) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.group.v1beta1.MsgCreateProposal.address":
		return x.Address != ""
	case "cosmos.group.v1beta1.MsgCreateProposal.proposers":
		return len(x.Proposers) != 0
	case "cosmos.group.v1beta1.MsgCreateProposal.metadata":
		return len(x.Metadata) != 0
	case "cosmos.group.v1beta1.MsgCreateProposal.msgs":
		return len(x.Msgs) != 0
	case "cosmos.group.v1beta1.MsgCreateProposal.exec":
		return x.Exec != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgCreateProposal"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgCreateProposal does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateProposal) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.group.v1beta1.MsgCreateProposal.address":
		x.Address = ""
	case "cosmos.group.v1beta1.MsgCreateProposal.proposers":
		x.Proposers = nil
	case "cosmos.group.v1beta1.MsgCreateProposal.metadata":
		x.Metadata = nil
	case "cosmos.group.v1beta1.MsgCreateProposal.msgs":
		x.Msgs = nil
	case "cosmos.group.v1beta1.MsgCreateProposal.exec":
		x.Exec = 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgCreateProposal"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgCreateProposal does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgCreateProposal) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.group.v1beta1.MsgCreateProposal.address":
		value := x.Address
		return protoreflect.ValueOfString(value)
	case "cosmos.group.v1beta1.MsgCreateProposal.proposers":
		if len(x.Proposers) == 0 {
			return protoreflect.ValueOfList(&_MsgCreateProposal_2_list{})
		}
		listValue := &_MsgCreateProposal_2_list{list: &x.Proposers}
		return protoreflect.ValueOfList(listValue)
	case "cosmos.group.v1beta1.MsgCreateProposal.metadata":
		value := x.Metadata
		return protoreflect.ValueOfBytes(value)
	case "cosmos.group.v1beta1.MsgCreateProposal.msgs":
		if len(x.Msgs) == 0 {
			return protoreflect.ValueOfList(&_MsgCreateProposal_4_list{})
		}
		listValue := &_MsgCreateProposal_4_list{list: &x.Msgs}
		return protoreflect.ValueOfList(listValue)
	case "cosmos.group.v1beta1.MsgCreateProposal.exec":
		value := x.Exec
		return protoreflect.ValueOfEnum((protoreflect.EnumNumber)(value))
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgCreateProposal"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgCreateProposal does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateProposal) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.group.v1beta1.MsgCreateProposal.address":
		x.Address = value.Interface().(string)
	case "cosmos.group.v1beta1.MsgCreateProposal.proposers":
		lv := value.List()
		clv := lv.(*_MsgCreateProposal_2_list)
		x.Proposers = *clv.list
	case "cosmos.group.v1beta1.MsgCreateProposal.metadata":
		x.Metadata = value.Bytes()
	case "cosmos.group.v1beta1.MsgCreateProposal.msgs":
		lv := value.List()
		clv := lv.(*_MsgCreateProposal_4_list)
		x.Msgs = *clv.list
	case "cosmos.group.v1beta1.MsgCreateProposal.exec":
		x.Exec = (Exec)(value.Enum())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgCreateProposal"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgCreateProposal does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateProposal) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.group.v1beta1.MsgCreateProposal.proposers":
		if x.Proposers == nil {
			x.Proposers = []string{}
		}
		value := &_MsgCreateProposal_2_list{list: &x.Proposers}
		return protoreflect.ValueOfList(value)
	case "cosmos.group.v1beta1.MsgCreateProposal.msgs":
		if x.Msgs == nil {
			x.Msgs = []*anypb.Any{}
		}
		value := &_MsgCreateProposal_4_list{list: &x.Msgs}
		return protoreflect.ValueOfList(value)
	case "cosmos.group.v1beta1.MsgCreateProposal.address":
		panic(fmt.Errorf("field address of message cosmos.group.v1beta1.MsgCreateProposal is not mutable"))
	case "cosmos.group.v1beta1.MsgCreateProposal.metadata":
		panic(fmt.Errorf("field metadata of message cosmos.group.v1beta1.MsgCreateProposal is not mutable"))
	case "cosmos.group.v1beta1.MsgCreateProposal.exec":
		panic(fmt.Errorf("field exec of message cosmos.group.v1beta1.MsgCreateProposal is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgCreateProposal"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgCreateProposal does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgCreateProposal) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.group.v1beta1.MsgCreateProposal.address":
		return protoreflect.ValueOfString("")
	case "cosmos.group.v1beta1.MsgCreateProposal.proposers":
		list := []string{}
		return protoreflect.ValueOfList(&_MsgCreateProposal_2_list{list: &list})
	case "cosmos.group.v1beta1.MsgCreateProposal.metadata":
		return protoreflect.ValueOfBytes(nil)
	case "cosmos.group.v1beta1.MsgCreateProposal.msgs":
		list := []*anypb.Any{}
		return protoreflect.ValueOfList(&_MsgCreateProposal_4_list{list: &list})
	case "cosmos.group.v1beta1.MsgCreateProposal.exec":
		return protoreflect.ValueOfEnum(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgCreateProposal"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgCreateProposal does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgCreateProposal) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.group.v1beta1.MsgCreateProposal", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgCreateProposal) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateProposal) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_MsgCreateProposal) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgCreateProposal) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgCreateProposal)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.Address)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if len(x.Proposers) > 0 {
			for _, s := range x.Proposers {
				l = len(s)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		l = len(x.Metadata)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if len(x.Msgs) > 0 {
			for _, e := range x.Msgs {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.Exec != 0 {
			n += 1 + runtime.Sov(uint64(x.Exec))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*MsgCreateProposal)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.Exec != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Exec))
			i--
			dAtA[i] = 0x28
		}
		if len(x.Msgs) > 0 {
			for iNdEx := len(x.Msgs) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Msgs[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0x22
			}
		}
		if len(x.Metadata) > 0 {
			i -= len(x.Metadata)
			copy(dAtA[i:], x.Metadata)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Metadata)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.Proposers) > 0 {
			for iNdEx := len(x.Proposers) - 1; iNdEx >= 0; iNdEx-- {
				i -= len(x.Proposers[iNdEx])
				copy(dAtA[i:], x.Proposers[iNdEx])
				i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Proposers[iNdEx])))
				i--
				dAtA[i] = 0x12
			}
		}
		if len(x.Address) > 0 {
			i -= len(x.Address)
			copy(dAtA[i:], x.Address)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Address)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*MsgCreateProposal)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgCreateProposal: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgCreateProposal: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Address = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Proposers", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Proposers = append(x.Proposers, string(dAtA[iNdEx:postIndex]))
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Metadata = append(x.Metadata[:0], dAtA[iNdEx:postIndex]...)
				if x.Metadata == nil {
					x.Metadata = []byte{}
				}
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Msgs", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Msgs = append(x.Msgs, &anypb.Any{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Msgs[len(x.Msgs)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 5:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Exec", wireType)
				}
				x.Exec = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Exec |= Exec(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_MsgCreateProposalResponse             protoreflect.MessageDescriptor
	fd_MsgCreateProposalResponse_proposal_id protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_group_v1beta1_tx_proto_init()
	md_MsgCreateProposalResponse = File_cosmos_group_v1beta1_tx_proto.Messages().ByName("MsgCreateProposalResponse")
	fd_MsgCreateProposalResponse_proposal_id = md_MsgCreateProposalResponse.Fields().ByName("proposal_id")
}

var _ protoreflect.Message = (*fastReflection_MsgCreateProposalResponse)(nil)

type fastReflection_MsgCreateProposalResponse MsgCreateProposalResponse

func (x *MsgCreateProposalResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgCreateProposalResponse)(x)
}

func (x *MsgCreateProposalResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_group_v1beta1_tx_proto_msgTypes[17]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgCreateProposalResponse_messageType fastReflection_MsgCreateProposalResponse_messageType
var _ protoreflect.MessageType = fastReflection_MsgCreateProposalResponse_messageType{}

type fastReflection_MsgCreateProposalResponse_messageType struct{}

func (x fastReflection_MsgCreateProposalResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgCreateProposalResponse)(nil)
}
func (x fastReflection_MsgCreateProposalResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgCreateProposalResponse)
}
func (x fastReflection_MsgCreateProposalResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgCreateProposalResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgCreateProposalResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgCreateProposalResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgCreateProposalResponse) Type() protoreflect.MessageType {
	return _fastReflection_MsgCreateProposalResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgCreateProposalResponse) New() protoreflect.Message {
	return new(fastReflection_MsgCreateProposalResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgCreateProposalResponse) Interface() protoreflect.ProtoMessage {
	return (*MsgCreateProposalResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgCreateProposalResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.ProposalId != uint64(0) {
		value := protoreflect.ValueOfUint64(x.ProposalId)
		if !f(fd_MsgCreateProposalResponse_proposal_id, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_MsgCreateProposalResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.group.v1beta1.MsgCreateProposalResponse.proposal_id":
		return x.ProposalId != uint64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgCreateProposalResponse"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgCreateProposalResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateProposalResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.group.v1beta1.MsgCreateProposalResponse.proposal_id":
		x.ProposalId = uint64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgCreateProposalResponse"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgCreateProposalResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgCreateProposalResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.group.v1beta1.MsgCreateProposalResponse.proposal_id":
		value := x.ProposalId
		return protoreflect.ValueOfUint64(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgCreateProposalResponse"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgCreateProposalResponse does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateProposalResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.group.v1beta1.MsgCreateProposalResponse.proposal_id":
		x.ProposalId = value.Uint()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgCreateProposalResponse"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgCreateProposalResponse does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateProposalResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.group.v1beta1.MsgCreateProposalResponse.proposal_id":
		panic(fmt.Errorf("field proposal_id of message cosmos.group.v1beta1.MsgCreateProposalResponse is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgCreateProposalResponse"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgCreateProposalResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgCreateProposalResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.group.v1beta1.MsgCreateProposalResponse.proposal_id":
		return protoreflect.ValueOfUint64(uint64(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgCreateProposalResponse"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgCreateProposalResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgCreateProposalResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.group.v1beta1.MsgCreateProposalResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgCreateProposalResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateProposalResponse) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_MsgCreateProposalResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgCreateProposalResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgCreateProposalResponse)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.ProposalId != 0 {
			n += 1 + runtime.Sov(uint64(x.ProposalId))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*MsgCreateProposalResponse)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.ProposalId != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.ProposalId))
			i--
			dAtA[i] = 0x8
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*MsgCreateProposalResponse)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgCreateProposalResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgCreateProposalResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ProposalId", wireType)
				}
				x.ProposalId = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.ProposalId |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_MsgVote             protoreflect.MessageDescriptor
	fd_MsgVote_proposal_id protoreflect.FieldDescriptor
	fd_MsgVote_voter       protoreflect.FieldDescriptor
	fd_MsgVote_choice      protoreflect.FieldDescriptor
	fd_MsgVote_metadata    protoreflect.FieldDescriptor
	fd_MsgVote_exec        protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_group_v1beta1_tx_proto_init()
	md_MsgVote = File_cosmos_group_v1beta1_tx_proto.Messages().ByName("MsgVote")
	fd_MsgVote_proposal_id = md_MsgVote.Fields().ByName("proposal_id")
	fd_MsgVote_voter = md_MsgVote.Fields().ByName("voter")
	fd_MsgVote_choice = md_MsgVote.Fields().ByName("choice")
	fd_MsgVote_metadata = md_MsgVote.Fields().ByName("metadata")
	fd_MsgVote_exec = md_MsgVote.Fields().ByName("exec")
}

var _ protoreflect.Message = (*fastReflection_MsgVote)(nil)

type fastReflection_MsgVote MsgVote

func (x *MsgVote) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgVote)(x)
}

func (x *MsgVote) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_group_v1beta1_tx_proto_msgTypes[18]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgVote_messageType fastReflection_MsgVote_messageType
var _ protoreflect.MessageType = fastReflection_MsgVote_messageType{}

type fastReflection_MsgVote_messageType struct{}

func (x fastReflection_MsgVote_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgVote)(nil)
}
func (x fastReflection_MsgVote_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgVote)
}
func (x fastReflection_MsgVote_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgVote
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgVote) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgVote
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgVote) Type() protoreflect.MessageType {
	return _fastReflection_MsgVote_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgVote) New() protoreflect.Message {
	return new(fastReflection_MsgVote)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgVote) Interface() protoreflect.ProtoMessage {
	return (*MsgVote)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgVote) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.ProposalId != uint64(0) {
		value := protoreflect.ValueOfUint64(x.ProposalId)
		if !f(fd_MsgVote_proposal_id, value) {
			return
		}
	}
	if x.Voter != "" {
		value := protoreflect.ValueOfString(x.Voter)
		if !f(fd_MsgVote_voter, value) {
			return
		}
	}
	if x.Choice != 0 {
		value := protoreflect.ValueOfEnum((protoreflect.EnumNumber)(x.Choice))
		if !f(fd_MsgVote_choice, value) {
			return
		}
	}
	if len(x.Metadata) != 0 {
		value := protoreflect.ValueOfBytes(x.Metadata)
		if !f(fd_MsgVote_metadata, value) {
			return
		}
	}
	if x.Exec != 0 {
		value := protoreflect.ValueOfEnum((protoreflect.EnumNumber)(x.Exec))
		if !f(fd_MsgVote_exec, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_MsgVote) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.group.v1beta1.MsgVote.proposal_id":
		return x.ProposalId != uint64(0)
	case "cosmos.group.v1beta1.MsgVote.voter":
		return x.Voter != ""
	case "cosmos.group.v1beta1.MsgVote.choice":
		return x.Choice != 0
	case "cosmos.group.v1beta1.MsgVote.metadata":
		return len(x.Metadata) != 0
	case "cosmos.group.v1beta1.MsgVote.exec":
		return x.Exec != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgVote"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgVote does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgVote) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.group.v1beta1.MsgVote.proposal_id":
		x.ProposalId = uint64(0)
	case "cosmos.group.v1beta1.MsgVote.voter":
		x.Voter = ""
	case "cosmos.group.v1beta1.MsgVote.choice":
		x.Choice = 0
	case "cosmos.group.v1beta1.MsgVote.metadata":
		x.Metadata = nil
	case "cosmos.group.v1beta1.MsgVote.exec":
		x.Exec = 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgVote"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgVote does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgVote) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.group.v1beta1.MsgVote.proposal_id":
		value := x.ProposalId
		return protoreflect.ValueOfUint64(value)
	case "cosmos.group.v1beta1.MsgVote.voter":
		value := x.Voter
		return protoreflect.ValueOfString(value)
	case "cosmos.group.v1beta1.MsgVote.choice":
		value := x.Choice
		return protoreflect.ValueOfEnum((protoreflect.EnumNumber)(value))
	case "cosmos.group.v1beta1.MsgVote.metadata":
		value := x.Metadata
		return protoreflect.ValueOfBytes(value)
	case "cosmos.group.v1beta1.MsgVote.exec":
		value := x.Exec
		return protoreflect.ValueOfEnum((protoreflect.EnumNumber)(value))
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgVote"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgVote does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgVote) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.group.v1beta1.MsgVote.proposal_id":
		x.ProposalId = value.Uint()
	case "cosmos.group.v1beta1.MsgVote.voter":
		x.Voter = value.Interface().(string)
	case "cosmos.group.v1beta1.MsgVote.choice":
		x.Choice = (Choice)(value.Enum())
	case "cosmos.group.v1beta1.MsgVote.metadata":
		x.Metadata = value.Bytes()
	case "cosmos.group.v1beta1.MsgVote.exec":
		x.Exec = (Exec)(value.Enum())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgVote"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgVote does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgVote) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.group.v1beta1.MsgVote.proposal_id":
		panic(fmt.Errorf("field proposal_id of message cosmos.group.v1beta1.MsgVote is not mutable"))
	case "cosmos.group.v1beta1.MsgVote.voter":
		panic(fmt.Errorf("field voter of message cosmos.group.v1beta1.MsgVote is not mutable"))
	case "cosmos.group.v1beta1.MsgVote.choice":
		panic(fmt.Errorf("field choice of message cosmos.group.v1beta1.MsgVote is not mutable"))
	case "cosmos.group.v1beta1.MsgVote.metadata":
		panic(fmt.Errorf("field metadata of message cosmos.group.v1beta1.MsgVote is not mutable"))
	case "cosmos.group.v1beta1.MsgVote.exec":
		panic(fmt.Errorf("field exec of message cosmos.group.v1beta1.MsgVote is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgVote"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgVote does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgVote) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.group.v1beta1.MsgVote.proposal_id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "cosmos.group.v1beta1.MsgVote.voter":
		return protoreflect.ValueOfString("")
	case "cosmos.group.v1beta1.MsgVote.choice":
		return protoreflect.ValueOfEnum(0)
	case "cosmos.group.v1beta1.MsgVote.metadata":
		return protoreflect.ValueOfBytes(nil)
	case "cosmos.group.v1beta1.MsgVote.exec":
		return protoreflect.ValueOfEnum(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgVote"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgVote does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgVote) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.group.v1beta1.MsgVote", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgVote) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgVote) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_MsgVote) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgVote) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgVote)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.ProposalId != 0 {
			n += 1 + runtime.Sov(uint64(x.ProposalId))
		}
		l = len(x.Voter)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Choice != 0 {
			n += 1 + runtime.Sov(uint64(x.Choice))
		}
		l = len(x.Metadata)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Exec != 0 {
			n += 1 + runtime.Sov(uint64(x.Exec))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*MsgVote)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.Exec != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Exec))
			i--
			dAtA[i] = 0x28
		}
		if len(x.Metadata) > 0 {
			i -= len(x.Metadata)
			copy(dAtA[i:], x.Metadata)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Metadata)))
			i--
			dAtA[i] = 0x22
		}
		if x.Choice != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Choice))
			i--
			dAtA[i] = 0x18
		}
		if len(x.Voter) > 0 {
			i -= len(x.Voter)
			copy(dAtA[i:], x.Voter)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Voter)))
			i--
			dAtA[i] = 0x12
		}
		if x.ProposalId != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.ProposalId))
			i--
			dAtA[i] = 0x8
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*MsgVote)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgVote: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgVote: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ProposalId", wireType)
				}
				x.ProposalId = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.ProposalId |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Voter", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Voter = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Choice", wireType)
				}
				x.Choice = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Choice |= Choice(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Metadata = append(x.Metadata[:0], dAtA[iNdEx:postIndex]...)
				if x.Metadata == nil {
					x.Metadata = []byte{}
				}
				iNdEx = postIndex
			case 5:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Exec", wireType)
				}
				x.Exec = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Exec |= Exec(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_MsgVoteResponse protoreflect.MessageDescriptor
)

func init() {
	file_cosmos_group_v1beta1_tx_proto_init()
	md_MsgVoteResponse = File_cosmos_group_v1beta1_tx_proto.Messages().ByName("MsgVoteResponse")
}

var _ protoreflect.Message = (*fastReflection_MsgVoteResponse)(nil)

type fastReflection_MsgVoteResponse MsgVoteResponse

func (x *MsgVoteResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgVoteResponse)(x)
}

func (x *MsgVoteResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_group_v1beta1_tx_proto_msgTypes[19]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgVoteResponse_messageType fastReflection_MsgVoteResponse_messageType
var _ protoreflect.MessageType = fastReflection_MsgVoteResponse_messageType{}

type fastReflection_MsgVoteResponse_messageType struct{}

func (x fastReflection_MsgVoteResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgVoteResponse)(nil)
}
func (x fastReflection_MsgVoteResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgVoteResponse)
}
func (x fastReflection_MsgVoteResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgVoteResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgVoteResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgVoteResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgVoteResponse) Type() protoreflect.MessageType {
	return _fastReflection_MsgVoteResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgVoteResponse) New() protoreflect.Message {
	return new(fastReflection_MsgVoteResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgVoteResponse) Interface() protoreflect.ProtoMessage {
	return (*MsgVoteResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgVoteResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_MsgVoteResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgVoteResponse"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgVoteResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgVoteResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgVoteResponse"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgVoteResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgVoteResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgVoteResponse"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgVoteResponse does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgVoteResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgVoteResponse"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgVoteResponse does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgVoteResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgVoteResponse"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgVoteResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgVoteResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgVoteResponse"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgVoteResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgVoteResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.group.v1beta1.MsgVoteResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgVoteResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgVoteResponse) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_MsgVoteResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgVoteResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgVoteResponse)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*MsgVoteResponse)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*MsgVoteResponse)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgVoteResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgVoteResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_MsgExec             protoreflect.MessageDescriptor
	fd_MsgExec_proposal_id protoreflect.FieldDescriptor
	fd_MsgExec_signer      protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_group_v1beta1_tx_proto_init()
	md_MsgExec = File_cosmos_group_v1beta1_tx_proto.Messages().ByName("MsgExec")
	fd_MsgExec_proposal_id = md_MsgExec.Fields().ByName("proposal_id")
	fd_MsgExec_signer = md_MsgExec.Fields().ByName("signer")
}

var _ protoreflect.Message = (*fastReflection_MsgExec)(nil)

type fastReflection_MsgExec MsgExec

func (x *MsgExec) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgExec)(x)
}

func (x *MsgExec) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_group_v1beta1_tx_proto_msgTypes[20]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgExec_messageType fastReflection_MsgExec_messageType
var _ protoreflect.MessageType = fastReflection_MsgExec_messageType{}

type fastReflection_MsgExec_messageType struct{}

func (x fastReflection_MsgExec_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgExec)(nil)
}
func (x fastReflection_MsgExec_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgExec)
}
func (x fastReflection_MsgExec_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgExec
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgExec) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgExec
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgExec) Type() protoreflect.MessageType {
	return _fastReflection_MsgExec_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgExec) New() protoreflect.Message {
	return new(fastReflection_MsgExec)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgExec) Interface() protoreflect.ProtoMessage {
	return (*MsgExec)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgExec) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.ProposalId != uint64(0) {
		value := protoreflect.ValueOfUint64(x.ProposalId)
		if !f(fd_MsgExec_proposal_id, value) {
			return
		}
	}
	if x.Signer != "" {
		value := protoreflect.ValueOfString(x.Signer)
		if !f(fd_MsgExec_signer, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_MsgExec) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.group.v1beta1.MsgExec.proposal_id":
		return x.ProposalId != uint64(0)
	case "cosmos.group.v1beta1.MsgExec.signer":
		return x.Signer != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgExec"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgExec does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgExec) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.group.v1beta1.MsgExec.proposal_id":
		x.ProposalId = uint64(0)
	case "cosmos.group.v1beta1.MsgExec.signer":
		x.Signer = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgExec"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgExec does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgExec) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.group.v1beta1.MsgExec.proposal_id":
		value := x.ProposalId
		return protoreflect.ValueOfUint64(value)
	case "cosmos.group.v1beta1.MsgExec.signer":
		value := x.Signer
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgExec"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgExec does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgExec) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.group.v1beta1.MsgExec.proposal_id":
		x.ProposalId = value.Uint()
	case "cosmos.group.v1beta1.MsgExec.signer":
		x.Signer = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgExec"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgExec does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgExec) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.group.v1beta1.MsgExec.proposal_id":
		panic(fmt.Errorf("field proposal_id of message cosmos.group.v1beta1.MsgExec is not mutable"))
	case "cosmos.group.v1beta1.MsgExec.signer":
		panic(fmt.Errorf("field signer of message cosmos.group.v1beta1.MsgExec is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgExec"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgExec does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgExec) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.group.v1beta1.MsgExec.proposal_id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "cosmos.group.v1beta1.MsgExec.signer":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgExec"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgExec does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgExec) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.group.v1beta1.MsgExec", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgExec) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgExec) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_MsgExec) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgExec) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgExec)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.ProposalId != 0 {
			n += 1 + runtime.Sov(uint64(x.ProposalId))
		}
		l = len(x.Signer)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*MsgExec)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.Signer) > 0 {
			i -= len(x.Signer)
			copy(dAtA[i:], x.Signer)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Signer)))
			i--
			dAtA[i] = 0x12
		}
		if x.ProposalId != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.ProposalId))
			i--
			dAtA[i] = 0x8
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*MsgExec)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgExec: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgExec: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ProposalId", wireType)
				}
				x.ProposalId = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.ProposalId |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Signer = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_MsgExecResponse protoreflect.MessageDescriptor
)

func init() {
	file_cosmos_group_v1beta1_tx_proto_init()
	md_MsgExecResponse = File_cosmos_group_v1beta1_tx_proto.Messages().ByName("MsgExecResponse")
}

var _ protoreflect.Message = (*fastReflection_MsgExecResponse)(nil)

type fastReflection_MsgExecResponse MsgExecResponse

func (x *MsgExecResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgExecResponse)(x)
}

func (x *MsgExecResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_group_v1beta1_tx_proto_msgTypes[21]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgExecResponse_messageType fastReflection_MsgExecResponse_messageType
var _ protoreflect.MessageType = fastReflection_MsgExecResponse_messageType{}

type fastReflection_MsgExecResponse_messageType struct{}

func (x fastReflection_MsgExecResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgExecResponse)(nil)
}
func (x fastReflection_MsgExecResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgExecResponse)
}
func (x fastReflection_MsgExecResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgExecResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgExecResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgExecResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgExecResponse) Type() protoreflect.MessageType {
	return _fastReflection_MsgExecResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgExecResponse) New() protoreflect.Message {
	return new(fastReflection_MsgExecResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgExecResponse) Interface() protoreflect.ProtoMessage {
	return (*MsgExecResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgExecResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_MsgExecResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgExecResponse"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgExecResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgExecResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgExecResponse"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgExecResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgExecResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgExecResponse"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgExecResponse does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgExecResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgExecResponse"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgExecResponse does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgExecResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgExecResponse"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgExecResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgExecResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.group.v1beta1.MsgExecResponse"))
		}
		panic(fmt.Errorf("message cosmos.group.v1beta1.MsgExecResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgExecResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.group.v1beta1.MsgExecResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgExecResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgExecResponse) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_MsgExecResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgExecResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgExecResponse)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*MsgExecResponse)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*MsgExecResponse)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgExecResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgExecResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: cosmos/group/v1beta1/tx.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Exec defines modes of execution of a proposal on creation or on new vote.
type Exec int32

const (
	// An empty value means that there should be a separate
	// MsgExec request for the proposal to execute.
	Exec_EXEC_UNSPECIFIED Exec = 0
	// Try to execute the proposal immediately.
	// If the proposal is not allowed per the DecisionPolicy,
	// the proposal will still be open and could
	// be executed at a later point.
	Exec_EXEC_TRY Exec = 1
)

// Enum value maps for Exec.
var (
	Exec_name = map[int32]string{
		0: "EXEC_UNSPECIFIED",
		1: "EXEC_TRY",
	}
	Exec_value = map[string]int32{
		"EXEC_UNSPECIFIED": 0,
		"EXEC_TRY":         1,
	}
)

func (x Exec) Enum() *Exec {
	p := new(Exec)
	*p = x
	return p
}

func (x Exec) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Exec) Descriptor() protoreflect.EnumDescriptor {
	return file_cosmos_group_v1beta1_tx_proto_enumTypes[0].Descriptor()
}

func (Exec) Type() protoreflect.EnumType {
	return &file_cosmos_group_v1beta1_tx_proto_enumTypes[0]
}

func (x Exec) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Exec.Descriptor instead.
func (Exec) EnumDescriptor() ([]byte, []int) {
	return file_cosmos_group_v1beta1_tx_proto_rawDescGZIP(), []int{0}
}

// MsgCreateGroup is the Msg/CreateGroup request type.
type MsgCreateGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// admin is the account address of the group admin.
	Admin string `protobuf:"bytes,1,opt,name=admin,proto3" json:"admin,omitempty"`
	// members defines the group members.
	Members []*Member `protobuf:"bytes,2,rep,name=members,proto3" json:"members,omitempty"`
	// metadata is any arbitrary metadata to attached to the group.
	Metadata []byte `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *MsgCreateGroup) Reset() {
	*x = MsgCreateGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_group_v1beta1_tx_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgCreateGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgCreateGroup) ProtoMessage() {}

// Deprecated: Use MsgCreateGroup.ProtoReflect.Descriptor instead.
func (*MsgCreateGroup) Descriptor() ([]byte, []int) {
	return file_cosmos_group_v1beta1_tx_proto_rawDescGZIP(), []int{0}
}

func (x *MsgCreateGroup) GetAdmin() string {
	if x != nil {
		return x.Admin
	}
	return ""
}

func (x *MsgCreateGroup) GetMembers() []*Member {
	if x != nil {
		return x.Members
	}
	return nil
}

func (x *MsgCreateGroup) GetMetadata() []byte {
	if x != nil {
		return x.Metadata
	}
	return nil
}

// MsgCreateGroupResponse is the Msg/CreateGroup response type.
type MsgCreateGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// group_id is the unique ID of the newly created group.
	GroupId uint64 `protobuf:"varint,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
}

func (x *MsgCreateGroupResponse) Reset() {
	*x = MsgCreateGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_group_v1beta1_tx_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgCreateGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgCreateGroupResponse) ProtoMessage() {}

// Deprecated: Use MsgCreateGroupResponse.ProtoReflect.Descriptor instead.
func (*MsgCreateGroupResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_group_v1beta1_tx_proto_rawDescGZIP(), []int{1}
}

func (x *MsgCreateGroupResponse) GetGroupId() uint64 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

// MsgUpdateGroupMembers is the Msg/UpdateGroupMembers request type.
type MsgUpdateGroupMembers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// admin is the account address of the group admin.
	Admin string `protobuf:"bytes,1,opt,name=admin,proto3" json:"admin,omitempty"`
	// group_id is the unique ID of the group.
	GroupId uint64 `protobuf:"varint,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	// member_updates is the list of members to update,
	// set weight to 0 to remove a member.
	MemberUpdates []*Member `protobuf:"bytes,3,rep,name=member_updates,json=memberUpdates,proto3" json:"member_updates,omitempty"`
}

func (x *MsgUpdateGroupMembers) Reset() {
	*x = MsgUpdateGroupMembers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_group_v1beta1_tx_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgUpdateGroupMembers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgUpdateGroupMembers) ProtoMessage() {}

// Deprecated: Use MsgUpdateGroupMembers.ProtoReflect.Descriptor instead.
func (*MsgUpdateGroupMembers) Descriptor() ([]byte, []int) {
	return file_cosmos_group_v1beta1_tx_proto_rawDescGZIP(), []int{2}
}

func (x *MsgUpdateGroupMembers) GetAdmin() string {
	if x != nil {
		return x.Admin
	}
	return ""
}

func (x *MsgUpdateGroupMembers) GetGroupId() uint64 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *MsgUpdateGroupMembers) GetMemberUpdates() []*Member {
	if x != nil {
		return x.MemberUpdates
	}
	return nil
}

// MsgUpdateGroupMembersResponse is the Msg/UpdateGroupMembers response type.
type MsgUpdateGroupMembersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MsgUpdateGroupMembersResponse) Reset() {
	*x = MsgUpdateGroupMembersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_group_v1beta1_tx_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgUpdateGroupMembersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgUpdateGroupMembersResponse) ProtoMessage() {}

// Deprecated: Use MsgUpdateGroupMembersResponse.ProtoReflect.Descriptor instead.
func (*MsgUpdateGroupMembersResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_group_v1beta1_tx_proto_rawDescGZIP(), []int{3}
}

// MsgUpdateGroupAdmin is the Msg/UpdateGroupAdmin request type.
type MsgUpdateGroupAdmin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// admin is the current account address of the group admin.
	Admin string `protobuf:"bytes,1,opt,name=admin,proto3" json:"admin,omitempty"`
	// group_id is the unique ID of the group.
	GroupId uint64 `protobuf:"varint,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	// new_admin is the group new admin account address.
	NewAdmin string `protobuf:"bytes,3,opt,name=new_admin,json=newAdmin,proto3" json:"new_admin,omitempty"`
}

func (x *MsgUpdateGroupAdmin) Reset() {
	*x = MsgUpdateGroupAdmin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_group_v1beta1_tx_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgUpdateGroupAdmin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgUpdateGroupAdmin) ProtoMessage() {}

// Deprecated: Use MsgUpdateGroupAdmin.ProtoReflect.Descriptor instead.
func (*MsgUpdateGroupAdmin) Descriptor() ([]byte, []int) {
	return file_cosmos_group_v1beta1_tx_proto_rawDescGZIP(), []int{4}
}

func (x *MsgUpdateGroupAdmin) GetAdmin() string {
	if x != nil {
		return x.Admin
	}
	return ""
}

func (x *MsgUpdateGroupAdmin) GetGroupId() uint64 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *MsgUpdateGroupAdmin) GetNewAdmin() string {
	if x != nil {
		return x.NewAdmin
	}
	return ""
}

// MsgUpdateGroupAdminResponse is the Msg/UpdateGroupAdmin response type.
type MsgUpdateGroupAdminResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MsgUpdateGroupAdminResponse) Reset() {
	*x = MsgUpdateGroupAdminResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_group_v1beta1_tx_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgUpdateGroupAdminResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgUpdateGroupAdminResponse) ProtoMessage() {}

// Deprecated: Use MsgUpdateGroupAdminResponse.ProtoReflect.Descriptor instead.
func (*MsgUpdateGroupAdminResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_group_v1beta1_tx_proto_rawDescGZIP(), []int{5}
}

// MsgUpdateGroupMetadata is the Msg/UpdateGroupMetadata request type.
type MsgUpdateGroupMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// admin is the account address of the group admin.
	Admin string `protobuf:"bytes,1,opt,name=admin,proto3" json:"admin,omitempty"`
	// group_id is the unique ID of the group.
	GroupId uint64 `protobuf:"varint,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	// metadata is the updated group's metadata.
	Metadata []byte `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *MsgUpdateGroupMetadata) Reset() {
	*x = MsgUpdateGroupMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_group_v1beta1_tx_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgUpdateGroupMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgUpdateGroupMetadata) ProtoMessage() {}

// Deprecated: Use MsgUpdateGroupMetadata.ProtoReflect.Descriptor instead.
func (*MsgUpdateGroupMetadata) Descriptor() ([]byte, []int) {
	return file_cosmos_group_v1beta1_tx_proto_rawDescGZIP(), []int{6}
}

func (x *MsgUpdateGroupMetadata) GetAdmin() string {
	if x != nil {
		return x.Admin
	}
	return ""
}

func (x *MsgUpdateGroupMetadata) GetGroupId() uint64 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *MsgUpdateGroupMetadata) GetMetadata() []byte {
	if x != nil {
		return x.Metadata
	}
	return nil
}

// MsgUpdateGroupMetadataResponse is the Msg/UpdateGroupMetadata response type.
type MsgUpdateGroupMetadataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MsgUpdateGroupMetadataResponse) Reset() {
	*x = MsgUpdateGroupMetadataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_group_v1beta1_tx_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgUpdateGroupMetadataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgUpdateGroupMetadataResponse) ProtoMessage() {}

// Deprecated: Use MsgUpdateGroupMetadataResponse.ProtoReflect.Descriptor instead.
func (*MsgUpdateGroupMetadataResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_group_v1beta1_tx_proto_rawDescGZIP(), []int{7}
}

// MsgCreateGroupPolicy is the Msg/CreateGroupPolicy request type.
type MsgCreateGroupPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// admin is the account address of the group admin.
	Admin string `protobuf:"bytes,1,opt,name=admin,proto3" json:"admin,omitempty"`
	// group_id is the unique ID of the group.
	GroupId uint64 `protobuf:"varint,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	// metadata is any arbitrary metadata attached to the group policy.
	Metadata []byte `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// decision_policy specifies the group policy's decision policy.
	DecisionPolicy *anypb.Any `protobuf:"bytes,4,opt,name=decision_policy,json=decisionPolicy,proto3" json:"decision_policy,omitempty"`
}

func (x *MsgCreateGroupPolicy) Reset() {
	*x = MsgCreateGroupPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_group_v1beta1_tx_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgCreateGroupPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgCreateGroupPolicy) ProtoMessage() {}

// Deprecated: Use MsgCreateGroupPolicy.ProtoReflect.Descriptor instead.
func (*MsgCreateGroupPolicy) Descriptor() ([]byte, []int) {
	return file_cosmos_group_v1beta1_tx_proto_rawDescGZIP(), []int{8}
}

func (x *MsgCreateGroupPolicy) GetAdmin() string {
	if x != nil {
		return x.Admin
	}
	return ""
}

func (x *MsgCreateGroupPolicy) GetGroupId() uint64 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *MsgCreateGroupPolicy) GetMetadata() []byte {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *MsgCreateGroupPolicy) GetDecisionPolicy() *anypb.Any {
	if x != nil {
		return x.DecisionPolicy
	}
	return nil
}

// MsgCreateGroupPolicyResponse is the Msg/CreateGroupPolicy response type.
type MsgCreateGroupPolicyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// address is the account address of the newly created group policy.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *MsgCreateGroupPolicyResponse) Reset() {
	*x = MsgCreateGroupPolicyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_group_v1beta1_tx_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgCreateGroupPolicyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgCreateGroupPolicyResponse) ProtoMessage() {}

// Deprecated: Use MsgCreateGroupPolicyResponse.ProtoReflect.Descriptor instead.
func (*MsgCreateGroupPolicyResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_group_v1beta1_tx_proto_rawDescGZIP(), []int{9}
}

func (x *MsgCreateGroupPolicyResponse) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

// MsgUpdateGroupPolicyAdmin is the Msg/UpdateGroupPolicyAdmin request type.
type MsgUpdateGroupPolicyAdmin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// admin is the account address of the group admin.
	Admin string `protobuf:"bytes,1,opt,name=admin,proto3" json:"admin,omitempty"`
	// address is the account address of the group policy.
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// new_admin is the new group policy admin.
	NewAdmin string `protobuf:"bytes,3,opt,name=new_admin,json=newAdmin,proto3" json:"new_admin,omitempty"`
}

func (x *MsgUpdateGroupPolicyAdmin) Reset() {
	*x = MsgUpdateGroupPolicyAdmin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_group_v1beta1_tx_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgUpdateGroupPolicyAdmin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgUpdateGroupPolicyAdmin) ProtoMessage() {}

// Deprecated: Use MsgUpdateGroupPolicyAdmin.ProtoReflect.Descriptor instead.
func (*MsgUpdateGroupPolicyAdmin) Descriptor() ([]byte, []int) {
	return file_cosmos_group_v1beta1_tx_proto_rawDescGZIP(), []int{10}
}

func (x *MsgUpdateGroupPolicyAdmin) GetAdmin() string {
	if x != nil {
		return x.Admin
	}
	return ""
}

func (x *MsgUpdateGroupPolicyAdmin) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *MsgUpdateGroupPolicyAdmin) GetNewAdmin() string {
	if x != nil {
		return x.NewAdmin
	}
	return ""
}

// MsgUpdateGroupPolicyAdminResponse is the Msg/UpdateGroupPolicyAdmin response type.
type MsgUpdateGroupPolicyAdminResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MsgUpdateGroupPolicyAdminResponse) Reset() {
	*x = MsgUpdateGroupPolicyAdminResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_group_v1beta1_tx_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgUpdateGroupPolicyAdminResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgUpdateGroupPolicyAdminResponse) ProtoMessage() {}

// Deprecated: Use MsgUpdateGroupPolicyAdminResponse.ProtoReflect.Descriptor instead.
func (*MsgUpdateGroupPolicyAdminResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_group_v1beta1_tx_proto_rawDescGZIP(), []int{11}
}

// MsgUpdateGroupPolicyDecisionPolicy is the Msg/UpdateGroupPolicyDecisionPolicy request type.
type MsgUpdateGroupPolicyDecisionPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// admin is the account address of the group admin.
	Admin string `protobuf:"bytes,1,opt,name=admin,proto3" json:"admin,omitempty"`
	// address is the account address of group policy.
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// decision_policy is the updated group policy's decision policy.
	DecisionPolicy *anypb.Any `protobuf:"bytes,3,opt,name=decision_policy,json=decisionPolicy,proto3" json:"decision_policy,omitempty"`
}

func (x *MsgUpdateGroupPolicyDecisionPolicy) Reset() {
	*x = MsgUpdateGroupPolicyDecisionPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_group_v1beta1_tx_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgUpdateGroupPolicyDecisionPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgUpdateGroupPolicyDecisionPolicy) ProtoMessage() {}

// Deprecated: Use MsgUpdateGroupPolicyDecisionPolicy.ProtoReflect.Descriptor instead.
func (*MsgUpdateGroupPolicyDecisionPolicy) Descriptor() ([]byte, []int) {
	return file_cosmos_group_v1beta1_tx_proto_rawDescGZIP(), []int{12}
}

func (x *MsgUpdateGroupPolicyDecisionPolicy) GetAdmin() string {
	if x != nil {
		return x.Admin
	}
	return ""
}

func (x *MsgUpdateGroupPolicyDecisionPolicy) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *MsgUpdateGroupPolicyDecisionPolicy) GetDecisionPolicy() *anypb.Any {
	if x != nil {
		return x.DecisionPolicy
	}
	return nil
}

// MsgUpdateGroupPolicyDecisionPolicyResponse is the Msg/UpdateGroupPolicyDecisionPolicy response type.
type MsgUpdateGroupPolicyDecisionPolicyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MsgUpdateGroupPolicyDecisionPolicyResponse) Reset() {
	*x = MsgUpdateGroupPolicyDecisionPolicyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_group_v1beta1_tx_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgUpdateGroupPolicyDecisionPolicyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgUpdateGroupPolicyDecisionPolicyResponse) ProtoMessage() {}

// Deprecated: Use MsgUpdateGroupPolicyDecisionPolicyResponse.ProtoReflect.Descriptor instead.
func (*MsgUpdateGroupPolicyDecisionPolicyResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_group_v1beta1_tx_proto_rawDescGZIP(), []int{13}
}

// MsgUpdateGroupPolicyMetadata is the Msg/UpdateGroupPolicyMetadata request type.
type MsgUpdateGroupPolicyMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// admin is the account address of the group admin.
	Admin string `protobuf:"bytes,1,opt,name=admin,proto3" json:"admin,omitempty"`
	// address is the account address of group policy.
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// metadata is the updated group policy metadata.
	Metadata []byte `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *MsgUpdateGroupPolicyMetadata) Reset() {
	*x = MsgUpdateGroupPolicyMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_group_v1beta1_tx_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgUpdateGroupPolicyMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgUpdateGroupPolicyMetadata) ProtoMessage() {}

// Deprecated: Use MsgUpdateGroupPolicyMetadata.ProtoReflect.Descriptor instead.
func (*MsgUpdateGroupPolicyMetadata) Descriptor() ([]byte, []int) {
	return file_cosmos_group_v1beta1_tx_proto_rawDescGZIP(), []int{14}
}

func (x *MsgUpdateGroupPolicyMetadata) GetAdmin() string {
	if x != nil {
		return x.Admin
	}
	return ""
}

func (x *MsgUpdateGroupPolicyMetadata) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *MsgUpdateGroupPolicyMetadata) GetMetadata() []byte {
	if x != nil {
		return x.Metadata
	}
	return nil
}

// MsgUpdateGroupPolicyMetadataResponse is the Msg/UpdateGroupPolicyMetadata response type.
type MsgUpdateGroupPolicyMetadataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MsgUpdateGroupPolicyMetadataResponse) Reset() {
	*x = MsgUpdateGroupPolicyMetadataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_group_v1beta1_tx_proto_msgTypes[15]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgUpdateGroupPolicyMetadataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgUpdateGroupPolicyMetadataResponse) ProtoMessage() {}

// Deprecated: Use MsgUpdateGroupPolicyMetadataResponse.ProtoReflect.Descriptor instead.
func (*MsgUpdateGroupPolicyMetadataResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_group_v1beta1_tx_proto_rawDescGZIP(), []int{15}
}

// MsgCreateProposal is the Msg/CreateProposal request type.
type MsgCreateProposal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// address is the account address of group policy.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// proposers are the account addresses of the proposers.
	// Proposers signatures will be counted as yes votes.
	Proposers []string `protobuf:"bytes,2,rep,name=proposers,proto3" json:"proposers,omitempty"`
	// metadata is any arbitrary metadata to attached to the proposal.
	Metadata []byte `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// msgs is a list of Msgs that will be executed if the proposal passes.
	Msgs []*anypb.Any `protobuf:"bytes,4,rep,name=msgs,proto3" json:"msgs,omitempty"`
	// exec defines the mode of execution of the proposal,
	// whether it should be executed immediately on creation or not.
	// If so, proposers signatures are considered as Yes votes.
	Exec Exec `protobuf:"varint,5,opt,name=exec,proto3,enum=cosmos.group.v1beta1.Exec" json:"exec,omitempty"`
}

func (x *MsgCreateProposal) Reset() {
	*x = MsgCreateProposal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_group_v1beta1_tx_proto_msgTypes[16]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgCreateProposal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgCreateProposal) ProtoMessage() {}

// Deprecated: Use MsgCreateProposal.ProtoReflect.Descriptor instead.
func (*MsgCreateProposal) Descriptor() ([]byte, []int) {
	return file_cosmos_group_v1beta1_tx_proto_rawDescGZIP(), []int{16}
}

func (x *MsgCreateProposal) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *MsgCreateProposal) GetProposers() []string {
	if x != nil {
		return x.Proposers
	}
	return nil
}

func (x *MsgCreateProposal) GetMetadata() []byte {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *MsgCreateProposal) GetMsgs() []*anypb.Any {
	if x != nil {
		return x.Msgs
	}
	return nil
}

func (x *MsgCreateProposal) GetExec() Exec {
	if x != nil {
		return x.Exec
	}
	return Exec_EXEC_UNSPECIFIED
}

// MsgCreateProposalResponse is the Msg/CreateProposal response type.
type MsgCreateProposalResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// proposal is the unique ID of the proposal.
	ProposalId uint64 `protobuf:"varint,1,opt,name=proposal_id,json=proposalId,proto3" json:"proposal_id,omitempty"`
}

func (x *MsgCreateProposalResponse) Reset() {
	*x = MsgCreateProposalResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_group_v1beta1_tx_proto_msgTypes[17]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgCreateProposalResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgCreateProposalResponse) ProtoMessage() {}

// Deprecated: Use MsgCreateProposalResponse.ProtoReflect.Descriptor instead.
func (*MsgCreateProposalResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_group_v1beta1_tx_proto_rawDescGZIP(), []int{17}
}

func (x *MsgCreateProposalResponse) GetProposalId() uint64 {
	if x != nil {
		return x.ProposalId
	}
	return 0
}

// MsgVote is the Msg/Vote request type.
type MsgVote struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// proposal is the unique ID of the proposal.
	ProposalId uint64 `protobuf:"varint,1,opt,name=proposal_id,json=proposalId,proto3" json:"proposal_id,omitempty"`
	// voter is the voter account address.
	Voter string `protobuf:"bytes,2,opt,name=voter,proto3" json:"voter,omitempty"`
	// choice is the voter's choice on the proposal.
	Choice Choice `protobuf:"varint,3,opt,name=choice,proto3,enum=cosmos.group.v1beta1.Choice" json:"choice,omitempty"`
	// metadata is any arbitrary metadata to attached to the vote.
	Metadata []byte `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// exec defines whether the proposal should be executed
	// immediately after voting or not.
	Exec Exec `protobuf:"varint,5,opt,name=exec,proto3,enum=cosmos.group.v1beta1.Exec" json:"exec,omitempty"`
}

func (x *MsgVote) Reset() {
	*x = MsgVote{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_group_v1beta1_tx_proto_msgTypes[18]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgVote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgVote) ProtoMessage() {}

// Deprecated: Use MsgVote.ProtoReflect.Descriptor instead.
func (*MsgVote) Descriptor() ([]byte, []int) {
	return file_cosmos_group_v1beta1_tx_proto_rawDescGZIP(), []int{18}
}

func (x *MsgVote) GetProposalId() uint64 {
	if x != nil {
		return x.ProposalId
	}
	return 0
}

func (x *MsgVote) GetVoter() string {
	if x != nil {
		return x.Voter
	}
	return ""
}

func (x *MsgVote) GetChoice() Choice {
	if x != nil {
		return x.Choice
	}
	return Choice_CHOICE_UNSPECIFIED
}

func (x *MsgVote) GetMetadata() []byte {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *MsgVote) GetExec() Exec {
	if x != nil {
		return x.Exec
	}
	return Exec_EXEC_UNSPECIFIED
}

// MsgVoteResponse is the Msg/Vote response type.
type MsgVoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MsgVoteResponse) Reset() {
	*x = MsgVoteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_group_v1beta1_tx_proto_msgTypes[19]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgVoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgVoteResponse) ProtoMessage() {}

// Deprecated: Use MsgVoteResponse.ProtoReflect.Descriptor instead.
func (*MsgVoteResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_group_v1beta1_tx_proto_rawDescGZIP(), []int{19}
}

// MsgExec is the Msg/Exec request type.
type MsgExec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// proposal is the unique ID of the proposal.
	ProposalId uint64 `protobuf:"varint,1,opt,name=proposal_id,json=proposalId,proto3" json:"proposal_id,omitempty"`
	// signer is the account address used to execute the proposal.
	Signer string `protobuf:"bytes,2,opt,name=signer,proto3" json:"signer,omitempty"`
}

func (x *MsgExec) Reset() {
	*x = MsgExec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_group_v1beta1_tx_proto_msgTypes[20]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgExec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgExec) ProtoMessage() {}

// Deprecated: Use MsgExec.ProtoReflect.Descriptor instead.
func (*MsgExec) Descriptor() ([]byte, []int) {
	return file_cosmos_group_v1beta1_tx_proto_rawDescGZIP(), []int{20}
}

func (x *MsgExec) GetProposalId() uint64 {
	if x != nil {
		return x.ProposalId
	}
	return 0
}

func (x *MsgExec) GetSigner() string {
	if x != nil {
		return x.Signer
	}
	return ""
}

// MsgExecResponse is the Msg/Exec request type.
type MsgExecResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MsgExecResponse) Reset() {
	*x = MsgExecResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_group_v1beta1_tx_proto_msgTypes[21]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgExecResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgExecResponse) ProtoMessage() {}

// Deprecated: Use MsgExecResponse.ProtoReflect.Descriptor instead.
func (*MsgExecResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_group_v1beta1_tx_proto_rawDescGZIP(), []int{21}
}

var File_cosmos_group_v1beta1_tx_proto protoreflect.FileDescriptor

var file_cosmos_group_v1beta1_tx_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2f, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x74, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x14, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x20, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2f,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x9a, 0x01, 0x0a, 0x0e, 0x4d, 0x73, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x2e, 0x0a, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xd2, 0xb4, 0x2d, 0x14, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52,
	0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x3c, 0x0a, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x07, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x22, 0x33, 0x0a, 0x16, 0x4d, 0x73, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x49, 0x64, 0x22, 0xad, 0x01, 0x0a, 0x15, 0x4d, 0x73, 0x67, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12,
	0x2e, 0x0a, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18,
	0xd2, 0xb4, 0x2d, 0x14, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x12,
	0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x49, 0x0a, 0x0e, 0x6d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x0d, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x73, 0x22, 0x1f, 0x0a, 0x1d, 0x4d, 0x73, 0x67, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x97, 0x01, 0x0a, 0x13, 0x4d, 0x73, 0x67, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x2e,
	0x0a, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xd2,
	0xb4, 0x2d, 0x14, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x19,
	0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x35, 0x0a, 0x09, 0x6e, 0x65, 0x77,
	0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xd2, 0xb4,
	0x2d, 0x14, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x08, 0x6e, 0x65, 0x77, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x22, 0x1d, 0x0a, 0x1b, 0x4d, 0x73, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x7f, 0x0a, 0x16, 0x4d, 0x73, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x2e, 0x0a, 0x05, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xd2, 0xb4, 0x2d, 0x14, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x52, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x22, 0x20, 0x0a, 0x1e, 0x4d, 0x73, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0xd6, 0x01, 0x0a, 0x14, 0x4d, 0x73, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x2e, 0x0a, 0x05, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xd2, 0xb4, 0x2d, 0x14,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x52, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x51, 0x0a, 0x0f, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e,
	0x79, 0x42, 0x12, 0xca, 0xb4, 0x2d, 0x0e, 0x44, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x0e, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x3a, 0x04, 0x88, 0xa0, 0x1f, 0x00, 0x22, 0x52, 0x0a, 0x1c, 0x4d,
	0x73, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xd2, 0xb4,
	0x2d, 0x14, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22,
	0xb6, 0x01, 0x0a, 0x19, 0x4d, 0x73, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x2e, 0x0a,
	0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xd2, 0xb4,
	0x2d, 0x14, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x32, 0x0a,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18,
	0xd2, 0xb4, 0x2d, 0x14, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x35, 0x0a, 0x09, 0x6e, 0x65, 0x77, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xd2, 0xb4, 0x2d, 0x14, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x08,
	0x6e, 0x65, 0x77, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x22, 0x23, 0x0a, 0x21, 0x4d, 0x73, 0x67, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xe1, 0x01,
	0x0a, 0x22, 0x4d, 0x73, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x44, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x12, 0x2e, 0x0a, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x18, 0xd2, 0xb4, 0x2d, 0x14, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x05, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x12, 0x32, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xd2, 0xb4, 0x2d, 0x14, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x51, 0x0a, 0x0f, 0x64, 0x65, 0x63, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x42, 0x12, 0xca, 0xb4, 0x2d, 0x0e, 0x44, 0x65, 0x63,
	0x69, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x0e, 0x64, 0x65, 0x63,
	0x69, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x3a, 0x04, 0x88, 0xa0, 0x1f,
	0x00, 0x22, 0x2c, 0x0a, 0x2a, 0x4d, 0x73, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x44, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x9e, 0x01, 0x0a, 0x1c, 0x4d, 0x73, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x2e, 0x0a, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x18, 0xd2, 0xb4, 0x2d, 0x14, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x12, 0x32, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x18, 0xd2, 0xb4, 0x2d, 0x14, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x22, 0x26, 0x0a, 0x24, 0x4d, 0x73, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xe1, 0x01, 0x0a, 0x11, 0x4d, 0x73, 0x67,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x12, 0x32,
	0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x18, 0xd2, 0xb4, 0x2d, 0x14, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x73,
	0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x28, 0x0a, 0x04,
	0x6d, 0x73, 0x67, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79,
	0x52, 0x04, 0x6d, 0x73, 0x67, 0x73, 0x12, 0x2e, 0x0a, 0x04, 0x65, 0x78, 0x65, 0x63, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x45, 0x78, 0x65, 0x63,
	0x52, 0x04, 0x65, 0x78, 0x65, 0x63, 0x3a, 0x04, 0x88, 0xa0, 0x1f, 0x00, 0x22, 0x3c, 0x0a, 0x19,
	0x4d, 0x73, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x6f,
	0x70, 0x6f, 0x73, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a,
	0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x49, 0x64, 0x22, 0xdc, 0x01, 0x0a, 0x07, 0x4d,
	0x73, 0x67, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73,
	0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x70, 0x72, 0x6f,
	0x70, 0x6f, 0x73, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x05, 0x76, 0x6f, 0x74, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xd2, 0xb4, 0x2d, 0x14, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x52, 0x05, 0x76, 0x6f, 0x74, 0x65, 0x72, 0x12, 0x34, 0x0a, 0x06, 0x63, 0x68, 0x6f, 0x69, 0x63,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43,
	0x68, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x06, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x2e, 0x0a, 0x04, 0x65, 0x78, 0x65,
	0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x45,
	0x78, 0x65, 0x63, 0x52, 0x04, 0x65, 0x78, 0x65, 0x63, 0x22, 0x11, 0x0a, 0x0f, 0x4d, 0x73, 0x67,
	0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5c, 0x0a, 0x07,
	0x4d, 0x73, 0x67, 0x45, 0x78, 0x65, 0x63, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x70, 0x6f,
	0x73, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x70, 0x72,
	0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x06, 0x73, 0x69, 0x67, 0x6e,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xd2, 0xb4, 0x2d, 0x14, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x52, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x22, 0x11, 0x0a, 0x0f, 0x4d, 0x73,
	0x67, 0x45, 0x78, 0x65, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2a, 0x2a, 0x0a,
	0x04, 0x45, 0x78, 0x65, 0x63, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x58, 0x45, 0x43, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x45,
	0x58, 0x45, 0x43, 0x5f, 0x54, 0x52, 0x59, 0x10, 0x01, 0x32, 0xfd, 0x09, 0x0a, 0x03, 0x4d, 0x73,
	0x67, 0x12, 0x61, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x12, 0x24, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x1a, 0x2c, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4d, 0x73,
	0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x76, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x2b, 0x2e, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x4d, 0x73, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x1a, 0x33, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4d,
	0x73, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x70, 0x0a, 0x10,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x12, 0x29, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x1a, 0x31, 0x2e, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x79,
	0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x2c, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x1a, 0x34, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x73, 0x0a, 0x11, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x2a,
	0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x1a, 0x32, 0x2e, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x4d, 0x73, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x82,
	0x01, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x2f, 0x2e, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x4d, 0x73, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x1a, 0x37, 0x2e, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x4d, 0x73, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x9d, 0x01, 0x0a, 0x1f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x44, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x38, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4d,
	0x73, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x44, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x1a, 0x40, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x44, 0x65, 0x63,
	0x69, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x8b, 0x01, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x32, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x3a, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x6a, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x6f,
	0x73, 0x61, 0x6c, 0x12, 0x27, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x1a, 0x2f, 0x2e, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f,
	0x70, 0x6f, 0x73, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a,
	0x04, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x1d, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67,
	0x56, 0x6f, 0x74, 0x65, 0x1a, 0x25, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x56,
	0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x04, 0x45,
	0x78, 0x65, 0x63, 0x12, 0x1d, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x45, 0x78,
	0x65, 0x63, 0x1a, 0x25, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x45, 0x78, 0x65,
	0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xd9, 0x01, 0x0a, 0x18, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x07, 0x54, 0x78, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x47, 0x58, 0xaa, 0x02, 0x14, 0x43, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x56, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0xca, 0x02, 0x14, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5c, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xe2, 0x02, 0x20, 0x43, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x5c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x16, 0x43,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x3a, 0x3a, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x3a, 0x3a, 0x56, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cosmos_group_v1beta1_tx_proto_rawDescOnce sync.Once
	file_cosmos_group_v1beta1_tx_proto_rawDescData = file_cosmos_group_v1beta1_tx_proto_rawDesc
)

func file_cosmos_group_v1beta1_tx_proto_rawDescGZIP() []byte {
	file_cosmos_group_v1beta1_tx_proto_rawDescOnce.Do(func() {
		file_cosmos_group_v1beta1_tx_proto_rawDescData = protoimpl.X.CompressGZIP(file_cosmos_group_v1beta1_tx_proto_rawDescData)
	})
	return file_cosmos_group_v1beta1_tx_proto_rawDescData
}

var file_cosmos_group_v1beta1_tx_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cosmos_group_v1beta1_tx_proto_msgTypes = make([]protoimpl.MessageInfo, 22)
var file_cosmos_group_v1beta1_tx_proto_goTypes = []interface{}{
	(Exec)(0),                                          // 0: cosmos.group.v1beta1.Exec
	(*MsgCreateGroup)(nil),                             // 1: cosmos.group.v1beta1.MsgCreateGroup
	(*MsgCreateGroupResponse)(nil),                     // 2: cosmos.group.v1beta1.MsgCreateGroupResponse
	(*MsgUpdateGroupMembers)(nil),                      // 3: cosmos.group.v1beta1.MsgUpdateGroupMembers
	(*MsgUpdateGroupMembersResponse)(nil),              // 4: cosmos.group.v1beta1.MsgUpdateGroupMembersResponse
	(*MsgUpdateGroupAdmin)(nil),                        // 5: cosmos.group.v1beta1.MsgUpdateGroupAdmin
	(*MsgUpdateGroupAdminResponse)(nil),                // 6: cosmos.group.v1beta1.MsgUpdateGroupAdminResponse
	(*MsgUpdateGroupMetadata)(nil),                     // 7: cosmos.group.v1beta1.MsgUpdateGroupMetadata
	(*MsgUpdateGroupMetadataResponse)(nil),             // 8: cosmos.group.v1beta1.MsgUpdateGroupMetadataResponse
	(*MsgCreateGroupPolicy)(nil),                       // 9: cosmos.group.v1beta1.MsgCreateGroupPolicy
	(*MsgCreateGroupPolicyResponse)(nil),               // 10: cosmos.group.v1beta1.MsgCreateGroupPolicyResponse
	(*MsgUpdateGroupPolicyAdmin)(nil),                  // 11: cosmos.group.v1beta1.MsgUpdateGroupPolicyAdmin
	(*MsgUpdateGroupPolicyAdminResponse)(nil),          // 12: cosmos.group.v1beta1.MsgUpdateGroupPolicyAdminResponse
	(*MsgUpdateGroupPolicyDecisionPolicy)(nil),         // 13: cosmos.group.v1beta1.MsgUpdateGroupPolicyDecisionPolicy
	(*MsgUpdateGroupPolicyDecisionPolicyResponse)(nil), // 14: cosmos.group.v1beta1.MsgUpdateGroupPolicyDecisionPolicyResponse
	(*MsgUpdateGroupPolicyMetadata)(nil),               // 15: cosmos.group.v1beta1.MsgUpdateGroupPolicyMetadata
	(*MsgUpdateGroupPolicyMetadataResponse)(nil),       // 16: cosmos.group.v1beta1.MsgUpdateGroupPolicyMetadataResponse
	(*MsgCreateProposal)(nil),                          // 17: cosmos.group.v1beta1.MsgCreateProposal
	(*MsgCreateProposalResponse)(nil),                  // 18: cosmos.group.v1beta1.MsgCreateProposalResponse
	(*MsgVote)(nil),                                    // 19: cosmos.group.v1beta1.MsgVote
	(*MsgVoteResponse)(nil),                            // 20: cosmos.group.v1beta1.MsgVoteResponse
	(*MsgExec)(nil),                                    // 21: cosmos.group.v1beta1.MsgExec
	(*MsgExecResponse)(nil),                            // 22: cosmos.group.v1beta1.MsgExecResponse
	(*Member)(nil),                                     // 23: cosmos.group.v1beta1.Member
	(*anypb.Any)(nil),                                  // 24: google.protobuf.Any
	(Choice)(0),                                        // 25: cosmos.group.v1beta1.Choice
}
var file_cosmos_group_v1beta1_tx_proto_depIdxs = []int32{
	23, // 0: cosmos.group.v1beta1.MsgCreateGroup.members:type_name -> cosmos.group.v1beta1.Member
	23, // 1: cosmos.group.v1beta1.MsgUpdateGroupMembers.member_updates:type_name -> cosmos.group.v1beta1.Member
	24, // 2: cosmos.group.v1beta1.MsgCreateGroupPolicy.decision_policy:type_name -> google.protobuf.Any
	24, // 3: cosmos.group.v1beta1.MsgUpdateGroupPolicyDecisionPolicy.decision_policy:type_name -> google.protobuf.Any
	24, // 4: cosmos.group.v1beta1.MsgCreateProposal.msgs:type_name -> google.protobuf.Any
	0,  // 5: cosmos.group.v1beta1.MsgCreateProposal.exec:type_name -> cosmos.group.v1beta1.Exec
	25, // 6: cosmos.group.v1beta1.MsgVote.choice:type_name -> cosmos.group.v1beta1.Choice
	0,  // 7: cosmos.group.v1beta1.MsgVote.exec:type_name -> cosmos.group.v1beta1.Exec
	1,  // 8: cosmos.group.v1beta1.Msg.CreateGroup:input_type -> cosmos.group.v1beta1.MsgCreateGroup
	3,  // 9: cosmos.group.v1beta1.Msg.UpdateGroupMembers:input_type -> cosmos.group.v1beta1.MsgUpdateGroupMembers
	5,  // 10: cosmos.group.v1beta1.Msg.UpdateGroupAdmin:input_type -> cosmos.group.v1beta1.MsgUpdateGroupAdmin
	7,  // 11: cosmos.group.v1beta1.Msg.UpdateGroupMetadata:input_type -> cosmos.group.v1beta1.MsgUpdateGroupMetadata
	9,  // 12: cosmos.group.v1beta1.Msg.CreateGroupPolicy:input_type -> cosmos.group.v1beta1.MsgCreateGroupPolicy
	11, // 13: cosmos.group.v1beta1.Msg.UpdateGroupPolicyAdmin:input_type -> cosmos.group.v1beta1.MsgUpdateGroupPolicyAdmin
	13, // 14: cosmos.group.v1beta1.Msg.UpdateGroupPolicyDecisionPolicy:input_type -> cosmos.group.v1beta1.MsgUpdateGroupPolicyDecisionPolicy
	15, // 15: cosmos.group.v1beta1.Msg.UpdateGroupPolicyMetadata:input_type -> cosmos.group.v1beta1.MsgUpdateGroupPolicyMetadata
	17, // 16: cosmos.group.v1beta1.Msg.CreateProposal:input_type -> cosmos.group.v1beta1.MsgCreateProposal
	19, // 17: cosmos.group.v1beta1.Msg.Vote:input_type -> cosmos.group.v1beta1.MsgVote
	21, // 18: cosmos.group.v1beta1.Msg.Exec:input_type -> cosmos.group.v1beta1.MsgExec
	2,  // 19: cosmos.group.v1beta1.Msg.CreateGroup:output_type -> cosmos.group.v1beta1.MsgCreateGroupResponse
	4,  // 20: cosmos.group.v1beta1.Msg.UpdateGroupMembers:output_type -> cosmos.group.v1beta1.MsgUpdateGroupMembersResponse
	6,  // 21: cosmos.group.v1beta1.Msg.UpdateGroupAdmin:output_type -> cosmos.group.v1beta1.MsgUpdateGroupAdminResponse
	8,  // 22: cosmos.group.v1beta1.Msg.UpdateGroupMetadata:output_type -> cosmos.group.v1beta1.MsgUpdateGroupMetadataResponse
	10, // 23: cosmos.group.v1beta1.Msg.CreateGroupPolicy:output_type -> cosmos.group.v1beta1.MsgCreateGroupPolicyResponse
	12, // 24: cosmos.group.v1beta1.Msg.UpdateGroupPolicyAdmin:output_type -> cosmos.group.v1beta1.MsgUpdateGroupPolicyAdminResponse
	14, // 25: cosmos.group.v1beta1.Msg.UpdateGroupPolicyDecisionPolicy:output_type -> cosmos.group.v1beta1.MsgUpdateGroupPolicyDecisionPolicyResponse
	16, // 26: cosmos.group.v1beta1.Msg.UpdateGroupPolicyMetadata:output_type -> cosmos.group.v1beta1.MsgUpdateGroupPolicyMetadataResponse
	18, // 27: cosmos.group.v1beta1.Msg.CreateProposal:output_type -> cosmos.group.v1beta1.MsgCreateProposalResponse
	20, // 28: cosmos.group.v1beta1.Msg.Vote:output_type -> cosmos.group.v1beta1.MsgVoteResponse
	22, // 29: cosmos.group.v1beta1.Msg.Exec:output_type -> cosmos.group.v1beta1.MsgExecResponse
	19, // [19:30] is the sub-list for method output_type
	8,  // [8:19] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_cosmos_group_v1beta1_tx_proto_init() }
func file_cosmos_group_v1beta1_tx_proto_init() {
	if File_cosmos_group_v1beta1_tx_proto != nil {
		return
	}
	file_cosmos_group_v1beta1_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_cosmos_group_v1beta1_tx_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgCreateGroup); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_group_v1beta1_tx_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgCreateGroupResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_group_v1beta1_tx_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgUpdateGroupMembers); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_group_v1beta1_tx_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgUpdateGroupMembersResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_group_v1beta1_tx_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgUpdateGroupAdmin); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_group_v1beta1_tx_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgUpdateGroupAdminResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_group_v1beta1_tx_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgUpdateGroupMetadata); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_group_v1beta1_tx_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgUpdateGroupMetadataResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_group_v1beta1_tx_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgCreateGroupPolicy); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_group_v1beta1_tx_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgCreateGroupPolicyResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_group_v1beta1_tx_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgUpdateGroupPolicyAdmin); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_group_v1beta1_tx_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgUpdateGroupPolicyAdminResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_group_v1beta1_tx_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgUpdateGroupPolicyDecisionPolicy); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_group_v1beta1_tx_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgUpdateGroupPolicyDecisionPolicyResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_group_v1beta1_tx_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgUpdateGroupPolicyMetadata); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_group_v1beta1_tx_proto_msgTypes[15].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgUpdateGroupPolicyMetadataResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_group_v1beta1_tx_proto_msgTypes[16].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgCreateProposal); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_group_v1beta1_tx_proto_msgTypes[17].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgCreateProposalResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_group_v1beta1_tx_proto_msgTypes[18].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgVote); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_group_v1beta1_tx_proto_msgTypes[19].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgVoteResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_group_v1beta1_tx_proto_msgTypes[20].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgExec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_group_v1beta1_tx_proto_msgTypes[21].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgExecResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cosmos_group_v1beta1_tx_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   22,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cosmos_group_v1beta1_tx_proto_goTypes,
		DependencyIndexes: file_cosmos_group_v1beta1_tx_proto_depIdxs,
		EnumInfos:         file_cosmos_group_v1beta1_tx_proto_enumTypes,
		MessageInfos:      file_cosmos_group_v1beta1_tx_proto_msgTypes,
	}.Build()
	File_cosmos_group_v1beta1_tx_proto = out.File
	file_cosmos_group_v1beta1_tx_proto_rawDesc = nil
	file_cosmos_group_v1beta1_tx_proto_goTypes = nil
	file_cosmos_group_v1beta1_tx_proto_depIdxs = nil
}
