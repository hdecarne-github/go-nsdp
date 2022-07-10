// message_header.go
//
// Copyright (C) 2022 Holger de Carne
//
// This software may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.
//
package nsdp

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"strings"
)

// Header NSDP message header
type Header struct {
	Version       ProtoVersion
	Operation     OperationCode
	Result        OperationResult
	Unknown1      uint32
	HostAddress   net.HardwareAddr
	DeviceAddress net.HardwareAddr
	Unknown2      uint16
	Sequence      uint16
	Signature     Signature
	Unknown3      uint32
}

type ProtoVersion uint8

const (
	ProtoVersion1 ProtoVersion = 0x01 // Only known version
)

type OperationCode uint8

const (
	ReadRequest   OperationCode = 0x01
	ReadResponse  OperationCode = 0x02
	WriteRequest  OperationCode = 0x03
	WriteResponse OperationCode = 0x04
)

type OperationResult uint16

type Signature uint32

const (
	NSDPSignature Signature = 0x4e534450
)

func newHeader(operation OperationCode) Header {
	return Header{
		Version:       ProtoVersion1,
		Operation:     operation,
		HostAddress:   make([]byte, 6),
		DeviceAddress: make([]byte, 6),
		Signature:     NSDPSignature,
	}
}

func (h *Header) writeString(builder *strings.Builder) {
	fmt.Fprintf(builder, "Header: %02xh %02xh %04xh %08xh %s %s %04xh %04xh %08xh", h.Version, h.Operation, h.Result, h.Unknown1, h.HostAddress.String(), h.DeviceAddress.String(), h.Unknown2, h.Sequence, h.Signature)
}

func (h *Header) marshalBuffer(buffer *bytes.Buffer) {
	buffer.WriteByte(byte(h.Version))
	buffer.WriteByte(byte(h.Operation))
	binary.Write(buffer, binary.BigEndian, h.Result)
	binary.Write(buffer, binary.BigEndian, h.Unknown1)
	buffer.Write(h.HostAddress)
	buffer.Write(h.DeviceAddress)
	binary.Write(buffer, binary.BigEndian, h.Unknown2)
	binary.Write(buffer, binary.BigEndian, h.Sequence)
	binary.Write(buffer, binary.BigEndian, h.Signature)
	binary.Write(buffer, binary.BigEndian, h.Unknown3)
}

func unmarshalHeaderBuffer(buffer *bytes.Buffer) (*Header, error) {
	header := &Header{
		HostAddress:   make([]byte, 6),
		DeviceAddress: make([]byte, 6),
	}
	version, err := buffer.ReadByte()
	if err != nil {
		return nil, fmt.Errorf("error while decoding proto version; cause: %v", err)
	}
	if version != uint8(ProtoVersion1) {
		return nil, fmt.Errorf("unrecognized proto version: %02xh", version)
	}
	header.Version = ProtoVersion(version)
	operation, err := buffer.ReadByte()
	if err != nil {
		return nil, fmt.Errorf("error while decoding operation code; cause: %v", err)
	}
	if operation != uint8(ReadRequest) && operation != uint8(ReadResponse) && operation != uint8(WriteRequest) && operation != uint8(WriteResponse) {
		return nil, fmt.Errorf("unrecognized operation code: %04xh", operation)
	}
	header.Operation = OperationCode(operation)
	err = binary.Read(buffer, binary.BigEndian, &header.Result)
	if err != nil {
		return nil, fmt.Errorf("error while decoding result code; cause: %v", err)
	}
	err = binary.Read(buffer, binary.BigEndian, &header.Unknown1)
	if err != nil {
		return nil, fmt.Errorf("error while decoding unknown1; cause: %v", err)
	}
	err = binary.Read(buffer, binary.BigEndian, header.HostAddress)
	if err != nil {
		return nil, fmt.Errorf("error while decoding host address; cause: %v", err)
	}
	err = binary.Read(buffer, binary.BigEndian, header.DeviceAddress)
	if err != nil {
		return nil, fmt.Errorf("error while decoding device address; cause: %v", err)
	}
	err = binary.Read(buffer, binary.BigEndian, &header.Unknown2)
	if err != nil {
		return nil, fmt.Errorf("error while decoding unknown2; cause: %v", err)
	}
	err = binary.Read(buffer, binary.BigEndian, &header.Sequence)
	if err != nil {
		return nil, fmt.Errorf("error while decoding sequence; cause: %v", err)
	}
	var signature uint32
	err = binary.Read(buffer, binary.BigEndian, &signature)
	if err != nil {
		return nil, fmt.Errorf("error while decoding signature; cause: %v", err)
	}
	if signature != uint32(NSDPSignature) {
		return nil, fmt.Errorf("unrecognized signature: %08xh", signature)
	}
	header.Signature = Signature(signature)
	err = binary.Read(buffer, binary.BigEndian, &header.Unknown3)
	if err != nil {
		return nil, fmt.Errorf("error while decoding unknown3; cause: %v", err)
	}
	return header, nil
}
