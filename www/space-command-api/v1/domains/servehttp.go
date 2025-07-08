package verboten

import (
	"net/http"

	"github.com/reiver/go-erorr"
	"github.com/reiver/go-http500"
	"github.com/reiver/go-json"

	"github.com/reiver/space-command/srv/http"
)

const path string = "/space-command-api/v1/domains"

func init() {
	var handler http.Handler = http.HandlerFunc(serveHTTP)

	err := httpsrv.Mux.HandlePath(handler, path)
	if nil != err {
		e := erorr.Errorf("problem registering http-handler with path-mux for path %q: %w", path, err)
		panic(e)
	}
}

func serveHTTP(responsewriter http.ResponseWriter, request *http.Request) {

	if nil == responsewriter {
		return
	}

	var temp map[string]any = map[string]any{
		"domains": []string{
			"bing.link",
			"bong.app",
			"bang.ooo",
			"its.xyz",
			"boomerang.social",
		},
	}

	bytes, err := json.Marshal(temp)
	if nil != err {
		http500.InternalServerError(responsewriter, request)
		return
	}

	responsewriter.Header().Add("Content-Type", "application/json")
	responsewriter.Write(bytes)
}
