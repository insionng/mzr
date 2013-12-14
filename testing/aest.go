package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"os"
)

var publicKey = []byte{0x06, 0x07, 0x01, 0x02, 0x03, 0x04, 0x05, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0x08, 0x09, 0x0a, 0x00}

func main() {
	//需要去加密的字符串
	plaintext := []byte("My name is Insion")
	//如果传入加密串的话，plaint就是传入的字符串
	if len(os.Args) > 1 {
		plaintext = []byte(os.Args[1])
	}

	//aes的加密字符串
	key := "astaxie12798akljzmknm.ahkjkljl;k"
	if len(os.Args) > 2 {
		key = os.Args[2]
	}

	fmt.Println(len(key))

	// 创建加密算法aes
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		fmt.Printf("Error: NewCipher(%d bytes) = %s", len(key), err)
		os.Exit(-1)
	}

	//加密字符串
	cfb := cipher.NewCFBEncrypter(c, publicKey)
	ciphertext := make([]byte, len(plaintext))

	cfb.XORKeyStream(ciphertext, plaintext)
	s := fmt.Sprintf("%x", ciphertext)
	fmt.Println(s)

	// 解密字符串
	cfbdec := cipher.NewCFBDecrypter(c, publicKey)
	plaintextCopy := make([]byte, len(s))

	cfbdec.XORKeyStream(plaintextCopy, ciphertext)
	p := fmt.Sprintf("%s", plaintextCopy)
	fmt.Println(p)
}
