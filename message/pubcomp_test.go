// Copyright (c) 2014 Dataence, LLC. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package message

import (
	"testing"

	"github.com/dataence/assert"
)

func TestPubcompMessageFields(t *testing.T) {
	msg := NewPubcompMessage()

	msg.SetPacketId(100)

	assert.Equal(t, true, 100, msg.PacketId())
}

func TestPubcompMessageDecode(t *testing.T) {
	msgBytes := []byte{
		byte(PUBCOMP << 4),
		2,
		0, // packet ID MSB (0)
		7, // packet ID LSB (7)
	}

	msg := NewPubcompMessage()
	n, err := msg.Decode(msgBytes)

	assert.NoError(t, true, err, "Error decoding message.")
	assert.Equal(t, true, len(msgBytes), n, "Error decoding message.")
	assert.Equal(t, true, PUBCOMP, msg.Type(), "Error decoding message.")
	assert.Equal(t, true, 7, msg.PacketId(), "Error decoding message.")
}

// test insufficient bytes
func TestPubcompMessageDecode2(t *testing.T) {
	msgBytes := []byte{
		byte(PUBCOMP << 4),
		2,
		7, // packet ID LSB (7)
	}

	msg := NewPubcompMessage()
	_, err := msg.Decode(msgBytes)

	assert.Error(t, true, err)
}

func TestPubcompMessageEncode(t *testing.T) {
	msgBytes := []byte{
		byte(PUBCOMP << 4),
		2,
		0, // packet ID MSB (0)
		7, // packet ID LSB (7)
	}

	msg := NewPubcompMessage()
	msg.SetPacketId(7)

	dst := make([]byte, 10)
	n, err := msg.Encode(dst)

	assert.NoError(t, true, err, "Error decoding message.")
	assert.Equal(t, true, len(msgBytes), n, "Error decoding message.")
	assert.Equal(t, true, msgBytes, dst[:n], "Error decoding message.")
}

// test to ensure encoding and decoding are the same
// decode, encode, and decode again
func TestPubcompDecodeEncodeEquiv(t *testing.T) {
	msgBytes := []byte{
		byte(PUBCOMP << 4),
		2,
		0, // packet ID MSB (0)
		7, // packet ID LSB (7)
	}

	msg := NewPubcompMessage()
	n, err := msg.Decode(msgBytes)

	assert.NoError(t, true, err, "Error decoding message.")
	assert.Equal(t, true, len(msgBytes), n, "Error decoding message.")

	dst := make([]byte, 100)
	n2, err := msg.Encode(dst)

	assert.NoError(t, true, err, "Error decoding message.")
	assert.Equal(t, true, len(msgBytes), n2, "Error decoding message.")
	assert.Equal(t, true, msgBytes, dst[:n2], "Error decoding message.")

	n3, err := msg.Decode(dst)

	assert.NoError(t, true, err, "Error decoding message.")
	assert.Equal(t, true, len(msgBytes), n3, "Error decoding message.")
}
