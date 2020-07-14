package main

import (
	"github.com/jmckee46/deployer/flaw"
)

func initCertbotEnv() flaw.Flaw {
	// set up cerbot-env
	// certbotBytes, err := exec.Command(
	// 	"certbot",
	// 	"certonly",
	// 	"--dns-route53",
	// 	"--agree-tos",
	// 	"--config-dir",
	// 	"certbot/files",
	// 	"--domains",
	// 	"myapptest.net",
	// 	"--domains",
	// 	"*.myapptest.net",
	// 	"--email",
	// 	"jmckee3@mac.com",
	// 	"--logs-dir",
	// 	"certbot/log-files",
	// 	"--no-eff-email",
	// 	"--preferred-challenges",
	// 	"dns",
	// 	"--renew-by-default",
	// 	"--work-dir",
	// 	"work-dir",
	// 	"--test-cert",
	// ).Output()

	// fmt.Println("certbotString:", string(certbotBytes))
	// if err != nil {
	// 	return flaw.From(err).Wrap("cannot set up certbot")
	// }

	// certbot
	//   certonly
	//     --dns-route53
	//     --agree-tos
	//     --config-dir certbot/files
	//     --domains myapptest.net
	//     --domains *.myapptest.net
	//     --email jmckee3@mac.com
	//     --logs-dir certbot/log-files
	//     --no-eff-email
	//     --preferred-challenges dns
	//     --renew-by-default
	//     --work-dir work-dir
	//     --test-cert

	return nil
}
