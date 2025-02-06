package types

import (
	"github.com/socketbasehq/socketbase/pkg/pkg/server"
	"go.uber.org/fx"
)

type Module struct {
	fx.Out

	Routes server.RouteGroup `group:"routes"`
}
