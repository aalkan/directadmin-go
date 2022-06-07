package license

import (
	"fmt"
	"github.com/aalkan/directadmin-go"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"time"
)

const baseURL = "https://www.directadmin.com/clients/api"

type client struct {
	baseURL    string
	config     *directadmin.Config
	httpClient *http.Client
}

func NewClient(config *directadmin.Config) *client {
	return &client{
		baseURL: baseURL,
		config:  config,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *client) request(method, endpoint string, params interface{}, response map[string]string) error {
	request, err := http.NewRequest(method, baseURL+endpoint, nil)
	if err != nil {
		return err
	}
	request.SetBasicAuth(c.config.Username, c.config.Password)
	if params != nil {
		query := request.URL.Query()
		reflectVal := reflect.ValueOf(params).Elem()
		for i := 0; i < reflectVal.NumField(); i++ {
			value := reflectVal.Field(i).Interface().(string)
			if value != "" {
				query.Add(reflectVal.Type().Field(i).Tag.Get("json"), value)
			}
		}
		request.URL.RawQuery = query.Encode()
	}

	resp, err := c.httpClient.Do(request)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status code: %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	data , err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	a, err := url.ParseQuery(string(data))
	if err != nil {
		return err
	}
	for key, val := range a {
			response[key]= val[0]
	}
	return nil
}
