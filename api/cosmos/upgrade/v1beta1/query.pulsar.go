package upgradev1beta1

import (
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_QueryCurrentPlanRequest protoreflect.MessageDescriptor
)

func init() {
	file_cosmos_upgrade_v1beta1_query_proto_init()
	md_QueryCurrentPlanRequest = File_cosmos_upgrade_v1beta1_query_proto.Messages().ByName("QueryCurrentPlanRequest")
}

var _ protoreflect.Message = (*fastReflection_QueryCurrentPlanRequest)(nil)

type fastReflection_QueryCurrentPlanRequest QueryCurrentPlanRequest

func (x *QueryCurrentPlanRequest) ProtoReflect() protoreflect.Message {
	return (*fastReflection_QueryCurrentPlanRequest)(x)
}

func (x *QueryCurrentPlanRequest) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_upgrade_v1beta1_query_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_QueryCurrentPlanRequest_messageType fastReflection_QueryCurrentPlanRequest_messageType
var _ protoreflect.MessageType = fastReflection_QueryCurrentPlanRequest_messageType{}

type fastReflection_QueryCurrentPlanRequest_messageType struct{}

func (x fastReflection_QueryCurrentPlanRequest_messageType) Zero() protoreflect.Message {
	return (*fastReflection_QueryCurrentPlanRequest)(nil)
}
func (x fastReflection_QueryCurrentPlanRequest_messageType) New() protoreflect.Message {
	return new(fastReflection_QueryCurrentPlanRequest)
}
func (x fastReflection_QueryCurrentPlanRequest_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryCurrentPlanRequest
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_QueryCurrentPlanRequest) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryCurrentPlanRequest
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_QueryCurrentPlanRequest) Type() protoreflect.MessageType {
	return _fastReflection_QueryCurrentPlanRequest_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_QueryCurrentPlanRequest) New() protoreflect.Message {
	return new(fastReflection_QueryCurrentPlanRequest)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_QueryCurrentPlanRequest) Interface() protoreflect.ProtoMessage {
	return (*QueryCurrentPlanRequest)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_QueryCurrentPlanRequest) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
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
func (x *fastReflection_QueryCurrentPlanRequest) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.upgrade.v1beta1.QueryCurrentPlanRequest"))
		}
		panic(fmt.Errorf("message cosmos.upgrade.v1beta1.QueryCurrentPlanRequest does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryCurrentPlanRequest) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.upgrade.v1beta1.QueryCurrentPlanRequest"))
		}
		panic(fmt.Errorf("message cosmos.upgrade.v1beta1.QueryCurrentPlanRequest does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_QueryCurrentPlanRequest) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.upgrade.v1beta1.QueryCurrentPlanRequest"))
		}
		panic(fmt.Errorf("message cosmos.upgrade.v1beta1.QueryCurrentPlanRequest does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_QueryCurrentPlanRequest) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.upgrade.v1beta1.QueryCurrentPlanRequest"))
		}
		panic(fmt.Errorf("message cosmos.upgrade.v1beta1.QueryCurrentPlanRequest does not contain field %s", fd.FullName()))
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
func (x *fastReflection_QueryCurrentPlanRequest) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.upgrade.v1beta1.QueryCurrentPlanRequest"))
		}
		panic(fmt.Errorf("message cosmos.upgrade.v1beta1.QueryCurrentPlanRequest does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_QueryCurrentPlanRequest) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.upgrade.v1beta1.QueryCurrentPlanRequest"))
		}
		panic(fmt.Errorf("message cosmos.upgrade.v1beta1.QueryCurrentPlanRequest does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_QueryCurrentPlanRequest) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.upgrade.v1beta1.QueryCurrentPlanRequest", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_QueryCurrentPlanRequest) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryCurrentPlanRequest) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_QueryCurrentPlanRequest) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_QueryCurrentPlanRequest) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*QueryCurrentPlanRequest)
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
		x := input.Message.Interface().(*QueryCurrentPlanRequest)
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
		x := input.Message.Interface().(*QueryCurrentPlanRequest)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryCurrentPlanRequest: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryCurrentPlanRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
	md_QueryCurrentPlanResponse      protoreflect.MessageDescriptor
	fd_QueryCurrentPlanResponse_plan protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_upgrade_v1beta1_query_proto_init()
	md_QueryCurrentPlanResponse = File_cosmos_upgrade_v1beta1_query_proto.Messages().ByName("QueryCurrentPlanResponse")
	fd_QueryCurrentPlanResponse_plan = md_QueryCurrentPlanResponse.Fields().ByName("plan")
}

var _ protoreflect.Message = (*fastReflection_QueryCurrentPlanResponse)(nil)

type fastReflection_QueryCurrentPlanResponse QueryCurrentPlanResponse

func (x *QueryCurrentPlanResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_QueryCurrentPlanResponse)(x)
}

func (x *QueryCurrentPlanResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_upgrade_v1beta1_query_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_QueryCurrentPlanResponse_messageType fastReflection_QueryCurrentPlanResponse_messageType
var _ protoreflect.MessageType = fastReflection_QueryCurrentPlanResponse_messageType{}

type fastReflection_QueryCurrentPlanResponse_messageType struct{}

func (x fastReflection_QueryCurrentPlanResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_QueryCurrentPlanResponse)(nil)
}
func (x fastReflection_QueryCurrentPlanResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_QueryCurrentPlanResponse)
}
func (x fastReflection_QueryCurrentPlanResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryCurrentPlanResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_QueryCurrentPlanResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryCurrentPlanResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_QueryCurrentPlanResponse) Type() protoreflect.MessageType {
	return _fastReflection_QueryCurrentPlanResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_QueryCurrentPlanResponse) New() protoreflect.Message {
	return new(fastReflection_QueryCurrentPlanResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_QueryCurrentPlanResponse) Interface() protoreflect.ProtoMessage {
	return (*QueryCurrentPlanResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_QueryCurrentPlanResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Plan != nil {
		value := protoreflect.ValueOfMessage(x.Plan.ProtoReflect())
		if !f(fd_QueryCurrentPlanResponse_plan, value) {
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
func (x *fastReflection_QueryCurrentPlanResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.upgrade.v1beta1.QueryCurrentPlanResponse.plan":
		return x.Plan != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.upgrade.v1beta1.QueryCurrentPlanResponse"))
		}
		panic(fmt.Errorf("message cosmos.upgrade.v1beta1.QueryCurrentPlanResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryCurrentPlanResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.upgrade.v1beta1.QueryCurrentPlanResponse.plan":
		x.Plan = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.upgrade.v1beta1.QueryCurrentPlanResponse"))
		}
		panic(fmt.Errorf("message cosmos.upgrade.v1beta1.QueryCurrentPlanResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_QueryCurrentPlanResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.upgrade.v1beta1.QueryCurrentPlanResponse.plan":
		value := x.Plan
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.upgrade.v1beta1.QueryCurrentPlanResponse"))
		}
		panic(fmt.Errorf("message cosmos.upgrade.v1beta1.QueryCurrentPlanResponse does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_QueryCurrentPlanResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.upgrade.v1beta1.QueryCurrentPlanResponse.plan":
		x.Plan = value.Message().Interface().(*Plan)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.upgrade.v1beta1.QueryCurrentPlanResponse"))
		}
		panic(fmt.Errorf("message cosmos.upgrade.v1beta1.QueryCurrentPlanResponse does not contain field %s", fd.FullName()))
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
func (x *fastReflection_QueryCurrentPlanResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.upgrade.v1beta1.QueryCurrentPlanResponse.plan":
		if x.Plan == nil {
			x.Plan = new(Plan)
		}
		return protoreflect.ValueOfMessage(x.Plan.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.upgrade.v1beta1.QueryCurrentPlanResponse"))
		}
		panic(fmt.Errorf("message cosmos.upgrade.v1beta1.QueryCurrentPlanResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_QueryCurrentPlanResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.upgrade.v1beta1.QueryCurrentPlanResponse.plan":
		m := new(Plan)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.upgrade.v1beta1.QueryCurrentPlanResponse"))
		}
		panic(fmt.Errorf("message cosmos.upgrade.v1beta1.QueryCurrentPlanResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_QueryCurrentPlanResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.upgrade.v1beta1.QueryCurrentPlanResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_QueryCurrentPlanResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryCurrentPlanResponse) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_QueryCurrentPlanResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_QueryCurrentPlanResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*QueryCurrentPlanResponse)
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
		if x.Plan != nil {
			l = options.Size(x.Plan)
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
		x := input.Message.Interface().(*QueryCurrentPlanResponse)
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
		if x.Plan != nil {
			encoded, err := options.Marshal(x.Plan)
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
		x := input.Message.Interface().(*QueryCurrentPlanResponse)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryCurrentPlanResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryCurrentPlanResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Plan", wireType)
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
				if x.Plan == nil {
					x.Plan = &Plan{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Plan); err != nil {
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
	md_QueryAppliedPlanRequest      protoreflect.MessageDescriptor
	fd_QueryAppliedPlanRequest_name protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_upgrade_v1beta1_query_proto_init()
	md_QueryAppliedPlanRequest = File_cosmos_upgrade_v1beta1_query_proto.Messages().ByName("QueryAppliedPlanRequest")
	fd_QueryAppliedPlanRequest_name = md_QueryAppliedPlanRequest.Fields().ByName("name")
}

var _ protoreflect.Message = (*fastReflection_QueryAppliedPlanRequest)(nil)

type fastReflection_QueryAppliedPlanRequest QueryAppliedPlanRequest

func (x *QueryAppliedPlanRequest) ProtoReflect() protoreflect.Message {
	return (*fastReflection_QueryAppliedPlanRequest)(x)
}

func (x *QueryAppliedPlanRequest) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_upgrade_v1beta1_query_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_QueryAppliedPlanRequest_messageType fastReflection_QueryAppliedPlanRequest_messageType
var _ protoreflect.MessageType = fastReflection_QueryAppliedPlanRequest_messageType{}

type fastReflection_QueryAppliedPlanRequest_messageType struct{}

func (x fastReflection_QueryAppliedPlanRequest_messageType) Zero() protoreflect.Message {
	return (*fastReflection_QueryAppliedPlanRequest)(nil)
}
func (x fastReflection_QueryAppliedPlanRequest_messageType) New() protoreflect.Message {
	return new(fastReflection_QueryAppliedPlanRequest)
}
func (x fastReflection_QueryAppliedPlanRequest_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryAppliedPlanRequest
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_QueryAppliedPlanRequest) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryAppliedPlanRequest
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_QueryAppliedPlanRequest) Type() protoreflect.MessageType {
	return _fastReflection_QueryAppliedPlanRequest_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_QueryAppliedPlanRequest) New() protoreflect.Message {
	return new(fastReflection_QueryAppliedPlanRequest)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_QueryAppliedPlanRequest) Interface() protoreflect.ProtoMessage {
	return (*QueryAppliedPlanRequest)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_QueryAppliedPlanRequest) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Name != "" {
		value := protoreflect.ValueOfString(x.Name)
		if !f(fd_QueryAppliedPlanRequest_name, value) {
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
func (x *fastReflection_QueryAppliedPlanRequest) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.upgrade.v1beta1.QueryAppliedPlanRequest.name":
		return x.Name != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.upgrade.v1beta1.QueryAppliedPlanRequest"))
		}
		panic(fmt.Errorf("message cosmos.upgrade.v1beta1.QueryAppliedPlanRequest does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryAppliedPlanRequest) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.upgrade.v1beta1.QueryAppliedPlanRequest.name":
		x.Name = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.upgrade.v1beta1.QueryAppliedPlanRequest"))
		}
		panic(fmt.Errorf("message cosmos.upgrade.v1beta1.QueryAppliedPlanRequest does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_QueryAppliedPlanRequest) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.upgrade.v1beta1.QueryAppliedPlanRequest.name":
		value := x.Name
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.upgrade.v1beta1.QueryAppliedPlanRequest"))
		}
		panic(fmt.Errorf("message cosmos.upgrade.v1beta1.QueryAppliedPlanRequest does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_QueryAppliedPlanRequest) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.upgrade.v1beta1.QueryAppliedPlanRequest.name":
		x.Name = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.upgrade.v1beta1.QueryAppliedPlanRequest"))
		}
		panic(fmt.Errorf("message cosmos.upgrade.v1beta1.QueryAppliedPlanRequest does not contain field %s", fd.FullName()))
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
func (x *fastReflection_QueryAppliedPlanRequest) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.upgrade.v1beta1.QueryAppliedPlanRequest.name":
		panic(fmt.Errorf("field name of message cosmos.upgrade.v1beta1.QueryAppliedPlanRequest is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.upgrade.v1beta1.QueryAppliedPlanRequest"))
		}
		panic(fmt.Errorf("message cosmos.upgrade.v1beta1.QueryAppliedPlanRequest does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_QueryAppliedPlanRequest) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.upgrade.v1beta1.QueryAppliedPlanRequest.name":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.upgrade.v1beta1.QueryAppliedPlanRequest"))
		}
		panic(fmt.Errorf("message cosmos.upgrade.v1beta1.QueryAppliedPlanRequest does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_QueryAppliedPlanRequest) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.upgrade.v1beta1.QueryAppliedPlanRequest", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_QueryAppliedPlanRequest) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryAppliedPlanRequest) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_QueryAppliedPlanRequest) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_QueryAppliedPlanRequest) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*QueryAppliedPlanRequest)
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
		l = len(x.Name)
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
		x := input.Message.Interface().(*QueryAppliedPlanRequest)
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
		if len(x.Name) > 0 {
			i -= len(x.Name)
			copy(dAtA[i:], x.Name)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Name)))
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
		x := input.Message.Interface().(*QueryAppliedPlanRequest)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryAppliedPlanRequest: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryAppliedPlanRequest: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
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
				x.Name = string(dAtA[iNdEx:postIndex])
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
	md_QueryAppliedPlanResponse        protoreflect.MessageDescriptor
	fd_QueryAppliedPlanResponse_height protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_upgrade_v1beta1_query_proto_init()
	md_QueryAppliedPlanResponse = File_cosmos_upgrade_v1beta1_query_proto.Messages().ByName("QueryAppliedPlanResponse")
	fd_QueryAppliedPlanResponse_height = md_QueryAppliedPlanResponse.Fields().ByName("height")
}

var _ protoreflect.Message = (*fastReflection_QueryAppliedPlanResponse)(nil)

type fastReflection_QueryAppliedPlanResponse QueryAppliedPlanResponse

func (x *QueryAppliedPlanResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_QueryAppliedPlanResponse)(x)
}

func (x *QueryAppliedPlanResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_upgrade_v1beta1_query_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_QueryAppliedPlanResponse_messageType fastReflection_QueryAppliedPlanResponse_messageType
var _ protoreflect.MessageType = fastReflection_QueryAppliedPlanResponse_messageType{}

type fastReflection_QueryAppliedPlanResponse_messageType struct{}

func (x fastReflection_QueryAppliedPlanResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_QueryAppliedPlanResponse)(nil)
}
func (x fastReflection_QueryAppliedPlanResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_QueryAppliedPlanResponse)
}
func (x fastReflection_QueryAppliedPlanResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryAppliedPlanResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_QueryAppliedPlanResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryAppliedPlanResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_QueryAppliedPlanResponse) Type() protoreflect.MessageType {
	return _fastReflection_QueryAppliedPlanResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_QueryAppliedPlanResponse) New() protoreflect.Message {
	return new(fastReflection_QueryAppliedPlanResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_QueryAppliedPlanResponse) Interface() protoreflect.ProtoMessage {
	return (*QueryAppliedPlanResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_QueryAppliedPlanResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Height != int64(0) {
		value := protoreflect.ValueOfInt64(x.Height)
		if !f(fd_QueryAppliedPlanResponse_height, value) {
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
func (x *fastReflection_QueryAppliedPlanResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.upgrade.v1beta1.QueryAppliedPlanResponse.height":
		return x.Height != int64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.upgrade.v1beta1.QueryAppliedPlanResponse"))
		}
		panic(fmt.Errorf("message cosmos.upgrade.v1beta1.QueryAppliedPlanResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryAppliedPlanResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.upgrade.v1beta1.QueryAppliedPlanResponse.height":
		x.Height = int64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.upgrade.v1beta1.QueryAppliedPlanResponse"))
		}
		panic(fmt.Errorf("message cosmos.upgrade.v1beta1.QueryAppliedPlanResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_QueryAppliedPlanResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.upgrade.v1beta1.QueryAppliedPlanResponse.height":
		value := x.Height
		return protoreflect.ValueOfInt64(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.upgrade.v1beta1.QueryAppliedPlanResponse"))
		}
		panic(fmt.Errorf("message cosmos.upgrade.v1beta1.QueryAppliedPlanResponse does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_QueryAppliedPlanResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.upgrade.v1beta1.QueryAppliedPlanResponse.height":
		x.Height = value.Int()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.upgrade.v1beta1.QueryAppliedPlanResponse"))
		}
		panic(fmt.Errorf("message cosmos.upgrade.v1beta1.QueryAppliedPlanResponse does not contain field %s", fd.FullName()))
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
func (x *fastReflection_QueryAppliedPlanResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.upgrade.v1beta1.QueryAppliedPlanResponse.height":
		panic(fmt.Errorf("field height of message cosmos.upgrade.v1beta1.QueryAppliedPlanResponse is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.upgrade.v1beta1.QueryAppliedPlanResponse"))
		}
		panic(fmt.Errorf("message cosmos.upgrade.v1beta1.QueryAppliedPlanResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_QueryAppliedPlanResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.upgrade.v1beta1.QueryAppliedPlanResponse.height":
		return protoreflect.ValueOfInt64(int64(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.upgrade.v1beta1.QueryAppliedPlanResponse"))
		}
		panic(fmt.Errorf("message cosmos.upgrade.v1beta1.QueryAppliedPlanResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_QueryAppliedPlanResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.upgrade.v1beta1.QueryAppliedPlanResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_QueryAppliedPlanResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryAppliedPlanResponse) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_QueryAppliedPlanResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_QueryAppliedPlanResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*QueryAppliedPlanResponse)
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
		if x.Height != 0 {
			n += 1 + runtime.Sov(uint64(x.Height))
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
		x := input.Message.Interface().(*QueryAppliedPlanResponse)
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
		if x.Height != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Height))
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
		x := input.Message.Interface().(*QueryAppliedPlanResponse)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryAppliedPlanResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryAppliedPlanResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
				}
				x.Height = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Height |= int64(b&0x7F) << shift
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
	md_QueryUpgradedConsensusStateRequest             protoreflect.MessageDescriptor
	fd_QueryUpgradedConsensusStateRequest_last_height protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_upgrade_v1beta1_query_proto_init()
	md_QueryUpgradedConsensusStateRequest = File_cosmos_upgrade_v1beta1_query_proto.Messages().ByName("QueryUpgradedConsensusStateRequest")
	fd_QueryUpgradedConsensusStateRequest_last_height = md_QueryUpgradedConsensusStateRequest.Fields().ByName("last_height")
}

var _ protoreflect.Message = (*fastReflection_QueryUpgradedConsensusStateRequest)(nil)

type fastReflection_QueryUpgradedConsensusStateRequest QueryUpgradedConsensusStateRequest

func (x *QueryUpgradedConsensusStateRequest) ProtoReflect() protoreflect.Message {
	return (*fastReflection_QueryUpgradedConsensusStateRequest)(x)
}

func (x *QueryUpgradedConsensusStateRequest) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_upgrade_v1beta1_query_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_QueryUpgradedConsensusStateRequest_messageType fastReflection_QueryUpgradedConsensusStateRequest_messageType
var _ protoreflect.MessageType = fastReflection_QueryUpgradedConsensusStateRequest_messageType{}

type fastReflection_QueryUpgradedConsensusStateRequest_messageType struct{}

func (x fastReflection_QueryUpgradedConsensusStateRequest_messageType) Zero() protoreflect.Message {
	return (*fastReflection_QueryUpgradedConsensusStateRequest)(nil)
}
func (x fastReflection_QueryUpgradedConsensusStateRequest_messageType) New() protoreflect.Message {
	return new(fastReflection_QueryUpgradedConsensusStateRequest)
}
func (x fastReflection_QueryUpgradedConsensusStateRequest_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryUpgradedConsensusStateRequest
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_QueryUpgradedConsensusStateRequest) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryUpgradedConsensusStateRequest
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_QueryUpgradedConsensusStateRequest) Type() protoreflect.MessageType {
	return _fastReflection_QueryUpgradedConsensusStateRequest_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_QueryUpgradedConsensusStateRequest) New() protoreflect.Message {
	return new(fastReflection_QueryUpgradedConsensusStateRequest)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_QueryUpgradedConsensusStateRequest) Interface() protoreflect.ProtoMessage {
	return (*QueryUpgradedConsensusStateRequest)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_QueryUpgradedConsensusStateRequest) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.LastHeight != int64(0) {
		value := protoreflect.ValueOfInt64(x.LastHeight)
		if !f(fd_QueryUpgradedConsensusStateRequest_last_height, value) {
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
func (x *fastReflection_QueryUpgradedConsensusStateRequest) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.upgrade.v1beta1.QueryUpgradedConsensusStateRequest.last_height":
		return x.LastHeight != int64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.upgrade.v1beta1.QueryUpgradedConsensusStateRequest"))
		}
		panic(fmt.Errorf("message cosmos.upgrade.v1beta1.QueryUpgradedConsensusStateRequest does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryUpgradedConsensusStateRequest) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.upgrade.v1beta1.QueryUpgradedConsensusStateRequest.last_height":
		x.LastHeight = int64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.upgrade.v1beta1.QueryUpgradedConsensusStateRequest"))
		}
		panic(fmt.Errorf("message cosmos.upgrade.v1beta1.QueryUpgradedConsensusStateRequest does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_QueryUpgradedConsensusStateRequest) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.upgrade.v1beta1.QueryUpgradedConsensusStateRequest.last_height":
		value := x.LastHeight
		return protoreflect.ValueOfInt64(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.upgrade.v1beta1.QueryUpgradedConsensusStateRequest"))
		}
		panic(fmt.Errorf("message cosmos.upgrade.v1beta1.QueryUpgradedConsensusStateRequest does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_QueryUpgradedConsensusStateRequest) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.upgrade.v1beta1.QueryUpgradedConsensusStateRequest.last_height":
		x.LastHeight = value.Int()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.upgrade.v1beta1.QueryUpgradedConsensusStateRequest"))
		}
		panic(fmt.Errorf("message cosmos.upgrade.v1beta1.QueryUpgradedConsensusStateRequest does not contain field %s", fd.FullName()))
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
func (x *fastReflection_QueryUpgradedConsensusStateRequest) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.upgrade.v1beta1.QueryUpgradedConsensusStateRequest.last_height":
		panic(fmt.Errorf("field last_height of message cosmos.upgrade.v1beta1.QueryUpgradedConsensusStateRequest is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.upgrade.v1beta1.QueryUpgradedConsensusStateRequest"))
		}
		panic(fmt.Errorf("message cosmos.upgrade.v1beta1.QueryUpgradedConsensusStateRequest does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_QueryUpgradedConsensusStateRequest) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.upgrade.v1beta1.QueryUpgradedConsensusStateRequest.last_height":
		return protoreflect.ValueOfInt64(int64(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.upgrade.v1beta1.QueryUpgradedConsensusStateRequest"))
		}
		panic(fmt.Errorf("message cosmos.upgrade.v1beta1.QueryUpgradedConsensusStateRequest does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_QueryUpgradedConsensusStateRequest) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.upgrade.v1beta1.QueryUpgradedConsensusStateRequest", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_QueryUpgradedConsensusStateRequest) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryUpgradedConsensusStateRequest) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_QueryUpgradedConsensusStateRequest) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_QueryUpgradedConsensusStateRequest) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*QueryUpgradedConsensusStateRequest)
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
		if x.LastHeight != 0 {
			n += 1 + runtime.Sov(uint64(x.LastHeight))
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
		x := input.Message.Interface().(*QueryUpgradedConsensusStateRequest)
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
		if x.LastHeight != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.LastHeight))
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
		x := input.Message.Interface().(*QueryUpgradedConsensusStateRequest)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryUpgradedConsensusStateRequest: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryUpgradedConsensusStateRequest: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field LastHeight", wireType)
				}
				x.LastHeight = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.LastHeight |= int64(b&0x7F) << shift
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
	md_QueryUpgradedConsensusStateResponse                          protoreflect.MessageDescriptor
	fd_QueryUpgradedConsensusStateResponse_upgraded_consensus_state protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_upgrade_v1beta1_query_proto_init()
	md_QueryUpgradedConsensusStateResponse = File_cosmos_upgrade_v1beta1_query_proto.Messages().ByName("QueryUpgradedConsensusStateResponse")
	fd_QueryUpgradedConsensusStateResponse_upgraded_consensus_state = md_QueryUpgradedConsensusStateResponse.Fields().ByName("upgraded_consensus_state")
}

var _ protoreflect.Message = (*fastReflection_QueryUpgradedConsensusStateResponse)(nil)

type fastReflection_QueryUpgradedConsensusStateResponse QueryUpgradedConsensusStateResponse

func (x *QueryUpgradedConsensusStateResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_QueryUpgradedConsensusStateResponse)(x)
}

func (x *QueryUpgradedConsensusStateResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_upgrade_v1beta1_query_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_QueryUpgradedConsensusStateResponse_messageType fastReflection_QueryUpgradedConsensusStateResponse_messageType
var _ protoreflect.MessageType = fastReflection_QueryUpgradedConsensusStateResponse_messageType{}

type fastReflection_QueryUpgradedConsensusStateResponse_messageType struct{}

func (x fastReflection_QueryUpgradedConsensusStateResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_QueryUpgradedConsensusStateResponse)(nil)
}
func (x fastReflection_QueryUpgradedConsensusStateResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_QueryUpgradedConsensusStateResponse)
}
func (x fastReflection_QueryUpgradedConsensusStateResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryUpgradedConsensusStateResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_QueryUpgradedConsensusStateResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryUpgradedConsensusStateResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_QueryUpgradedConsensusStateResponse) Type() protoreflect.MessageType {
	return _fastReflection_QueryUpgradedConsensusStateResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_QueryUpgradedConsensusStateResponse) New() protoreflect.Message {
	return new(fastReflection_QueryUpgradedConsensusStateResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_QueryUpgradedConsensusStateResponse) Interface() protoreflect.ProtoMessage {
	return (*QueryUpgradedConsensusStateResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_QueryUpgradedConsensusStateResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.UpgradedConsensusState) != 0 {
		value := protoreflect.ValueOfBytes(x.UpgradedConsensusState)
		if !f(fd_QueryUpgradedConsensusStateResponse_upgraded_consensus_state, value) {
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
func (x *fastReflection_QueryUpgradedConsensusStateResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.upgrade.v1beta1.QueryUpgradedConsensusStateResponse.upgraded_consensus_state":
		return len(x.UpgradedConsensusState) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.upgrade.v1beta1.QueryUpgradedConsensusStateResponse"))
		}
		panic(fmt.Errorf("message cosmos.upgrade.v1beta1.QueryUpgradedConsensusStateResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryUpgradedConsensusStateResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.upgrade.v1beta1.QueryUpgradedConsensusStateResponse.upgraded_consensus_state":
		x.UpgradedConsensusState = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.upgrade.v1beta1.QueryUpgradedConsensusStateResponse"))
		}
		panic(fmt.Errorf("message cosmos.upgrade.v1beta1.QueryUpgradedConsensusStateResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_QueryUpgradedConsensusStateResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.upgrade.v1beta1.QueryUpgradedConsensusStateResponse.upgraded_consensus_state":
		value := x.UpgradedConsensusState
		return protoreflect.ValueOfBytes(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.upgrade.v1beta1.QueryUpgradedConsensusStateResponse"))
		}
		panic(fmt.Errorf("message cosmos.upgrade.v1beta1.QueryUpgradedConsensusStateResponse does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_QueryUpgradedConsensusStateResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.upgrade.v1beta1.QueryUpgradedConsensusStateResponse.upgraded_consensus_state":
		x.UpgradedConsensusState = value.Bytes()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.upgrade.v1beta1.QueryUpgradedConsensusStateResponse"))
		}
		panic(fmt.Errorf("message cosmos.upgrade.v1beta1.QueryUpgradedConsensusStateResponse does not contain field %s", fd.FullName()))
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
func (x *fastReflection_QueryUpgradedConsensusStateResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.upgrade.v1beta1.QueryUpgradedConsensusStateResponse.upgraded_consensus_state":
		panic(fmt.Errorf("field upgraded_consensus_state of message cosmos.upgrade.v1beta1.QueryUpgradedConsensusStateResponse is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.upgrade.v1beta1.QueryUpgradedConsensusStateResponse"))
		}
		panic(fmt.Errorf("message cosmos.upgrade.v1beta1.QueryUpgradedConsensusStateResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_QueryUpgradedConsensusStateResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.upgrade.v1beta1.QueryUpgradedConsensusStateResponse.upgraded_consensus_state":
		return protoreflect.ValueOfBytes(nil)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.upgrade.v1beta1.QueryUpgradedConsensusStateResponse"))
		}
		panic(fmt.Errorf("message cosmos.upgrade.v1beta1.QueryUpgradedConsensusStateResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_QueryUpgradedConsensusStateResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.upgrade.v1beta1.QueryUpgradedConsensusStateResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_QueryUpgradedConsensusStateResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryUpgradedConsensusStateResponse) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_QueryUpgradedConsensusStateResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_QueryUpgradedConsensusStateResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*QueryUpgradedConsensusStateResponse)
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
		l = len(x.UpgradedConsensusState)
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
		x := input.Message.Interface().(*QueryUpgradedConsensusStateResponse)
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
		if len(x.UpgradedConsensusState) > 0 {
			i -= len(x.UpgradedConsensusState)
			copy(dAtA[i:], x.UpgradedConsensusState)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.UpgradedConsensusState)))
			i--
			dAtA[i] = 0x12
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
		x := input.Message.Interface().(*QueryUpgradedConsensusStateResponse)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryUpgradedConsensusStateResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryUpgradedConsensusStateResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field UpgradedConsensusState", wireType)
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
				x.UpgradedConsensusState = append(x.UpgradedConsensusState[:0], dAtA[iNdEx:postIndex]...)
				if x.UpgradedConsensusState == nil {
					x.UpgradedConsensusState = []byte{}
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
	md_QueryModuleVersionsRequest             protoreflect.MessageDescriptor
	fd_QueryModuleVersionsRequest_module_name protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_upgrade_v1beta1_query_proto_init()
	md_QueryModuleVersionsRequest = File_cosmos_upgrade_v1beta1_query_proto.Messages().ByName("QueryModuleVersionsRequest")
	fd_QueryModuleVersionsRequest_module_name = md_QueryModuleVersionsRequest.Fields().ByName("module_name")
}

var _ protoreflect.Message = (*fastReflection_QueryModuleVersionsRequest)(nil)

type fastReflection_QueryModuleVersionsRequest QueryModuleVersionsRequest

func (x *QueryModuleVersionsRequest) ProtoReflect() protoreflect.Message {
	return (*fastReflection_QueryModuleVersionsRequest)(x)
}

func (x *QueryModuleVersionsRequest) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_upgrade_v1beta1_query_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_QueryModuleVersionsRequest_messageType fastReflection_QueryModuleVersionsRequest_messageType
var _ protoreflect.MessageType = fastReflection_QueryModuleVersionsRequest_messageType{}

type fastReflection_QueryModuleVersionsRequest_messageType struct{}

func (x fastReflection_QueryModuleVersionsRequest_messageType) Zero() protoreflect.Message {
	return (*fastReflection_QueryModuleVersionsRequest)(nil)
}
func (x fastReflection_QueryModuleVersionsRequest_messageType) New() protoreflect.Message {
	return new(fastReflection_QueryModuleVersionsRequest)
}
func (x fastReflection_QueryModuleVersionsRequest_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryModuleVersionsRequest
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_QueryModuleVersionsRequest) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryModuleVersionsRequest
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_QueryModuleVersionsRequest) Type() protoreflect.MessageType {
	return _fastReflection_QueryModuleVersionsRequest_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_QueryModuleVersionsRequest) New() protoreflect.Message {
	return new(fastReflection_QueryModuleVersionsRequest)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_QueryModuleVersionsRequest) Interface() protoreflect.ProtoMessage {
	return (*QueryModuleVersionsRequest)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_QueryModuleVersionsRequest) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.ModuleName != "" {
		value := protoreflect.ValueOfString(x.ModuleName)
		if !f(fd_QueryModuleVersionsRequest_module_name, value) {
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
func (x *fastReflection_QueryModuleVersionsRequest) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.upgrade.v1beta1.QueryModuleVersionsRequest.module_name":
		return x.ModuleName != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.upgrade.v1beta1.QueryModuleVersionsRequest"))
		}
		panic(fmt.Errorf("message cosmos.upgrade.v1beta1.QueryModuleVersionsRequest does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryModuleVersionsRequest) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.upgrade.v1beta1.QueryModuleVersionsRequest.module_name":
		x.ModuleName = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.upgrade.v1beta1.QueryModuleVersionsRequest"))
		}
		panic(fmt.Errorf("message cosmos.upgrade.v1beta1.QueryModuleVersionsRequest does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_QueryModuleVersionsRequest) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.upgrade.v1beta1.QueryModuleVersionsRequest.module_name":
		value := x.ModuleName
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.upgrade.v1beta1.QueryModuleVersionsRequest"))
		}
		panic(fmt.Errorf("message cosmos.upgrade.v1beta1.QueryModuleVersionsRequest does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_QueryModuleVersionsRequest) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.upgrade.v1beta1.QueryModuleVersionsRequest.module_name":
		x.ModuleName = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.upgrade.v1beta1.QueryModuleVersionsRequest"))
		}
		panic(fmt.Errorf("message cosmos.upgrade.v1beta1.QueryModuleVersionsRequest does not contain field %s", fd.FullName()))
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
func (x *fastReflection_QueryModuleVersionsRequest) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.upgrade.v1beta1.QueryModuleVersionsRequest.module_name":
		panic(fmt.Errorf("field module_name of message cosmos.upgrade.v1beta1.QueryModuleVersionsRequest is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.upgrade.v1beta1.QueryModuleVersionsRequest"))
		}
		panic(fmt.Errorf("message cosmos.upgrade.v1beta1.QueryModuleVersionsRequest does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_QueryModuleVersionsRequest) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.upgrade.v1beta1.QueryModuleVersionsRequest.module_name":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.upgrade.v1beta1.QueryModuleVersionsRequest"))
		}
		panic(fmt.Errorf("message cosmos.upgrade.v1beta1.QueryModuleVersionsRequest does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_QueryModuleVersionsRequest) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.upgrade.v1beta1.QueryModuleVersionsRequest", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_QueryModuleVersionsRequest) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryModuleVersionsRequest) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_QueryModuleVersionsRequest) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_QueryModuleVersionsRequest) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*QueryModuleVersionsRequest)
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
		l = len(x.ModuleName)
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
		x := input.Message.Interface().(*QueryModuleVersionsRequest)
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
		if len(x.ModuleName) > 0 {
			i -= len(x.ModuleName)
			copy(dAtA[i:], x.ModuleName)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ModuleName)))
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
		x := input.Message.Interface().(*QueryModuleVersionsRequest)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryModuleVersionsRequest: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryModuleVersionsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ModuleName", wireType)
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
				x.ModuleName = string(dAtA[iNdEx:postIndex])
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

var _ protoreflect.List = (*_QueryModuleVersionsResponse_1_list)(nil)

type _QueryModuleVersionsResponse_1_list struct {
	list *[]*ModuleVersion
}

func (x *_QueryModuleVersionsResponse_1_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_QueryModuleVersionsResponse_1_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_QueryModuleVersionsResponse_1_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*ModuleVersion)
	(*x.list)[i] = concreteValue
}

func (x *_QueryModuleVersionsResponse_1_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*ModuleVersion)
	*x.list = append(*x.list, concreteValue)
}

func (x *_QueryModuleVersionsResponse_1_list) AppendMutable() protoreflect.Value {
	v := new(ModuleVersion)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_QueryModuleVersionsResponse_1_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_QueryModuleVersionsResponse_1_list) NewElement() protoreflect.Value {
	v := new(ModuleVersion)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_QueryModuleVersionsResponse_1_list) IsValid() bool {
	return x.list != nil
}

var (
	md_QueryModuleVersionsResponse                 protoreflect.MessageDescriptor
	fd_QueryModuleVersionsResponse_module_versions protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_upgrade_v1beta1_query_proto_init()
	md_QueryModuleVersionsResponse = File_cosmos_upgrade_v1beta1_query_proto.Messages().ByName("QueryModuleVersionsResponse")
	fd_QueryModuleVersionsResponse_module_versions = md_QueryModuleVersionsResponse.Fields().ByName("module_versions")
}

var _ protoreflect.Message = (*fastReflection_QueryModuleVersionsResponse)(nil)

type fastReflection_QueryModuleVersionsResponse QueryModuleVersionsResponse

func (x *QueryModuleVersionsResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_QueryModuleVersionsResponse)(x)
}

func (x *QueryModuleVersionsResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_upgrade_v1beta1_query_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_QueryModuleVersionsResponse_messageType fastReflection_QueryModuleVersionsResponse_messageType
var _ protoreflect.MessageType = fastReflection_QueryModuleVersionsResponse_messageType{}

type fastReflection_QueryModuleVersionsResponse_messageType struct{}

func (x fastReflection_QueryModuleVersionsResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_QueryModuleVersionsResponse)(nil)
}
func (x fastReflection_QueryModuleVersionsResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_QueryModuleVersionsResponse)
}
func (x fastReflection_QueryModuleVersionsResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryModuleVersionsResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_QueryModuleVersionsResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryModuleVersionsResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_QueryModuleVersionsResponse) Type() protoreflect.MessageType {
	return _fastReflection_QueryModuleVersionsResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_QueryModuleVersionsResponse) New() protoreflect.Message {
	return new(fastReflection_QueryModuleVersionsResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_QueryModuleVersionsResponse) Interface() protoreflect.ProtoMessage {
	return (*QueryModuleVersionsResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_QueryModuleVersionsResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.ModuleVersions) != 0 {
		value := protoreflect.ValueOfList(&_QueryModuleVersionsResponse_1_list{list: &x.ModuleVersions})
		if !f(fd_QueryModuleVersionsResponse_module_versions, value) {
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
func (x *fastReflection_QueryModuleVersionsResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.upgrade.v1beta1.QueryModuleVersionsResponse.module_versions":
		return len(x.ModuleVersions) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.upgrade.v1beta1.QueryModuleVersionsResponse"))
		}
		panic(fmt.Errorf("message cosmos.upgrade.v1beta1.QueryModuleVersionsResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryModuleVersionsResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.upgrade.v1beta1.QueryModuleVersionsResponse.module_versions":
		x.ModuleVersions = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.upgrade.v1beta1.QueryModuleVersionsResponse"))
		}
		panic(fmt.Errorf("message cosmos.upgrade.v1beta1.QueryModuleVersionsResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_QueryModuleVersionsResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.upgrade.v1beta1.QueryModuleVersionsResponse.module_versions":
		if len(x.ModuleVersions) == 0 {
			return protoreflect.ValueOfList(&_QueryModuleVersionsResponse_1_list{})
		}
		listValue := &_QueryModuleVersionsResponse_1_list{list: &x.ModuleVersions}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.upgrade.v1beta1.QueryModuleVersionsResponse"))
		}
		panic(fmt.Errorf("message cosmos.upgrade.v1beta1.QueryModuleVersionsResponse does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_QueryModuleVersionsResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.upgrade.v1beta1.QueryModuleVersionsResponse.module_versions":
		lv := value.List()
		clv := lv.(*_QueryModuleVersionsResponse_1_list)
		x.ModuleVersions = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.upgrade.v1beta1.QueryModuleVersionsResponse"))
		}
		panic(fmt.Errorf("message cosmos.upgrade.v1beta1.QueryModuleVersionsResponse does not contain field %s", fd.FullName()))
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
func (x *fastReflection_QueryModuleVersionsResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.upgrade.v1beta1.QueryModuleVersionsResponse.module_versions":
		if x.ModuleVersions == nil {
			x.ModuleVersions = []*ModuleVersion{}
		}
		value := &_QueryModuleVersionsResponse_1_list{list: &x.ModuleVersions}
		return protoreflect.ValueOfList(value)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.upgrade.v1beta1.QueryModuleVersionsResponse"))
		}
		panic(fmt.Errorf("message cosmos.upgrade.v1beta1.QueryModuleVersionsResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_QueryModuleVersionsResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.upgrade.v1beta1.QueryModuleVersionsResponse.module_versions":
		list := []*ModuleVersion{}
		return protoreflect.ValueOfList(&_QueryModuleVersionsResponse_1_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.upgrade.v1beta1.QueryModuleVersionsResponse"))
		}
		panic(fmt.Errorf("message cosmos.upgrade.v1beta1.QueryModuleVersionsResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_QueryModuleVersionsResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.upgrade.v1beta1.QueryModuleVersionsResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_QueryModuleVersionsResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryModuleVersionsResponse) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_QueryModuleVersionsResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_QueryModuleVersionsResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*QueryModuleVersionsResponse)
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
		if len(x.ModuleVersions) > 0 {
			for _, e := range x.ModuleVersions {
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
		x := input.Message.Interface().(*QueryModuleVersionsResponse)
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
		if len(x.ModuleVersions) > 0 {
			for iNdEx := len(x.ModuleVersions) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.ModuleVersions[iNdEx])
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
				dAtA[i] = 0xa
			}
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
		x := input.Message.Interface().(*QueryModuleVersionsResponse)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryModuleVersionsResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryModuleVersionsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ModuleVersions", wireType)
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
				x.ModuleVersions = append(x.ModuleVersions, &ModuleVersion{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.ModuleVersions[len(x.ModuleVersions)-1]); err != nil {
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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: cosmos/upgrade/v1beta1/query.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// QueryCurrentPlanRequest is the request type for the Query/CurrentPlan RPC
// method.
type QueryCurrentPlanRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *QueryCurrentPlanRequest) Reset() {
	*x = QueryCurrentPlanRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_upgrade_v1beta1_query_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryCurrentPlanRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryCurrentPlanRequest) ProtoMessage() {}

// Deprecated: Use QueryCurrentPlanRequest.ProtoReflect.Descriptor instead.
func (*QueryCurrentPlanRequest) Descriptor() ([]byte, []int) {
	return file_cosmos_upgrade_v1beta1_query_proto_rawDescGZIP(), []int{0}
}

// QueryCurrentPlanResponse is the response type for the Query/CurrentPlan RPC
// method.
type QueryCurrentPlanResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// plan is the current upgrade plan.
	Plan *Plan `protobuf:"bytes,1,opt,name=plan,proto3" json:"plan,omitempty"`
}

func (x *QueryCurrentPlanResponse) Reset() {
	*x = QueryCurrentPlanResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_upgrade_v1beta1_query_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryCurrentPlanResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryCurrentPlanResponse) ProtoMessage() {}

// Deprecated: Use QueryCurrentPlanResponse.ProtoReflect.Descriptor instead.
func (*QueryCurrentPlanResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_upgrade_v1beta1_query_proto_rawDescGZIP(), []int{1}
}

func (x *QueryCurrentPlanResponse) GetPlan() *Plan {
	if x != nil {
		return x.Plan
	}
	return nil
}

// QueryCurrentPlanRequest is the request type for the Query/AppliedPlan RPC
// method.
type QueryAppliedPlanRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name is the name of the applied plan to query for.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *QueryAppliedPlanRequest) Reset() {
	*x = QueryAppliedPlanRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_upgrade_v1beta1_query_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryAppliedPlanRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryAppliedPlanRequest) ProtoMessage() {}

// Deprecated: Use QueryAppliedPlanRequest.ProtoReflect.Descriptor instead.
func (*QueryAppliedPlanRequest) Descriptor() ([]byte, []int) {
	return file_cosmos_upgrade_v1beta1_query_proto_rawDescGZIP(), []int{2}
}

func (x *QueryAppliedPlanRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// QueryAppliedPlanResponse is the response type for the Query/AppliedPlan RPC
// method.
type QueryAppliedPlanResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// height is the block height at which the plan was applied.
	Height int64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
}

func (x *QueryAppliedPlanResponse) Reset() {
	*x = QueryAppliedPlanResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_upgrade_v1beta1_query_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryAppliedPlanResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryAppliedPlanResponse) ProtoMessage() {}

// Deprecated: Use QueryAppliedPlanResponse.ProtoReflect.Descriptor instead.
func (*QueryAppliedPlanResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_upgrade_v1beta1_query_proto_rawDescGZIP(), []int{3}
}

func (x *QueryAppliedPlanResponse) GetHeight() int64 {
	if x != nil {
		return x.Height
	}
	return 0
}

// QueryUpgradedConsensusStateRequest is the request type for the Query/UpgradedConsensusState
// RPC method.
//
// Deprecated: Do not use.
type QueryUpgradedConsensusStateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// last height of the current chain must be sent in request
	// as this is the height under which next consensus state is stored
	LastHeight int64 `protobuf:"varint,1,opt,name=last_height,json=lastHeight,proto3" json:"last_height,omitempty"`
}

func (x *QueryUpgradedConsensusStateRequest) Reset() {
	*x = QueryUpgradedConsensusStateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_upgrade_v1beta1_query_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryUpgradedConsensusStateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryUpgradedConsensusStateRequest) ProtoMessage() {}

// Deprecated: Use QueryUpgradedConsensusStateRequest.ProtoReflect.Descriptor instead.
func (*QueryUpgradedConsensusStateRequest) Descriptor() ([]byte, []int) {
	return file_cosmos_upgrade_v1beta1_query_proto_rawDescGZIP(), []int{4}
}

func (x *QueryUpgradedConsensusStateRequest) GetLastHeight() int64 {
	if x != nil {
		return x.LastHeight
	}
	return 0
}

// QueryUpgradedConsensusStateResponse is the response type for the Query/UpgradedConsensusState
// RPC method.
//
// Deprecated: Do not use.
type QueryUpgradedConsensusStateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Since: cosmos-sdk 0.43
	UpgradedConsensusState []byte `protobuf:"bytes,2,opt,name=upgraded_consensus_state,json=upgradedConsensusState,proto3" json:"upgraded_consensus_state,omitempty"`
}

func (x *QueryUpgradedConsensusStateResponse) Reset() {
	*x = QueryUpgradedConsensusStateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_upgrade_v1beta1_query_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryUpgradedConsensusStateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryUpgradedConsensusStateResponse) ProtoMessage() {}

// Deprecated: Use QueryUpgradedConsensusStateResponse.ProtoReflect.Descriptor instead.
func (*QueryUpgradedConsensusStateResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_upgrade_v1beta1_query_proto_rawDescGZIP(), []int{5}
}

func (x *QueryUpgradedConsensusStateResponse) GetUpgradedConsensusState() []byte {
	if x != nil {
		return x.UpgradedConsensusState
	}
	return nil
}

// QueryModuleVersionsRequest is the request type for the Query/ModuleVersions
// RPC method.
//
// Since: cosmos-sdk 0.43
type QueryModuleVersionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// module_name is a field to query a specific module
	// consensus version from state. Leaving this empty will
	// fetch the full list of module versions from state
	ModuleName string `protobuf:"bytes,1,opt,name=module_name,json=moduleName,proto3" json:"module_name,omitempty"`
}

func (x *QueryModuleVersionsRequest) Reset() {
	*x = QueryModuleVersionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_upgrade_v1beta1_query_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryModuleVersionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryModuleVersionsRequest) ProtoMessage() {}

// Deprecated: Use QueryModuleVersionsRequest.ProtoReflect.Descriptor instead.
func (*QueryModuleVersionsRequest) Descriptor() ([]byte, []int) {
	return file_cosmos_upgrade_v1beta1_query_proto_rawDescGZIP(), []int{6}
}

func (x *QueryModuleVersionsRequest) GetModuleName() string {
	if x != nil {
		return x.ModuleName
	}
	return ""
}

// QueryModuleVersionsResponse is the response type for the Query/ModuleVersions
// RPC method.
//
// Since: cosmos-sdk 0.43
type QueryModuleVersionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// module_versions is a list of module names with their consensus versions.
	ModuleVersions []*ModuleVersion `protobuf:"bytes,1,rep,name=module_versions,json=moduleVersions,proto3" json:"module_versions,omitempty"`
}

func (x *QueryModuleVersionsResponse) Reset() {
	*x = QueryModuleVersionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_upgrade_v1beta1_query_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryModuleVersionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryModuleVersionsResponse) ProtoMessage() {}

// Deprecated: Use QueryModuleVersionsResponse.ProtoReflect.Descriptor instead.
func (*QueryModuleVersionsResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_upgrade_v1beta1_query_proto_rawDescGZIP(), []int{7}
}

func (x *QueryModuleVersionsResponse) GetModuleVersions() []*ModuleVersion {
	if x != nil {
		return x.ModuleVersions
	}
	return nil
}

var File_cosmos_upgrade_v1beta1_query_proto protoreflect.FileDescriptor

var file_cosmos_upgrade_v1beta1_query_proto_rawDesc = []byte{
	0x0a, 0x22, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x75, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65,
	0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x75, 0x70, 0x67,
	0x72, 0x61, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2f, 0x75, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2f, 0x75, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x19, 0x0a, 0x17, 0x51, 0x75, 0x65, 0x72, 0x79, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x4c, 0x0a, 0x18, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x04, 0x70, 0x6c, 0x61, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x75,
	0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50,
	0x6c, 0x61, 0x6e, 0x52, 0x04, 0x70, 0x6c, 0x61, 0x6e, 0x22, 0x2d, 0x0a, 0x17, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x32, 0x0a, 0x18, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x49, 0x0a, 0x22,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x64, 0x43, 0x6f, 0x6e,
	0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x48, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x3a, 0x02, 0x18, 0x01, 0x22, 0x69, 0x0a, 0x23, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75,
	0x73, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38,
	0x0a, 0x18, 0x75, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x65,
	0x6e, 0x73, 0x75, 0x73, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x16, 0x75, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e,
	0x73, 0x75, 0x73, 0x53, 0x74, 0x61, 0x74, 0x65, 0x3a, 0x02, 0x18, 0x01, 0x4a, 0x04, 0x08, 0x01,
	0x10, 0x02, 0x22, 0x3d, 0x0a, 0x1a, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x22, 0x6d, 0x0a, 0x1b, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x4e, 0x0a, 0x0f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2e, 0x75, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x0e, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x32, 0xdc, 0x05, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x9e, 0x01, 0x0a, 0x0b, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x2f, 0x2e, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2e, 0x75, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x75, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2c, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x26, 0x12, 0x24, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x75,
	0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x12, 0xa5, 0x01, 0x0a, 0x0b,
	0x41, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x2f, 0x2e, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x75, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x65,
	0x64, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x75, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41, 0x70, 0x70, 0x6c, 0x69,
	0x65, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x33,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2d, 0x12, 0x2b, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f,
	0x75, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f,
	0x61, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x64, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x2f, 0x7b, 0x6e, 0x61,
	0x6d, 0x65, 0x7d, 0x12, 0xdc, 0x01, 0x0a, 0x16, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x64,
	0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x3a,
	0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x75, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x55, 0x70, 0x67,
	0x72, 0x61, 0x64, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3b, 0x2e, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2e, 0x75, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65,
	0x64, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x49, 0x88, 0x02, 0x01, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x40, 0x12, 0x3e, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x75, 0x70, 0x67, 0x72,
	0x61, 0x64, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x75, 0x70, 0x67, 0x72,
	0x61, 0x64, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x2f, 0x7b, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x7d, 0x12, 0xaa, 0x01, 0x0a, 0x0e, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x32, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x75,
	0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2e, 0x75, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2f,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x29, 0x12, 0x27, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f,
	0x75, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f,
	0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x42,
	0xea, 0x01, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x75,
	0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x0a,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x46, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x75, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x2f, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x75, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x55, 0x58, 0xaa, 0x02, 0x16, 0x43, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2e, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x2e, 0x56, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0xca, 0x02, 0x16, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5c, 0x55, 0x70, 0x67,
	0x72, 0x61, 0x64, 0x65, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xe2, 0x02, 0x22, 0x43,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5c, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x5c, 0x56, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x18, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x3a, 0x3a, 0x55, 0x70, 0x67, 0x72,
	0x61, 0x64, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cosmos_upgrade_v1beta1_query_proto_rawDescOnce sync.Once
	file_cosmos_upgrade_v1beta1_query_proto_rawDescData = file_cosmos_upgrade_v1beta1_query_proto_rawDesc
)

func file_cosmos_upgrade_v1beta1_query_proto_rawDescGZIP() []byte {
	file_cosmos_upgrade_v1beta1_query_proto_rawDescOnce.Do(func() {
		file_cosmos_upgrade_v1beta1_query_proto_rawDescData = protoimpl.X.CompressGZIP(file_cosmos_upgrade_v1beta1_query_proto_rawDescData)
	})
	return file_cosmos_upgrade_v1beta1_query_proto_rawDescData
}

var file_cosmos_upgrade_v1beta1_query_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_cosmos_upgrade_v1beta1_query_proto_goTypes = []interface{}{
	(*QueryCurrentPlanRequest)(nil),             // 0: cosmos.upgrade.v1beta1.QueryCurrentPlanRequest
	(*QueryCurrentPlanResponse)(nil),            // 1: cosmos.upgrade.v1beta1.QueryCurrentPlanResponse
	(*QueryAppliedPlanRequest)(nil),             // 2: cosmos.upgrade.v1beta1.QueryAppliedPlanRequest
	(*QueryAppliedPlanResponse)(nil),            // 3: cosmos.upgrade.v1beta1.QueryAppliedPlanResponse
	(*QueryUpgradedConsensusStateRequest)(nil),  // 4: cosmos.upgrade.v1beta1.QueryUpgradedConsensusStateRequest
	(*QueryUpgradedConsensusStateResponse)(nil), // 5: cosmos.upgrade.v1beta1.QueryUpgradedConsensusStateResponse
	(*QueryModuleVersionsRequest)(nil),          // 6: cosmos.upgrade.v1beta1.QueryModuleVersionsRequest
	(*QueryModuleVersionsResponse)(nil),         // 7: cosmos.upgrade.v1beta1.QueryModuleVersionsResponse
	(*Plan)(nil),                                // 8: cosmos.upgrade.v1beta1.Plan
	(*ModuleVersion)(nil),                       // 9: cosmos.upgrade.v1beta1.ModuleVersion
}
var file_cosmos_upgrade_v1beta1_query_proto_depIdxs = []int32{
	8, // 0: cosmos.upgrade.v1beta1.QueryCurrentPlanResponse.plan:type_name -> cosmos.upgrade.v1beta1.Plan
	9, // 1: cosmos.upgrade.v1beta1.QueryModuleVersionsResponse.module_versions:type_name -> cosmos.upgrade.v1beta1.ModuleVersion
	0, // 2: cosmos.upgrade.v1beta1.Query.CurrentPlan:input_type -> cosmos.upgrade.v1beta1.QueryCurrentPlanRequest
	2, // 3: cosmos.upgrade.v1beta1.Query.AppliedPlan:input_type -> cosmos.upgrade.v1beta1.QueryAppliedPlanRequest
	4, // 4: cosmos.upgrade.v1beta1.Query.UpgradedConsensusState:input_type -> cosmos.upgrade.v1beta1.QueryUpgradedConsensusStateRequest
	6, // 5: cosmos.upgrade.v1beta1.Query.ModuleVersions:input_type -> cosmos.upgrade.v1beta1.QueryModuleVersionsRequest
	1, // 6: cosmos.upgrade.v1beta1.Query.CurrentPlan:output_type -> cosmos.upgrade.v1beta1.QueryCurrentPlanResponse
	3, // 7: cosmos.upgrade.v1beta1.Query.AppliedPlan:output_type -> cosmos.upgrade.v1beta1.QueryAppliedPlanResponse
	5, // 8: cosmos.upgrade.v1beta1.Query.UpgradedConsensusState:output_type -> cosmos.upgrade.v1beta1.QueryUpgradedConsensusStateResponse
	7, // 9: cosmos.upgrade.v1beta1.Query.ModuleVersions:output_type -> cosmos.upgrade.v1beta1.QueryModuleVersionsResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_cosmos_upgrade_v1beta1_query_proto_init() }
func file_cosmos_upgrade_v1beta1_query_proto_init() {
	if File_cosmos_upgrade_v1beta1_query_proto != nil {
		return
	}
	file_cosmos_upgrade_v1beta1_upgrade_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_cosmos_upgrade_v1beta1_query_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryCurrentPlanRequest); i {
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
		file_cosmos_upgrade_v1beta1_query_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryCurrentPlanResponse); i {
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
		file_cosmos_upgrade_v1beta1_query_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryAppliedPlanRequest); i {
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
		file_cosmos_upgrade_v1beta1_query_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryAppliedPlanResponse); i {
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
		file_cosmos_upgrade_v1beta1_query_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryUpgradedConsensusStateRequest); i {
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
		file_cosmos_upgrade_v1beta1_query_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryUpgradedConsensusStateResponse); i {
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
		file_cosmos_upgrade_v1beta1_query_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryModuleVersionsRequest); i {
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
		file_cosmos_upgrade_v1beta1_query_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryModuleVersionsResponse); i {
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
			RawDescriptor: file_cosmos_upgrade_v1beta1_query_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cosmos_upgrade_v1beta1_query_proto_goTypes,
		DependencyIndexes: file_cosmos_upgrade_v1beta1_query_proto_depIdxs,
		MessageInfos:      file_cosmos_upgrade_v1beta1_query_proto_msgTypes,
	}.Build()
	File_cosmos_upgrade_v1beta1_query_proto = out.File
	file_cosmos_upgrade_v1beta1_query_proto_rawDesc = nil
	file_cosmos_upgrade_v1beta1_query_proto_goTypes = nil
	file_cosmos_upgrade_v1beta1_query_proto_depIdxs = nil
}
