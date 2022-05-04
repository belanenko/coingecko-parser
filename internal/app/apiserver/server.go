package apiserver

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/belanenko/coingecko-parser/internal/app/model"
	"github.com/belanenko/coingecko-parser/internal/app/parser"
	"github.com/belanenko/coingecko-parser/internal/app/store"
	"github.com/gorilla/mux"
)

var (
	errIncorrectCurrency = errors.New("error incorrect currency")
)

type server struct {
	gecko  parser.Parser
	router *mux.Router
	store  store.Store
}

func newServer(store store.Store, gecko parser.Parser) *server {
	s := &server{
		gecko:  gecko,
		router: mux.NewRouter(),
		store:  store,
	}

	s.configureRouter()
	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) parseCurrenciesHistory() error {
	if s.gecko.Len() <= 0 {
		return errors.New("currencies not found")
	}
	for _, currency := range s.gecko.CurrenciesList() {
		h, err := s.gecko.GetPriceHistoryPeriod(currency, "365")
		if err != nil {
			fmt.Fprintf(os.Stdout, "%s | %v\n", currency, err.Error())
			continue
		}
		if err := s.store.History().Add(currency, h); err != nil {
			log.Fatal(err)
		}
	}

	return nil
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/getHistory", s.handleGetHistory())
}

func (s *server) handleGetHistory() http.HandlerFunc {
	type request struct {
		Currency string `json:"name"`
	}
	type answer struct {
		Currency string          `json:"currency"`
		History  []model.History `json:"history"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		history, err := s.store.History().GetHistory(req.Currency)
		if err != nil {
			if err == store.ErrNoRows {
				s.error(w, r, http.StatusUnprocessableEntity, errIncorrectCurrency)
			} else {
				s.error(w, r, http.StatusInternalServerError, errIncorrectCurrency)
			}
			return
		}

		if history == nil {
			s.error(w, r, http.StatusUnprocessableEntity, store.ErrNoRows)
		}

		s.respond(w, r, http.StatusOK, answer{
			Currency: req.Currency,
			History:  history,
		})
	}
}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
