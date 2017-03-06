# Octoship

```go
package main

import (
    "fmt"
    "github.com/whytheplatypus/octoship"
)

func main() {
    token := "a github token"
    // a team to check membership for
    team := 1337

    u, err := octoship.Check(token, team)
    success := "is"

    if err != nil {
        success = "is not"
    }

    if u == nil {
        fmt.Printf("No user was found for this token")
    } else {

        fmt.Printf("User %s %s a member of team %d", u.Email, success, team)
    }
}

```



`import github.com/whytheplatypus/octoship`




# Functions


## [Check](/target/check.go#L23)
```go
func Check(token string, team int) (u *github.User, err error)
```
Check determines if a user, identified by the token, is a member of the given
team. If they are not both an error and any user information that was attainable
is returned.






# Types


## [ErrNotMember](/target/check.go#L13)
ErrNotMember is returned if the authenticated user is not a member of the given
team of if something has gone wrong in the request to github.

```go
type ErrNotMember struct {
    Errors []error
}
```







### [Error](/target/check.go#L17)
```go
func (e *ErrNotMember) Error() string
```










# Packages



