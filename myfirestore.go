package main

import (
	"context"
	"firebase.google.com/go"
	"google.golang.org/api/option"
	"log"
	"time"
	"github.com/aws/aws-lambda-go/lambda"
)

type iratto struct {
	Type string
	Created time.Time
}

type authData struct {
	Type                    string `json:"type"`
	ProjectID               string `json:"project_id"`
	PrivateKeyID            string `json:"private_key_id"`
	PrivateKey              string `json:"private_key"`
	ClientEmail             string `json:"client_email"`
	ClientID                string `json:"client_id"`
	AuthURI                 string `json:"auth_uri"`
	TokenURI                string `json:"token_uri"`
	AuthProviderX509CertURL string `json:"auth_provider_x509_cert_url"`
	ClientX509CertURL       string `json:"client_x509_cert_url"`
}

func main() {
	lambda.Start(hello)
}

func hello(event IotButtonEvent) {
	ctx := context.Background()

	opt := option.WithCredentialsFile("./firebase_secret_key.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	_, _, err = client.Collection("users").Add(ctx, iratto{
		Type:   event.DeviceEvent.ButtonClicked.ClickType,
		Created: time.Now(),
	})
	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
	}

}
