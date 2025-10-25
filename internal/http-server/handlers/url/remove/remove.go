package remove

import (
	"log/slog"
	"net/http"
	resp "shortli/internal/lib/api/response"
	"shortli/internal/lib/logger/sl"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type Response struct {
	resp.Response
}

//go:generate go run github.com/vektra/mockery/v2@v2.53.5 --name=URLRemover
type URLRemover interface {
	DeleteURL(alias string) error
}

func New(log *slog.Logger, urlRemover URLRemover) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.url.remove.New"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		alias := chi.URLParam(r, "alias")
		if alias == "" {
			log.Info("alias is empty")
			render.JSON(w, r, resp.Error("invalid request"))
			return
		}

		err := urlRemover.DeleteURL(alias)
		if err != nil {
			log.Error("failed to delete alias", sl.Err(err))
			render.JSON(w, r, resp.Error("failed to delete alias"))
			return
		}

		log.Info("alias added", slog.String("alias", alias))

		render.JSON(w, r, Response{
			Response: resp.OK(),
		})
	}
}
