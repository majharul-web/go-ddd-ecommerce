package cmd

import (
	"ecommerce/config"
	"ecommerce/rest"
)

func Serve() {
	// Load configuration
	conf := config.GetConfig()
	// Start the REST server
	rest.Start(conf)

}
