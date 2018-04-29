package cxtgo

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/ratelimit"
)

func TestBaseOpts(t *testing.T) {
	r := require.New(t)

	params := Params{"t": "t"}
	raw := true
	name := "test"
	id := "10"
	userAgent := "test"
	rl := ratelimit.New(2)
	key := "key"
	secret := "secret"
	toggle := true
	w := os.Stdout

	opts := []BaseOpt{
		WithCustomParams(params),
		WithIncludeRaw(raw),
		WithName(name),
		WithRatelimit(rl),
		WithID(id),
		WithUserAgent(userAgent),
		WithAPIKey(key),
		WithAPISecret(secret),
		WithDebug(toggle),
		WithDebuglogger(w),
	}
	base := NewBase(opts...)
	r.Equal(params, base.CustomParams)
	r.Equal(raw, base.Raw)
	r.Equal(name, base.Name)
	r.Equal(id, base.ID)
	r.Equal(userAgent, base.UserAgent)
	r.Equal(rl, base.Ratelimit)
	r.Equal(key, base.APIKEY)
	r.Equal(secret, base.APISecret)
	r.Equal(toggle, base.Debug)
	r.Equal(w, base.DebugLog)
}
