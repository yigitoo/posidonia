package main

import (
	"fmt"

	"github.com/yigitoo/posidonia/server/lib"
)

func main() {
	api := lib.SetupApi()

	fmt.Printf("\n\n-------------------------------------\nLink: http://localhost%s/\nAuthor: yigitoo<Yiğit GÜMÜŞ gumusyigit101@gmail.com>\n-------------------------------------\n\n\n", lib.PORT)
	fmt.Println("")
	api.Run(lib.PORT)
}
