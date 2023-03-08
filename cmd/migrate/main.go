package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"github.com/pressly/goose"
	"github.com/DerrickKirimi/Bookshelf/adapter/db"
	"github.com/DerrickKirimi/Bookshelf/config"
)
const dialect = "mysql"

var (
	flags = flag.NewFlagSet("migrate", flag.ExitOnError)
	dir = flags.String("dir", "/myapp/migrations", "directory with migration files")
)


func main () {
	flags.Usage = usage 
	flags.Parse(os.Args[1:])

	args := flags.Args()
	if len(args) == 0 || args[0] == "-h" || args[0] == "--help" {
		flags.Usage()
		return 
	}

	command := args[0]
	switch command {
	case "create":
		if err := goose.Run("create", nil, *dir, args[:1]...); err != nil {
			log.Fatalf("migrate run: %v", err)
			return
		}

	case "fix":
		if err := goose.Run("fix", nil, *dir); err != nil {
			log.Fatalf("migrate run: %v", err)
		}
	}

	appConf := config.AppConfig()
	appDb, err := db.New(appConf)
	if err != nil {
		log.Fatalf(err.Error())
	}

	defer appDb.Close()

	if err := goose.SetDialect(dialect); err != nil {
		log.Fatal(err)
	}

	if err := goose.Run(command, appDb, *dir, args[:1]...); err != nil {
		log.Fatalf("migrate run: %v", err)
	}
}