package main

import (
	"fmt"

	"github.com/ixismail/orbit/internals/infrastructure"
)

func main() {

	infrastructure.ConfigureOrbit()

	config := infrastructure.LoadConfig()

	fmt.Println("Configuration loaded successfully:", config)

}