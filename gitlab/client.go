package gitlab

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/Frontware/GitLabBack/config"
)

const (
	defaultBaseURL = "https://gitlab.com/"
	apiVersionPath = "api/v4/"
	userAgent      = "GitLabBack"
)

// Client basic struct for GitLab client
type Client struct {
	client *http.Client

	baseURL *url.URL

	token string

	UserAgent string
}

// New creates new GitLab client
func New(conf *config.Config) *Client {
	if conf == nil {
		return nil
	}

	c := &Client{
		client: http.DefaultClient,
		token:  conf.Token,
	}

	if conf.BaseURL != "" {
		c.setBaseURL(conf.BaseURL)
	} else {
		c.setBaseURL(defaultBaseURL)
	}

	return c
}

// setBaseURL assigns the base url to client
func (c *Client) setBaseURL(urlStr string) error {
	if !strings.HasSuffix(urlStr, "/") {
		urlStr += "/"
	}

	baseURL, err := url.Parse(urlStr)
	if err != nil {
		return err
	}

	if !strings.HasSuffix(baseURL.Path, apiVersionPath) {
		baseURL.Path += apiVersionPath
	}

	c.baseURL = baseURL

	return nil
}

// NewRequest creates a new http request for GitLab api
func (c *Client) NewRequest(method, path string, query map[string]string) (*http.Request, error) {
	u := *c.baseURL

	unescaped, err := url.PathUnescape(path)
	if err != nil {
		return nil, err
	}

	u.Path = u.Path + unescaped

	if query != nil {
		q := u.Query()
		for k, v := range query {
			q.Set(k, v)
		}
		u.RawQuery = q.Encode()
	}

	req, err := http.NewRequest(method, u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Private-Token", c.token)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", userAgent)

	return req, nil
}

// Do requests the api and store response in struct
func (c *Client) Do(req *http.Request, v interface{}) (err error) {

	if v == nil {
		return nil
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, v)

	return err
}
