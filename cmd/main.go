package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/illiafox/mailbase/database"
	"github.com/illiafox/mailbase/server"
	"github.com/illiafox/mailbase/util/config"
	"github.com/illiafox/mailbase/util/multiwriter"
	"log"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"
)

//nolint:funlen
func main() {
	defer debug.FreeOSMemory()

	// // Log
	file, err := os.Create("log.txt")
	if err != nil {
		panic(err)
	}

	log.SetOutput(multiwriter.NewMultiWriter(file, os.Stderr))

	// // Flags
	configPath := flag.String("conf", "config.toml", "config file location\nEx: -conf my_conf.json")

	format := config.TOML

	flag.Func("format",
		fmt.Sprintf("config format, default 'json' (available: %s)\nEx: -format yaml", config.Available),
		config.FlagString(&format),
	)

	flag.Parse()

	// // Parsing config
	conf, err := config.ReadConfig(*configPath, format)
	if err != nil {
		log.Println("Parsing config:", err)
		return
	}

	log.Print("Initializing connections")
	db, err := database.NewDatabase(conf)
	if err != nil {
		log.Println("New Database:", err)
		return
	}

	// Clear old sessions
	err = db.MySQL.Session.Clear(7)
	if err != nil {
		log.Println("Clearing sessions:", err)
		return
	}

	serv := server.Init(db, conf)

	sig := make(chan os.Signal, 1)

	// If you have better solution, please suggest it in the issue or contact me https://t.me/ebashu_gerych
	defer func() {
		close(sig)
		log.Print("Shutting down server")
		err = serv.Shutdown(context.Background())
		if err != nil {
			log.Println(fmt.Errorf("shutdown: server: %w", err))
		}
		ok := true
		log.Print("Closing connections")
		for _, err = range db.Close() {
			if err != nil {
				ok = false
				log.Println(err)
			}
		}
		if ok {
			log.Println("Database has closed successfully")
		} else {
			log.Println("Database has closed with errors")
		}
	}()

	go func() {
		if conf.Host.HTTP {
			log.Printf("Server started [HTTP] at 127.0.0.1:" + conf.Host.Port)
			err = serv.ListenAndServe() // HTTP
		} else {
			log.Printf("Server started [HTTPS] at 127.0.0.1:" + conf.Host.Port)
			err = serv.ListenAndServeTLS(conf.Host.Cert, conf.Host.Key) // HTTPS
		}
		select {
		case <-sig:
		default:
			log.Println(fmt.Errorf("server: unforeseeable stop: %w", err))
			sig <- nil
		}
	}()

	// Catch interrupt
	// work not properly
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGHUP)
	<-sig
}
