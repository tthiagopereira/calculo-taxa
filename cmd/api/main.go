package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gitub.com/thiagopereira7/taxas/internal/entity"
	"net/http"
)

func main() {
	// chi
	//r := chi.NewRouter()
	//r.Use(middleware.Logger)
	//r.Get("/order", orderHandler)
	//
	//err := http.ListenAndServe(":8888", r)
	//if err != nil {
	//	return
	//}
	e := echo.New()
	// Configurar o Logger com as opções desejadas

	e.Use(middleware.Logger())

	e.GET("/order", OrderHandler)

	e.Logger.Fatal(e.Start(":1323"))
}

func OrderHandler(c echo.Context) error {
	order, err := entity.NewOrder("123", 10, 1)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	errPrice := order.CalculateFinalPrice()

	if errPrice != nil {
		return c.String(http.StatusInternalServerError, errPrice.Error())
	}

	return c.JSON(http.StatusCreated, order)
}

//func orderHandler(w http.ResponseWriter, r *http.Request) {
//
//	order, err := entity.NewOrder("123", 20.55, 1.1)
//	if err != nil {
//		panic(err)
//	}
//	error := order.CalculateFinalPrice()
//	if error != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//	}
//
//	json.NewEncoder(w).Encode(order)
//}
