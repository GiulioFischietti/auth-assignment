package crypto_test

import (
	"testing"

	"auth-service/crypto"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateToken(t *testing.T) {

	token, err := crypto.GenerateToken()

	require.NoError(
		t,
		err,
	)

	assert.NotEmpty(
		t,
		token,
	)

}

func TestGenerateTokenIsUnique(t *testing.T) {

	token1, err := crypto.GenerateToken()

	require.NoError(
		t,
		err,
	)

	token2, err := crypto.GenerateToken()

	require.NoError(
		t,
		err,
	)

	assert.NotEqual(
		t,
		token1,
		token2,
	)

}
