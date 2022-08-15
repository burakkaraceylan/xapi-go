package client

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/burakkaraceylan/xapi-go/pkg/resources/about"
	"github.com/burakkaraceylan/xapi-go/pkg/resources/documents"
	"github.com/burakkaraceylan/xapi-go/pkg/resources/statement"
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
		return nil, nil, fmt.Errorf("failed to marshal: %w", err)
	}

	str := string(b)
	lrs_req.Content = &str

	req, err := lrs_req.Init()

	if err != nil {
		return nil, nil, fmt.Errorf("failed to init request: %w", err)
	}

	resp, err := lrs.sendRequest(req)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to send request: %w", err)
	}

	// If we used PUT we don't expect a return value
	if statement.ID != nil {
		return nil, resp, nil
	}

	// If we used POST we expect an array of uuid strings
	var idList []string

	b, err = io.ReadAll(resp.Response.Body)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to read response body: %w", err)
	}

	err = json.Unmarshal(b, &idList)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return idList, resp, nil

}

// SaveStatements is used to save multiple statement to the record store
func (lrs *RemoteLRS) SaveStatements(statements []statement.Statement) ([]string, *Response, error) {
	lrs_req := lrs.newRequest("POST", "statements", nil, nil, nil)

	b, err := json.Marshal(statements)

	if err != nil {
		return nil, nil, fmt.Errorf("marshaling failed: %w", err)
	}

	str := string(b)
	lrs_req.Content = &str

	req, err := lrs_req.Init()

	if err != nil {
		return nil, nil, fmt.Errorf("failed to init request: %w", err)
	}

	resp, err := lrs.sendRequest(req)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to send request: %w", err)
	}

	var idList []string

	b, err = io.ReadAll(resp.Response.Body)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to read response body: %w", err)
	}

	err = json.Unmarshal(b, &idList)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return idList, resp, nil

}

// GetStatement is used to fetch a single statement from record store
func (lrs *RemoteLRS) GetStatement(id string) (*statement.Statement, *Response, error) {
	lrs_request := lrs.newRequest("GET", "statements", nil, &map[string]string{"statementId": id}, nil)

	req, err := lrs_request.Init()

	if err != nil {
		return nil, nil, fmt.Errorf("failed to init request: %w", err)
	}

	lrs_resp, err := lrs.sendRequest(req)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to send request: %w", err)
	}

	if lrs_resp.Response.StatusCode != 200 {
		return nil, lrs_resp, nil
	}

	statement := &statement.Statement{}

	if err := lrs_resp.Bind(statement); err != nil {
		return nil, nil, fmt.Errorf("failed to bind response: %w", err)
	}

	return statement, lrs_resp, nil
}

// GetVoidedStatement is used to fetch a single voided statement from record store
func (lrs *RemoteLRS) GetVoidedStatement(id string) (*statement.Statement, *Response, error) {
	lrs_request := lrs.newRequest("GET", "statements", nil, &map[string]string{"voidedStatementId": id}, nil)

	req, err := lrs_request.Init()

	if err != nil {
		return nil, nil, fmt.Errorf("failed to init request: %w", err)
	}

	lrs_resp, err := lrs.sendRequest(req)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to send request: %w", err)
	}

	if lrs_resp.Response.StatusCode != 200 {
		return nil, lrs_resp, nil
	}

	statement := &statement.Statement{}

	if err := lrs_resp.Bind(statement); err != nil {
		return nil, nil, fmt.Errorf("failed to bind response: %w", err)
	}

	return statement, lrs_resp, nil
}

// QueryParams represents query parameters of a statement
type StatementQueryParams struct {
	StatementID       *string
	VoidedStatementId *string
	Agent             *statement.Agent
	Verb              *statement.Verb
	Activity          *statement.Activity
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
func (q *StatementQueryParams) Map() map[string]string {
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
func (lrs *RemoteLRS) QueryStatements(params ...*StatementQueryParams) (*statement.StatementResult, *Response, error) {

	var query_params map[string]string

	if len(params) > 0 {
		query_params = params[0].Map()
	}

	lrs_request := lrs.newRequest("GET", "statements", nil, &query_params, nil)
	req, err := lrs_request.Init()

	if err != nil {
		return nil, nil, fmt.Errorf("failed to init request: %w", err)
	}

	lrs_resp, err := lrs.sendRequest(req)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to send request: %w", err)
	}

	result := &statement.StatementResult{}

	if err := lrs_resp.Bind(result); err != nil {
		return nil, nil, fmt.Errorf("failed to bin response: %w", err)
	}

	return result, lrs_resp, nil
}

// About is used to fetch information about the LRS
func (lrs *RemoteLRS) About() (*about.About, error) {
	lrs_request := lrs.newRequest("GET", "about", nil, nil, nil)

	req, err := lrs_request.Init()

	if err != nil {
		return nil, fmt.Errorf("failed to init request: %w", err)
	}

	lrs_resp, err := lrs.sendRequest(req)

	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	about := &about.About{}

	if err := lrs_resp.Bind(about); err != nil {
		return nil, fmt.Errorf("failed to bind response %w", err)
	}

	return about, nil

}

// GetStateIds request parameters
type GetStateIdsOptionalParams struct {
	Registration *string
	Since        *time.Time
}

// GetStateIds is used to fetch the state ids
func (lrs *RemoteLRS) GetStateIds(activity statement.Activity, agent statement.Agent, params ...*GetStateIdsOptionalParams) ([]string, *Response, error) {
	var opt *GetStateIdsOptionalParams

	if len(params) == 1 {
		if params[0] == nil {
			return nil, nil, errors.New("optional parameters can't be a nil pointer")
		}

		opt = params[0]

	}

	if len(params) > 1 {
		return nil, nil, errors.New("too many arguments")
	}

	query_params := make(map[string]string)

	query_params["activityId"] = activity.ID
	query_params["agent"] = agent.ToJSON()

	if opt != nil {
		if opt.Registration != nil {
			query_params["registration"] = *opt.Registration
		}

		if opt.Since != nil {
			query_params["since"] = opt.Since.Format(time.RFC3339Nano)
		}
	}

	lrs_request := lrs.newRequest("GET", "activities/state", nil, &query_params, nil)
	req, err := lrs_request.Init()

	if err != nil {
		return nil, nil, fmt.Errorf("failed to init request: %w", err)
	}

	lrs_resp, err := lrs.sendRequest(req)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to send request: %w", err)
	}

	if lrs_resp.Response.StatusCode != 200 {
		return nil, lrs_resp, nil
	}

	var idList []string

	if err := lrs_resp.Bind(&idList); err != nil {
		return nil, nil, fmt.Errorf("failed to bind response: %w", err)
	}

	return idList, lrs_resp, nil
}

// GetState request parameters
type GetStateOptionalParams struct {
	Registration *string
}

// GetState is used to fetch a state
func (lrs *RemoteLRS) GetState(activity statement.Activity, agent statement.Agent, stateID string, params ...*GetStateOptionalParams) (*documents.StateDocument, *Response, error) {
	var opt *GetStateOptionalParams

	if len(params) == 1 {
		if params[0] == nil {
			return nil, nil, errors.New("optional parameters can't be a nil pointer")
		}

		opt = params[0]

	}

	if len(params) > 1 {
		return nil, nil, errors.New("too many arguments")
	}

	query_params := make(map[string]string)

	query_params["activityId"] = activity.ID
	query_params["stateId"] = stateID
	query_params["agent"] = agent.ToJSON()

	if opt != nil {
		if opt.Registration != nil {
			query_params["registration"] = *opt.Registration
		}
	}

	if len(query_params["stateId"]) == 0 {
		return nil, nil, errors.New("stateId can't be null")
	}

	lrs_request := lrs.newRequest("GET", "activities/state", nil, &query_params, nil)
	req, err := lrs_request.Init()

	if err != nil {
		return nil, nil, fmt.Errorf("failed to init request: %w", err)
	}

	lrs_resp, err := lrs.sendRequest(req)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to send request: %w", err)
	}

	if lrs_resp.Response.StatusCode != 200 {
		return nil, lrs_resp, nil
	}

	content, err := io.ReadAll(lrs_resp.Response.Body)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to read response body: %w", err)
	}

	doc := documents.StateDocument{
		Agent:    agent,
		Activity: activity,
		Document: documents.Document{
			ID:      stateID,
			Content: content,
		},
	}

	if opt != nil {
		if opt.Registration != nil {
			query_params["registration"] = *opt.Registration
		}
	}

	if ts := lrs_resp.Response.Header.Get("last-modified"); len(ts) > 0 {
		t, err := time.Parse("Mon, 2 Jan 2006 15:04:05 GMT", ts)

		if err != nil {
			return nil, nil, fmt.Errorf("failed to parse time: %w", err)
		}

		doc.Timestamp = t
	}

	if ct := lrs_resp.Response.Header.Get("content-type"); len(ct) > 0 {
		doc.ContentType = ct
	}

	if etag := lrs_resp.Response.Header.Get("etag"); len(etag) > 0 {
		doc.Etag = etag
	}

	return &doc, lrs_resp, nil
}

// SaveState is used to save a state document to the LRS
func (lrs *RemoteLRS) SaveState(state *documents.StateDocument) (*documents.StateDocument, *Response, error) {
	if state == nil {
		return nil, nil, errors.New("argument can't be nil")
	}

	content := string(state.Content)

	headers := make(map[string]string)

	if len(state.ContentType) > 0 {
		headers["Content-Type"] = state.ContentType
	} else {
		headers["Content-Type"] = "application/octet-stream"
	}

	if len(state.Etag) > 0 {
		headers["If-Match"] = state.Etag
	}

	query_params := make(map[string]string)
	query_params["activityId"] = state.Activity.ID
	query_params["stateId"] = state.ID
	query_params["agent"] = state.Agent.ToJSON()

	lrs_request := lrs.newRequest("PUT", "activities/state", &headers, &query_params, &content)

	req, err := lrs_request.Init()

	if err != nil {
		return nil, nil, fmt.Errorf("failed to init request: %w", err)
	}

	lrs_resp, err := lrs.sendRequest(req)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to send request: %w", err)
	}

	return state, lrs_resp, nil
}

// DeleteState is used to delete a state (if stateId is provided) or all states related to agent/activity/registration
func (lrs *RemoteLRS) DeleteState(state *documents.StateDocument) (*Response, error) {
	if state == nil {
		return nil, errors.New("argument can't be nil")
	}

	params := make(map[string]string)

	b, err := json.Marshal(state.Agent)

	if err != nil {
		return nil, fmt.Errorf("failed to marshal: %w", err)
	}

	params["activityId"] = state.Activity.ID
	params["agent"] = string(b)

	if len(state.ID) > 0 {
		params["stateId"] = state.ID
	}

	if state.Registration != nil {
		params["registration"] = *state.Registration
	}

	headers := make(map[string]string)

	if len(state.ContentType) > 0 {
		headers["Content-Type"] = state.ContentType
	} else {
		headers["Content-Type"] = "application/octet-stream"
	}

	if len(state.Etag) > 0 {
		headers["If-Match"] = state.Etag
	}

	lrs_request := lrs.newRequest("DELETE", "activities/state", &headers, &params, nil)

	req, err := lrs_request.Init()

	if err != nil {
		return nil, fmt.Errorf("failed to init request: %w", err)
	}

	lrs_resp, err := lrs.sendRequest(req)

	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	return lrs_resp, nil
}

// Optional params for GetActivityProfileIds
type GetActivityProfileIdsOptionalParams struct {
	Since *time.Time
}

// GetActivityProfileIds is used to fetch the activity profile ids
func (lrs *RemoteLRS) GetActivityProfileIds(activity statement.Activity, params ...*GetActivityProfileIdsOptionalParams) ([]string, *Response, error) {
	var opt *GetActivityProfileIdsOptionalParams

	if len(params) == 1 {
		if params[0] == nil {
			return nil, nil, errors.New("optional parameters can't be a nil pointer")
		}

		opt = params[0]

	}

	if len(params) > 1 {
		return nil, nil, errors.New("too many arguments")
	}

	query_params := make(map[string]string)

	query_params["activityId"] = activity.ID

	if opt != nil {
		if opt.Since != nil {
			query_params["since"] = opt.Since.Format(time.RFC3339Nano)
		}
	}

	lrs_request := lrs.newRequest("GET", "activities/profile", nil, &query_params, nil)
	req, err := lrs_request.Init()

	if err != nil {
		return nil, nil, fmt.Errorf("failed to init request: %w", err)
	}

	lrs_resp, err := lrs.sendRequest(req)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to send request: %w", err)
	}

	if lrs_resp.Response.StatusCode != 200 {
		return nil, lrs_resp, nil
	}

	var idList []string

	if err := lrs_resp.Bind(&idList); err != nil {
		return nil, nil, fmt.Errorf("failed to bind response: %w", err)
	}

	return idList, lrs_resp, nil
}

// GetActivityProfile is used to fetch an actity profile
func (lrs *RemoteLRS) GetActivityProfile(activity statement.Activity, profileID string) (*documents.ActivityDocument, *Response, error) {

	query_params := make(map[string]string)

	query_params["profileId"] = profileID
	query_params["activityId"] = activity.ID

	lrs_request := lrs.newRequest("GET", "activities/profile", nil, &query_params, nil)
	req, err := lrs_request.Init()

	if err != nil {
		return nil, nil, fmt.Errorf("failed to init request: %w", err)
	}

	lrs_resp, err := lrs.sendRequest(req)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to send request: %w", err)
	}

	if lrs_resp.Response.StatusCode != 200 {
		return nil, lrs_resp, nil
	}

	content, err := io.ReadAll(lrs_resp.Response.Body)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to read body: %w", err)
	}

	doc := documents.ActivityDocument{
		Activity: activity,
		Document: documents.Document{
			ID:      profileID,
			Content: content,
		},
	}

	if ts := lrs_resp.Response.Header.Get("last-modified"); len(ts) > 0 {
		t, err := time.Parse("Mon, 2 Jan 2006 15:04:05 GMT", ts)

		if err != nil {
			return nil, nil, fmt.Errorf("failed to parse time: %w", err)
		}

		doc.Timestamp = t
	}

	if ct := lrs_resp.Response.Header.Get("content-type"); len(ct) > 0 {
		doc.ContentType = ct
	}

	if etag := lrs_resp.Response.Header.Get("etag"); len(etag) > 0 {
		doc.Etag = etag
	}

	return &doc, lrs_resp, nil
}

// SaveActivityProfile is used to save an activity profile document to the LRS
func (lrs *RemoteLRS) SaveActivityProfile(profile *documents.ActivityDocument) (*documents.ActivityDocument, *Response, error) {

	if profile == nil {
		return nil, nil, errors.New("argument can't be nil")
	}

	content := string(profile.Content)

	headers := make(map[string]string)

	if len(profile.ContentType) > 0 {
		headers["Content-Type"] = profile.ContentType
	} else {
		headers["Content-Type"] = "application/octet-stream"
	}

	if len(profile.Etag) > 0 {
		headers["If-Match"] = profile.Etag
	}

	params := make(map[string]string)
	params["activityId"] = profile.Activity.ID
	params["profileId"] = profile.ID

	lrs_request := lrs.newRequest("PUT", "activities/profile", &headers, &params, &content)

	req, err := lrs_request.Init()

	if err != nil {
		return nil, nil, fmt.Errorf("failed to init request: %w", err)
	}

	lrs_resp, err := lrs.sendRequest(req)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to send request: %w", err)
	}

	return profile, lrs_resp, nil
}

// DeleteActivityProfile is used to delete a an activty profile
func (lrs *RemoteLRS) DeleteActivityProfile(profile *documents.ActivityDocument) (*Response, error) {

	if profile == nil {
		return nil, errors.New("argument can't be nil")
	}

	params := make(map[string]string)

	params["activityId"] = profile.Activity.ID
	params["profileId"] = profile.ID

	headers := make(map[string]string)

	if len(profile.Etag) > 0 {
		headers["If-Match"] = profile.Etag
	}

	lrs_request := lrs.newRequest("DELETE", "activities/profile", &headers, &params, nil)

	req, err := lrs_request.Init()

	if err != nil {
		return nil, fmt.Errorf("failed to init request: %w", err)
	}

	lrs_resp, err := lrs.sendRequest(req)

	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	return lrs_resp, nil
}

// Optional params for GetAgentProfileIds
type GetAgentProfileIdsoptionalParams struct {
	Agent statement.Agent
	Since *time.Time
}

// GetAgentProfileIds is used to fetch the agent profile ids
func (lrs *RemoteLRS) GetAgentProfileIds(agent statement.Agent, params ...*GetAgentProfileIdsoptionalParams) ([]string, *Response, error) {
	var opt *GetAgentProfileIdsoptionalParams

	if len(params) == 1 {
		if params[0] == nil {
			return nil, nil, errors.New("optional parameters can't be a nil pointer")
		}

		opt = params[0]

	}

	if len(params) > 1 {
		return nil, nil, errors.New("too many arguments")
	}

	query_params := make(map[string]string)

	query_params["agent"] = agent.ToJSON()

	if opt != nil {
		if opt.Since != nil {
			query_params["since"] = opt.Since.Format(time.RFC3339Nano)
		}
	}

	lrs_request := lrs.newRequest("GET", "agents/profile", nil, &query_params, nil)
	req, err := lrs_request.Init()

	if err != nil {
		return nil, nil, fmt.Errorf("failed to init request: %w", err)
	}

	lrs_resp, err := lrs.sendRequest(req)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to send request: %w", err)
	}

	if lrs_resp.Response.StatusCode != 200 {
		return nil, lrs_resp, nil
	}

	var idList []string

	if err := lrs_resp.Bind(&idList); err != nil {
		return nil, nil, fmt.Errorf("failed to bind response: %w", err)
	}

	return idList, lrs_resp, nil
}

// GetAgentProfile is used to fetch an actity profile
func (lrs *RemoteLRS) GetAgentProfile(agent statement.Agent, profileID string) (*documents.AgentDocument, *Response, error) {

	query_params := make(map[string]string)

	query_params["profileId"] = profileID
	query_params["agent"] = agent.ToJSON()

	lrs_request := lrs.newRequest("GET", "agents/profile", nil, &query_params, nil)
	req, err := lrs_request.Init()

	if err != nil {
		return nil, nil, fmt.Errorf("failed to init request: %w", err)
	}

	lrs_resp, err := lrs.sendRequest(req)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to send request: %w", err)
	}

	if lrs_resp.Response.StatusCode != 200 {
		return nil, lrs_resp, nil
	}

	content, err := io.ReadAll(lrs_resp.Response.Body)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to read body: %w", err)
	}

	doc := documents.AgentDocument{
		Agent: agent,
		Document: documents.Document{
			ID:      profileID,
			Content: content,
		},
	}

	if ts := lrs_resp.Response.Header.Get("last-modified"); len(ts) > 0 {
		t, err := time.Parse("Mon, 2 Jan 2006 15:04:05 GMT", ts)

		if err != nil {
			return nil, nil, fmt.Errorf("failed to parse time: %w", err)
		}

		doc.Timestamp = t
	}

	if ct := lrs_resp.Response.Header.Get("content-type"); len(ct) > 0 {
		doc.ContentType = ct
	}

	if etag := lrs_resp.Response.Header.Get("etag"); len(etag) > 0 {
		doc.Etag = etag
	}

	return &doc, lrs_resp, nil
}

// SaveAgentProfile is used to save an agent profile document to the LRS
func (lrs *RemoteLRS) SaveAgentProfile(profile *documents.AgentDocument) (*documents.AgentDocument, *Response, error) {

	if profile == nil {
		return nil, nil, errors.New("argument can't be nil")
	}

	content := string(profile.Content)

	headers := make(map[string]string)

	if len(profile.ContentType) > 0 {
		headers["Content-Type"] = profile.ContentType
	} else {
		headers["Content-Type"] = "application/octet-stream"
	}

	if len(profile.Etag) > 0 {
		headers["If-Match"] = profile.Etag
	}

	params := make(map[string]string)

	b, err := json.Marshal(profile.Agent)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal %w", err)
	}

	params["agent"] = string(b)
	params["profileId"] = profile.ID

	lrs_request := lrs.newRequest("PUT", "agents/profile", &headers, &params, &content)

	req, err := lrs_request.Init()

	if err != nil {
		return nil, nil, fmt.Errorf("failed to init request: %w", err)
	}

	lrs_resp, err := lrs.sendRequest(req)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to send request: %w", err)
	}

	return profile, lrs_resp, nil
}

// DeleteAgentProfile is used to delete a an activty profile
func (lrs *RemoteLRS) DeleteAgentProfile(profile *documents.AgentDocument) (*Response, error) {

	if profile == nil {
		return nil, errors.New("argument can't be nil")
	}

	params := make(map[string]string)

	b, err := json.Marshal(profile.Agent)

	if err != nil {
		return nil, fmt.Errorf("failed to marshal %w", err)
	}

	params["agent"] = string(b)
	params["profileId"] = profile.ID

	headers := make(map[string]string)

	if len(profile.Etag) > 0 {
		headers["If-Match"] = profile.Etag
	}

	lrs_request := lrs.newRequest("DELETE", "agents/profile", &headers, &params, nil)

	req, err := lrs_request.Init()

	if err != nil {
		return nil, fmt.Errorf("failed to init request: %w", err)
	}

	lrs_resp, err := lrs.sendRequest(req)

	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	return lrs_resp, nil
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
		lrs.Auth = fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(lrs.Username+":"+lrs.Password)))
		return &lrs, nil
	}

	return nil, errors.New("unknown error")
}
