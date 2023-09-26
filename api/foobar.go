package api

import (
	"net/http"

	"github.com/antiphp/gonew/internal/render"
	lctx "github.com/hamba/logger/v2/ctx"
	jsoniter "github.com/json-iterator/go"
)

func (s *Server) handleFoobar() http.HandlerFunc {
	type foobarRequest struct {
		Msg string `json:"input"`
	}

	type foobarResponse struct {
		Msg string `json:"output"`
	}

	return func(rw http.ResponseWriter, req *http.Request) {
		request := foobarRequest{}
		switch {
		case req.Header.Get("Content-Length") == "", req.Header.Get("Content-Length") == "0":
		default:
			if err := jsoniter.NewDecoder(req.Body).Decode(&request); err != nil {
				s.log.Debug("Could not unmarshal body", lctx.Err(err))
				render.JSONError(rw, http.StatusBadRequest, "invalid json")
				return
			}
		}

		msg, err := s.app.Foobar(req.Context(), request.Msg)
		if err != nil {
			s.log.Error("Could not foobar", lctx.Err(err))
			render.JSONError(rw, http.StatusInternalServerError, "could not foobar")
			return
		}

		res := foobarResponse{Msg: msg}
		render.JSON(rw, http.StatusOK, &res)
	}
}
