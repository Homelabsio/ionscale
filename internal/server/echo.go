package server

import (
	"fmt"
	"github.com/hashicorp/go-hclog"
	"github.com/labstack/echo/v4"
	"runtime"
	"time"
)

func EchoLogger(logger hclog.Logger) echo.MiddlewareFunc {
	httpLogger := logger.Named("http")
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			if !httpLogger.IsTrace() {
				return next(c)
			}

			request := c.Request()
			response := c.Response()
			start := time.Now()
			if err = next(c); err != nil {
				c.Error(err)
			}

			httpLogger.Trace("finished server http call",
				"http.code", response.Status,
				"http.method", request.Method,
				"http.uri", request.RequestURI,
				"http.start_time", start.Format(time.RFC3339),
				"http.duration", time.Since(start))

			return
		}
	}
}

func EchoRecover(logger hclog.Logger) echo.MiddlewareFunc {
	httpLogger := logger.Named("http")
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			defer func() {
				if r := recover(); r != nil {
					err, ok := r.(error)
					if !ok {
						err = fmt.Errorf("%v", r)
					}
					stack := make([]byte, 4<<10) // 4 KB
					length := runtime.Stack(stack, false)
					httpLogger.Error("panic handling request", "err", err, "stack", string(stack[:length]))
					c.Error(err)
				}
			}()
			return next(c)
		}
	}
}
