package Api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"secondchallange/Internal/Models"
	"secondchallange/Internal/adapter/Stream"
	"secondchallange/Internal/service"
	"secondchallange/config"
)

type TransactionController struct {
	Configurations     config.Configurations
	TransactionService service.IService
}

func Request(ts service.IService, config config.Configurations) {
	c := &TransactionController{
		Configurations:     config,
		TransactionService: ts,
	}
	r := chi.NewRouter()
	r.Get("/transactions", c.GetAllTransaction)
	r.Post("/Addtransaction", c.CreateTransaction)
	http.ListenAndServe("localhost:9090", r)

}

func (c *TransactionController) GetAllTransaction(w http.ResponseWriter, r *http.Request) {
	res, err := c.TransactionService.List(context.Background())
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (c *TransactionController) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var transModel Models.Transaction

	if err := json.NewDecoder(r.Body).Decode(&transModel); err != nil {
		fmt.Fprintf(w, err.Error())
	}
	Stream.NewProduce(&transModel, c.Configurations)
	res, err := c.TransactionService.Create(context.Background(), &transModel)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)

}
