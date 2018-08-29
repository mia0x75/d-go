package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

// 服务器密码加密方案，因为需要解密进行连接，所有使用双向加密
// key := []byte("EBvB4DnNLIVWSgARns6Y5anaHqZ9J6nK")
// plaintext := []byte("Chan9eme!bah@")
// ciphertext, _ := hex.DecodeString("da0b932fa0ac44da6e0d993388f424f2b5b29e2fc0e915024ab3c63f1adf9940708df1cfd4813970f91f5f64233fa292f75f936ce66da2c52115ff449c1deefc")
// c, _ := utils.EncryptCBC(key, plaintext)
// p, _ := utils.DecryptCBC(key, ciphertext)
// fmt.Println(hex.EncodeToString(c))
// fmt.Println(string(p))

// 用户密码加密方案，因为只需要验证，所以采用类似哈希的方式
// "golang.org/x/crypto/scrypt"
// salt := []byte{0xc8, 0x28, 0xf2, 0x58, 0xa7, 0x6a, 0xad, 0x7b}
// dk, _ := scrypt.Key([]byte("1234567890abcedfghij"), salt, 1<<15, 8, 1, 32)
// fmt.Printf("%d - %v", len(dk), dk)

func PKCS7Pad(data []byte) []byte {
	padding := aes.BlockSize - len(data)%aes.BlockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padtext...)
}

func PKCS7UPad(data []byte) []byte {
	padLength := int(data[len(data)-1])
	return data[:len(data)-padLength]
}

func EncryptCBC(plaintext, key []byte) ([]byte, error) {
	plaintext = PKCS7Pad(plaintext)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	cbc := cipher.NewCBCEncrypter(block, iv)
	cbc.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)
	return ciphertext, nil
}

func DecryptCBC(ciphertext, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(ciphertext) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	cbc := cipher.NewCBCDecrypter(block, iv)
	cbc.CryptBlocks(ciphertext, ciphertext)
	plaintext := ciphertext
	return PKCS7UPad(plaintext), nil
}
