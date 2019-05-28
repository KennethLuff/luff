package main

import (
	"fmt"
	"luff/pkg/setting"
	"luff/routers"
	"net/http"
)

func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
