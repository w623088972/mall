package middleware

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"myself/mall/conf"
	"myself/mall/errno"
	"myself/mall/redis"

	"github.com/beijibeijing/viper"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/goroom/rand"
	"github.com/sirupsen/logrus"
)

//Claims struct
type Claims struct {
	UserId    int
	UserIdStr string
	OToken    string
}

//AuthMiddleware  gin.HandlerFunc 处理需要登录验证的所有请求
func AuthMiddleware(group string) gin.HandlerFunc {
	return func(c *gin.Context) {
		language := c.MustGet("language").(string)
		// Parse the json web token.
		claims, err := ParseRequest(c)
		if err != nil {
			conf.SendResponse(c, errno.ErrTokenInvalid, "AuthMiddleware ParseRequest failed. err is "+err.Error(), nil, language)
			c.Abort()
			return
		}
		c.Set("userId", claims.UserId)

		//处理版本号
		//version := c.MustGet("version").(string)
		versionS := c.GetString("version")
		version, _ := strconv.Atoi(versionS)
		if version < viper.GetInt("version") {
			conf.LOG.Self.WithFields(logrus.Fields{
				"c.version":            version,
				"viper.GetInt.version": viper.GetInt("version"),
			}).Debug("AuthMiddleware")

			conf.SendResponse(c, errno.InternalServerError, "AuthMiddleware ParseRequest failed.", nil, language)
			c.Abort()
			return
		}

		if registerSource := c.Request.Header.Get("register_source"); registerSource != "" {
			c.Set("register_source", registerSource)
		}

		c.Next()
	}
}

// ParseRequest gets the token from the header and
// pass it to the Parse function to parses the token.
func ParseRequest(c *gin.Context) (*Claims, error) {
	//header := c.Request.Header.Get("Authorization")

	// Load the jwt secret from config
	secret := viper.GetString("jwt_secret")

	//if len(header) == 0 {
	//return &Claims{}, ErrMissingHeader
	//}

	var t, requestId string
	// Parse the header to get the token part.
	//fmt.Sscanf(header, "Bearer %s", &t)

	t = c.Request.Header.Get("token")
	if t == "" && c.Request.Method == http.MethodGet {
		t = c.Query("token")
	} else if t == "" && (c.Request.Method == http.MethodPost || c.Request.Method == http.MethodPut) {
		t = c.PostForm("token")
	}

	requestId = c.MustGet("requestId").(string)

	return Parse(t, secret, requestId)
}

// Parse validates the token with the specified secret,
// and returns the context if the token was valid.
func Parse(tokenString string, secret string, requestId string) (*Claims, error) {
	claims := &Claims{}

	// Parse the token.
	token, err := jwt.Parse(tokenString, secretFunc(secret))
	// Parse error.
	if err != nil {
		conf.LOG.Self.WithFields(logrus.Fields{
			"requestId":   requestId,
			"tokenString": tokenString,
		}).Debug("AuthMiddleware-ParseRequest-Parse-jwt.Parse")

		return claims, err
		// Read the token if it's valid.
	} else if tokenClaims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		claims.UserId = int(tokenClaims["userId"].(float64))
		claims.OToken = tokenClaims["oToken"].(string)
		_, ok := tokenClaims["userIdStr"]
		if ok {
			claims.UserIdStr = tokenClaims["userIdStr"].(string)
		} else {
			claims.UserIdStr = strconv.Itoa(int(tokenClaims["userId"].(float64)))
		}

		//验证是否失效
		redisKey := viper.GetString("project.name") + ":uid:jwtToken"
		oToken, err := redis.HGet(redisKey, claims.UserIdStr)
		if err != nil {
			conf.LOG.Self.WithFields(logrus.Fields{
				"requestId":   requestId,
				"tokenString": tokenString,
			}).Debug("AuthMiddleware-ParseRequest-Parse-redis.HGet")
			return claims, err
		}
		if oToken != claims.OToken {
			conf.LOG.Self.WithFields(logrus.Fields{
				"requestId":     requestId,
				"tokenString":   tokenString,
				"oToken":        oToken,
				"claims.OToken": claims.OToken,
			}).Debug("AuthMiddleware-ParseRequest-Parse oToken != claims.OToken")
			return claims, ErrUToken
		}

		return claims, nil

		// Other errors.
	} else {
		return claims, err
	}
}

//CreatToken 生成token Sign signs the context with the specified secret.
func CreatToken(userId int, loginType string, requestId string) (tokenString string, err error) {
	// Load the jwt secret from the Gin config if the secret isn't specified.
	secret := viper.GetString("jwt_secret")

	//随机生成验证码
	rd := rand.GetRand()
	randCode := rd.String(6, rand.RST_NUMBER)
	oToken := time.Now().Format("20060102150405") + randCode

	// The token content.
	userIDStr := loginType + strconv.Itoa(userId)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":    userId,
		"userIdStr": userIDStr,
		"loginType": loginType,
		"oToken":    oToken,
		"nbf":       time.Now().Unix(),
		"iat":       time.Now().Unix(),
	})

	//存储uToken
	redisKey := viper.GetString("projectName") + ":uid:jwtToken"
	err = redis.HSet(redisKey, userIDStr, oToken)
	if err != nil {
		return
	}

	// Sign the token with the specified secret.
	tokenString, err = token.SignedString([]byte(secret))

	conf.LOG.Self.WithFields(logrus.Fields{
		"requestId":   requestId,
		"userId":      userId,
		"userIdStr":   userIDStr,
		"loginType":   loginType,
		"oToken":      oToken,
		"tokenString": tokenString,
	}).Trace("CreatToken")

	return
}

//jwt 方法
var (
	// ErrMissingHeader means the `Authorization` header was empty.
	ErrMissingHeader = errors.New("The length of the `Authorization` header is zero")
	ErrUToken        = errors.New("The token not match")
)

// secretFunc validates the secret format.
func secretFunc(secret string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		// Make sure the `alg` is what we except.
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(secret), nil
	}
}

func IsValidToken(host, token, secret string) (bool, error) {
	_, err := Parse(host, token, secret)
	if err != nil {
		return false, err
	}
	return true, nil
}
