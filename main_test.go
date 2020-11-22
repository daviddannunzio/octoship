package octoship_test

import (
	"context"
	"flag"
	"os"
	"testing"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

var token string
var team int64

func TestMain(m *testing.M) {
	flag.StringVar(&token, "token", "", "a valid personal github token for use in testing")
	flag.Parse()

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	// list all repositories for the authenticated user
	teams, _, err := client.Teams.ListUserTeams(ctx, nil)

	if err != nil || len(teams) < 1 {
		panic("no teams for token")
	}

	team = *teams[0].ID

	os.Exit(m.Run())
}
