package app

import (
	"github.com/kevinmichaelchen/go-fx-gin-skip-log/internal/app/handler"
	"go.uber.org/fx"
)

var Module = fx.Options(
	handler.Module,
)
