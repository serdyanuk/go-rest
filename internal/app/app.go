package app

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/serdyanuk/go-rest/config"
	"github.com/serdyanuk/go-rest/internal/app/handler"
	"github.com/serdyanuk/go-rest/internal/app/middleware"
	"github.com/serdyanuk/go-rest/internal/app/store"
	"github.com/serdyanuk/go-rest/internal/app/store/sqlstore"
	"github.com/serdyanuk/go-rest/internal/pkg/logger"
	"github.com/sirupsen/logrus"
)

var (
	doneSignal      = make(chan os.Signal, 1)
	shutdownTimeout = time.Second * 10
)

type App struct {
	config     *config.Config
	logger     *logger.Logger
	httpServer *http.Server
	router     *httprouter.Router
	store      store.Store
	errCh      chan error
}

func New(config *config.Config) *App {
	router := httprouter.New()
	httpServer := &http.Server{
		Addr:    config.Port,
		Handler: router,
	}
	return &App{
		logger:     logger.Get(),
		config:     config,
		httpServer: httpServer,
		router:     router,
	}
}

func (a *App) Run() {
	signal.Notify(doneSignal, syscall.SIGINT, syscall.SIGTERM)

	a.logger.WithField("MODE", a.config.Mode).Info()

	// database
	a.runStore()

	// http
	a.router.POST("/signup", middleware.ErrorHandler(handler.Signup(a.store.User())))
	go a.runHttpServer()

	select {
	case err := <-a.errCh:
		a.logger.Error(err)
		return
	case <-doneSignal:
		a.shutdown(shutdownTimeout)
	}
}

func (a *App) runStore() {
	s, err := sqlstore.New(a.config.DBConfig)
	if err != nil {
		a.logger.Fatal(err)
	}
	a.store = s
}

func (a *App) runHttpServer() {
	a.logger.WithField("ADDR", a.httpServer.Addr).Info("HTTP running")
	a.errCh <- a.httpServer.ListenAndServe()
}

func (a *App) shutdown(timeout time.Duration) {
	a.logger.Info("App gracefuly shutdown...")

	ctx, _ := context.WithTimeout(context.Background(), timeout)
	err := a.httpServer.Shutdown(ctx)
	if err != nil {
		logrus.Error(err)
	}
}
