package crypto

import "crypto/rc4"

func CipherDropBytes(cipher *rc4.Cipher, length int) {
	dropped := make([]byte, length)
	cipher.XORKeyStream(dropped[:], dropped[:])
}
