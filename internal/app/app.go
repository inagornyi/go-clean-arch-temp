package app

import (
	"errors"
	"go-clean-arch-temp/cmd/config"
	"go-clean-arch-temp/pkg/httpserver"
	"go-clean-arch-temp/pkg/mariadb"
	"go-clean-arch-temp/pkg/rabbitmq"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	cfg "github.com/inagornyi/go-config"
)

func Run() error {
	log.Println("[app]: up")
	defer func() {
		log.Println("[app]: down")
	}()

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	dir = filepath.Clean(filepath.Join(dir, ""))

	cfg, err := cfg.NewConfig(config.Config{}, dir, "config.yml")
	if err != nil {
		return err
	}

	db, err := mariadb.NewConnection(cfg.MariaDB.User, cfg.MariaDB.Password, cfg.MariaDB.Name)
	if err != nil {
		return err
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatal(err)
		}
		log.Println("[mariadb]: down")
	}()
	log.Println("[mariadb]: up")

	driver, err := mysql.WithInstance(db.DB(), &mysql.Config{})
	if err != nil {
		log.Fatal(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://../../migrations",
		"mysql",
		driver,
	)
	if err != nil {
		log.Fatal(err)
	}
	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatal(err)
	}

	rmq, err := rabbitmq.NewConnection(cfg.RabbitMQ.URL)
	if err != nil {
		return err
	}
	log.Println("[rabitmq]: up")
	defer func() {
		if err := rmq.Shutdown(); err != nil {
			log.Fatal(err)
		}
		log.Println("[rabitmq]: down")
	}()

	server := httpserver.NewHttpServer()
	err = server.Run()
	if err != nil {
		return err
	}
	defer func() {
		if err := server.Shutdown(); err != nil {
			log.Fatal(err)
		}
		log.Println("[server]: down")
	}()
	log.Println("[server]: up")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case <-quit:
		break
	case <-rmq.Notify():
		break
	case <-server.Notify():
		break
	}

	return nil
}
