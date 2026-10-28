package correlation

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"github.com/gin-gonic/gin"
)

type correlationKey struct{}

const HeaderCorrelationID = "X-Correlation-ID"

func GenerateID() string {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return fmt.Sprintf("req-%d", randInt())
	}
	return hex.EncodeToString(bytes)
}

func randInt() int64 {
	return 1000000000
}

func WithCorrelationID(ctx context.Context, correlationID string) context.Context {
	return context.WithValue(ctx, correlationKey{}, correlationID)
}

func FromContext(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	if val, ok := ctx.Value(correlationKey{}).(string); ok {
		return val
	}
	return ""
}

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		corrID := c.GetHeader(HeaderCorrelationID)
		if corrID == "" {
			corrID = GenerateID()
		}

		c.Header(HeaderCorrelationID, corrID)

		ctx := WithCorrelationID(c.Request.Context(), corrID)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
