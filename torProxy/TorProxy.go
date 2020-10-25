package torProxy

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/proxy"
	"net/http"
	"net/url"

	"gotor/loggerFactory"
	"io/ioutil"
)

const (
	NOT_URL_PROVIDED = "Please, provide a valid url"
)

type TorProxy struct {
	client *http.Client
	logger *log.Logger
}

func NewTorProxy(torProxyUrl string) (*TorProxy, error) {
	logger := loggerFactory.NewLogger()
	tbProxyURL, err := url.Parse(torProxyUrl)

	if err != nil {
		logger.Fatal("Failed to parse proxy URL: %v\n", err)
		return nil, err
	}

	tbDialer, err := proxy.FromURL(tbProxyURL, proxy.Direct)
	if err != nil {
		logger.Fatal("Failed to obtain proxy dialer: %v\n", err)
		return nil, err
	}

	return &TorProxy{
		logger: logger,
		client: &http.Client{Transport: &http.Transport{Dial: tbDialer.Dial}},
	}, nil

}

func (tb *TorProxy) Get(url string) (string, error) {

	if url == "" {
		return "", errors.New(NOT_URL_PROVIDED)
	}

	resp, err := tb.client.Get(url)

	if err != nil {
		tb.logger.Fatal("Failed to issue GET request: %v\n", err)
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		tb.logger.Fatal("Failed to read the body: %v\n", err)
		return "", err
	}

	return string(body), nil
}
