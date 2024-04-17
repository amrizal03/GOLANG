package productcontroller

import (
	"go-jwt-mux/helper"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	data := []map[string]interface{}{
		{
			"id":           1,
			"nama_product": "Kemeja",
			"stok":         10,
		},
		{
			"id":           2,
			"nama_product": "Celana",
			"stok":         20,
		},
		{
			"id":           3,
			"nama_product": "Sepatu",
			"stok":         30,
		},
	}

	helper.ResponseJSON(w, http.StatusOK, data)
}
