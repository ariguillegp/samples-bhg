package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ariguillegp/samples-bhg/shodan/shodan"
)

func main() {
	apiKey := os.Getenv("SHODAN_API_KEY")
	s := shodan.New(apiKey)
	info, err := s.APIInfo()
	if err != nil {
		log.Panicln(err)
	}

	fmt.Printf(
		"Query Credits: %d\nScan Credits: %d\n\n",
		info.QueryCredits,
		info.ScanCredits)
}
