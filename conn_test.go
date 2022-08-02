// conn_test.go
//
// Copyright (C) 2022 Holger de Carne
//
// This software may be modified and distributed under the terms
// of the MIT license. See the LICENSE file for details.

package nsdp

import (
	"testing"

	"github.com/stretchr/testify/require"
)

//const connTestTarget string = IPv4BroadcastTarget

const connTestTarget string = "127.0.0.1:63322"

func TestConnSendReceiveMessageBroadcast(t *testing.T) {
	responder, err := NewTestResponder(connTestTarget)
	require.Nil(t, err)
	defer responder.Stop()
	responder.AddResponses(
		"0102000000000000bcd07432b8dc6cb0ce1c8394000099d14e534450000000000001000847533130384576330003000773776974636831000400066cb0ce1c839400050000000600040a01000300070004ffff0000000800040a010001000b000100000d0007322e30362e3137000e0000000f0001010c0000030105000c0000030200000c0000030304000c0000030400000c0000030504000c0000030600000c0000030700000c0000030800001000003101000000011b86e2c2000000000d159e3800000000000000000000000000000000000000000000000000000000000000001000003102000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000100000310300000000039bd6ce000000000874984f000000000000000000000000000000000000000000000000000000000000000010000031040000000000133f340000000000cf6d03000000000000000000000000000000000000000000000000000000000000000010000031050000000009668768000000010afa8d1d0000000000000000000000000000000000000000000000000000000000000000100000310600000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000010000031070000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001000003108000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000ffff0000",
		"0102000000000000bcd07432b8dce4f4c6ffa7a2000099d14e53445000000000000100084753313038457633000300077377697463683200040006e4f4c6ffa7a200050000000600040a01000400070004ffff0000000800040a010001000b000100000d0007322e30362e3137000e0000000f0001010c0000030105000c0000030205000c0000030302000c0000030404000c0000030500000c0000030600000c0000030700000c0000030800001000003101000000009d57dcbf000000000e10739f0000000000000000000000000000000000000000000000000000000000000000100000310200000000091cf6760000000028dfe4ca000000000000000000000000000000000000000000000000000000000000000010000031030000000005a930200000000081ccfd9a000000000000000000000000000000000000000000000000000000000000000010000031040000000000c2ebb8000000000cd0177800000000000000000000000000000000000000000000000000000000000000001000003105000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000100000310600000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000010000031070000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001000003108000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000ffff0000")
	err = responder.Start()
	require.Nil(t, err)
	conn, err := NewConn(connTestTarget, true)
	require.Nil(t, err)
	defer conn.Close()
	conn.ReceiveDeviceLimit = 2
	msg := prepareTestMessage()
	responses, err := conn.SendReceiveMessage(msg)
	require.Nil(t, err)
	require.Equal(t, 2, len(responses))
}

func TestConnSendReceiveMessageUnicast(t *testing.T) {
	responder, err := NewTestResponder(connTestTarget)
	require.Nil(t, err)
	defer responder.Stop()
	responder.AddResponses("0102000000000000bcd07432b8dce4f4c6ffa7a200001a414e53445000000000000100084753313038457633000300077377697463683200040006e4f4c6ffa7a200050000000600040a01000400070004ffff0000000800040a010001000b000100000d0007322e30362e3137000e0000000f0001010c0000030105000c0000030205000c0000030302000c0000030404000c0000030500000c0000030600000c0000030700000c0000030800001000003101000000009d55f306000000000e100c210000000000000000000000000000000000000000000000000000000000000000100000310200000000091c99ed0000000028ddfe8b000000000000000000000000000000000000000000000000000000000000000010000031030000000005a92fe00000000081cb4ea2000000000000000000000000000000000000000000000000000000000000000010000031040000000000c2e89b000000000cce6b8c00000000000000000000000000000000000000000000000000000000000000001000003105000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000100000310600000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000010000031070000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001000003108000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000ffff0000")
	err = responder.Start()
	require.Nil(t, err)
	conn, err := NewConn(connTestTarget, true)
	require.Nil(t, err)
	defer conn.Close()
	msg := prepareTestMessage()
	msg.Header.DeviceAddress = []byte{0xe4, 0xf4, 0xc6, 0xff, 0xa7, 0xa2}
	responses, err := conn.SendReceiveMessage(msg)
	require.Nil(t, err)
	require.Equal(t, 1, len(responses))
}

func prepareTestMessage() *Message {
	message := NewMessage(ReadRequest)
	message.AppendTLV(EmptyDeviceModel())
	message.AppendTLV(EmptyDeviceName())
	message.AppendTLV(EmptyDeviceMAC())
	message.AppendTLV(EmptyDeviceLocation())
	message.AppendTLV(EmptyDeviceIP())
	message.AppendTLV(EmptyDeviceNetmask())
	message.AppendTLV(EmptyRouterIP())
	message.AppendTLV(EmptyDHCPMode())
	message.AppendTLV(EmptyPortStatus())
	message.AppendTLV(EmptyPortStatistic())
	message.AppendTLV(EmptyFWVersionSlot1())
	message.AppendTLV(EmptyFWVersionSlot2())
	message.AppendTLV(EmptyNextFWSlot())
	return message
}
