package auth

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	authpb "github.com/jonDufty/recipes/auth/rpc/authpb"
	"github.com/jonDufty/recipes/config"
	"github.com/stretchr/testify/require"
)

func TestTwirpServer(t *testing.T) {
	testApp := NewTestApp(config.NewAuthConfig())
	testApp.InitServers()
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

	testApp := NewTestApp(config.NewAuthConfig())
	testApp.InitServers()
	testServer := httptest.NewServer(testApp.Http)
	defer testServer.Close()

	t.Run("Check base route", func(t *testing.T) {
		client := http.Client{}
		resp, err := client.Get(testServer.URL + "/system/healthcheck")
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)

	})
}
