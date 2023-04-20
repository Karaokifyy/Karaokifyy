package main

import (
	env_vars "github.com/Karaokifyy/BackEnd/src/helpers/env_vars"
	router "github.com/Karaokifyy/BackEnd/src/routes"
)

// primary entry point for back-end
func main() {
	env_vars.GetEnvVars()
	router.Run()
}
