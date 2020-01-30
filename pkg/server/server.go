package server

import (
	"context"
	"github.com/kevin-monteiro/http-stub-provider/internal/stub"
	"github.com/kevin-monteiro/http-stub-provider/pkg/types"
	"os"
	"os/signal"
	"syscall"
	"time"

	"log"
	"net/http"
	"strconv"
)

func StartStubServer(params *types.Server) {
	addr := params.Addr + ":" + strconv.Itoa(params.Port)
	router := stub.CreateRouter(params.StubPath)

	srv := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Print("Server Started : ", addr)

	<-done
	log.Print("Server Stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
	log.Print("Server Exited Properly")
}
