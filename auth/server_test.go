package auth

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	authpb "github.com/jonDufty/recipes/auth/rpc/authpb"
	"github.com/jonDufty/recipes/config"
	"github.com/stretchr/testify/require"
)

func newTestDB() *TestApp {
	cfg := config.NewAuthConfig()
	testApp := NewTestApp(cfg)

	testApp.InitDB()

	return testApp
}

func TestSignup(t *testing.T) {
	testApp := newTestDB()
	testApp.InitServers()

	err := testApp.App.DB.Ping()
	require.NoError(t, err)

	testServer := httptest.NewServer(testApp.Http)

	ctx := context.Background()

	defer testServer.Close()
	defer testApp.Closer()

	type testCase struct {
		name        string
		Input       *authpb.SignupRequest
		Expected    *authpb.LoginResponse
		ShouldError bool
	}

	tests := []testCase{
		{
			name: "Test basic signup",
			Input: &authpb.SignupRequest{
				Email:    "test@example.com",
				Password: "password",
				FullName: "Test User1",
			},
			Expected: &authpb.LoginResponse{
				User: &authpb.User{
					Email:    "test@example.com",
					UserId:   int64(1),
					FullName: "Test User1",
				},
			},
			ShouldError: false,
		},
		{
			name: "Test second signup",
			Input: &authpb.SignupRequest{
				Email:    "test2@example.com",
				Password: "password",
				FullName: "Test User2",
			},
			Expected: &authpb.LoginResponse{
				User: &authpb.User{
					Email:    "test2@example.com",
					UserId:   int64(2),
					FullName: "Test User2",
				},
			},
			ShouldError: false,
		},
		{
			name: "Test signup with empty email",
			Input: &authpb.SignupRequest{
				Email:    "",
				Password: "password",
				FullName: "Test User3",
			},
			Expected:    nil,
			ShouldError: true,
		},
		{
			name: "Test signup with empty email",
			Input: &authpb.SignupRequest{
				Email:    "test@example.com",
				Password: "",
				FullName: "Test User4",
			},
			Expected:    nil,
			ShouldError: true,
		},
	}

	client := authpb.NewAuthProtobufClient(testServer.URL+"/auth", &http.Client{})

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			resp, err := client.SignupWithPassword(ctx, tc.Input)
			if tc.ShouldError {
				require.Error(t, err)
				return
			}
			fmt.Printf("%v", resp)
			require.NoError(t, err)
			require.Equal(t, tc.Expected.User.Email, resp.User.Email)
			require.Equal(t, tc.Expected.User.FullName, resp.User.FullName)
			require.Equal(t, tc.Expected.User.UserId, resp.User.UserId)
		})
	}
}
