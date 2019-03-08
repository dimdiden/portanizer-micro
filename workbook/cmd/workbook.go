package main

import (
	"os"
	"net/http"
	"syscall"
	"fmt"

	"github.com/kelseyhightower/envconfig"

	"github.com/dimdiden/portanizer-micro/workbook"
	"github.com/dimdiden/portanizer-micro/workbook/gormdb"
	workbooksvc "github.com/dimdiden/portanizer-micro/workbook/implementation"
	"github.com/dimdiden/portanizer-micro/workbook/transport"
	httptransport "github.com/dimdiden/portanizer-micro/workbook/transport/http"
	"github.com/jinzhu/gorm"

	"os/signal"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type config struct {
	DatabaseURL string `envconfig:"DATABASE_URL"`
	HTTPAddr string `envconfig:"HTTP_ADDR"`
}

func main() {
	var cfg config
	envconfig.MustProcess("", &cfg)

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = level.NewFilter(logger, level.AllowDebug())
		logger = log.With(logger,
			"svc", "workbook",
			"ts", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}
	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	var postRepository workbook.PostRepository
	var tagRepository workbook.TagRepository
	{
		var db *gorm.DB
		var err error
		db, err = gorm.Open("postgres", cfg.DatabaseURL)
		level.Info(logger).Log("msg", "trying to connect with options: "+cfg.DatabaseURL)
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}
		err = gormdb.RunMigration(db)
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}
		postRepository = gormdb.NewPostRepository(db, logger)
		tagRepository = gormdb.NewTagRepository(db, logger)
	}
	level.Info(logger).Log("msg", "database connected")

	var postSvc workbook.PostService
	var tagSvc workbook.TagService
	{
		postSvc = workbooksvc.NewPostService(postRepository, logger)
		tagSvc = workbooksvc.NewTagService(tagRepository, logger)
	}

	var h http.Handler
	{
		postEndpoints := transport.MakePostEndpoints(postSvc)
		tagEndpoints := transport.MakeTagEndpoints(tagSvc)
		h = httptransport.NewService(postEndpoints, tagEndpoints, logger)
	}

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		level.Info(logger).Log("transport", "HTTP", "addr", cfg.HTTPAddr)
		server := &http.Server{
			Addr:    cfg.HTTPAddr,
			Handler: h,
		}
		errs <- server.ListenAndServe()
	}()

	level.Error(logger).Log("exit", <-errs)
}
