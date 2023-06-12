package helper

import (
	"context"
	"errors"
	"fmt"
	jwtv5 "github.com/golang-jwt/jwt/v5"
	"strconv"
	"time"
)

type MyCustomClaims struct {
	Data interface{}
	jwtv5.RegisteredClaims
}
type Hjwt struct {
	Key []byte `json:"key"`
}

const (
	bearerWord string = "Bearer"

	// bearerFormat authorization token format
	bearerFormat string = "Bearer %s"

	// authorizationKey holds the key used to store the JWT Token in the request tokenHeader.
	authorizationKey string = "Authorization"

	// reason holds the error reason.
	reason string = "UNAUTHORIZED"
)

func NewJwt(key string) *Hjwt {
	return &Hjwt{Key: []byte(key)}
}
func (h *Hjwt) GetToken(c context.Context, id int64, data interface{}) (tokenStr string, err error) {
	// Create claims with multiple fields populated
	claims := MyCustomClaims{
		Data: data,
		RegisteredClaims: jwtv5.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwtv5.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwtv5.NewNumericDate(time.Now()),
			NotBefore: jwtv5.NewNumericDate(time.Now()),
			Issuer:    "example-kratos",
			Subject:   "study",
			Audience:  []string{"somebody_else"},
			ID:        strconv.FormatInt(id, 10),
		},
	}

	token := jwtv5.NewWithClaims(jwtv5.SigningMethodHS256, claims)
	tokenStr, err = token.SignedString(h.Key)
	return
}

func (h *Hjwt) ParamToken(c context.Context, token string) (*MyCustomClaims, error) {
	tc, err := jwtv5.ParseWithClaims(token, &MyCustomClaims{}, func(token *jwtv5.Token) (interface{}, error) {
		return h.Key, nil
	}, jwtv5.WithLeeway(5*time.Second))

	if err != nil {
		return nil, err
	}

	if tc.Valid {
		d, ok := tc.Claims.(*MyCustomClaims)
		if !ok {
			return nil, err
		}
		return d, err
	} else if errors.Is(err, jwtv5.ErrTokenMalformed) {
		fmt.Println("That's not even a token")
		return nil, err
	} else if errors.Is(err, jwtv5.ErrTokenSignatureInvalid) {
		return nil, err
		// Invalid signature     fmt.Println("Invalid signature")
	} else if errors.Is(err, jwtv5.ErrTokenExpired) || errors.Is(err, jwtv5.ErrTokenNotValidYet) {
		// Token is either expired or not active yet
		fmt.Println("Timing is everything")
		return nil, err
	} else {
		fmt.Println("Couldn't handle this token:", err)
		return nil, err
	}
}

func (h *Hjwt) Refresh(c context.Context, token string, cl jwtv5.Claims) {
	// 解析JWT token，并获取其中的payload
	tc, err := jwtv5.Parse(token, func(token *jwtv5.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwtv5.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret_key"), nil
	})
	if err != nil {
		fmt.Println(err)
	}
	claims, ok := tc.Claims.(*MyCustomClaims)
	if !ok || !tc.Valid {
		fmt.Println("invalid token")
	}

	// 检查token是否需要刷新
	expTime := claims.ExpiresAt.Unix()

	if expTime-time.Now().Unix() > 30 { // 如果token距离过期还有30秒以上，则无需刷新
		fmt.Println("no need to refresh token")
		return
	}
	jwtR := claims.RegisteredClaims
	jwtR.ExpiresAt = jwtv5.NewNumericDate(time.Now().Add(24 * time.Hour))
	// jwtR.IssuedAt= jwtv5.NewNumericDate(time.Now())
	// jwtR.NotBefore= jwtv5.NewNumericDate(time.Now())
	// 生成新的JWT token并输出
	newToken := jwtv5.NewWithClaims(jwtv5.SigningMethodHS256, MyCustomClaims{
		Data:             claims.Data,
		RegisteredClaims: jwtR,
	})
	newTokenString, err := newToken.SignedString(h.Key)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(newTokenString)

}
