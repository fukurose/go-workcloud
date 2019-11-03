package workcloud

import (
	"regexp"
)

const (
	tokenRegrexp = "<meta content=\"(.+?)\" name=\"csrf-token\" />"
	titleRegexp  = "<title>(.+?)</title>"
)

// get csrf token.
func getToken(doc string) string {
	tokenRegexp := regexp.MustCompile(tokenRegrexp)
	return tokenRegexp.FindStringSubmatch(doc)[1]
}

// isLogin return true only when you are logged in.
func isLogin(doc string) bool {
	title := regexp.MustCompile(titleRegexp).FindStringSubmatch(doc)[1]
	return title != "ログイン" && title != "Login"
}
