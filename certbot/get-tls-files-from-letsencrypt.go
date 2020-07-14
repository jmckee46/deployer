package certbot

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/jmckee46/deployer/flaw"
)

// GetTLSFilesFromLetsencrypt gets tls files from Let's Encrypt
func GetTLSFilesFromLetsencrypt() flaw.Flaw {
	fmt.Println("getting files from lets encrypt...")
	domain := os.Getenv("DE_DOMAIN")

	_, err := exec.Command(
		"certbot",
		"certonly",
		"--dns-route53",
		"--agree-tos",
		"--config-dir",
		"certbot/files",
		"--domains",
		domain,
		"--domains",
		"*."+domain,
		"--email",
		os.Getenv("DE_EMAIL"),
		"--logs-dir",
		"certbot/log-files",
		"--no-eff-email",
		"--preferred-challenges",
		"dns",
		"--renew-by-default",
		"--work-dir",
		"work-dir",
		"--test-cert",
	).Output()

	// fmt.Println("certbotString:", string(certbotBytes))
	if err != nil {
		return flaw.From(err).Wrap("cannot get tls files from lets encrypt")
	}

	return nil
}
