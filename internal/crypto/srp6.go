package crypto

import (
	"crypto/sha1"
	"math/big"
	"strings"

	"github.com/davidaburns/menethil-core/internal/bytes"
)

const (
	PROOF_DATA_SIZE = 16
	PROOF_SIZE = 20
	SALT_SIZE = 32
	VERIFIER_SIZE = 32
	KEY_SIZE = 32
	LARGE_PRIME_SIZE = 32
	SESSION_KEY_SIZE = 40
	GENERATOR = 7
)

var (
	xorHash = []byte {
		0xDD, 0x7B, 0xB0, 0x3A, 0x38, 0xAC, 0x73, 0x11,
		0x03, 0x98, 0x7C, 0x5A, 0x50, 0x6F, 0xCA, 0x96,
		0x6C, 0x7B, 0xC2, 0xA7,
	}

	largeSafePrime = bytes.ReverseBytes([]byte {
		0x89, 0x4B, 0x64, 0x5E, 0x89, 0xE1, 0x53, 0x5B,
		0xBD, 0xAD, 0x5B, 0x8B, 0x29, 0x06, 0x50, 0x53,
		0x08, 0x01, 0xB1, 0x8E, 0xBF, 0xBF, 0x5E, 0x8F,
		0xAB, 0x3C, 0x82, 0x87, 0x2A, 0x3E, 0x9B, 0xB7,
	})

	n = bytes.BytesToBigIntLittleEndian(largeSafePrime)
	g = big.NewInt(GENERATOR)
	k = big.NewInt(3)
)

func SRP6ClientChallengeProof(
	username string,
	salt []byte,
	clientPublicKey []byte,
	serverPublicKey []byte,
	sessionKey []byte,
) []byte {
	hashedUsername := sha1.Sum([]byte(strings.ToUpper(username)))

	proof := sha1.New()
	proof.Write(xorHash)
	proof.Write(hashedUsername[:])
	proof.Write(salt)
	proof.Write(clientPublicKey)
	proof.Write(serverPublicKey)
	proof.Write(sessionKey)

	return proof.Sum(nil)
}

func SRP6ServerChallengeProof(
	clientPublicKey []byte,
	clientProof []byte,
	sessionKey []byte,
) []byte {
	proof := sha1.New()

	proof.Write(clientPublicKey)
	proof.Write(clientProof)
	proof.Write(sessionKey)

	return proof.Sum(nil)
}

func SRP6ReconnectProof(
	username string,
	clientData []byte,
	serverData []byte,
	sessionKey []byte,
) []byte {
	proof := sha1.New()

	proof.Write([]byte(strings.ToUpper(username)))
	proof.Write(clientData)
	proof.Write(serverData)
	proof.Write(sessionKey)

	return proof.Sum(nil)
}

func SRP6WorldProof(
	username string,
	clientSeed []byte,
	serverSeed []byte,
	sessionKey []byte,
) []byte {
	proof := sha1.New()

	proof.Write([]byte(strings.ToUpper(username)))
	proof.Write([]byte{0, 0, 0, 0})
	proof.Write(clientSeed)
	proof.Write(serverSeed)
	proof.Write(sessionKey)

	return proof.Sum(nil)
}
