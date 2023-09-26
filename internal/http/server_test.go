package http_test

import (
	"context"
	"io"
	"net/http"
	"testing"
	"time"

	httpx "github.com/antiphp/gonew/internal/http"
	"github.com/hamba/testutils/retry"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestServer(t *testing.T) {
	var handlerCalled bool
	h := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		handlerCalled = true
	})

	srv := httpx.NewServer(context.Background(), ":55555", h)
	srv.Serve(func(err error) {
		require.NoError(t, err)
	})
	t.Cleanup(func() {
		_ = srv.Close()
	})

	var res *http.Response
	var err error
	retry.Run(t, func(t *retry.SubT) {
		res, err = http.DefaultClient.Get("http://localhost:55555/")
		require.NoError(t, err)
	})
	t.Cleanup(func() {
		_, _ = io.Copy(io.Discard, res.Body)
		_ = res.Body.Close()
	})

	err = srv.Shutdown(time.Second)
	require.NoError(t, err)

	assert.Equal(t, res.StatusCode, http.StatusOK)
	assert.True(t, handlerCalled)
}
