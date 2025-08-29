package redirect_test

import (
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http/httptest"
	"shortli/internal/http-server/handlers/redirect/mocks"
	"shortli/internal/lib/api"
	"shortli/internal/lib/logger/handlers/slogdiscard"
	"testing"

	"shortli/internal/http-server/handlers/redirect"
)

func TestRedirectHandler(t *testing.T) {
	tests := []struct {
		name      string
		alias     string
		url       string
		respError string
		mockError error
	}{
		{
			name:  "Success",
			alias: "test_alias",
			url:   "https://example.com",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			urlGetterMock := mocks.NewURLGetter(t)

			if tt.respError == "" || tt.mockError != nil {
				urlGetterMock.On("GetURL", tt.alias).
					Return(tt.url, tt.mockError).Once()
			}

			r := chi.NewRouter()
			r.Get("/{alias}", redirect.New(slogdiscard.NewDiscardLogger(), urlGetterMock))

			ts := httptest.NewServer(r)
			defer ts.Close()

			redirectedToURL, err := api.GetRedirect(ts.URL + "/" + tt.alias)
			require.NoError(t, err)

			// Check the final URL after redirection.
			assert.Equal(t, tt.url, redirectedToURL)
		})
	}
}
