package main

import (
	"fmt"

	natsSrvTest "github.com/nats-io/nats-server/v2/test"
)

func main() {
	s := natsSrvTest.RunRandClientPortServer()
	defer s.Shutdown()

	fmt.Printf("Total Subscribers: %v \n", s.GlobalAccount().TotalSubs())
}
