package main

import (
	"log"

	"github.com/ITheCorgi/EscrumGolang/internal/app"
)

const srvCfgPath = "configs/serverconfig.yml"

func main() {
	if err := app.Run(srvCfgPath); err != nil {
		log.Println(err.Error())
	}
}
