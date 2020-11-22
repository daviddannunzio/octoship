package octoship

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// ErrNotMember is returned if the authenticated user is not a member of the given team
// of if something has gone wrong in the request to github.
type ErrNotMember struct {
	Errors []error
}

func (e *ErrNotMember) Error() string {
	return fmt.Sprintf("user is not a member of the team : %v", e.Errors)
}

// Check determines if a user, identified by the token, is a member of the given team.
// If they are not both an error and any user information that was attainable is returned.
func Check(token string, team int) (u *github.User, err error) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	u, _, e := client.Users.Get(ctx, "")

	if u == nil {
		return nil, &ErrNotMember{
			Errors: []error{e},
		}
	}

	// list all repositories for the authenticated user
	// ok, _, er := client.Organizations.IsTeamMember(ctx, team, *u.Login)
	membership, _, er := client.Teams.GetTeamMembership(ctx, int64(team), *u.Login)

	if membership == nil || (*membership.Role != "member" && *membership.Role != "maintainer") {
		err = &ErrNotMember{
			Errors: []error{
				er,
				e,
			},
		}

	}

	return u, err
}
