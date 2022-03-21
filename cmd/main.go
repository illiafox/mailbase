package main

import (
	"flag"
	"fmt"
	"log"
	"mailbase/database"
	"mailbase/server"
	"mailbase/util/config"
	"os"
	"os/signal"
	"runtime/debug"
)

func main() {
	defer debug.FreeOSMemory()

	// // Flags
	configPath := flag.String("conf", "config.json", "config file location\nEx: -conf my_conf.json")

	format := config.TOML

	flag.Func("format", fmt.Sprintf("config format, default 'json' (available: %s)\nEx: -format yaml", config.Available), func(s string) error {
		if s != "" {
			format = config.FormatMap[s]
			if format == 0 { // Config formats start from 1
				return fmt.Errorf("unknown format '%s'\n", s)
			}
		}
		return nil
	})

	flag.Parse()

	// // Parsing config
	conf, err := config.ReadConfig(*configPath, format)
	if err != nil {
		fmt.Println("Parsing config: ", err)
		return
	}

	db, err := database.NewDatabase(conf)
	if err != nil {
		fmt.Println("New Database: ", err)
		return
	}

	// If you have better solution, please suggest it in the issue or contact me https://t.me/ebashu_gerych
	defer func() {
		ok := true
		for _, err = range db.Close() {
			if err != nil {
				ok = false
				log.Println(err)
			}
		}
		if ok {
			fmt.Println()
			log.Println("Database has closed successfully")
		}
	}()

	// Clear old sessions
	err = db.MySQL.ClearSessions(7)
	if err != nil {
		fmt.Println("Clearing sessions: ", err)
		return
	}

	sig := make(chan os.Signal)

	// // Handling
	go server.Init(db, conf, sig)

	// Catch interrupt
	signal.Notify(sig, os.Interrupt)
	<-sig
}
