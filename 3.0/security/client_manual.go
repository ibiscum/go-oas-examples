package main

import (
	"context"
	"flag"
	"log"

	"golang.org/x/oauth2/clientcredentials"
)

func main() {

	clientID := flag.String("client-id", "test_client_1", "client id from your client")
	clientSecret := flag.String("client-secret", "test_secret", "client secret from your client")
	flag.Parse()

	config := clientcredentials.Config{
		ClientID:     *clientID,
		ClientSecret: *clientSecret,
		TokenURL:     "http://localhost:8080/v1/oauth/tokens",
		Scopes:       []string{"read_write"},
	}

	//client := config.Client(context.Background())
	//log.Printf("%#v", client.Transport)

	token, err := config.Token(context.Background())
	if err != nil {
		log.Fatalf("token: %#v", err)
	}

	log.Printf("token: %#v", token)
}
