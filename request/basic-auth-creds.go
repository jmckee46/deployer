package request

// BasicAuthCredentials returns the request basic auth creds
func (r *Request) BasicAuthCredentials() (string, string, bool) {
	return r.BasicAuth()
}
