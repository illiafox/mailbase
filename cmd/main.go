package main

import (
	"flag"
	"fmt"
	"mailbase/util/config"
	"os"
)

func main() {

	// // Flags
	configPath := flag.String("conf", "config.json", "config file location path")

	format := config.JSON

	flag.Func("format", fmt.Sprintf("config format (available: %s)", config.Available), func(s string) error {
		format := config.FormatMap[s]
		if format == 0 { // Config formats start from 1
			return fmt.Errorf("unknown format '%s'\n", s)
		}
		return nil
	})

	flag.Parse()

	// // Parsing config
	conf, err := config.ReadConfig(*configPath, format)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	println(conf)
}
