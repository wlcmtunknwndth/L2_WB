package main

import "net/http"

type Handler struct {
	b *Broker
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	var bodyReq []byte
	data := h.b.GetFromChan(string(bodyReq))
	w.Write(data)
}

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	var reqOrder []byte
	h.b.SendToChan("save", reqOrder)
}
