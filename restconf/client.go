package restconf

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	Username   string
	Password   string
	Host       string
	HttpClient *http.Client
}

func NewClient(username, password, host string) (*Client, error) {
	hostURL, err := url.Parse(host)
	if err != nil {
		return nil, err
	}

	if hostURL.Scheme == "" {
		hostURL.Scheme = "https"
	}

	if hostURL.Port() == "" {
		hostURL.Host = fmt.Sprintf("%s:%s", hostURL.Host, "443")
	}

	normalizedHost := fmt.Sprintf("%s://%s", hostURL.Scheme, hostURL.Host)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			return net.DialTimeout(network, addr, 5*time.Second)
		},
	}

	httpClient := &http.Client{
		Transport: tr,
		Timeout:   10 * time.Second,
	}

	client := &Client{
		Username:   username,
		Password:   password,
		Host:       normalizedHost,
		HttpClient: httpClient,
	}

	return client, nil
}

func (c *Client) CreateConfigBlock(ctx context.Context, path, content string) error {
	url := fmt.Sprintf("%s/%s", c.Host, path)
	reqBody := []byte(content)

	req, err := http.NewRequestWithContext(ctx, "PUT", url, bytes.NewReader(reqBody))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/yang-data+json")
	req.SetBasicAuth(c.Username, c.Password)

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}

func (c *Client) ReadConfigBlock(ctx context.Context, path string) (string, error) {
	url := fmt.Sprintf("%s/%s", c.Host, path)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Accept", "application/yang-data+json")
	req.SetBasicAuth(c.Username, c.Password)

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(bodyBytes), nil
}

func (c *Client) UpdateConfigBlock(ctx context.Context, path, content string) error {
	url := fmt.Sprintf("%s/%s", c.Host, path)
	reqBody := []byte(content)

	req, err := http.NewRequestWithContext(ctx, "PUT", url, bytes.NewReader(reqBody))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/yang-data+json")
	req.SetBasicAuth(c.Username, c.Password)

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}

func (c *Client) DeleteConfigBlock(ctx context.Context, path string) error {
	url := fmt.Sprintf("%s/%s", c.Host, path)

	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return err
	}

	req.SetBasicAuth(c.Username, c.Password)

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}
