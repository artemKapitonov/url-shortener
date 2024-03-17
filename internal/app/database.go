package app

import (
	"flag"
)

func getTypeOfStorageByFlag() string {
	name := flag.String("db", "postgres", "")
	flag.Parse()

	return *name
}
