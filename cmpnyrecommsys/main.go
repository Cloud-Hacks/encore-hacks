package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/afzal442/encore-hacks/pkg/echo"
	"github.com/afzal442/encore-hacks/v1/ipad"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil && os.Getenv("ENV") != "production" {
		log.Fatalf("Error loading .env file.")
	}

	//database.ConnectDatabase(database.Config())

	e := echo.Engine()

	e.GET("/", echo.OnlineCheck)

	v1 := e.Group("/v1")

	ipad.Routes(v1)

	go func() {
		if err := e.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))); err != nil && !errors.Is(err, http.ErrServerClosed) {
			e.Logger.Fatal("Shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
