package travis

import (
	"testing"

	"gopkg.in/yaml.v2"
)

func TestMarshal(t *testing.T) {
	DefaultTravisConfig.Env.Global = []string{
		"GIT_BRANCH=master",
		"GITHUB_REPO=tools/godep",
	}
	data, err := yaml.Marshal(DefaultTravisConfig)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(data))
}
