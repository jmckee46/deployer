package flaw

import "regexp"

func stripPathname(pathname string) string {
	rgx := regexp.MustCompile(".+/jmckee46/deployer/(.+)")

	pathMatches := rgx.FindAllStringSubmatch(pathname, -1)

	if pathMatches == nil {
		return pathname
	}

	return pathMatches[0][1]
}
