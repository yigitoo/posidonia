package main

import (
	"fmt"
	"strconv"
)

func main() {
	api := SetupApi()

	fmt.Printf("Link: https://localhost:%s/\n", strconv.Itoa(int(iPORT)))
	api.Run(PORT)
}
