package server

import (
	"fmt"
	"net/http"

	"github.com/CryptoLarry/ssm/ssmapi"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/snowzach/gorestapi/store"
)

// SSMGet fetches array of ssms by txid
func (s *Server) SSMGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Get the TXID
		txid := chi.URLParam(r, "txid")
		if txid == "" {
			render.Render(w, r, ErrInvalidRequest(fmt.Errorf("Invalid TXID")))
			return
		}
		bs, err := s.SSMStore.SSMGetByTXID(r.Context(), txid)
		if err == store.ErrNotFound {
			render.Render(w, r, ErrNotFound)
			return
		} else if err != nil {
			render.Render(w, r, ErrInvalidRequest(err))
			return
		}

		render.JSON(w, r, bs)
	}

}

// SSMSave saves ssm to DB
func (s *Server) SSMSave() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var b = new(ssmapi.Ssmitem)
		if err := render.DecodeJSON(r.Body, &b); err != nil {
			render.Render(w, r, ErrInvalidRequest(err))
			return
		}
		ssm, err := s.SSMStore.SSMPost(r.Context(), b)
		if err != nil {
			render.Render(w, r, ErrStandardError(err))
			return
		}
		render.JSON(w, r, ssm)
	}
}
