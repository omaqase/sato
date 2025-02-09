package payment

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncryption(t *testing.T) {
	tests := []struct {
		name string
		data string
	}{
		{
			name: "encrypt_card_number",
			data: "4242424242424242",
		},
		{
			name: "encrypt_cvv",
			data: "123",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encrypted := Encrypt([]byte(tt.data))
			assert.NotEmpty(t, encrypted)
			assert.NotEqual(t, tt.data, string(encrypted))
		})
	}
}
