package lito

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"xorm.io/xorm"
)

type Lito struct {
	server *fiber.App
	db     *xorm.Engine
	port   int
}

func New(db *xorm.Engine, port int) (*Lito, error) {
	var lito Lito
	lito.server = fiber.New()
	lito.port = port

	if db == nil {
		return &lito, errors.New("database connection not provided")
	}

	return &lito, nil
}

func (l *Lito) Run() error {
	fmt.Println("hello from lito")
	return nil
}
