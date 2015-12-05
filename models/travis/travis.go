package travis

type TravisConfig struct {
	Language string   `yaml:"language"`
	Go       []string `yaml:"go"`

	Env struct {
		Global []string `yaml:"global"`
	} `yaml:"env"`

	Notifications struct {
		Email struct {
			Recipients []string `yaml:"recipients"`
			OnSuccess  string   `yaml:"on_success"` // change, always, never
			OnFailure  string   `yaml:"on_failure"`
		} `yaml:"email"`
	} `yaml:"notifications"`
	BeforeInstall []string `yaml:"before_install"`
	Script        []string `yaml:"script"`
	AfterSuccess  []string `yaml:"after_success"`
}

var DefaultTravisConfig = TravisConfig{
	Language: "go",
	Go:       []string{"1.5"},
	Script:   []string{"bash -x build.sh"},
}

func init() {
	DefaultTravisConfig.Notifications.Email.OnSuccess = "always"
	DefaultTravisConfig.Notifications.Email.OnFailure = "always"
}
