package apiserver

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/belanenko/coingecko-parser/internal/app/geckoparser"
	"github.com/belanenko/coingecko-parser/internal/app/store"
	"github.com/gorilla/mux"
)

type APIServer struct {
	db     *store.Store
	config *Config
	router *mux.Router
	gecko  *geckoparser.GeckoParser
}

func New(config *Config, gecko *geckoparser.GeckoParser) *APIServer {
	return &APIServer{
		config: config,
		router: mux.NewRouter(),
		gecko:  gecko,
	}
}

func (r *APIServer) configureStore() error {
	db := store.NewStore(r.config.Store)
	if err := db.Open(); err != nil {
		return err
	}
	r.db = db
	return nil
}

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/getHistory", s.getHistoryHandle())
}

func (s *APIServer) Start() error {
	if err := s.configureStore(); err != nil {
		return err
	}
	if err := s.parseCurrenciesHistoryToStore(); err != nil {
		return err
	}

	s.configureRouter()

	return http.ListenAndServe(s.config.BindAddress, s.router)
}

func (s *APIServer) parseCurrenciesHistoryToStore() error {
	for _, currency := range s.gecko.Wallets {
		history, err := s.gecko.GetPriceHistoryPeriod(currency, "365")
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			continue
		}
		if err := s.db.History().Add(currency, history); err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			continue
		}
	}
	return nil
}

func (s *APIServer) getHistoryHandle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		currency := r.FormValue("currency")
		if currency == "" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		id, err := s.db.History().GetCurrencyId(currency)
		if err != nil {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		if id == -1 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		list, err := s.db.History().GetHistory("bitcoin")
		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
			return
		}
		w.WriteHeader(http.StatusOK)
		b, err := json.Marshal(list)
		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
			return
		}
		w.Write(b)
	}
}
