package main

import (
	"log"
	"net/http"
	"os"

	"github.com/shuryak/api-wrappers/internal/api"
	"github.com/shuryak/api-wrappers/internal/api/calc"
	"github.com/shuryak/api-wrappers/internal/api/greetings"
	"github.com/shuryak/api-wrappers/pkg/router"
)

func main() {
	logger := log.New(os.Stdout, "[LOG] ", log.Lshortfile)

	r := router.New(logger)

	calcH := calc.NewCalc(logger)

	r.Add(
		router.POST("/calc", calcH.Handle).SetErrHandler(api.ErrHandler),
	)

	r.Add(
		router.NewGroup("/greetings",
			router.POST("/hello", greetings.Hello),
			router.NewGroup("/sub",
				router.POST("/hello", greetings.Hello),
				router.GET("/hello", greetings.HelloGet),
			),
		),
	)

	srv := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: r,
	}

	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
