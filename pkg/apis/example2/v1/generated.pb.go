/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: k8s.io/apiserver/pkg/apis/example2/v1/generated.proto

package v1

import (
	fmt "fmt"

	io "io"

	proto "github.com/gogo/protobuf/proto"

	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
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

func (m *ReplicaSet) Reset()      { *m = ReplicaSet{} }
func (*ReplicaSet) ProtoMessage() {}
func (*ReplicaSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0d8f6d73eb5bf83, []int{0}
}
func (m *ReplicaSet) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ReplicaSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *ReplicaSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplicaSet.Merge(m, src)
}
func (m *ReplicaSet) XXX_Size() int {
	return m.Size()
}
func (m *ReplicaSet) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplicaSet.DiscardUnknown(m)
}

var xxx_messageInfo_ReplicaSet proto.InternalMessageInfo

func (m *ReplicaSetSpec) Reset()      { *m = ReplicaSetSpec{} }
func (*ReplicaSetSpec) ProtoMessage() {}
func (*ReplicaSetSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0d8f6d73eb5bf83, []int{1}
}
func (m *ReplicaSetSpec) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ReplicaSetSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *ReplicaSetSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplicaSetSpec.Merge(m, src)
}
func (m *ReplicaSetSpec) XXX_Size() int {
	return m.Size()
}
func (m *ReplicaSetSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplicaSetSpec.DiscardUnknown(m)
}

var xxx_messageInfo_ReplicaSetSpec proto.InternalMessageInfo

func (m *ReplicaSetStatus) Reset()      { *m = ReplicaSetStatus{} }
func (*ReplicaSetStatus) ProtoMessage() {}
func (*ReplicaSetStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0d8f6d73eb5bf83, []int{2}
}
func (m *ReplicaSetStatus) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ReplicaSetStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *ReplicaSetStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplicaSetStatus.Merge(m, src)
}
func (m *ReplicaSetStatus) XXX_Size() int {
	return m.Size()
}
func (m *ReplicaSetStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplicaSetStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ReplicaSetStatus proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ReplicaSet)(nil), "k8s.io.apiserver.pkg.apis.example2.v1.ReplicaSet")
	proto.RegisterType((*ReplicaSetSpec)(nil), "k8s.io.apiserver.pkg.apis.example2.v1.ReplicaSetSpec")
	proto.RegisterType((*ReplicaSetStatus)(nil), "k8s.io.apiserver.pkg.apis.example2.v1.ReplicaSetStatus")
}

func init() {
	proto.RegisterFile("k8s.io/apiserver/pkg/apis/example2/v1/generated.proto", fileDescriptor_c0d8f6d73eb5bf83)
}

var fileDescriptor_c0d8f6d73eb5bf83 = []byte{
	// 388 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xc1, 0x6a, 0xe2, 0x40,
	0x18, 0xc7, 0x13, 0xd7, 0x15, 0x99, 0x15, 0x91, 0x9c, 0xc4, 0xc3, 0xb8, 0x08, 0x82, 0x87, 0xdd,
	0x99, 0x55, 0xd6, 0xdd, 0xd2, 0x53, 0xc9, 0xb5, 0x94, 0x42, 0x3c, 0x14, 0x7a, 0x69, 0xc7, 0xf8,
	0x35, 0xa6, 0x1a, 0x33, 0x64, 0x26, 0xa1, 0xbd, 0xf5, 0x11, 0xfa, 0x18, 0x7d, 0x14, 0x8f, 0x1e,
	0x3d, 0x49, 0x4d, 0x5f, 0xa4, 0x38, 0x49, 0x13, 0xaa, 0x96, 0xda, 0x5b, 0xfe, 0x93, 0xf9, 0xfd,
	0xbe, 0x3f, 0x1f, 0x83, 0xfa, 0x93, 0x23, 0x41, 0x5c, 0x9f, 0x32, 0xee, 0x0a, 0x08, 0x22, 0x08,
	0x28, 0x9f, 0x38, 0x2a, 0x51, 0xb8, 0x63, 0x1e, 0x9f, 0x42, 0x8f, 0x46, 0x5d, 0xea, 0xc0, 0x0c,
	0x02, 0x26, 0x61, 0x44, 0x78, 0xe0, 0x4b, 0xdf, 0x68, 0x27, 0x18, 0xc9, 0x30, 0xc2, 0x27, 0x8e,
	0x4a, 0xe4, 0x0d, 0x23, 0x51, 0xb7, 0xf1, 0xdb, 0x71, 0xe5, 0x38, 0x1c, 0x12, 0xdb, 0xf7, 0xa8,
	0xe3, 0x3b, 0x3e, 0x55, 0xf4, 0x30, 0xbc, 0x51, 0x49, 0x05, 0xf5, 0x95, 0x58, 0x1b, 0x7f, 0xf3,
	0x32, 0x1e, 0xb3, 0xc7, 0xee, 0x0c, 0x82, 0xfb, 0xbc, 0x8f, 0x07, 0x92, 0xed, 0xe9, 0xd2, 0xa0,
	0x1f, 0x51, 0x41, 0x38, 0x93, 0xae, 0x07, 0x3b, 0xc0, 0xbf, 0xcf, 0x00, 0x61, 0x8f, 0xc1, 0x63,
	0xdb, 0x5c, 0xeb, 0xa9, 0x80, 0x90, 0x05, 0x7c, 0xea, 0xda, 0x6c, 0x00, 0xd2, 0xb8, 0x46, 0xe5,
	0x4d, 0xa5, 0x11, 0x93, 0xac, 0xae, 0xff, 0xd4, 0x3b, 0x3f, 0x7a, 0x7f, 0x48, 0xbe, 0x96, 0xcc,
	0x9c, 0x6f, 0x66, 0x73, 0x9b, 0x44, 0x5d, 0x72, 0x3e, 0xbc, 0x05, 0x5b, 0x9e, 0x81, 0x64, 0xa6,
	0x31, 0x5f, 0x35, 0xb5, 0x78, 0xd5, 0x44, 0xf9, 0x99, 0x95, 0x59, 0x8d, 0x0b, 0x54, 0x14, 0x1c,
	0xec, 0x7a, 0x41, 0xd9, 0xfb, 0xe4, 0xa0, 0xa5, 0x93, 0xbc, 0xe2, 0x80, 0x83, 0x6d, 0x56, 0xd2,
	0x11, 0xc5, 0x4d, 0xb2, 0x94, 0xd0, 0xb8, 0x42, 0x25, 0x21, 0x99, 0x0c, 0x45, 0xfd, 0x9b, 0x52,
	0xff, 0xff, 0xba, 0x5a, 0xe1, 0x66, 0x35, 0x95, 0x97, 0x92, 0x6c, 0xa5, 0xda, 0xd6, 0x31, 0xaa,
	0xbe, 0xaf, 0x61, 0x74, 0x50, 0x39, 0x48, 0x4e, 0x84, 0xda, 0xd6, 0x77, 0xb3, 0x12, 0xaf, 0x9a,
	0xe5, 0xf4, 0x96, 0xb0, 0xb2, 0xbf, 0xad, 0x13, 0x54, 0xdb, 0x9e, 0x63, 0xfc, 0xda, 0xa1, 0x6b,
	0xe9, 0xe4, 0x3d, 0x06, 0xf3, 0x74, 0xbe, 0xc6, 0xda, 0x62, 0x8d, 0xb5, 0xe5, 0x1a, 0x6b, 0x0f,
	0x31, 0xd6, 0xe7, 0x31, 0xd6, 0x17, 0x31, 0xd6, 0x97, 0x31, 0xd6, 0x9f, 0x63, 0xac, 0x3f, 0xbe,
	0x60, 0xed, 0xb2, 0x7d, 0xd0, 0xd3, 0x7f, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xf4, 0xd7, 0x7c, 0xa3,
	0x22, 0x03, 0x00, 0x00,
}

func (m *ReplicaSet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReplicaSet) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ReplicaSet) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Status.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.Spec.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.ObjectMeta.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *ReplicaSetSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReplicaSetSpec) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ReplicaSetSpec) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Replicas != nil {
		i = encodeVarintGenerated(dAtA, i, uint64(*m.Replicas))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ReplicaSetStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReplicaSetStatus) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ReplicaSetStatus) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	i = encodeVarintGenerated(dAtA, i, uint64(m.Replicas))
	i--
	dAtA[i] = 0x8
	return len(dAtA) - i, nil
}

func encodeVarintGenerated(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenerated(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ReplicaSet) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ObjectMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Spec.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Status.Size()
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *ReplicaSetSpec) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Replicas != nil {
		n += 1 + sovGenerated(uint64(*m.Replicas))
	}
	return n
}

func (m *ReplicaSetStatus) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovGenerated(uint64(m.Replicas))
	return n
}

func sovGenerated(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *ReplicaSet) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ReplicaSet{`,
		`ObjectMeta:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.ObjectMeta), "ObjectMeta", "v1.ObjectMeta", 1), `&`, ``, 1) + `,`,
		`Spec:` + strings.Replace(strings.Replace(this.Spec.String(), "ReplicaSetSpec", "ReplicaSetSpec", 1), `&`, ``, 1) + `,`,
		`Status:` + strings.Replace(strings.Replace(this.Status.String(), "ReplicaSetStatus", "ReplicaSetStatus", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ReplicaSetSpec) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ReplicaSetSpec{`,
		`Replicas:` + valueToStringGenerated(this.Replicas) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ReplicaSetStatus) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ReplicaSetStatus{`,
		`Replicas:` + fmt.Sprintf("%v", this.Replicas) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringGenerated(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *ReplicaSet) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: ReplicaSet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReplicaSet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ObjectMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Spec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Spec.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Status.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenerated
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
func (m *ReplicaSetSpec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: ReplicaSetSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReplicaSetSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Replicas", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Replicas = &v
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenerated
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
func (m *ReplicaSetStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: ReplicaSetStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReplicaSetStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Replicas", wireType)
			}
			m.Replicas = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Replicas |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenerated
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
func skipGenerated(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenerated
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
					return 0, ErrIntOverflowGenerated
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
					return 0, ErrIntOverflowGenerated
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
				return 0, ErrInvalidLengthGenerated
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenerated
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenerated
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenerated        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerated          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenerated = fmt.Errorf("proto: unexpected end of group")
)
