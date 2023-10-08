package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Joaovitordebrito/wex-purchase-service/src/config"
	"github.com/Joaovitordebrito/wex-purchase-service/src/router"
)

func main() {
	config.LoadConfig()

	fmt.Printf("Api running on port %d", config.Port)

	r := router.Generate()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
