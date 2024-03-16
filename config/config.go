package config

import (
	"flag"

	"github.com/vharitonsky/iniflags"
)

var (
	Port   = flag.Int("port", 8081, "Port for running application")
	ApiURL = flag.String("url", "http://localhost:8082", "URL for api")
)

func init() {
	iniflags.Parse()
}
