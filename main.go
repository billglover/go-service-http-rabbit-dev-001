package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/nebhale/client-go/bindings"
)

func main() {
	log := log.New(os.Stdout, "", log.Lmicroseconds|log.LUTC|log.Ldate|log.Lshortfile)

	if err := run(log); err != nil {
		log.Println("startup", err)
		os.Exit(1)
	}
}

func run(log *log.Logger) error {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	b := bindings.FromServiceBindingRoot()
	b = bindings.Filter(b, "rmq")
	if len(b) != 1 {
		return fmt.Errorf("incorrect number of RabbitMQ bindings: %d", len(b))
	}

	u, ok := bindings.Get(b[0], "url")
	if !ok {
		return fmt.Errorf("no URL in binding")
	}

	fmt.Printf("%+v", b)

	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Something worked, don't ask how."))
		w.Write([]byte("Trying to bind to:" + u))
	})

	err := http.ListenAndServe(":"+port, r)
	return err
}
