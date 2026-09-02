package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/vignesh/online-ordering/internal/api"
	"github.com/vignesh/online-ordering/internal/config"
	"github.com/vignesh/online-ordering/internal/db"
	"github.com/vignesh/online-ordering/internal/repository"
	"github.com/vignesh/online-ordering/internal/service"
)

func main() {
	cfg := config.Load()

	dbPool, err := db.New(cfg)
	if err != nil {
		panic(err)
	}

	repo := repository.NewOrderRepository(dbPool)
	customerRepo := repository.NewCustomerRepository(dbPool)

	wp := service.NewWorkerPool(3)

	orderService := service.NewOrderService(repo, *wp)
	_ = service.NewCustomerService(customerRepo)

	handler := api.NewOrderHandler(orderService)

	router := api.NewRouter(handler)

	if err := http.ListenAndServe(":"+cfg.Port, router); err != nil {
		log.Fatal(err)
	}
	fmt.Print("server is up and running!")
}
