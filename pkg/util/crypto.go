/*
@Time : 2020/3/19 17:00
@Author : FB
@File : crypto
@Software: GoLand
*/
package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"net/url"
	"os"
	"strings"
)

// AES加密
func AesEncrypt(orig string, key string) string {
	var block cipher.Block
	var k []byte

	// 转成字节数组
	origData := []byte(orig)
	// 分组秘钥
	// NewCipher该函数限制了输入k的长度必须为16, 24或者32
	if strings.EqualFold(key, "") {
		// 默认使用192bit的秘钥
		k = GetDefault192BitKey()
		block, _ = aes.NewCipher(k)
	} else {
		k = []byte(key)
		block, _ = aes.NewCipher(k)
	}
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 补全码
	origData = pKCS7Padding(origData, blockSize)
	// 加密模式
	blockMode := cipher.NewCBCEncrypter(block, k[:blockSize])
	// 创建数组
	cryted := make([]byte, len(origData))
	// 加密
	blockMode.CryptBlocks(cryted, origData)
	return base64.StdEncoding.EncodeToString(cryted)
}

// AES解密
func AesDecrypt(cryted string, key string) string {
	var block cipher.Block
	var k []byte

	// 转成字节数组
	crytedByte, _ := base64.StdEncoding.DecodeString(cryted)
	// 分组秘钥
	// NewCipher该函数限制了输入k的长度必须为16, 24或者32
	if strings.EqualFold(key, "") {
		// 默认使用192bit的秘钥 24位
		k = GetDefault192BitKey()
		block, _ = aes.NewCipher(k)
	} else {
		k = []byte(key)
		block, _ = aes.NewCipher(k)
	}
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 加密模式
	blockMode := cipher.NewCBCDecrypter(block, k[:blockSize])
	// 创建数组
	orig := make([]byte, len(crytedByte))
	// 解密
	blockMode.CryptBlocks(orig, crytedByte)
	// 去补全码
	orig = pKCS7UnPadding(orig)
	return string(orig)
}

// 添加补码
// AES加密数据块分组长度必须为128bit(byte[16])，密钥长度可以是128bit(byte[16])、192bit(byte[24])、256bit(byte[32])中的任意一个。
func pKCS7Padding(ciphertext []byte, blocksize int) []byte {
	padding := blocksize - len(ciphertext)%blocksize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// 去除补码
func pKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// 获取128bit的秘钥
func GetDefault128BitKey() []byte {
	return []byte{0x2b, 0x7e, 0x15, 0x16, 0x28, 0xae, 0xd2, 0xa6, 0xab, 0xf7, 0x15, 0x88, 0x09, 0xcf, 0x4f, 0x3c}
}

// 获取192bit的秘钥
func GetDefault192BitKey() []byte {
	return []byte{
		0x8e, 0x73, 0xb0, 0xf7, 0xda, 0x0e, 0x64, 0x52, 0xc8, 0x10, 0xf3, 0x2b, 0x80, 0x90, 0x79, 0xe5,
		0x62, 0xf8, 0xea, 0xd2, 0x52, 0x2c, 0x6b, 0x7b,
	}
}

// 获取256bit的秘钥
func GetDefault256BitKey() []byte {
	return []byte{
		0x60, 0x3d, 0xeb, 0x10, 0x15, 0xca, 0x71, 0xbe, 0x2b, 0x73, 0xae, 0xf0, 0x85, 0x7d, 0x77, 0x81,
		0x1f, 0x35, 0x2c, 0x07, 0x3b, 0x61, 0x08, 0xd7, 0x2d, 0x98, 0x10, 0xa3, 0x09, 0x14, 0xdf, 0xf4,
	}
}

//**********Base64**********//
// base64解码
func Base64Decode(data string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(data)
}

// base64编码
func Base64EncodeByte(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// base64编码
func Base64Encode(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}

// MD5编码
func MD5Encode(data string) (sign string) {
	h := md5.New() // 初始化HASH对象
	h.Write([]byte(data))
	sign = hex.EncodeToString(h.Sum(nil))
	return
}

// MD5编码
func MD5EncodeByte(data []byte) (sign string) {
	h := md5.New() // 初始化HASH对象
	h.Write(data)
	sign = hex.EncodeToString(h.Sum(nil))
	return
}

var (
	private_path string
	public_path  string
)

// path 私钥和公钥的存储目录
// Key目录下如已有秘钥文件，则尽量不要调用该方法。
// 测试时生成秘钥文件时，请指定其他路径
func NewRSA(path ...string) {

	// 生成1024位的私钥
	pk, _ := rsa.GenerateKey(rand.Reader, 1024)
	keyOut, _ := os.Create(private_path)
	pem.Encode(keyOut, &pem.Block{Type: "RAS PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(pk)})
	keyOut.Close()

	// 根据私钥 生成公钥
	publicKey := &pk.PublicKey
	derPkix, _ := x509.MarshalPKIXPublicKey(publicKey)

	PublickeyOut, _ := os.Create(public_path)
	pem.Encode(PublickeyOut, &pem.Block{Type: "PUBLIC KEY", Bytes: derPkix})
	PublickeyOut.Close()
}

// 加密
// 1s可执行26000次
func RsaEncrypt(data, publicKey []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	// 解析PKIX格式的公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	// 加密
	return rsa.EncryptPKCS1v15(rand.Reader, pub, data)
}

// 解密
// 1s可执行1600次
func RsaDecrypt(data, privateKey []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	// 解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 解密
	return rsa.DecryptPKCS1v15(rand.Reader, priv, data)
}

// 获取秘钥
func GetPrivateKey() (data []byte, err error) {
	f, err := os.Open(private_path)
	if err != nil {
		return
	}
	data, err = ioutil.ReadAll(f)

	return
}

// 获取公钥
func GetPublicKey() (data []byte, err error) {
	f, err := os.Open(public_path)
	if err != nil {
		return
	}
	data, err = ioutil.ReadAll(f)

	return
}

// 解密
func RSADecode(data string) (str string, err error) {
	if strings.Contains(data, "%") {
		data, err = url.QueryUnescape(data)
		if err != nil {
			err = errors.New("请求参数错误！")
			return
		}
	}

	crypt, err := Base64Decode(data)
	if err != nil {
		err = errors.New("请求参数错误！")
		return
	}
	private, err := GetPrivateKey()
	if err != nil {
		err = errors.New("服务器错误！")
		return
	}
	b, err := RsaDecrypt([]byte(crypt), private)
	if err != nil {
		err = errors.New("请求参数错误！")
		return
	}

	str = string(b)

	return
}

// 加密
func RSAEncode(data string) (str string, err error) {
	public, err := GetPublicKey()
	if err != nil {
		err = errors.New("服务器错误！")
		return
	}
	ResourceID, err := RsaEncrypt([]byte(data), public)
	if err != nil {
		err = errors.New("服务器错误！")
		return
	}
	tmp := base64.StdEncoding.EncodeToString(ResourceID)
	str = url.QueryEscape(tmp)

	return
}

// SHA256编码
func SHA256(data []byte) string {
	h := sha256.New()
	h.Write(data)
	return string(h.Sum(nil))
}
