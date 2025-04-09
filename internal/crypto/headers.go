package crypto

import (
	"crypto/rc4"
	"sync"

	"github.com/davidaburns/menethil-core/internal/errors"
)

const (
	SIZE_FIELD_MAX_VALUE = 0x7FFFFF
	LARGE_HEADER_THRESHOLD = 0x7FFF
	LARGE_HEADER_FLAG = 0x80
)

type HeaderCrypto struct {
	decryptCipher *rc4.Cipher
	encrypCipher *rc4.Cipher
	decryptMutex sync.Mutex
	encryptMutex sync.Mutex
}

func NewHeaderCrypto(sessionKey []byte) (*HeaderCrypto, error) {
	return NewHeaderCryptoWithKeys(sessionKey, WRATH_DECRYPT_KEY, WRATH_ENCRYPT_KEY)
}

func NewHeaderCryptoWithKeys(sessionKey []byte, decryptKey []byte, encryptKey []byte) (*HeaderCrypto, error) {
	decryptCipher, err := rc4.NewCipher(GenerateKey(sessionKey, decryptKey))
	if err != nil {
		return nil, err
	}

	encryptCipher, err := rc4.NewCipher(GenerateKey(sessionKey, encryptKey))
	if err != nil {
		return nil, err
	}

	header := &HeaderCrypto{
		decryptCipher: decryptCipher,
		encrypCipher: encryptCipher,
	}

	CipherDropBytes(header.decryptCipher, 1024)
	CipherDropBytes(header.encrypCipher, 1024)

	return header, nil
}

func (h *HeaderCrypto) Encrypt(data []byte) error {
	if h.encrypCipher == nil {
		return errors.ErrCryptoNotInitialized
	}

	h.encryptMutex.Lock()
	h.encrypCipher.XORKeyStream(data, data)
	h.encryptMutex.Unlock()

	return nil
}

func (h *HeaderCrypto) Decrypt(data []byte) error {
	if h.decryptCipher == nil {
		return errors.ErrCryptoNotInitialized
	}

	h.decryptMutex.Lock()
	h.decryptCipher.XORKeyStream(data, data)
	h.decryptMutex.Unlock()

	return nil
}

func (h *HeaderCrypto) EncodeHeader(opcode uint16, size uint32) ([]byte, error) {
	size += 2
	if size > SIZE_FIELD_MAX_VALUE {
		return nil, errors.ErrHeaderSizeTooLarge
	}

	var header []byte
	if size > LARGE_HEADER_THRESHOLD {
		header = []byte {
			byte(size >> 16) | LARGE_HEADER_FLAG,
			byte(size >> 8),
			byte(size),
			byte(opcode),
			byte(opcode >> 8),
		}
	} else {
		header = []byte {
			byte(size >> 8),
			byte(size),
			byte(opcode),
			byte(opcode >> 8),
		}
	}

	if h.encrypCipher != nil {
		if err := h.Encrypt(header); err != nil {
			return nil, err
		}
	}

	return header, nil
}
