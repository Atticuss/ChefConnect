package liveview

import (
	"errors"
	"os"

	"github.com/brendonmatos/golive"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/atticuss/chefconnect/controllers"
	"github.com/atticuss/chefconnect/services"
)

var service *services.Service

type liveViewController struct {
	Service services.Service
	Config  Config
	App     *fiber.App
}

// Config defines the... configuration?
type Config struct {
	Port   string
	Domain string
	Logger *zerolog.Logger
	IsProd bool

	// UTC a boolean stating whether to use UTC time zone or local.
	UTC bool
}

// NewLiveViewController configures a controller for handling request/response logic via
// goliveview, a websocket based framework.
func NewLiveViewController(svc *services.Service, config *Config) controllers.Controller {
	lv := liveViewController{
		Service: *svc,
		Config:  *config,
	}

	service = svc
	return &lv
}

/*
<!-- Font Awesome -->
		<link
		  href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/5.15.1/css/all.min.css"
		  rel="stylesheet"
		/>
		<!-- Google Fonts -->
		<link
		  href="https://fonts.googleapis.com/css?family=Roboto:300,400,500,700&display=swap"
		  rel="stylesheet"
		/>
		<!-- MDB -->
		<link
		  href="https://cdnjs.cloudflare.com/ajax/libs/mdb-ui-kit/3.9.0/mdb.min.css"
		  rel="stylesheet"
		/>
*/

func (lvCtlr *liveViewController) SetupController() error {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if !lvCtlr.Config.IsProd {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	log.Logger = log.Output(
		zerolog.ConsoleWriter{
			Out:     os.Stderr,
			NoColor: false,
		},
	)

	lvCtlr.App = fiber.New()
	liveServer := golive.NewServer()

	lvCtlr.App.Get("/", liveServer.CreateHTMLHandler(NewRoot, golive.PageContent{
		Lang:  "us",
		Title: "ChefConnect",
		Head: `<!-- Font Awesome -->
		<link
		  href="/static/all.min.css"
		  rel="stylesheet"
		/>
		<!-- Google Fonts -->
		<link
		  href="/static/gfonts.css"
		  rel="stylesheet"
		/>
		<!-- MDB -->
		<link
		  href="/static/mdb.min.css"
		  rel="stylesheet"
		/>
		`,
	}))

	lvCtlr.App.Get("/ws", websocket.New(liveServer.HandleWSRequest))
	lvCtlr.App.Static("/static", "./controllers/liveview/static")

	return nil
}

func (lvCtlr *liveViewController) Run() error {
	err := lvCtlr.App.Listen(":3000")
	return err
}

func (lvCtlr *liveViewController) Stop() error {
	return errors.New("not implemented")
}
