package thread

import (
	"crypto/rand"
	"testing"
	"time"

	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewLibp2pIdentity(t *testing.T) {

}

func TestLibp2pIdentity_MarshalBinary(t *testing.T) {

}

func TestLibp2pIdentity_UnmarshalBinary(t *testing.T) {

}

func TestLibp2pIdentity_Sign(t *testing.T) {

}

func TestLibp2pIdentity_GetPublic(t *testing.T) {

}

func TestLibp2pIdentity_Decrypt(t *testing.T) {

}

func TestLibp2pIdentity_Token(t *testing.T) {
	aud, err := makeLibp2pIdentity(t).GetPublic().DID()
	require.NoError(t, err)
	tk, err := makeLibp2pIdentity(t).Token(aud, time.Minute)
	require.NoError(t, err)
	assert.NotEmpty(t, tk)
}

func TestLibp2pPubKey_Validate(t *testing.T) {
	i1 := makeLibp2pIdentity(t).GetPublic() // The key receiving the token, i.e., the audience
	aud, err := i1.DID()
	require.NoError(t, err)
	i2 := makeLibp2pIdentity(t) // The identity that will be validated by the audience
	tk, err := i2.Token(aud, time.Minute)
	require.NoError(t, err)
	k, doc, err := i1.Validate(tk)
	require.NoError(t, err)
	assert.True(t, k.Equals(i2.GetPublic()))
	assert.NotEmpty(t, doc)
}

func TestLibp2pIdentity_Equals(t *testing.T) {

}

func TestNewLibp2pPubKey(t *testing.T) {

}

func TestLibp2pPubKey_MarshalBinary(t *testing.T) {

}

func TestLibp2pPubKey_UnmarshalBinary(t *testing.T) {

}

func TestLibp2pPubKey_String(t *testing.T) {

}

func TestLibp2pPubKey_UnmarshalString(t *testing.T) {

}

func TestLibp2pPubKey_Encrypt(t *testing.T) {

}

func TestLibp2pPubKey_Hash(t *testing.T) {

}

func TestLibp2pPubKey_DID(t *testing.T) {
	i := makeLibp2pIdentity(t)
	id, err := i.GetPublic().DID()
	require.NoError(t, err)
	assert.NotEmpty(t, id)
}

func TestLibp2pPubKey_Equals(t *testing.T) {

}

func makeLibp2pIdentity(t *testing.T) Identity {
	sk, _, err := crypto.GenerateEd25519Key(rand.Reader)
	require.NoError(t, err)
	return NewLibp2pIdentity(sk)
}