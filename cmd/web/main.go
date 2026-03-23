package main

import (
	"flag"
	"log/slog"
	"os"
	"text/template"

	"github.com/go-playground/form/v4"
	"github.com/rudimuliawan/snippetbox-gin/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type application struct {
	logger        *slog.Logger
	snippets      *models.SnippetModel
	templateCache map[string]*template.Template
	formDecoder   *form.Decoder
}

func main() {
	addr := flag.String("addr", ":8080", "HTTP network address")
	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MySQL data source name")
	flag.Parse()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	db, err := openDB(*dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	templateCache, err := newTemplateCache()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	app := &application{
		logger:        logger,
		snippets:      &models.SnippetModel{DB: db},
		templateCache: templateCache,
		formDecoder:   form.NewDecoder(),
	}

	r := app.Setup()

	logger.Info("starting server")

	err = r.Run(*addr)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	os.Exit(1)
}

func openDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		return nil, err
	}

	sqlDb, err := db.DB()
	if err != nil {
		return nil, err
	}

	err = sqlDb.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
