package awshttpsigningclient

import (
	"net/http"

	"github.com/jmckee46/deployer/http-client"
)

func New() (*httpclient.Client, *Credentials) {
	state := NewState()

	transport := Transport{
		State: state,
	}

	httpClient := &http.Client{
		Transport: http.RoundTripper(transport),
	}

	client := &httpclient.Client{
		Client: httpClient,
	}

	return client, state.Creds
}
