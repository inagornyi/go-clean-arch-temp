package app

import (
	"go-clean-arch-temp/cmd/config"
	"log"
	"os"
	"path/filepath"

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

	_, err = cfg.NewConfig(config.Config{}, dir, "config.yml")
	if err != nil {
		return err
	}

	return nil
}
