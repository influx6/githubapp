// Package apps defines a package which exposes a app interface which allows us
// easily talk to a github application which is hosted on a remote server and will contain
// all necessaary internal logic to talk to this giving repos on Github.
package apps

// contains given fields using with the app pkg
const (
	GITHUB              = "https://github.com/"
	GITHUBLOGIN         = "https://github.com/login/oauth"
	GITHUBAuthorization = "https://github.com/login/oauth/authorize"
)

// App defines a struct which embodies the underline Github Application access
// token details.
type App struct {
	Name   string `json:"name" toml:"name" yaml:"token"`
	Token  string `json:"token" toml:"token" yaml:"token"`
	Secret string `json:"secret" toml:"secret" yaml:"secret"`
}
