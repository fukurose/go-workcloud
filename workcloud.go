package workcloud

import (
	"errors"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"
)

const (
	loginURL = "https://workcloud.jp/login"
)

// WorkCloud is the http.Client for WorkCloud.
type WorkCloud struct {
	client *http.Client
	user   string
	pass   string
	agent  string
}

// New returns a new WorkCloud.
func New(user, pass string) *WorkCloud {
	jar, _ := cookiejar.New(nil)
	return &WorkCloud{
		client: &http.Client{Jar: jar},
		user:   user,
		pass:   pass,
		agent:  "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:62.0) Gecko/20100101 Firefox/62.0",
	}
}

// PunchTime is punch in / out at current time
func (workCloud *WorkCloud) PunchTime(key string) error {
	doc, err := workCloud.login()
	if err != nil {
		return err
	}

	if !isLogin(doc) {
		return errors.New("login failed")
	}

	// get token
	token := getToken(doc)
	if token == "" {
		return errors.New("Not find token for punch time")
	}

	// set up  check in / out info
	today := time.Now()
	values := url.Values{}
	values.Add("device_time", today.String())
	values.Add("return", "entry_panel")

	url := "https://workcloud.jp/timesheets/" + today.Format("2006-01-02") + "/work_hours/" + key + "/1/timestamp"
	workCloud.post(url, token, values)

	return nil
}

// login to workCloudcan.
func (workCloud *WorkCloud) login() (string, error) {
	doc, err := workCloud.get(loginURL)
	if err != nil {
		return "", err
	}

	// get token
	token := getToken(doc)
	if token == "" {
		return "", errors.New("Not find token for login")
	}

	// set up login info
	values := url.Values{}
	values.Add("authenticity_token", token)
	values.Add("commit", "login")
	values.Add("user[login]", workCloud.user)
	values.Add("user[password]", workCloud.pass)
	values.Add("user[remember_me]", "1")

	return workCloud.post(loginURL, "", values)
}
