package main

import (
	"crypto/aes"
	"crypto/cipher"

	"encoding/hex"
	"fmt"
	"github.com/insionng/veryhour/libs"
	"github.com/insionng/veryhour/utils"
)

func main() {
	//需要被加密的字符串
	plaintext := "哈哈肯定是假的！要知道，rì本的风俗是有了地位之后，都会换个威风的名字，甚至连姓氏都会换掉，小五郎这个名字没问题，可放在一个能率领上千骑兵的家老哈哈我成功啦功啦!哈哈我成功啦!哈哈我成功啦!哈哈我成功啦!哈哈我成功啦!"

	ae, e := AesEncrypt(plaintext, libs.AesKey, libs.AesPublicKey)
	ad, e2 := AesDecrypt(ae, libs.AesKey, libs.AesPublicKey)
	fmt.Println(e, e2)
	fmt.Println("---------------")
	fmt.Println("AE:", plaintext)
	fmt.Println("DE:", ad)

	rsa, e3 := utils.RsaEncrypt([]byte(libs.AesKey), libs.RsaPublicKey)
	fmt.Println(rsa, e3)
	y := fmt.Sprintf("%x", rsa)
	x, errr := hex.DecodeString(y)
	fmt.Println(errr)
	fmt.Println("YY>>>>>>>", y)
	fmt.Println("XX>>>>>>>", x)
	rsa2, e4 := utils.RsaDecrypt(x, libs.RsaPrivateKey)
	fmt.Println(">>>>>>>", rsa2, e4, "<<<<<<<<<<")
	fmt.Println("ssss>>>>", string(rsa2), "<<<<ssss")

}

func AesEncrypt(content string, privateKey string, publicKey string) (string, error) {

	if c, err := aes.NewCipher([]byte(privateKey)); err != nil {
		//fmt.Printf("Error: NewCipher(%d bytes) = %s", len(privateKey), err)
		return "", err
	} else {

		//加密字符串
		cfb := cipher.NewCFBEncrypter(c, []byte(publicKey))
		ciphertext := make([]byte, len(content))
		cfb.XORKeyStream(ciphertext, []byte(content))

		//s := fmt.Sprintf("%x", ciphertext)
		//fmt.Println(s)
		return string(ciphertext), err
	}

}

func AesDecrypt(ciphertext string, privateKey string, publicKey string) (string, error) {

	if c, err := aes.NewCipher([]byte(privateKey)); err != nil {
		//fmt.Printf("Error: NewCipher(%d bytes) = %s", len(privateKey), err)
		return "", err
	} else {
		cipherz := []byte(ciphertext)
		// 解密字符串
		cfbdec := cipher.NewCFBDecrypter(c, []byte(publicKey))
		contentCopy := make([]byte, len(cipherz))
		cfbdec.XORKeyStream(contentCopy, cipherz)

		//p := fmt.Sprintf("%s", contentCopy)
		//fmt.Println(p)
		return string(contentCopy), err
	}
}
