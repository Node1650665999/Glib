package crypto

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io"
)

//HashByMd5 实现 md5 加密
func HashByMd5(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

//HashBySha1 实现sha1散列加密
func HashBySha1(str string) string {
	hash := sha1.Sum([]byte(str))
	return fmt.Sprintf("%x", hash)
}

//HashBySha256 实现sha256散列加密
func HashBySha256(str string) string {
	hash := sha256.New()
	io.WriteString(hash, str)
	//return fmt.Sprintf("%x", hash.Sum(nil))
	return hex.EncodeToString(hash.Sum(nil))
}

//HashByHmac 利用哈希算法，以一个密钥和一个消息为输入，生成一个消息摘要作为输出
func HashByHmac(str string, key string) string  {
	hash := hmac.New(md5.New, []byte(key))
	hash.Write([]byte(str))
	return hex.EncodeToString(hash.Sum([]byte("")))
}

//EncodeByBcrypt 密码加密
func EncodeByBcrypt(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

//DecodeByBcrypt 密码比对
func DecodeByBcrypt(password string, hashed string) (match bool, err error) {
	if err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password)); err != nil {
		return false, errors.New("密码比对错误！")
	}
	return true, nil
}

//EncodeByBase64 实现base64编码
func EncodeByBase64(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

//DecodeByBase64 实现base64解码
func DecodeByBase64(str string) (string, error) {
	strDecode,err := base64.StdEncoding.DecodeString(str)
	return string(strDecode), err
}

//UrlEncode 实现Url编码
func UrlEncode(url string) string {
	return base64.URLEncoding.EncodeToString([]byte(url))
}

//UrlDecode 实现Url解码
func UrlDecode(url string) (string,error) {
	urlDecode, err := base64.URLEncoding.DecodeString(url)
	if err == nil {
		return  string(urlDecode), nil
	}
	return "", err
}
