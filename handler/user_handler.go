package handler

import (
	"encoding/json"
	"ginserver/pkg"
	"io"
	"net/http"
)

type UserHandler struct {
	Service pkg.UserService
}

func (uh *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	ret, err := uh.Service.GetAll(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.WriteHeader(http.StatusOK)
	data, err := json.Marshal(ret)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.Write(data)
}

func (uh *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	var input pkg.UserCreateInput
	err = json.Unmarshal(body, &input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	ret, err := uh.Service.CreateUser(r.Context(), input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.WriteHeader(http.StatusOK)
	data, err := json.Marshal(ret)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.Write(data)
}
