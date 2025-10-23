package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

type CustomResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *CustomResponseWriter) Write(data []byte) (n int, err error) {
	w.body.Write(data)
	return w.ResponseWriter.Write(data)
}

func LoggerMiddleware() gin.HandlerFunc {
	logPath := "internal/logs/http.log"

	logger := zerolog.New(&lumberjack.Logger{
		Filename:   logPath,
		MaxSize:    1, // megabytes
		MaxBackups: 3,
		MaxAge:     5,    //days
		Compress:   true, // disabled by default
		LocalTime:  true,
	}).With().Timestamp().Logger()

	return func(ctx *gin.Context) {
		start := time.Now()
		contentType := ctx.GetHeader("Content-Type")
		requestBody := make(map[string]any)
		var formFiles []map[string]any

		// multipart/form-data
		if strings.HasPrefix(contentType, "multipart/form-data") {
			if err := ctx.Request.ParseMultipartForm(32 << 20); err == nil && ctx.Request.MultipartForm != nil {
				// for value
				for key, values := range ctx.Request.MultipartForm.Value {
					if len(values) == 1 {
						requestBody[key] = values[0]
					} else {
						requestBody[key] = values
					}
				}

				// for file
				for field, files := range ctx.Request.MultipartForm.File {
					for _, file := range files {
						formFiles = append(formFiles, map[string]any{
							"field":        field,
							"filename":     file.Filename,
							"size":         formatFileSize(file.Size),
							"content_type": file.Header.Get("Content-Type"),
						})
					}
				}

				if len(formFiles) > 0 {
					requestBody["form_files"] = formFiles
				}
			}
			log.Println("multipart/form-data")
		} else {
			bodyRequest, err2 := io.ReadAll(ctx.Request.Body)
			if err2 != nil {
				logger.Error().Err(err2).Msg("Failed to read request body")
			}

			ctx.Request.Body = io.NopCloser(bytes.NewBuffer(bodyRequest))

			// application/json
			if strings.HasPrefix(contentType, "application/json") {
				_ = json.Unmarshal(bodyRequest, &requestBody)
			} else {
				// application/x-www-form-urlencoded
				values, _ := url.ParseQuery(string(bodyRequest))
				for key, value := range values {
					if len(value) > 0 {
						requestBody[key] = value[0]
					} else {
						requestBody[key] = value
					}
				}
			}
		}

		customWriter := &CustomResponseWriter{body: bytes.NewBufferString(""), ResponseWriter: ctx.Writer}

		ctx.Writer = customWriter

		ctx.Next()

		duration := time.Since(start)
		statusCode := ctx.Writer.Status()

		responseContentType := ctx.Writer.Header().Get("Content-Type")
		responseBodyRaw := customWriter.body.String()
		var responseBodyParsed interface{}

		if strings.HasPrefix(responseContentType, "image/") {
			responseBodyParsed = "[BINARY DATA]"
		} else if strings.HasPrefix(responseContentType, "application/json") ||
			strings.HasPrefix(strings.TrimSpace(responseBodyRaw), "{") ||
			strings.HasPrefix(strings.TrimSpace(responseBodyRaw), "[") {
			if err := json.Unmarshal([]byte(responseBodyRaw), &responseBodyParsed); err != nil {
				responseBodyParsed = responseBodyRaw
			}
		} else {
			responseBodyParsed = responseBodyRaw
		}

		log.Printf("%s", responseBodyRaw)

		logEvent := logger.Info()
		if statusCode >= 500 {
			logEvent = logger.Error()
		} else if statusCode >= 400 {
			logEvent = logger.Warn()
		}

		logEvent.Str("method", ctx.Request.Method).
			Str("path", ctx.Request.URL.Path).
			Str("ip", ctx.ClientIP()).
			Str("query", ctx.Request.URL.RawQuery).
			Str("user_agent", ctx.Request.UserAgent()).
			Str("referer", ctx.Request.Referer()).
			Str("proto", ctx.Request.Proto).
			Str("host", ctx.Request.Host).
			Str("remote_addr", ctx.ClientIP()).
			Str("request_uri", ctx.Request.RequestURI).
			Int64("content_length", ctx.Request.ContentLength).
			Interface("headers", ctx.Request.Header).
			Interface("request_body", requestBody).
			Int("status", ctx.Writer.Status()).
			Interface("response_body", responseBodyParsed).
			Int64("duration_ms", duration.Milliseconds()).
			Msg("HTTP request log")

	}
}

func formatFileSize(size int64) string {
	switch {
	case size >= 1<<20:
		return fmt.Sprintf("%.2f MB", float64(size)/(1<<20))
	case size >= 1<<10:
		return fmt.Sprintf("%.2f KB", float64(size)/(1<<10))
	default:
		return fmt.Sprintf("%d B", size)
	}
}
