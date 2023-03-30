package openai_test

import (
	. "github.com/leizhenpeng/go-api2d"
	"github.com/leizhenpeng/go-api2d/internal/test"
	"github.com/leizhenpeng/go-api2d/internal/test/checks"

	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

// TestListModels Tests the models endpoint of the API using the mocked server.
func TestListModels(t *testing.T) {
	server := test.NewTestServer()
	server.RegisterHandler("/v1/models", handleModelsEndpoint)
	// create the test server
	var err error
	ts := server.OpenAITestServer()
	ts.Start()
	defer ts.Close()

	config := DefaultConfig(test.GetTestToken())
	config.BaseURL = ts.URL + "/v1"
	client := NewClientWithConfig(config)
	ctx := context.Background()

	_, err = client.ListModels(ctx)
	checks.NoError(t, err, "ListModels error")
}

// handleModelsEndpoint Handles the models endpoint by the test server.
func handleModelsEndpoint(w http.ResponseWriter, _ *http.Request) {
	resBytes, _ := json.Marshal(ModelsList{})
	fmt.Fprintln(w, string(resBytes))
}
