package main

import (
	"fmt"

	"github.com/yigitoo/posidonia/server/lib"
)

func main() {
	api := lib.SetupApi()

	fmt.Printf("\n\n-------------------------------------\nLink: https://localhost%s/\n-------------------------------------\n\n\n", lib.PORT)
	api.Run(lib.PORT)
}
