package api_test

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/antiphp/gonew/api"
	"github.com/hamba/logger/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestServer_HandleFoobar(t *testing.T) {
	tests := []struct {
		name       string
		req        string
		msg        string
		resp       string
		err        error
		wantStatus int
		wantResp   string
	}{
		{
			name:       "handles success",
			req:        `{"input":"foobar"}`,
			msg:        "foobar",
			resp:       "foobar",
			wantResp:   `{"output":"foobar"}`,
			wantStatus: http.StatusOK,
		},
		{
			name:       "handles error",
			req:        `{"input":"foobar"}`,
			msg:        "foobar",
			err:        errors.New("test"),
			wantResp:   `{"errno":500,"error":"could not foobar"}`,
			wantStatus: http.StatusInternalServerError,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			mf := &mockFoobar{}
			mf.On("Foobar", test.msg).Return(test.resp, test.err).Once()

			srvUrl := setupTestServer(t, mf)

			req, err := http.NewRequest(http.MethodPost, srvUrl+"/foobar", strings.NewReader(test.req))
			require.NoError(t, err)

			resp, err := http.DefaultClient.Do(req)
			require.NoError(t, err)
			t.Cleanup(func() { _ = resp.Body.Close() })

			assert.Equal(t, test.wantStatus, resp.StatusCode)

			b, err := io.ReadAll(resp.Body)
			require.NoError(t, err)

			assert.JSONEq(t, test.wantResp, string(b))

			mf.AssertExpectations(t)
		})
	}
}

func setupTestServer(t *testing.T, app api.Foobar) string {
	t.Helper()

	apiSrv := api.NewServer(app, logger.New(io.Discard, logger.LogfmtFormat(), logger.Debug))

	server := httptest.NewServer(apiSrv)
	t.Cleanup(server.Close)

	return server.URL
}

type mockFoobar struct {
	mock.Mock
}

func (m *mockFoobar) Foobar(_ context.Context, msg string) (string, error) {
	args := m.Called(msg)
	return args.Get(0).(string), args.Error(1)
}
