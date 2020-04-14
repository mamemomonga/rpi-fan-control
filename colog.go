package main

import (
	"github.com/comail/colog"
	"log"
	"os"
)

func init() {
	colog.SetDefaultLevel(colog.LDebug)

	if os.Getenv("DEBUG") != "" {
		colog.SetMinLevel(colog.LTrace)
	} else {
		colog.SetMinLevel(colog.LWarning)
	}

	colors := true
	if os.Getenv("TERM") != "" {
		colors = false
	}

	colog.SetFormatter(&colog.StdFormatter{
		Colors: colors,
		Flag:   log.Ldate | log.Ltime,
	})

	colog.Register()
}
