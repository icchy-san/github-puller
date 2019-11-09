package main

import (
	"context"
	"errors"
	"log"
	"os"

	gh "puller/usecase"

	"github.com/google/go-github/v28/github"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
)

func main() {
	err := loadEnv()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	oauthClient := oauth2.NewClient(oauth2.NoContext, ts)
	gh := gh.GH{Client: github.NewClient(oauthClient)}

	pRepoNames := gh.GetPrivateRepositoryNames(ctx)
	log.Printf("%v", pRepoNames)
}

func loadEnv() error {
	err := godotenv.Load()
	if err != nil {
		return errors.New("Error loading .env file")
	}
	return nil
}
