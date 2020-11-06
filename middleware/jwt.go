package middleware

import (
	"errors"
	"log"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	rd "github.com/goroom/rand"
	"github.com/spf13/viper"

	"myself/mall/errno"
	"myself/mall/handler"
	redisDb "myself/mall/redis"
)

type Claims struct {
	UserId    int
	UserIdStr string
	UserToken string
}

func AuthMiddleware(group string) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := ParseRequest(c)
		if err != nil {
			handler.SendResponse(c, errno.ErrTokenInvalid, "AuthMiddleware ParseRequest failed. err is "+err.Error(), nil)
			c.Abort()
			return
		}
		c.Set("user_id", claims.UserId)
		source := c.Request.Header.Get("register_source")
		c.Set("register_source", source)

		c.Next()
	}
}

var (
	// ErrMissingHeader means the `Authorization` header was empty.
	ErrMissingHeader = errors.New("The length of the `Authorization` header is zero.")
	ErrUserToken     = errors.New("The token not match.")
)

// secretFunc validates the secret format.
func secretFunc(secret string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
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

// Parse validates the token with the specified secret,
// and returns the context if the token was valid.
func Parse(host, tokenString, secret string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.Parse(tokenString, secretFunc(secret))
	if err != nil {
		return claims, err
	}
	tokenClaims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("")
	}
	claims.UserId = int(tokenClaims["userId"].(float64))
	claims.UserToken = tokenClaims["userToken"].(string)
	_, ok = tokenClaims["userIdStr"]
	if ok {
		claims.UserIdStr = tokenClaims["userIdStr"].(string)
	} else {
		claims.UserIdStr = strconv.Itoa(int(tokenClaims["userId"].(float64)))
	}
	log.Println("claims", claims)

	rc := redisDb.RedisClient.Self.Get()
	defer func() {
		_ = rc.Close()
		log.Println("token Parse end redis ActiveCount:", redisDb.RedisClient.Self.ActiveCount())
	}()
	log.Println("token Parse start redis ActiveCount:", redisDb.RedisClient.Self.ActiveCount())

	//获取uToken
	redisKey := host + ":userId:jwtToken"
	userToken, err := redis.String(rc.Do("HGET", redisKey, claims.UserIdStr))

	if err != nil {
		return claims, err
	}
	if userToken != claims.UserToken {
		log.Println("redisKey", redisKey)
		log.Println("claims.UserId", claims.UserId)
		log.Println("claims.UserIdStr", claims.UserIdStr)
		log.Println("claims.UserToken", claims.UserToken)
		log.Println("userToken", userToken)
		return claims, ErrUserToken
	}
	return claims, nil
}

// ParseRequest gets the token from the header and
// pass it to the Parse function to parses the token.
func ParseRequest(c *gin.Context) (*Claims, error) {
	jwtSecret := viper.GetString("jwtSecretKey")
	t := c.Request.Header.Get("token")
	host := c.Request.Host
	log.Println("ParseRequest token:", t)

	return Parse(host, t, jwtSecret)
}

//NewToken 生成新的 jwt token
func NewToken(userId int, host, loginType string) (tokenString string, err error) {
	// Load the jwt secret from the Gin config if the secret isn't specified.
	jwtSecret := viper.GetString("jwtSecretKey")

	rand := rd.GetRand()
	randCode := rand.String(6, rd.RST_NUMBER)
	userToken := time.Now().Format("20060102150405") + randCode

	//The token content.
	userIdStr := loginType + strconv.Itoa(userId)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":    userId,
		"userIdStr": userIdStr,
		"loginType": loginType,
		"userToken": userToken,
		"nbf":       time.Now().Unix(),
		"iat":       time.Now().Unix(),
	})

	rc := redisDb.RedisClient.Self.Get()
	defer func() {
		_ = rc.Close()
		log.Println("NewToken end redis ActiveCount:", redisDb.RedisClient.Self.ActiveCount())
	}()
	log.Println("NewToken start redis ActiveCount:", redisDb.RedisClient.Self.ActiveCount())

	redisKey := host + ":userId:jwtToken"
	_, err = rc.Do("HSET", redisKey, userIdStr, userToken)
	if err != nil {
		return
	}

	// Sign the token with the specified secret.
	tokenString, err = token.SignedString([]byte(jwtSecret))
	log.Println("NewToken userId:", userId)
	log.Println("NewToken userIdStr:", userIdStr)
	log.Println("NewToken loginType:", loginType)
	log.Println("NewToken userToken:", userToken)
	log.Println("NewToken token:", tokenString)

	return
}
