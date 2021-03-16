package main

import "github.com/ITheCorgi/EscrumGolang/internal/app"

const serverConfigurationPath = "../../configs/serverconfig.yml"

func main() {
	app.Run(serverConfigurationPath)
}
