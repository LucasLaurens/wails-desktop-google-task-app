package server

import (
	"calandar-desktop-task/internal/config"
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/wailsapp/wails/v2/pkg/runtime"
	"golang.org/x/oauth2"
)

type Server struct {
	listener net.Listener
}

type AuthConfig struct {
	config *oauth2.Config
}

func Handle(config *oauth2.Config, ctx context.Context) *oauth2.Token {
	server := start()
	defer server.stop()
	server.init(config)
	authConfig := &AuthConfig{config: config}
	authConfig.oauth2Authorization(ctx)
	code := server.getCode()
	token := authConfig.getToken(code)
	return token
}

func start() *Server {
	url := config.GetConfig("URL")
	listener, err := net.Listen("tcp", url)

	if err != nil {
		log.Fatalf("%v", err)
	}

	return &Server{
		listener: listener,
	}
}

func (server *Server) stop() {
	server.listener.Close()
}

func (server Server) init(config *oauth2.Config) *oauth2.Config {
	redirectURL := fmt.Sprintf(
		"http://%v",
		server.listener.Addr().String(),
	)

	fmt.Printf(
		"The redirect url is : %s",
		redirectURL,
	)

	config.RedirectURL = redirectURL
	return config
}

func (authConfig *AuthConfig) oauth2Authorization(ctx context.Context) {
	authURL := authConfig.config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Println("Opening browser for authorization:", authURL)
	runtime.BrowserOpenURL(ctx, authURL)
}

func (server Server) getCode() string {
	codeChannel := make(chan string)

	go func() {
		http.HandleFunc(
			"/",
			func(response http.ResponseWriter, request *http.Request) {
				code := request.URL.Query().Get("code")
				fmt.Fprint(response, "You can close this window now.")
				codeChannel <- code
			},
		)

		_ = http.Serve(server.listener, nil)
	}()

	code := <-codeChannel
	close(codeChannel)

	if code == "" {
		log.Fatal("no authorization code received")
	}

	return code
}

func (authConfig *AuthConfig) getToken(code string) *oauth2.Token {
	token, err := authConfig.config.Exchange(
		context.Background(),
		code,
	)

	if err != nil {
		log.Fatalf("unable to retrieve token from web: %s", err)
	}

	return token
}
