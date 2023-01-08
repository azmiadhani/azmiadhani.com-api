package main

import (
	"fmt"
	"os"

	"github.com/azmiadhani/azmiadhani.com-api/internal/utils"
)

func main() {
	utils.InitEnv()
	fmt.Println(os.Getenv("ROUTER_PORT"))
}
