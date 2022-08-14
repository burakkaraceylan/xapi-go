package tests

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/burakkaraceylan/xapi-go/pkg/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type streamResponse struct {
	ID      int               `json:"id"`
	Args    map[string]string `json:"args"`
	Headers map[string]string `json:"headers"`
	Origin  string            `json:"origin"`
	URL     string            `json:"url"`
}

type RequestsTestSuite struct {
	suite.Suite
}

func (suite *RequestsTestSuite) TestRequest() {
	req := client.Request{}
	_, err := req.Init()
	assert.EqualError(suite.T(), err, "method can't be empty")

	req.Method = "GET"
	_, err = req.Init()
	assert.EqualError(suite.T(), err, "url can't be empty")

	req.URL = "https://httpbin.org/get"
	req.Headers = &map[string]string{"Foo": "Bar", "Foo2": "Bar2"}
	req.QueryParams = &map[string]string{"Foo": "Bar"}

	request, err := req.Init()

	assert.Nil(suite.T(), err)

	client := http.Client{}
	resp, err := client.Do(request)
	assert.Nil(suite.T(), err)

	b, err := io.ReadAll(resp.Body)
	assert.Nil(suite.T(), err)

	rs := streamResponse{}
	err = json.Unmarshal(b, &rs)
	assert.Nil(suite.T(), err)

	assert.Equal(suite.T(), resp.StatusCode, 200)
	assert.Contains(suite.T(), rs.Headers, "Foo")
	assert.Contains(suite.T(), rs.Args, "Foo")
}

func (suite *RequestsTestSuite) TestResponse() {
	req := client.Request{
		Method: "GET",
		URL:    "https://httpbin.org/get",
	}

	req2, err := req.Init()
	assert.Nil(suite.T(), err)

	httpc := http.Client{}

	resp, err := httpc.Do(req2)
	assert.Nil(suite.T(), err)

	response := client.Response{
		Status:   resp.StatusCode,
		Request:  req2,
		Response: resp,
	}

	expected := fmt.Sprintf("CODE: %d\nURL: %s\n", resp.StatusCode, req2.URL)
	assert.Equal(suite.T(), response.String(), expected)
}

func TestRequestsTestSuite(t *testing.T) {
	suite.Run(t, new(RequestsTestSuite))
}
