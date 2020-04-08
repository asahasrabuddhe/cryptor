// aes 256 cbc
package module

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
)

func AESCBCEncrypt(plain []byte, key []byte) (encoded []byte, err error) {
	var (
		iv        []byte
		c         cipher.Block
		bSize     int
		encrypter cipher.BlockMode
	)

	plain = PKCS5Padding(plain, aes.BlockSize)
	c, err = aes.NewCipher(key)
	if err != nil {
		return
	}

	// generate a random IV
	bSize = c.BlockSize()
	iv = make([]byte, bSize)
	_, err = rand.Read(iv)
	if err != nil {
		return
	}

	encoded = make([]byte, len(plain))
	encrypter = cipher.NewCBCEncrypter(c, iv)
	encrypter.CryptBlocks(encoded, plain)

	// attach the IV to be retrieved for
	encoded = append(encoded, iv...)
	return
}

func AESCBCDecrypt(encoded []byte, key []byte) (decoded []byte, err error) {
	var (
		iv        []byte
		c         cipher.Block
		decrypter cipher.BlockMode
	)

	// get the iv
	iv = encoded[len(encoded)-aes.BlockSize:]

	// remove padding from crypt text
	encoded = encoded[:len(encoded)-aes.BlockSize]

	c, err = aes.NewCipher(key)
	if err != nil {
		return
	}
	decoded = make([]byte, len(encoded))
	decrypter = cipher.NewCBCDecrypter(c, iv)
	decrypter.CryptBlocks(decoded, encoded)
	decoded = PKCS5UnPadding(decoded)
	return
}

func PKCS5Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}
