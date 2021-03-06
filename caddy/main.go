// By moving the application's package main logic into
// a package other than main, it becomes much easier to
// wrap caddy for custom builds that are go-gettable.
// https://caddy.community/t/my-wish-for-0-9-go-gettable-custom-builds/59?u=matt

package main

import "github.com/johnthethird/caddy/caddy/caddymain"

var run = caddymain.Run // replaced for tests

func main() {
	run()
}
