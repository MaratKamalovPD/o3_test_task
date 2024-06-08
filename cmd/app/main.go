package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"

	app "github.com/MaratKamalovPD/o3_test_task/internal/pkg/server"
	signal "github.com/MaratKamalovPD/o3_test_task/internal/pkg/server/delivery"
)

func main() {
	srv := new(app.Server)

	var shutdownCannel chan os.Signal = signal.GetShutdownChannel()

	go func() {
		log.Println("starting server")

		if err := srv.Run(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalln("something went wrong while starting the server, err=", err)
		}
	}()

	<-shutdownCannel

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatalln("something went wrong while shutting down the server, err=", err)
	}

	log.Println("server was successful shut down")
}
