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
		log.Println("Starting server on port")
		if err := srv.Run(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalln("Failed to start server: ", err)
		}
	}()

	<-shutdownCannel
	log.Println("Gracefully shutting down server")

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatalln("Failed to shutdown the server gracefully: ", err)
	}

	log.Println("Server shutdown is successful")
}
