package handler

import (
	"ginserver/pkg"
	"net/http"
)

type PongHandler struct {
	Service pkg.PongService
}

func (uh *PongHandler) Pong(w http.ResponseWriter, r *http.Request) {
	ret, err := uh.Service.Pong(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(ret))
}
