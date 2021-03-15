package main

import "github.com/ITheCorgi/EscrumGolang/app/server"

const serverConfigurationPath = "serverconfig"

func main() {
	server.Run(serverConfigurationPath)
}
