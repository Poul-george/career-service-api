package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Poul-george/go-api/api/config"
	cmdConfig "github.com/Poul-george/go-api/api/presentation/cmd/config"
)

func main() {
	c := config.GetServerConfig()
	e := cmdConfig.NewEchoServer()
	s := http.Server{
		Addr:        c.StartAddress,
		Handler:     e,
		IdleTimeout: time.Duration(c.IdleTimeout),
	}
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		fmt.Printf("s.ListenAndServe() Error = %v \n", err.Error())
		e.Logger.Fatal(err)
	}
}
