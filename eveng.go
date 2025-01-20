package evengsdk

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"github.com/hashicorp/go-cleanhttp"
	"github.com/hashicorp/go-retryablehttp"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	userAgent = "go-evengapi"
)

type Response struct {
	Code    json.Number `json:"code"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message"`
	Status  string      `json:"status"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Html5    string `json:"html5,omitempty"`
}

type Auth struct {
	Email    string `json:"email"`
	Folder   string `json:"folder"`
	Lab      string `json:"lab"`
	Lang     string `json:"lang"`
	Name     string `json:"name"`
	Role     string `json:"role"`
	Tenant   int    `json:"tenant"`
	Html5    int    `json:"html5"`
	Username string `json:"username"`
}

type Client struct {
	client                    *retryablehttp.Client
	baseURL                   *url.URL
	username, password, Html5 string
	cookie                    *http.Cookie
	UserAgent                 string
	Lab                       *LabService
	Node                      *NodeService
	Folder                    *FolderService
	Network                   *NetworkService
	mutex                     *sync.Mutex
}

func newClient() (*Client, error) {
	c := &Client{UserAgent: userAgent}

	c.client = &retryablehttp.Client{
		ErrorHandler: retryablehttp.PassthroughErrorHandler,
		HTTPClient:   cleanhttp.DefaultPooledClient(),
		RetryWaitMin: 100 * time.Millisecond,
		RetryWaitMax: 400 * time.Millisecond,
		RetryMax:     5,
		CheckRetry: func(ctx context.Context, resp *http.Response, err error) (bool, error) {
			return false, nil
		},
	}
	(c.client.HTTPClient.Transport).(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	c.Lab = &LabService{client: c}
	c.Node = &NodeService{client: c}
	c.Folder = &FolderService{client: c}
	c.Network = &NetworkService{client: c}
	c.mutex = &sync.Mutex{}
	return c, nil
}

// NewBasicAuthClient returns a new Client with basic auth
// Html5 is optional and can be set to "1" to enable Apache Guacamole
func NewBasicAuthClient(username, password, Html5, baseURL string) (*Client, error) {
	client, err := newClient()
	if err != nil {
		return nil, err
	}

	client.username = username
	client.password = password
	client.Html5 = Html5
	err = client.setBaseURL(baseURL)
	if err != nil {
		return nil, err
	}
	return client, client.login()
}

func (c *Client) login() error {
	login := &Login{
		Username: c.username,
		Password: c.password,
		Html5:    c.Html5,
	}
	body, _ := json.Marshal(login)
	everesp, resp, _ := c.Do(context.Background(), "POST", "api/auth/login", body)
	if everesp.Status != "success" {
		return errors.New("Login Failed")
	}
	c.cookie = resp.Cookies()[0]
	return nil
}

func (c *Client) GetAuth() (*Auth, error) {
	eve, _, err := c.Do(context.Background(), "GET", "api/auth", nil)
	if err != nil {
		return nil, err
	}
	data, err := json.Marshal(eve.Data)
	if err != nil {
		return nil, err
	}
	var auth Auth
	err = json.Unmarshal(data, &auth)
	if err != nil {
		return nil, err
	}
	return &auth, nil
}

func (c *Client) Do(ctx context.Context, method, url string, body []byte) (*Response, *http.Response, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	req, err := retryablehttp.NewRequest(method, c.baseURL.String()+url, bytes.NewBuffer(body))
	req.Close = true
	if err != nil {
		return &Response{Code: "0", Message: "Failed to create request"}, nil, err
	}
	if c.cookie != nil {
		req.AddCookie(c.cookie)
	}
	req.Header.Set("Content-Type", "application/json")
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return &Response{Code: "0", Message: "Failed to send request"}, nil, err
	}
	defer resp.Body.Close()
	var response Response
	bodystr, err := io.ReadAll(resp.Body)
	if err != nil {
		return &Response{Code: json.Number(strconv.Itoa(resp.StatusCode)), Message: resp.Status}, nil, err
	}
	err = json.Unmarshal(bodystr, &response)
	if err != nil {
		return &Response{Code: json.Number(strconv.Itoa(resp.StatusCode)), Message: resp.Status}, nil, err
	}
	if response.Status != "success" {
		return &response, resp, errors.New(response.Message)
	}
	if status, _ := response.Code.Int64(); !(200 <= status && status <= 300) {
		return &response, resp, errors.New(response.Message)
	}

	return &response, resp, nil
}

func (c *Client) BaseURL() *url.URL {
	u := *c.baseURL
	return &u
}

// setBaseURL sets the base URL for API requests to a custom endpoint.
func (c *Client) setBaseURL(urlStr string) error {
	if !strings.HasSuffix(urlStr, "/") {
		urlStr += "/"
	}
	baseURL, err := url.Parse(urlStr)
	if err != nil {
		return err
	}

	c.baseURL = baseURL

	return nil
}
