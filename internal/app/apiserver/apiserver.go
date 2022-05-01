package apiserver

import (
	"encoding/json"
	"net/http"

	"github.com/belanenko/coingecko-parser/internal/app/storage"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Apiserver struct {
	config  *Config
	logger  *logrus.Logger
	router  *mux.Router
	storage *storage.WalletsHistory
}

func New(config *Config) *Apiserver {
	return &Apiserver{
		config:  config,
		logger:  logrus.New(),
		router:  mux.NewRouter(),
		storage: storage.NewWalletsHistory(),
	}
}

func (s *Apiserver) configureRouter() {
	s.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		jsoned, err := json.Marshal(s.config.WalletsHistory.GetAllHistory())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(jsoned)
	})
}

func (s *Apiserver) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}

func (s *Apiserver) Run() error {
	s.configureRouter()
	if err := s.configureLogger(); err != nil {
		return err
	}
	s.logger.Infof("Log Level: %v", s.config.LogLevel)

	s.logger.Infof("Starting apiserver on port %s", s.config.BindAddress)
	return http.ListenAndServe(s.config.BindAddress, s.router)
}
