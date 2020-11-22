package octoship_test

import (
	"fmt"
	"testing"

	"github.com/daviddannunzio/octoship"
)

func ExampleCheck() {
	token := "a github token"
	// a team to check membership for
	team := int64(1337)

	u, err := octoship.Check(token, int(team))
	success := "is"

	if err != nil {
		success = "is not"
	}

	if u == nil {
		fmt.Printf("No user was found for this token")
	} else {

		fmt.Printf("User %s is %s a member of team %d", *u.Email, success, team)
	}
	//Output: No user was found for this token
}

func Test_Check(t *testing.T) {
	t.Logf("testing with user team %d", team)

	cases := map[string]struct {
		t  string
		tm int64
		e  error
	}{
		"good": {
			token,
			team,
			nil,
		},
		"bad_team": {
			token,
			123,
			nil,
		},
		"bad_token": {
			"bad_token",
			team,
			nil,
		},
	}

	for k, c := range cases {
		t.Logf("Running case %s", k)

		u, err := octoship.Check(c.t, int(c.tm))

		if err == nil && c.e != nil {
			t.Fatalf("case %s expected non-nil error", k)
		}

		t.Logf("User for case %s : %v", k, u)

		t.Logf("End case %s Run", k)
	}

}
