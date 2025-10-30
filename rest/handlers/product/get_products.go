package product

import (
	"go-practice-api/utilities"
	"net/http"
	"strconv"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	// get query parameters
	// get page from query parameters
	// get limit from query parameters

	reqQuery := r.URL.Query()

	pageAsStr := reqQuery.Get("page")
	limitAsStr := reqQuery.Get("limit")

	page, _ := strconv.ParseInt(pageAsStr, 10, 32)
	limit, _ := strconv.ParseInt(limitAsStr, 10, 32)

	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	productList, err := h.svc.List(page, limit)
	if err != nil {
		utilities.SendError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	cnt, err := h.svc.Count()
	if err != nil {
		utilities.SendError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	utilities.SendPage(w, productList, page, limit, cnt)
}
