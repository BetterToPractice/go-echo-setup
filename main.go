package main

import (
	"github.com/BetterToPractice/go-echo-setup/cmd"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @schemes http https
// @basePath /
func main() {
	cmd.Execute()
}
