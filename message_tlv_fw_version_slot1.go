// message_tlv_device_name.go
//
// Copyright (C) 2022 Holger de Carne
//
// This software may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.

package nsdp

import (
	"fmt"
)

type FWVersionSlot1 struct {
	Version string
}

func EmptyFWVersionSlot1() *FWVersionSlot1 {
	return NewFWVersionSlot1("")
}

func NewFWVersionSlot1(version string) *FWVersionSlot1 {
	return &FWVersionSlot1{Version: version}
}

func unmarshalFWVersionSlot1(bytes []byte) (*FWVersionSlot1, error) {
	return NewFWVersionSlot1(string(bytes)), nil
}

func (tlv *FWVersionSlot1) Type() Type {
	return TypeFWVersionSlot1
}

func (tlv *FWVersionSlot1) Length() uint16 {
	return uint16(len(tlv.Version))
}

func (tlv *FWVersionSlot1) Value() []byte {
	return []byte(tlv.Version)
}

func (tlv *FWVersionSlot1) String() string {
	return fmt.Sprintf("FWVersionSlot1(%04xh) '%s'", TypeFWVersionSlot1, tlv.Version)
}
