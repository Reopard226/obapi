package hooks

import (
	"context"
	"github.com/twitchtv/twirp"
	"go.uber.org/zap"
	"strings"
	"time"
)

var logger *zap.Logger

var reqStartTimestampKey = new(int)

func markReqStart(ctx context.Context) context.Context {
	return context.WithValue(ctx, reqStartTimestampKey, time.Now())
}

func getReqStart(ctx context.Context) (time.Time, bool) {
	t, ok := ctx.Value(reqStartTimestampKey).(time.Time)
	return t, ok
}

func NewLoggerHooks() *twirp.ServerHooks {
	logger, _ = zap.NewDevelopment()
	defer logger.Sync()
	sugar := logger.Sugar()

	hooks := &twirp.ServerHooks{}
	// RequestReceived: inc twirp.total.req_recv
	hooks.RequestReceived = func(ctx context.Context) (context.Context, error) {
		ctx = markReqStart(ctx)
		//stats.Inc("twirp.total.requests", 1, 1.0)
		sugar.Infow("Request received",
			// Structured context as loosely typed key-value pairs.
			"url", "T2etst",
			"attempt", 5,
			"backoff", time.Second,
		)

		return ctx, nil
	}

	// RequestRouted: inc twirp.<method>.req_recv
	hooks.RequestRouted = func(ctx context.Context) (context.Context, error) {
		_, ok := twirp.MethodName(ctx)
		if !ok {
			return ctx, nil
		}
		//stats.Inc("twirp."+sanitize(method)+".requests", 1, 1.0)
		return ctx, nil
	}

	// ResponseSent:
	// - inc twirp.total.responses
	// - inc twirp.<method>.responses
	// - inc twirp.status_codes.total.<code>
	// - inc twirp.status_codes.<method>.<code>
	// - time twirp.all_methods.response
	// - time twirp.<method>.response
	// - time twirp.status_codes.all_methods.<code>
	// - time twirp.status_codes.<method>.<code>
	hooks.ResponseSent = func(ctx context.Context) {
		// Three pieces of data to get, none are guaranteed to be present:
		// - time that the request started
		// - method that was called
		// - status code of response
		/*		var (
					start  time.Time
					method string
					status string

					haveStart  bool
					haveMethod bool
					haveStatus bool
				)

				start, haveStart = getReqStart(ctx)
				method, haveMethod = twirp.MethodName(ctx)
				status, haveStatus = twirp.StatusCode(ctx)

				method = sanitize(method)
				status = sanitize(status)

				stats.Inc("twirp.total.responses", 1, 1.0)

				if haveMethod {
					stats.Inc("twirp."+method+".responses", 1, 1.0)
				}
				if haveStatus {
					stats.Inc("twirp.status_codes.total."+status, 1, 1.0)
				}
				if haveMethod && haveStatus {
					stats.Inc("twirp.status_codes."+method+"."+status, 1, 1.0)
				}

				if haveStart {
					dur := time.Now().Sub(start)
					stats.TimingDuration("twirp.all_methods.response", dur, 1.0)

					if haveMethod {
						stats.TimingDuration("twirp."+method+".response", dur, 1.0)
					}
					if haveStatus {
						stats.TimingDuration("twirp.status_codes.all_methods."+status, dur, 1.0)
					}
					if haveMethod && haveStatus {
						stats.TimingDuration("twirp.status_codes."+method+"."+status, dur, 1.0)
					}
				}*/
	}
	return hooks
}

func sanitize(s string) string {
	return strings.Map(sanitizeRune, s)
}

func sanitizeRune(r rune) rune {
	switch {
	case 'a' <= r && r <= 'z':
		return r
	case '0' <= r && r <= '9':
		return r
	case 'A' <= r && r <= 'Z':
		return r
	default:
		return '_'
	}
}
