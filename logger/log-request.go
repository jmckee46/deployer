package logger

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

// LogRequest logs the request and can log the request body too
func LogRequest(tag string, r *http.Request) {
	// logging the body is handy for debugging, but
	// may be too much overhead for production...
	buf, bodyErr := ioutil.ReadAll(r.Body)
	if bodyErr != nil {
		log.Print("bodyErr ", bodyErr.Error())
	}
	rdr1 := bytes.NewBuffer(buf)
	rdr2 := ioutil.NopCloser(bytes.NewBuffer(buf))
	// fmt.Printf("BODY: %s\n", rdr1)
	r.Body = rdr2

	message := map[string]interface{}{
		"body":             rdr1.String(),
		"headers":          sanitizedHeaders(r.Header),
		"host":             r.Host,
		"method":           r.Method,
		"path":             r.URL.Path,
		"query-parameters": r.URL.Query(),
		"raw-query":        r.URL.RawQuery,
	}

	Info(tag, message)
}

func sanitizedHeaders(headers map[string][]string) map[string][]string {
	sanitizedHeaders := make(map[string][]string)

	for key, value := range headers {
		if key != "Authorization" {
			sanitizedHeaders[key] = value
		}
	}

	return sanitizedHeaders
}
