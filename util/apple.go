package util

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"net/http"

	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
	"github.com/lestrrat-go/jwx/jwk"
)

//解析identity_token的方法来验证
func ParseAppleToken(appleIdentityToken string) (*jwt.Token, error) {
	set, err := jwk.FetchHTTP("https://appleid.apple.com/auth/keys", jwk.WithHTTPClient(http.DefaultClient))
	if err != nil {
		return nil, err
	}

	var isSuccess bool
	var token *jwt.Token

	//苹果这个秘钥有点坑，返回了多个公钥，要对每一个公钥都去进行解析，有一个成功了就行，全失败才算失败，一开始项目只用了其中一个公钥，运气好每次都成功解析出来了，直到上生产环境才测出来这个bug
	for _, key := range set.Keys {
		pubKeyIfAce, _ := key.Materialize()
		pubKey, ok := pubKeyIfAce.(*rsa.PublicKey)
		if !ok {
			beego.Error(fmt.Errorf(`expected RSA public key from %s`, "https://appleid.apple.com/auth/keys"))
			continue
		}

		token, err = jwt.Parse(appleIdentityToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				beego.Error(fmt.Errorf("Unexpected signing method: %v", token.Header["alg"]))
			}
			return pubKey, nil
		})
		if err != nil {
			beego.Warning("Token Parse error:", err)
			continue
		}
		if !token.Valid {
			err = errors.New("token 无效")
			beego.Warning("token 无效")
			continue
		}
		isSuccess = true
		break
	}
	if isSuccess {
		return token, nil
	}

	return nil, errors.New("token 无效")
}
