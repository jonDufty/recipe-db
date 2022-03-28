package auth

import (
	"testing"

	"github.com/jonDufty/recipes/auth/models/user"
	"github.com/stretchr/testify/require"
)

type test struct {
	name  string
	input *user.User
	token string
}

func TestCreateJwt(t *testing.T) {

	tests := []test{
		{
			name:  "Test Basic Input",
			input: user.NewTest(),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := GenerateToken(tc.input)
			require.NoError(t, err)
		})
	}

}

func TestParseJwt(t *testing.T) {
	tests := []test{
		{
			name:  "Test TestUser input",
			input: user.NewTest(),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			token, _ := GenerateToken(tc.input)
			_, err := ParseToken(token)
			require.NoError(t, err)
		})
	}

	t.Run("Test Invalid Token", func(t *testing.T) {
		_, err := ParseToken("foo.bar.baz")
		require.Error(t, err)
	})
}
