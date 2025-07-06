package main

import (
	"context"

	"os"

	"github.com/arnokay/arnobot-shared/applog"
	echoControllers "github.com/arnokay/arnobot-shared/controllers/echo"
	"github.com/arnokay/arnobot-shared/middlewares"
	"github.com/arnokay/arnobot-shared/pkg/assert"
	sharedService "github.com/arnokay/arnobot-shared/service"
	"github.com/labstack/echo/v4"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"

	"github.com/arnokay/arnobot-api/internal/api"
	"github.com/arnokay/arnobot-api/internal/api/middleware"
	"github.com/arnokay/arnobot-api/internal/app/config"
	"github.com/arnokay/arnobot-api/internal/app/service"
)

const AppName = "api"

type application struct {
	logger applog.Logger

	cache     jetstream.KeyValue
	msgBroker *nats.Conn
	api       *echo.Echo

	services       *service.Services
	apiControllers echoControllers.Controller
	apiMiddlewares *middleware.Middlewares
}

func main() {
	var app application

	// load config
	cfg := config.Load()

	// load logger
	logger := applog.NewCharmLogger(os.Stdout, AppName, cfg.Global.LogLevel, nil)
  applog.SetDefault(logger)
	app.logger = logger

	// load message broker
	mbConn, _, cache := openMB()
	app.msgBroker = mbConn
	app.cache = cache

	// load services
	services := &service.Services{}
	services.AuthModule = sharedService.NewAuthModule(app.msgBroker)
	services.PlatformModule = sharedService.NewPlatformModuleIn(app.msgBroker)

	app.services = services

	// load middlewares
	app.apiMiddlewares = middleware.New(
		middlewares.NewAuthMiddleware(app.services.AuthModule),
	)

	// load api controllers
	app.apiControllers = &api.Controllers{
		PlatformController: api.NewPlatformController(
			app.services.PlatformModule,
			app.apiMiddlewares,
		),
	}

	app.Start()
}

func openMB() (*nats.Conn, jetstream.JetStream, jetstream.KeyValue) {
	nc, err := nats.Connect(config.Config.MB.URL)
	assert.NoError(err, "openMB: cannot open message broker connection")

	js, err := jetstream.New(nc)
	assert.NoError(err, "openMB: cannot create jetstream")
	kv, err := js.CreateKeyValue(context.Background(), jetstream.KeyValueConfig{
		Bucket: "default-api",
	})
	assert.NoError(err, "openMB: cannot create default kv store")

	return nc, js, kv
}
