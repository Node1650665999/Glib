package crypto_test

import (
	"encoding/base64"
	"github.com/node1650665999/Glib/crypto"
	"testing"
)

/*func TestAES(t *testing.T) {
	name := "tcl"
	encodeStr := crypto.EncodeByAES(name)
	decodeStr := crypto.DecodeByAES(encodeStr)
	t.Log(encodeStr, decodeStr)
}*/

func TestRSA(t *testing.T) {
	// 私钥生成
	//openssl genrsa -out rsa_private_key.pem 1024
	var privateKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEA7+lJjUhzAWDTeQPSnH5BWSd8nEKTMyBJB/8YRH+eKL70lf9s
PXRyv+61ZGQt4DfSYDx8cy2NwbnXiUuDLqwgeGZ0MjVHQgWfe3J/9xxkkbIxXyPB
Pyz1uQOk8PPAAIJkEqEoIDkC9x9xW9YWBGQiwy+ayzLYiayAW96N6LTkiZZrpmrw
6R82e+pGb+AiENTd+l9Th44279cn25zX5tSYvwbYxZ2VWobTCGQe2q4aR2dNbYgd
BVOt2IF1GaOZMwScdfxxV9GkBOuHPC+8E/6+CRGEXm7ZqLzdkzjKv3lgRNBSY0nS
qRpgT7D5DMNfNfzgtYv2EIpX6kQwhkhgL9p3CwIDAQABAoIBADcLwtmM3v5Y9gyV
KPTJLztCiR/dUqLvbHJOQIYu9d4JelsUQQSUvGN3ZN1E8xW4GSgFmNRghl8FwgN5
dP73dXfKoiyG9vOaEK6lZeEP/a9EQHnA3W0eZr7trCGw+8PiJw3zNh62lgyXhU36
ABE/3I4GVTD8WJy2HLl3hf33y1wZ2pDNL/Ho+z9crWsr1iQ+PcETGoMaaAiJ3Bs/
58Avw74FeATR0IH/KT8tFjFmgJv8jxtsAqhFwfbAZ66N/jKWlJAIte/npDUSnCzh
NJbFQ4u2/URBN7+tklV1DrPGuSpZv8M11vsgltWLw6JVo1nv0YtGYzWF5yZ2ZeUN
qRCrp4kCgYEA/07WtA2KHFuKO7Y3y0n8RJQhfUqf/sxQ7D0FSyFdOhISfIplW5ye
k14Q3lCwVD235el9jKrArsbArmVC9nYyHl98iNJG+Y2ZQagmd6F63CEGBg7ZNTh7
4S7z5o41EB0dvU5vi/2CTiy8QgtsokSzZuK+iMGGoYtdMRD2M9PNjV0CgYEA8I/D
wgXoMmF6kfAQZ8zQsyb/T/GhvfzYmrNPkc8gpbkj/p/wwBqPpe1CsqWpXXbGwref
v+7Vw1rrNhkED0tFfdsp/DnJ/unQvm2Wh7ZTsn46Ofg9Xo3R/B5pA9gZ9OsoyGCi
3+A5f9u26YhkEt+Pun+jWczGGwxkNVEQ5g8e54cCgYBHSz+/hexkYNeoNwk7loyA
phD4COfG4k1SuvOIeGetOLC64HbPb1wE8Qaq3kNvMtDwvhQWPPSTmeLikFpzsqvq
OWXwWzAArh7267raO1iwsfQZqvnS19QYHOF1J47/0fGlFIsnv4IszGdB1ije42pp
t7XXQJuU7vL2KbNm46WJ7QKBgAqrpdhGYM1TS5eLmX6xNBSuRybppe4CeC0shPwH
vv/63WDfAVPUGckXZBz+giu2KAzdDkX6NxsqPkKxC2AOS6/Qd+VLPu2Cu5Km08WD
TeUd+kE2BKrcCZNwWeIkxMn7YFy7BJ5/mK1WNp/XP/EiX4K7RKioD6WFgDBpPyGl
TA6jAoGBAPSyB6JcYkTwgUl52bvx2wQohJDHupwnQqkuFRPiY2km/GVdytiQG5Df
/jHcR5RNHTxY4BwvswbmfYrKzqilhrHPXjZ2ayhW6f29Bisk9lViAsDY3o/YXN1O
INWLZUfDn9q2yfndc09DZGmumtfPqMiolwIpa09k3bKRE/e8L5fo
-----END RSA PRIVATE KEY-----
`)

	// 公钥: 根据私钥生成
	//openssl rsa -in private.key -pubout -out public.pem
	var publicKey = []byte(`
-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA7+lJjUhzAWDTeQPSnH5B
WSd8nEKTMyBJB/8YRH+eKL70lf9sPXRyv+61ZGQt4DfSYDx8cy2NwbnXiUuDLqwg
eGZ0MjVHQgWfe3J/9xxkkbIxXyPBPyz1uQOk8PPAAIJkEqEoIDkC9x9xW9YWBGQi
wy+ayzLYiayAW96N6LTkiZZrpmrw6R82e+pGb+AiENTd+l9Th44279cn25zX5tSY
vwbYxZ2VWobTCGQe2q4aR2dNbYgdBVOt2IF1GaOZMwScdfxxV9GkBOuHPC+8E/6+
CRGEXm7ZqLzdkzjKv3lgRNBSY0nSqRpgT7D5DMNfNfzgtYv2EIpX6kQwhkhgL9p3
CwIDAQAB
-----END PUBLIC KEY-----
`)
	data, err := crypto.EncodeByRsa([]byte("hello world"), publicKey)
	if err != nil {
		t.Fatalf("encode err: %v", err)
	}
	t.Logf("encrypt data : %v",base64.StdEncoding.EncodeToString(data))

	origData, err := crypto.DecodeByRsa(data, privateKey)
	if err != nil {
		t.Fatalf("decode err: %v", err)
	}
	t.Logf("origin data : %v", string(origData))
}