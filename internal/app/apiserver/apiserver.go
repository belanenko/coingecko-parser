package apiserver

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Apiserver struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

func New(config *Config) *Apiserver {
	return &Apiserver{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *Apiserver) configureRouter() {
	s.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("main page"))
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

	return http.ListenAndServe(s.config.BindAddress, s.router)
}
