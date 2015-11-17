package github

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var TOKEN = os.Getenv("GITHUB_TOKEN")

func TestGetUser(t *testing.T) {
	gh := New(TOKEN)
	Convey("Should get user success", t, func() {
		user, err := gh.User()
		So(err, ShouldBeNil)
		t.Log(user)
	})

	Convey("Should get hooks", t, func() {
		hooks, err := gh.Hooks("codeskyblue", "fswatch")
		So(err, ShouldBeNil)
		for _, hook := range hooks {
			t.Log(hook)
		}
	})
}
