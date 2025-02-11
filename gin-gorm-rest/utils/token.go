package utils

import (
	"context"
	"errors"
	"fmt"
	"time"
	"vietanh/gin-gorm-rest/config"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func RefreshToken(c *gin.Context) {

}

func GenerateToken(ttl time.Duration, payload interface{}, secretJWTkey string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	now := time.Now().UTC()
	claim := token.Claims.(jwt.MapClaims)

	claim["sub"] = payload
	claim["exp"] = now.Add(ttl).Unix()
	claim["iat"] = now.Unix()
	claim["nbf"] = now.Unix()

	tokenString, err := token.SignedString([]byte(secretJWTkey))

	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}
	return tokenString, nil
}

func GenerateRefreshToken(timestamp int64, payload interface{}, secretJWTkey string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	now := time.Now().UTC()
	// Kiểm tra nếu timestamp nhỏ hơn thời gian hiện tại -> lỗi
	if timestamp <= now.Unix() {
		return "", fmt.Errorf("invalid expiration timestamp: must be in the future")
	}
	claim := token.Claims.(jwt.MapClaims)

	claim["sub"] = payload
	claim["exp"] = timestamp
	claim["iat"] = now.Unix()
	claim["nbf"] = now.Unix()

	tokenString, err := token.SignedString([]byte(secretJWTkey))

	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}
	return tokenString, nil
}
func ValidateRefreshToken(token string, signedJWTKey string) (interface{}, interface{}, error) {
	// 2️⃣ Giải mã token
	tkn, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", jwtToken.Header["alg"])
		}
		return []byte(signedJWTKey), nil
	})
	if err != nil {
		return nil, nil, fmt.Errorf("invalid token: %w", err)
	}

	// 3️⃣ Lấy claims từ token
	claims, ok := tkn.Claims.(jwt.MapClaims)
	if !ok || !tkn.Valid {
		return nil, nil, fmt.Errorf("invalid token claim")
	}

	// Trả về cả "sub" và "exp" nếu có
	if sub, exists := claims["sub"]; exists {
		if exp, expExists := claims["exp"]; expExists {
			return sub, exp, nil
		}
		return sub, nil, errors.New("exp not found in token")
	}
	// 5️⃣ trả về
	return nil, nil, fmt.Errorf("token does not contain subject")
}

func ValidateToken(token string, signedJWTKey string) (interface{}, error) {
	ctx := context.Background()

	// 🔹 Kiểm tra token trong Redis (Them tiền tố trước)
	redisToken := "Bearer " + token

	// 1️⃣ Kiểm tra token có bị thu hồi không trong Redis
	exists, err := config.RedisClient.Exists(ctx, redisToken).Result()
	if err != nil {
		return nil, fmt.Errorf("redis error: %w", err)
	}
	if exists > 0 { // Nếu token có trong Redis, nghĩa là nó đã bị thu hồi
		return nil, fmt.Errorf("token has been revoked")
	}

	// 2️⃣ Giải mã token
	tkn, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", jwtToken.Header["alg"])
		}
		return []byte(signedJWTKey), nil
	})
	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	// 3️⃣ Lấy claims từ token
	claims, ok := tkn.Claims.(jwt.MapClaims)
	if !ok || !tkn.Valid {
		return nil, fmt.Errorf("invalid token claim")
	}

	// 4️⃣ Trả về "sub" nếu có
	if sub, exists := claims["sub"]; exists {
		return sub, nil
	}
	return nil, fmt.Errorf("token does not contain subject")
}
