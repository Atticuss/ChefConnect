package rest

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// SetLogger initializes the logging middleware.
// simplified from example in: https://github.com/gin-contrib/logger
func (restCtrl *restController) setLogger(config Config) gin.HandlerFunc {
	var sublog zerolog.Logger
	if config.Logger == nil {
		sublog = log.Logger
	} else {
		sublog = *config.Logger
	}

	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		if raw != "" {
			path = path + "?" + raw
		}

		c.Next()
		end := time.Now()
		latency := end.Sub(start)
		if config.UTC {
			end = end.UTC()
		}

		msg := "Request"
		if len(c.Errors) > 0 {
			msg = c.Errors.String()
		}

		dumplogger := sublog.With().
			Int("status", c.Writer.Status()).
			Str("method", c.Request.Method).
			Str("path", path).
			Str("ip", c.ClientIP()).
			Dur("latency", latency).
			Str("user-agent", c.Request.UserAgent()).
			Logger()

		switch {
		case c.Writer.Status() >= http.StatusBadRequest && c.Writer.Status() < http.StatusInternalServerError:
			{
				dumplogger.Warn().
					Msg(msg)
			}
		case c.Writer.Status() >= http.StatusInternalServerError:
			{
				dumplogger.Error().
					Msg(msg)
			}
		default:
			dumplogger.Info().
				Msg(msg)
		}
	}
}
