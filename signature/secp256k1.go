package signature

import (
	"errors"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcec/v2/ecdsa"

	"github.com/ontio/ontology-crypto/ec"
)

func Secp256k1Sign(pri *ec.PrivateKey, hash []byte) ([]byte, error) {
	key, err := pri.PrivateKey.ECDH()
	if err != nil {
		return nil, err
	}
	privateKey, _ := btcec.PrivKeyFromBytes(key.Bytes())
	return ecdsa.SignCompact(privateKey, hash, false)
}

func Secp256k1Verify(pub *ec.PublicKey, hash []byte, sig []byte) bool {
	recKey, _, err := ecdsa.RecoverCompact(sig, hash)
	if err != nil {
		return false
	}
	publicKey, err := pub.PublicKey.ECDH()
	if err != nil {
		return false
	}
	parsePubKey, err := btcec.ParsePubKey(publicKey.Bytes())
	if err != nil {
		return false
	}
	return recKey.IsEqual(parsePubKey)
}

func ConvertToEthCompatible(sig []byte) ([]byte, error) {
	s, err := Deserialize(sig)
	if err != nil {
		return nil, err
	}

	t, ok := s.Value.([]byte)
	if !ok {
		return nil, errors.New("invalid signature type")
	}

	if len(t) != 65 {
		return nil, errors.New("invalid signature length")
	}

	v := t[0] - 27
	copy(t, t[1:])
	t[64] = v
	return t, nil
}
