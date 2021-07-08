package common

import (
	mycrypto "github.com/node1650665999/Glib/crypto"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

//Header 定义了JWT的头部信息,由两部分组成：加密的算法和类型
type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

//NewHeader 实例化Header
func NewHeader(alg, typ string) *Header {
	return &Header{Alg: alg, Typ: typ}
}

//PayLoad 定义了JWT中的有效信息
type PayLoad struct {
	//iss: jwt签发者
	Iss string `json:"iss"`

	//sub: jwt所面向的用户
	Sub string `json:"sub"`

	//aud: 接收jwt的一方
	Aud string `json:"aud"`

	//exp: jwt的过期时间，这个过期时间必须要大于签发时间(时间戳)
	Exp int64 `json:"exp"`

	//nbf: 定义在什么时间之前,该jwt都是不可用的(时间戳)
	Nbf int64 `json:"nbf"`

	//iat: jwt的签发时间(时间戳)
	Iat int64 `json:"iat"`

	//jti: jwt的唯一身份标识，主要用来作为一次性token,从而回避重放攻击。
	Jti string `json:"jti"`
}
//NewPayLoad 实例化一个PayLoad对象,他存储四个参数:过期时间、签发时间、不可用时间、token
func NewPayLoad(exp int64, iat int64, nbf int64,jti string) *PayLoad {
	return &PayLoad{Exp: exp, Iat: iat, Jti: jti}
}

var hmacKey = "MMW4n4slID"

//GenerateToken  生成一个token
func GenerateToken(expire int) string{
	//初始化header和payload
	currentTime := time.Now().Unix()
	exp := currentTime + int64(expire)
	jti := StrUuid(30)
	h := NewHeader("md5", "jwt")
	p := NewPayLoad(exp,currentTime,currentTime,jti)

	//签名
	headerAndPayLoad, sign := makeSign(h, p)
	return headerAndPayLoad + "." + sign
}

//VerifyToken 用来验证token是否合法
func VerifyToken(token string) error  {

	currentTime := time.Now().Unix()

	tokens := strings.Split(token, ".")
	if len(tokens) != 3 {
		return fmt.Errorf("token length is invalid")
	}

	h, err  := mycrypto.DecodeByBase64(tokens[0])
	if err != nil {
		return  err
	}

	p, err := mycrypto.DecodeByBase64(tokens[1])
	if err != nil {
		return  err
	}

	var header Header
	var payload PayLoad

	sign             := tokens[2]
	json.Unmarshal([]byte(h), &header)
	json.Unmarshal([]byte(p), &payload)

	if payload.Iat > currentTime {
		return fmt.Errorf("签发时间大于当前时间")
	}

	fmt.Println(payload.Exp, currentTime)

	if payload.Exp < currentTime {
		return fmt.Errorf("token 已过期")
	}

	if payload.Nbf > currentTime {
		return fmt.Errorf("token 还未生效")
	}

	_, sg := makeSign(&header, &payload)
	if sign != sg {
		return fmt.Errorf("签名错误")
	}

	return nil
}

//makeSign 用来生成签名信息
func makeSign(h *Header, p *PayLoad) (hp string, sign string) {
	header,_ := json.Marshal(h)
	payLoad,_:= json.Marshal(p)

	headerAndPayLoad := mycrypto.EncodeByBase64(string(header)) + "." +  mycrypto.EncodeByBase64(string(payLoad))

	return headerAndPayLoad, mycrypto.HashByHmac(headerAndPayLoad, hmacKey)
}

