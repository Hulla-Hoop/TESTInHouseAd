package handlers

import "net/http"

// example /category/create?name=Fresh

func (h *marketHandlers) CreateCategory(w http.ResponseWriter, r *http.Request) {
	reqID, ok := r.Context().Value("reqID").(string)
	if !ok {
		reqID = ""
	}

	name := r.URL.Query().Get("name")

	category, err := h.service.CreateCategory(reqID, name)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.Write(category)
	w.WriteHeader(http.StatusOK)
}
