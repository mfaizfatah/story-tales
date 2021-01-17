package middleware

import (
	"encoding/base64"
	"log"
	"net/http"
	"strings"

	"github.com/go-redis/redis"
	"github.com/mfaizfatah/story-tales/app/adapter"
	"github.com/mfaizfatah/story-tales/app/helpers/logger"
)

// CheckSession middleware
func CheckSession(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		header := r.Header.Get("Authorization")
		auth := strings.Split(header, " ")

		// valueRedis, err := adapter.GetClientRedis().Get(auth[1]).Result()
		valueRedis, err := adapter.UseRedis().Get(auth[1]).Result()
		if err != nil || err == redis.Nil {
			Response(ctx, w, http.StatusUnauthorized, "Session invalid")
			return
		}

		if strings.ToLower(auth[0]) != "bearer" {
			Response(ctx, w, http.StatusForbidden, "Bearer Token Format Error")
			return
		}
		log.Printf("LOG DEBUG :: VALUE REDIS() => %v", valueRedis)
		decryptData, err := base64.StdEncoding.DecodeString(valueRedis)
		if err != nil {
			ctx = logger.Logf(ctx, "Error while decrypt header => %v", err)
			Response(ctx, w, http.StatusUnauthorized, "Wrong Authorization")
			return
		}

		ctx = logger.Logf(ctx, "authorization => %v", string(decryptData))
		ctx = logger.Device(ctx, decryptData)

		req := r.WithContext(ctx)
		req.Header.Set("x-app-token", string(decryptData))

		next.ServeHTTP(w, req)
	})
}
