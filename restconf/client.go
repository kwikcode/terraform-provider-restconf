package restconf

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

type Client struct {
	Username   string
	Password   string
	DeviceHost string
	DevicePort string
	HttpClient *http.Client
}

func NewClient(username, password, deviceHost, devicePort string) (*Client, error) {
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
		DeviceHost: deviceHost,
		DevicePort: devicePort,
		HttpClient: httpClient,
	}

	return client, nil
}

func (c *Client) CreateConfigBlock(ctx context.Context, path, content string) error {
	url := fmt.Sprintf("https://%s:%s/restconf/data/%s", c.DeviceHost, c.DevicePort, path)
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
	url := fmt.Sprintf("https://%s:%s/restconf/data/%s", c.DeviceHost, c.DevicePort, path)

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

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(bodyBytes), nil
}

func (c *Client) UpdateConfigBlock(ctx context.Context, path, content string) error {
	url := fmt.Sprintf("https://%s:%s/restconf/data/%s", c.DeviceHost, c.DevicePort, path)
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
	url := fmt.Sprintf("https://%s:%s/restconf/data/%s", c.DeviceHost, c.DevicePort, path)

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
