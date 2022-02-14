package tests

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jonDufty/recipes/auth"
	authpb "github.com/jonDufty/recipes/auth/rpc/authpb"
	"github.com/jonDufty/recipes/config"
	"github.com/stretchr/testify/require"
)

type TestApp struct {
	App   *auth.App
	Http  http.Handler
	Twirp *auth.Server
}

func NewTestApp() *TestApp {
	cfg := config.NewAuthConfig()
	app := &auth.App{
		Config: cfg,
	}

	router := auth.NewRouter(app)
	twirpServer := auth.NewServer(app)

	return &TestApp{
		App:   app,
		Http:  router,
		Twirp: twirpServer,
	}
}

func TestTwirpServer(t *testing.T) {
	testApp := NewTestApp()

	testServer := httptest.NewServer(testApp.Http)
	defer testServer.Close()

	t.Run("Check SayHello route", func(t *testing.T) {
		client := authpb.NewAuthProtobufClient(testServer.URL+"/auth", &http.Client{})

		resp, err := client.SayHello(context.Background(), &authpb.SayHelloRequest{Name: "Jon"})
		require.NoError(t, err)
		require.Equal(t, resp.Greeting, "Hello Jon")
	})

}

func TestHttpRoutes(t *testing.T) {

	testApp := NewTestApp()

	testServer := httptest.NewServer(testApp.Http)
	defer testServer.Close()

	t.Run("Check base route", func(t *testing.T) {
		client := http.Client{}
		resp, err := client.Get(testServer.URL + "/")
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)

	})
}
