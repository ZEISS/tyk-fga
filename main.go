package main

import (
	"net/http"

	"github.com/TykTechnologies/tyk/ctx"
)

func FGAAuth(w http.ResponseWriter, r *http.Request) {
	conf := ctx.GetOASDefinition(r)
	conf.Paths.Find("/auth").Options.Extensions["x-fiber-authz-fga"] = "FGA"
}

func main() {}
