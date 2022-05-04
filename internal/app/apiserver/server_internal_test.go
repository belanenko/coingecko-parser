package apiserver

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/belanenko/coingecko-parser/internal/app/model"
	"github.com/belanenko/coingecko-parser/internal/app/parser/testparser"
	"github.com/belanenko/coingecko-parser/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestServer_HandleGetHistory(t *testing.T) {
	s := newServer(teststore.New(), testparser.New())
	s.store.History().Add(model.TestCurrencyName, []model.History{{Price: 1, Timestamp: "1"}})
	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]interface{}{
				"name": model.TestCurrencyName,
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "invalid",
			payload: map[string]interface{}{
				"name": "invalid",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPost, "/getHistory", b)
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}
