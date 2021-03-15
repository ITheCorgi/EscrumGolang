package main

import "github.com/ITheCorgi/EscrumGolang/app/server"

const serverConfigurationPath = "serverconfig.yml"

func main() {
	server.Run(serverConfigurationPath)
}
