package main
import (
	"log"
	"github.com/hashicorp/terraform/config"
	"fmt"
	"flag"
)

var dir = flag.String("dir", "", "Directory to validate")

func main() {
	if *dir == "" {
		log.Fatal("Directory must be specified")
	}

	fmt.Println("Valdating directory ", *dir)
	cfg, err := config.LoadDir(*dir)
	if err != nil {
		log.Fatalf("Loading: %s\n\nError: %s", *dir, err)
	}
	err = cfg.Validate()
	if err != nil {
		log.Fatalf("Validating: %s\n\nError: %s", *dir, err)
	}
}
