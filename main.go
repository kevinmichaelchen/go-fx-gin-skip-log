package main

import (
	"github.com/kevinmichaelchen/go-fx-gin-skip-log/internal/app"
	"go.uber.org/fx"
)

func main() {
	a := fx.New(
		app.Module,
	)
	a.Run()
}
