package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/klyakssa/go-musthave-shortener-tpl/internal/config"
	"github.com/klyakssa/go-musthave-shortener-tpl/internal/logger"
	"github.com/klyakssa/go-musthave-shortener-tpl/internal/model"
	"github.com/klyakssa/go-musthave-shortener-tpl/internal/repository"
)

type MyHandlerStruct struct {
	cfg    *config.Config
	Logger *logger.MyLogger
}

func NewMyHandler(cfg *config.Config, l *logger.MyLogger) *MyHandlerStruct {
	return &MyHandlerStruct{
		cfg:    cfg,
		Logger: l,
	}
}

func (h *MyHandlerStruct) ShortenHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	shrt, err := repository.Shorten(string(body))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = r.Body.Close(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Length", strconv.Itoa(len(h.cfg.WebConfig.BaseUrl+"/"+shrt)))
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(h.cfg.WebConfig.BaseUrl + "/" + shrt))
}

func (h *MyHandlerStruct) UnshortenHandler(w http.ResponseWriter, r *http.Request) {
	lng, err := repository.Unshorten(r.URL.Path[1:])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Add("Location", lng)
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func (h *MyHandlerStruct) NewShortenHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var req model.ShortenRequest
	if err = json.Unmarshal(body, &req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	shrt, err := repository.Shorten(req.URL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = r.Body.Close(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(model.ShortenResponse{
		Result: h.cfg.WebConfig.BaseUrl + "/" + shrt,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(http.StatusCreated)
	w.Write(data)
}
