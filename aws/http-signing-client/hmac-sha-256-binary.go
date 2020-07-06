package awshttpsigningclient

import (
	"crypto/hmac"
	"crypto/sha256"

	"github.com/jmckee46/deployer/flaw"
	"github.com/jmckee46/deployer/logger"
)

func hmacSHA256Binary(key []byte, content string) []byte {
	mac := hmac.New(sha256.New, key)
	_, err := mac.Write([]byte(content))

	if err != nil {
		flawError := flaw.From(err).Wrap("cannot hmacSHA256Binary")
		logger.Panic("hmac-sha256-binary", flawError)
	}

	return mac.Sum(nil)
}
