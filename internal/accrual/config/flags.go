package config

import (
	"flag"
)

type AppFlags struct {
	addr        *string // -a
	databaseURI *string // -d
}

func (p *AppFlags) Addr() string {
	return *p.addr
}

func (p *AppFlags) DatabaseURI() string {
	return *p.databaseURI
}

func parseFlags() AppFlags {
	var addr, databaseURI string

	flag.StringVar(&addr, "a", DefaultAddress, "Host IP address")
	flag.StringVar(&databaseURI, "d", DefaultDB, "Connection string for DB")
	flag.Parse()

	return AppFlags{addr: &addr, databaseURI: &databaseURI}

}
