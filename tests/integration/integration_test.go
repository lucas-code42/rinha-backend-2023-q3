package integration

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/lucas-code42/rinha-backend/infra"
	"github.com/lucas-code42/rinha-backend/internal/configs"
	"github.com/lucas-code42/rinha-backend/internal/domain"
	"github.com/lucas-code42/rinha-backend/internal/repository"
	"github.com/lucas-code42/rinha-backend/pkg/sql"
)

// TODO: better error logs
// TODO: abstract http resquest...?

var UUID string

func setupTestServer() *httptest.Server {
	configs.Init()
	db := sql.New()
	repo := repository.New(db.SqlClient)
	echo := echo.New()
	srv := infra.New(echo, repo)
	return httptest.NewServer(srv.SetupRouters())
}

func Test_LiveEndpoint(t *testing.T) {
	ts := setupTestServer()
	defer ts.Close()

	r, err := http.Get(fmt.Sprintf("%s/live", ts.URL))
	if err != nil {
		t.Errorf("expected nil got %d", err)
	}

	if r.StatusCode != http.StatusOK {
		t.Errorf("expected http status_code: %d got http status_code: %d", http.StatusOK, r.StatusCode)
	}
}

func Test_CreatePersonEndpoint(t *testing.T) {
	ts := setupTestServer()
	defer ts.Close()

	// TODO: use fuzztesting
	client := &http.Client{}
	req, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("%s/pessoas", ts.URL),
		strings.NewReader(`{
			"apelido": "j",
			"nome": "Jhon Doe",
			"stack": [
				"js"
			],
			"nascimento": "2001/10/11"
		}`),
	)

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

	value, ok := res.Header["Location"]
	if !ok {
		t.Error(err)
	}
	UUID = strings.Replace(value[0], "/pessoas/", "", 1)
}

func Test_CountPeopleEndpoint(t *testing.T) {
	ts := setupTestServer()
	defer ts.Close()

	client := &http.Client{}
	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("%s/contagem-pessoas", ts.URL),
		nil,
	)
	if err != nil {
		t.Error(err)
	}

	res, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Error(err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}

	var jsonBody map[string]interface{}
	if err = json.Unmarshal(body, &jsonBody); err != nil {
		t.Error(err)
	}

	fmt.Println(jsonBody)

	_, ok := jsonBody["totalRecords"]
	if !ok {
		t.Error("could not get 'totalRecords' key")
	}
}

func Test_GetPersonIdEndpoint(t *testing.T) {
	if UUID == "" {
		t.Error("UIID is empty")
	}

	ts := setupTestServer()
	defer ts.Close()

	client := &http.Client{}
	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("%s/pessoas/%s", ts.URL, UUID),
		nil,
	)
	if err != nil {
		t.Error(err)
	}

	res, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}

	var expectedBody domain.Pessoa
	if err = json.Unmarshal(body, &expectedBody); err != nil {
		t.Error(err)
	}
}

func Test_SearchPersonEndpoint(t *testing.T) {
	ts := setupTestServer()
	defer ts.Close()

	client := &http.Client{}
	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("%s/pessoas?t=js", ts.URL), // simple fake term
		nil,
	)
	if err != nil {
		t.Error(err)
	}

	res, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}

	var expectedResponde []domain.Pessoa
	if err = json.Unmarshal(body, &expectedResponde); err != nil {
		t.Error(err)
	}

	if len(expectedResponde) == 0 {
		t.Error("expected response but got empty result")
	}
}
