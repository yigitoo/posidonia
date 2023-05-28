package main

import (
	"fmt"

	"github.com/yigitoo/posidonia/server/lib"
)

func main() {
	api := lib.SetupApi()

	fmt.Printf("Link: https://localhost:%s/\n", lib.PORT)
	api.Run(lib.PORT)
}
