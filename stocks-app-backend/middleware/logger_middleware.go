package middlewares

import (
	"bytes"
	"encoding/json"
	"io"
	"time"

	"github.com/felipepalacio293/stocks-app/config"
	"github.com/gin-gonic/gin"
)

type ResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *ResponseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w *ResponseWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func LoggerMiddleware(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !cfg.EnableRequestLogs {
			c.Next()
			return
		}

		start := time.Now()

		var requestBody []byte
		if c.Request.Body != nil {
			requestBody, _ = io.ReadAll(c.Request.Body)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))
		}

		responseWriter := &ResponseWriter{
			ResponseWriter: c.Writer,
			body:           bytes.NewBufferString(""),
		}
		c.Writer = responseWriter

		c.Next()

		latency := time.Since(start)

		if cfg.Environment != "production" {
			logRequest := map[string]interface{}{
				"timestamp":    time.Now().Format(time.RFC3339),
				"method":       c.Request.Method,
				"path":         c.Request.URL.Path,
				"query":        c.Request.URL.RawQuery,
				"status":       c.Writer.Status(),
				"latency":      latency.String(),
				"request_body": string(requestBody),
				"response":     responseWriter.body.String(),
				"client_ip":    c.ClientIP(),
				"user_agent":   c.Request.UserAgent(),
			}
			logRequestJSON, _ := json.Marshal(gin.H{"request": logRequest})
			gin.DefaultWriter.Write([]byte(string(logRequestJSON) + "\n"))
		} else {
			logRequest := map[string]interface{}{
				"timestamp": time.Now().Format(time.RFC3339),
				"method":    c.Request.Method,
				"path":      c.Request.URL.Path,
				"status":    c.Writer.Status(),
				"latency":   latency.String(),
				"client_ip": c.ClientIP(),
			}
			logRequestJSON, _ := json.Marshal(gin.H{"request": logRequest})
			gin.DefaultWriter.Write([]byte(string(logRequestJSON) + "\n"))
		}
	}
}
