module github.com/aquaproj/aqua-installer

go 1.17

require (
	github.com/adrg/xdg v0.4.0
	github.com/mattn/go-shellwords v1.0.12
	// go-githubactions v1.0.0 requires Go 1.18, but GopherJS doesn't support Go 1.18 yet
	github.com/sethvargo/go-githubactions v0.5.2
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	golang.org/x/sys v0.0.0-20211025201205-69cdffdb9359 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)
