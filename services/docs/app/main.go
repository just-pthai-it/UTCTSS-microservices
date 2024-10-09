package main

import (
	"UTCTSS-microservices/common"
	_ "UTCTSS-microservices/docs"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	serveOpenAPIsDocs()
}

func serveOpenAPIsDocs() {
	if os.Getenv("ENV") != "local" || os.Getenv("SERVICE_NAME") != "docs" {
		return
	}

	cmd := exec.Command("swagger", "serve", "--port=2000", "--no-open", "./docs/swagger.json")
	fmt.Println(common.Purple + "Serve OpenAPIs docs!" + common.Reset)
	_, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(common.Red + "Serve OpenAPIs docs failed: " + err.Error() + common.Reset)
	}
}
