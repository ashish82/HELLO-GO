package logger

import (
	"HELLO-GO/constant"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mattn/go-isatty"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"time"
)

var Logger *logrus.Logger
var (
	green        = string([]byte{27, 91, 57, 55, 59, 52, 50, 109})
	white        = string([]byte{27, 91, 57, 48, 59, 52, 55, 109})
	yellow       = string([]byte{27, 91, 57, 55, 59, 52, 51, 109})
	red          = string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
	blue         = string([]byte{27, 91, 57, 55, 59, 52, 52, 109})
	magenta      = string([]byte{27, 91, 57, 55, 59, 52, 53, 109})
	cyan         = string([]byte{27, 91, 57, 55, 59, 52, 54, 109})
	reset        = string([]byte{27, 91, 48, 109})
	disableColor = false
)

type PlainFormatter struct {
	TimestampFormat string
}

func (f *PlainFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := fmt.Sprintf(entry.Time.Format(f.TimestampFormat))
	return []byte(fmt.Sprintf("%s %s %s\n", timestamp, entry.Level, entry.Message)), nil
}

func GetLogger() *logrus.Logger {
	if Logger == nil {
		f, err := os.OpenFile(constant.LogFilePath+"hello.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
		if err != nil {
			fmt.Printf("error opening file: %v", err)
		}
		Logger = logrus.New()
		Logger.SetNoLock()
		Logger.SetOutput(f)
		Logger.SetLevel(logrus.InfoLevel)
		pf := new(PlainFormatter)
		pf.TimestampFormat = "2006-01-02 15:04:05"
		Logger.SetFormatter(pf)
		Logger.SetReportCaller(true)

	}
	return Logger
}

func LoggerWithWriter(out io.Writer, notlogged ...string) gin.HandlerFunc {
	isTerm := true

	if w, ok := out.(*os.File); !ok ||
		(os.Getenv("TERM") == "dumb" || (!isatty.IsTerminal(w.Fd()) && !isatty.IsCygwinTerminal(w.Fd()))) ||
		disableColor {
		isTerm = false
	}

	var skip map[string]struct{}

	if length := len(notlogged); length > 0 {
		skip = make(map[string]struct{}, length)

		for _, path := range notlogged {
			skip[path] = struct{}{}
		}
	}

	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Log only when path is not being skipped
		if _, ok := skip[path]; !ok {
			// Stop timer
			end := time.Now()
			latency := end.Sub(start)

			latencyInMilliSec := latency.Seconds() * 1000.0

			clientIP := c.ClientIP()
			method := c.Request.Method
			statusCode := c.Writer.Status()
			var statusColor, methodColor, resetColor string
			if isTerm {
				statusColor = colorForStatus(statusCode)
				methodColor = colorForMethod(method)
				resetColor = reset
			}

			if raw != "" {
				path = path + "?" + raw
			}

			fmt.Fprintf(out, "%s - - [%v] \"%s%s%s %s HTTP/1.1\" %s%d%s 100 %.2f %.2f\n",
				clientIP,
				end.Format("02/Jan/2006:15:04:05 -0700"),
				methodColor, method, resetColor,
				path,
				// HTTP string
				statusColor, statusCode, resetColor,
				//bytes
				latencyInMilliSec,
				latencyInMilliSec, //This is supposed to be the time to commit the response
			)
		}
	}
}

func colorForStatus(code int) string {
	switch {
	case code >= http.StatusOK && code < http.StatusMultipleChoices:
		return green
	case code >= http.StatusMultipleChoices && code < http.StatusBadRequest:
		return white
	case code >= http.StatusBadRequest && code < http.StatusInternalServerError:
		return yellow
	default:
		return red
	}
}

func colorForMethod(method string) string {
	switch method {
	case "GET":
		return blue
	case "POST":
		return cyan
	case "PUT":
		return yellow
	case "DELETE":
		return red
	case "PATCH":
		return green
	case "HEAD":
		return magenta
	case "OPTIONS":
		return white
	default:
		return reset
	}
}
