README for gb-path
==================

`gb-path` is a plugin for `gb`, the project based build tool for Go. It helps 
you use tools that still rely on the `$GOPATH` environment variable to be set 
to interoperate with `gb`-based projects.

Installation
------------

First, install `gb`. You can find more information about this on http://getgb.io/, but a simple

	go get github.com/constabulary/gb/...

should be enough to get started. After that, run

	go get github.com/akrennmair/gb-path

and you're done. Running `gb path` in your shell should give you output like this:

	$ gb path
	GOPATH="/home/username/go:/home/username/go/vendor"; export GOPATH;

This output can be used in conjunction with `eval`. Running

	$ eval `gb path`

will set your `$GOPATH` environment variable to the path(s) that were 
determined by `gb`. This makes it easy to switch your `$GOPATH` between your 
`gb`-based projects while still being able to use tools that rely on `$GOPATH` 
to bet set.
