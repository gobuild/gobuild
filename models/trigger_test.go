package models

import "testing"

func TestTriggerTravisBuild(t *testing.T) {
	err := TriggerTravisBuild("tools", "godep", "master", "codeskyblue@gmail.com")
	if err != nil {
		t.Fatal(err)
	}
}
