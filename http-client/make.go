package httpclient

import (
	"net/http"

	"github.com/jmckee46/deployer/flaw"
)

func Make(req *http.Request) (*http.Response, flaw.Flaw) {
	client := FromPool()

	res, flawError := flaw.FromHTTPResponse(
		client.Do(req),
	)

	ToPool(client)

	if flawError != nil {
		return nil, flawError
	}

	return res, nil
}
