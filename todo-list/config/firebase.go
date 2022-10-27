package config

import (
	"context"

	firebase "firebase.google.com/go"

	"google.golang.org/api/option"
)

var FirebaseApp *firebase.App

func InitialFirebase(pathFile string) (*firebase.App, error) {
	opt := option.WithCredentialsFile(pathFile)
	return firebase.NewApp(context.Background(), nil, opt)
}
