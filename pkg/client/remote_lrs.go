package client

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/burakkaraceylan/xapi-go/pkg/resources/about"
	"github.com/burakkaraceylan/xapi-go/pkg/resources/statement"
	"github.com/burakkaraceylan/xapi-go/pkg/resources/statement/properties"
)

// RemoteLRS represents a connection to an LRS
type RemoteLRS struct {
	Endpoint string
	Version  string
	Username string
	Password string
	Auth     string
}

func (lrs *RemoteLRS) newRequest(method string, resource string, headers *map[string]string, params *map[string]string, content *string) *Request {
	url := lrs.Endpoint + resource

	lrs_req := Request{
		Method:      method,
		URL:         url,
		Headers:     headers,
		QueryParams: params,
		Content:     content,
	}

	return &lrs_req
}

func (lrs *RemoteLRS) sendRequest(req *http.Request) (*Response, error) {
	client := &http.Client{}

	req.Header.Add("X-Experience-API-Version", lrs.Version)
	req.Header.Add("Content-Type", "application/json")

	if len(lrs.Auth) > 0 {
		req.Header.Add("Authorization", lrs.Auth)
	}

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	return &Response{
		Status:   resp.StatusCode,
		Request:  req,
		Response: resp,
	}, nil
}

// SaveStatement is used to save a statement to the record store
func (lrs *RemoteLRS) SaveStatement(statement statement.Statement) ([]string, *Response, error) {
	lrs_req := lrs.newRequest("POST", "statements", nil, nil, nil)

	if statement.ID != nil && len(*statement.ID) != 0 {
		lrs_req.Method = "PUT"
		params := map[string]string{"statementId": *statement.ID}
		lrs_req.QueryParams = &params
	}

	b, err := json.Marshal(statement)

	if err != nil {
		return nil, nil, err
	}

	str := string(b)
	lrs_req.Content = &str

	req, err := lrs_req.Init()

	if err != nil {
		return nil, nil, err
	}

	resp, err := lrs.sendRequest(req)

	if err != nil {
		return nil, nil, err
	}

	// If we used PUT we don't expect a return value
	if statement.ID != nil {
		return nil, resp, nil
	}

	// If we used POST we expect an array of uuid strings
	var idList []string

	b, err = io.ReadAll(resp.Response.Body)

	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(b, &idList)

	if err != nil {
		return nil, nil, err
	}

	return idList, resp, nil

}

// SaveStatements is used to save multiple statement to the record store
func (lrs *RemoteLRS) SaveStatements(statements []statement.Statement) ([]string, *Response, error) {
	lrs_req := lrs.newRequest("POST", "statements", nil, nil, nil)

	b, err := json.Marshal(statements)

	if err != nil {
		return nil, nil, err
	}

	str := string(b)
	lrs_req.Content = &str

	req, err := lrs_req.Init()

	if err != nil {
		return nil, nil, err
	}

	resp, err := lrs.sendRequest(req)

	if err != nil {
		return nil, nil, err
	}

	var idList []string

	b, err = io.ReadAll(resp.Response.Body)

	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(b, &idList)

	if err != nil {
		return nil, nil, err
	}

	return idList, resp, nil

}

// GetStatement is used to fetch a single statement from record store
func (lrs *RemoteLRS) GetStatement(id string) (*statement.Statement, *Response, error) {
	lrs_request := lrs.newRequest("GET", "statements", nil, &map[string]string{"statementId": id}, nil)

	req, err := lrs_request.Init()

	if err != nil {
		return nil, nil, err
	}

	lrs_resp, err := lrs.sendRequest(req)

	if err != nil {
		return nil, nil, err
	}

	statement := &statement.Statement{}

	lrs_resp.Bind(statement)

	return statement, lrs_resp, nil
}

// TODO: Return a Statement struct not a Response
// GetVoidedStatement is used to fetch a single voided statement from record store
func (lrs *RemoteLRS) GetVoidedStatement(id string) (*Response, error) {
	lrs_request := lrs.newRequest("GET", "statements", nil, &map[string]string{"voidedStatementId": id}, nil)

	req, err := lrs_request.Init()

	if err != nil {
		return nil, err
	}

	lrs_resp, err := lrs.sendRequest(req)

	if err != nil {
		return nil, err
	}

	return lrs_resp, nil
}

// QueryParams represents query parameters of a statement
type QueryParams struct {
	StatementID       *string
	VoidedStatementId *string
	Agent             *properties.Actor
	Verb              *properties.Verb
	Activity          *properties.Object
	Registeration     *string
	RelatedActivities *bool
	RelatedAgents     *bool
	Since             *time.Time
	Until             *time.Time
	Limit             *int64
	Format            *string
	Attachments       *bool
	Ascending         *bool
}

// Map is used to generate a dictionary from a QueryParams object
func (q *QueryParams) Map() map[string]string {
	params := make(map[string]string)

	if q.StatementID != nil {
		params["statementId"] = *q.VoidedStatementId
	}

	if q.VoidedStatementId != nil {
		params["voidedStatementId"] = *q.StatementID
	}

	if q.Agent != nil {
		b, err := json.Marshal(*q.Agent)

		if err == nil {
			params["agent"] = string(b)
		}
	}

	if q.Verb != nil {
		params["verb"] = string(q.Verb.ID)
	}

	if q.Activity != nil {
		params["activity"] = string(q.Activity.ID)
	}

	if q.Registeration != nil {
		params["registeration"] = *q.Registeration
	}

	if q.RelatedActivities != nil {
		params["related_activities"] = strconv.FormatBool(*q.RelatedActivities)
	}

	if q.RelatedAgents != nil {
		params["related_agents"] = strconv.FormatBool(*q.RelatedAgents)
	}

	if q.Since != nil {
		params["since"] = q.Since.String()
	}

	if q.Until != nil {
		params["until"] = q.Until.String()
	}

	if q.Limit != nil {
		params["limit"] = strconv.FormatInt(*q.Limit, 10)
	}

	if q.Format != nil {
		if *q.Format == "ids" || *q.Format == "exact" || *q.Format == "canonical" {
			params["format"] = *q.Format
		}
	}

	if q.Attachments != nil {
		params["attachments"] = strconv.FormatBool(*q.Attachments)
	}

	if q.Ascending != nil {
		params["ascending"] = strconv.FormatBool(*q.Ascending)
	}

	return params

}

// QueryStatements is used to query the statements on the LRS
func (lrs *RemoteLRS) QueryStatements(params *QueryParams) (*statement.MoreStatements, error) {

	var query_params map[string]string

	if params != nil {
		query_params = params.Map()
	}

	lrs_request := lrs.newRequest("GET", "statements", nil, &query_params, nil)
	req, err := lrs_request.Init()

	if err != nil {
		return nil, err
	}

	lrs_resp, err := lrs.sendRequest(req)

	if err != nil {
		return nil, err
	}

	statements := &statement.MoreStatements{}

	if err := lrs_resp.Bind(statements); err != nil {
		return nil, err
	}

	return statements, nil
}

// About is used to fetch information about the LRS
func (lrs *RemoteLRS) About() (*about.About, error) {
	lrs_request := lrs.newRequest("GET", "about", nil, nil, nil)

	req, err := lrs_request.Init()

	if err != nil {
		return nil, err
	}

	lrs_resp, err := lrs.sendRequest(req)

	if err != nil {
		return nil, err
	}

	about := &about.About{}

	if err := lrs_resp.Bind(about); err != nil {
		return nil, err
	}

	return about, nil

}

// NewRemoteLRS is used to construct and initialize a RemoteLRS object
func NewRemoteLRS(endpoint string, version string, authentication ...string) (*RemoteLRS, error) {

	lrs := RemoteLRS{
		Endpoint: endpoint,
		Version:  version,
	}

	if len(authentication) == 0 {
		return nil, errors.New("authentication params not provided")
	}

	if len(authentication) == 3 {
		return nil, errors.New("too many authentication params")
	}

	if len(authentication) == 1 {
		lrs.Auth = authentication[0]
		return &lrs, nil
	}

	if len(authentication) == 2 {
		lrs.Username = authentication[0]
		lrs.Password = authentication[1]
		return &lrs, nil
	}

	return nil, errors.New("unknown error")
}
