package remove_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"shortli/internal/http-server/handlers/url/remove"
	"shortli/internal/http-server/handlers/url/remove/mocks"
	"shortli/internal/lib/logger/handlers/slogdiscard"
	"shortli/internal/storage"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDeleteHandler(t *testing.T) {
	cases := []struct {
		name      string
		alias     string
		respError string
		mockError error
	}{
		{
			name:      "delete non-existent alias",
			alias:     "nonexistent",
			respError: "failed to delete alias",
			mockError: storage.ErrURLNotFound,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			urlRemoverMock := mocks.NewURLRemover(t)

			if tc.respError == "" || tc.mockError != nil {
				urlRemoverMock.On("DeleteURL", tc.alias).
					Return(tc.mockError).
					Once()
			}

			r := chi.NewRouter()
			r.Delete("/{alias}", remove.New(slogdiscard.NewDiscardLogger(), urlRemoverMock))

			ts := httptest.NewServer(r)
			defer ts.Close()

			req, err := http.NewRequest("DELETE", ts.URL+"/"+tc.alias, nil)
			assert.NoError(t, err)

			rr := httptest.NewRecorder()
			r.ServeHTTP(rr, req)

			require.Equal(t, rr.Code, http.StatusOK)

			body := rr.Body.String()

			var resp remove.Response

			require.NoError(t, json.Unmarshal([]byte(body), &resp))

			require.Equal(t, tc.respError, resp.Error)
		})
	}
}
