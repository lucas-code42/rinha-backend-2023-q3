package integration

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/lucas-code42/rinha-backend/infra"
	"github.com/lucas-code42/rinha-backend/internal/configs"
	"github.com/lucas-code42/rinha-backend/internal/repository"
	"github.com/lucas-code42/rinha-backend/pkg/sql"
)

func setupServer() *httptest.Server {
	configs.Init()
	db := sql.New("test")
	repo := repository.New(db.SqlClient)
	echo := echo.New()
	srv := infra.New(echo, repo)
	return httptest.NewServer(srv.SetupRouters())
}

func TestLiveEndpoint(t *testing.T) {
	ts := setupServer()
	defer ts.Close()

	r, err := http.Get(fmt.Sprintf("%s/live", ts.URL))
	if err != nil {
		t.Errorf("expected nil got %d", err)
	}

	if r.StatusCode != http.StatusOK {
		t.Errorf("expected http status_code: %d got http status_code: %d", http.StatusOK, r.StatusCode)
	}
}

func TestCreatePersonEndpoint(t *testing.T) {
	ts := setupServer()
	defer ts.Close()

	// TODO: usar fuzztesting
	payload := strings.NewReader(`{
		"apelido": "petu",
		"nome": "yasmin",
		"stack": [
			"js"
		],
		"nascimento": "2001/01/01"
	}`)

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/pessoas", ts.URL), payload)

	if err != nil {
		t.Error(err)
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusCreated {
		t.Error(err)
	}

	_, err = io.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}
}
