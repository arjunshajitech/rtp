// SPDX-FileCopyrightText: 2023 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package rtp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayoutDelayExtensionTooSmall(t *testing.T) {
	t1 := PlayoutDelayExtension{}

	var rawData []byte

	err := t1.Unmarshal(rawData)
	assert.ErrorIs(t, err, errTooSmall)
}

func TestPlayoutDelayExtensionTooLarge(t *testing.T) {
	t1 := PlayoutDelayExtension{MinDelay: 1 << 12, MaxDelay: 1 << 12}

	_, err := t1.Marshal()
	assert.ErrorIs(t, err, errPlayoutDelayInvalidValue)
}

func TestPlayoutDelayExtension(t *testing.T) {
	t1 := PlayoutDelayExtension{}

	rawData := []byte{
		0x01, 0x01, 0x00,
	}

	err := t1.Unmarshal(rawData)
	assert.NoError(t, err)

	t2 := PlayoutDelayExtension{
		MinDelay: 1 << 4, MaxDelay: 1 << 8,
	}

	assert.Equal(t, t1, t2)

	dstData, _ := t2.Marshal()
	assert.Equal(t, dstData, rawData)
}

func TestPlayoutDelayExtensionExtraBytes(t *testing.T) {
	t1 := PlayoutDelayExtension{}

	rawData := []byte{
		0x01, 0x01, 0x00, 0xff, 0xff,
	}

	err := t1.Unmarshal(rawData)
	assert.NoError(t, err)

	t2 := PlayoutDelayExtension{
		MinDelay: 1 << 4, MaxDelay: 1 << 8,
	}

	assert.Equal(t, t1, t2)
}
