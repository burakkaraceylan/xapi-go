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

	"github.com/burakkaraceylan/xapi-go/pkg/resources"
	"github.com/burakkaraceylan/xapi-go/pkg/resources/about"
	activityprofile "github.com/burakkaraceylan/xapi-go/pkg/resources/activity_profile"
	"github.com/burakkaraceylan/xapi-go/pkg/resources/state"
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

	if err := lrs_resp.Bind(statement); err != nil {
		return nil, nil, err
	}

	return statement, lrs_resp, nil
}

// TODO: Return a Statement struct not a Response
// GetVoidedStatement is used to fetch a single voided statement from record store
func (lrs *RemoteLRS) GetVoidedStatement(id string) (*statement.Statement, *Response, error) {
	lrs_request := lrs.newRequest("GET", "statements", nil, &map[string]string{"voidedStatementId": id}, nil)

	req, err := lrs_request.Init()

	if err != nil {
		return nil, nil, err
	}

	lrs_resp, err := lrs.sendRequest(req)

	if err != nil {
		return nil, nil, err
	}

	statement := &statement.Statement{}

	if err := lrs_resp.Bind(statement); err != nil {
		return nil, nil, err
	}

	return statement, lrs_resp, nil
}

// QueryParams represents query parameters of a statement
type StatementQueryParams struct {
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
func (lrs *RemoteLRS) QueryStatements(params *StatementQueryParams) (*statement.StatementResult, *Response, error) {

	var query_params map[string]string

	if params != nil {
		query_params = params.Map()
	}

	lrs_request := lrs.newRequest("GET", "statements", nil, &query_params, nil)
	req, err := lrs_request.Init()

	if err != nil {
		return nil, nil, err
	}

	lrs_resp, err := lrs.sendRequest(req)

	if err != nil {
		return nil, nil, err
	}

	result := &statement.StatementResult{}

	if err := lrs_resp.Bind(result); err != nil {
		return nil, nil, err
	}

	return result, lrs_resp, nil
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

// GetStateIds request parameters
type GetStateIdsQueryParams struct {
	Activity     properties.Object
	Agent        properties.Actor
	Registration *string
	Since        *time.Time
}

func (q *GetStateIdsQueryParams) Map() map[string]string {
	params := make(map[string]string)

	params["activityId"] = q.Activity.ID

	b, err := json.Marshal(q.Agent)

	if err == nil {
		params["agent"] = string(b)
	}

	params["agent"] = string(b)

	if q.Registration != nil {
		params["registration"] = *q.Registration
	}

	if q.Since != nil {
		params["since"] = q.Since.Format(time.RFC3339Nano)
	}

	return params
}

// GetStateIds is used to fetch the state ids
func (lrs *RemoteLRS) GetStateIds(params *GetStateIdsQueryParams) ([]string, *Response, error) {

	var query_params map[string]string

	if params != nil {
		query_params = params.Map()
	}

	lrs_request := lrs.newRequest("GET", "activities/state", nil, &query_params, nil)
	req, err := lrs_request.Init()

	if err != nil {
		return nil, nil, err
	}

	lrs_resp, err := lrs.sendRequest(req)

	if err != nil {
		return nil, nil, err
	}

	if lrs_resp.Response.StatusCode != 200 {
		return nil, lrs_resp, nil
	}

	var idList []string

	if err := lrs_resp.Bind(&idList); err != nil {
		return nil, nil, err
	}

	return idList, lrs_resp, nil
}

// GetState request parameters
type GetStateQueryParams struct {
	Activity     properties.Object
	Agent        properties.Actor
	Registration *string
	StateID      string
}

func (q *GetStateQueryParams) Map() map[string]string {
	params := make(map[string]string)

	params["activityId"] = q.Activity.ID

	b, err := json.Marshal(q.Agent)

	if err == nil {
		params["agent"] = string(b)
	}

	params["agent"] = string(b)

	if q.Registration != nil {
		params["registration"] = *q.Registration
	}

	params["stateId"] = q.StateID

	return params
}

// GetState is used to fetch a state
func (lrs *RemoteLRS) GetState(params *GetStateQueryParams) (*state.StateDocument, *Response, error) {
	var query_params map[string]string

	if params != nil {
		query_params = params.Map()
	}

	if len(query_params["stateId"]) == 0 {
		return nil, nil, errors.New("stateId can't be null")
	}

	lrs_request := lrs.newRequest("GET", "activities/state", nil, &query_params, nil)
	req, err := lrs_request.Init()

	if err != nil {
		return nil, nil, err
	}

	lrs_resp, err := lrs.sendRequest(req)

	if err != nil {
		return nil, nil, err
	}

	if lrs_resp.Response.StatusCode != 200 {
		return nil, lrs_resp, nil
	}

	content, err := io.ReadAll(lrs_resp.Response.Body)

	if err != nil {
		return nil, nil, err
	}

	doc := state.StateDocument{
		Agent:    params.Agent,
		Activity: params.Activity,
		Document: resources.Document{
			ID:      params.StateID,
			Content: content,
		},
	}

	if params.Registration != nil {
		doc.Registration = params.Registration
	}

	if ts := lrs_resp.Response.Header.Get("last-modified"); len(ts) > 0 {
		t, err := time.Parse("Mon, 2 Jan 2006 15:04:05 GMT", ts)

		if err != nil {
			return nil, nil, err
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
func (lrs *RemoteLRS) SaveState(state *state.StateDocument) (*state.StateDocument, *Response, error) {

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

	params := GetStateQueryParams{
		Activity: state.Activity,
		StateID:  state.ID,
		Agent:    state.Agent,
	}

	pmap := params.Map()

	lrs_request := lrs.newRequest("PUT", "activities/state", &headers, &pmap, &content)

	req, err := lrs_request.Init()

	if err != nil {
		return nil, nil, err
	}

	lrs_resp, err := lrs.sendRequest(req)

	if err != nil {
		return nil, nil, err
	}

	return state, lrs_resp, nil
}

// DeleteState is used to delete a state (if stateId is provided) or all states related to agent/activity/registration
func (lrs *RemoteLRS) DeleteState(state *state.StateDocument) (*Response, error) {

	params := make(map[string]string)

	b, err := json.Marshal(state.Agent)

	if err != nil {
		return nil, err
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
		return nil, err
	}

	lrs_resp, err := lrs.sendRequest(req)

	if err != nil {
		return nil, err
	}

	return lrs_resp, nil
}

type GetActivityProfilesParams struct {
	Activity properties.Object
	Since    *time.Time
}

// GetActivityProfileIds is used to fetch the activity profile ids
func (lrs *RemoteLRS) GetActivityProfileIds(params *GetActivityProfilesParams) ([]string, *Response, error) {

	if *params.Activity.ObjectType != "Activity" {
		return nil, nil, errors.New("must be an activity")
	}

	query_params := make(map[string]string)

	query_params["activityId"] = params.Activity.ID

	if params.Since != nil {
		query_params["since"] = params.Since.Format(time.RFC3339Nano)
	}

	lrs_request := lrs.newRequest("GET", "activities/profile", nil, &query_params, nil)
	req, err := lrs_request.Init()

	if err != nil {
		return nil, nil, err
	}

	lrs_resp, err := lrs.sendRequest(req)

	if err != nil {
		return nil, nil, err
	}

	if lrs_resp.Response.StatusCode != 200 {
		return nil, lrs_resp, nil
	}

	var idList []string

	if err := lrs_resp.Bind(&idList); err != nil {
		return nil, nil, err
	}

	return idList, lrs_resp, nil
}

type GetActivityProfileParams struct {
	Activity  properties.Object
	ProfileID string
}

// GetActivityProfile is used to fetch an actity profile
func (lrs *RemoteLRS) GetActivityProfile(params *GetActivityProfileParams) (*activityprofile.ActivityDocument, *Response, error) {

	query_params := make(map[string]string)

	query_params["profileId"] = params.ProfileID
	query_params["activityId"] = params.Activity.ID

	lrs_request := lrs.newRequest("GET", "activities/profile", nil, &query_params, nil)
	req, err := lrs_request.Init()

	if err != nil {
		return nil, nil, err
	}

	lrs_resp, err := lrs.sendRequest(req)

	if err != nil {
		return nil, nil, err
	}

	if lrs_resp.Response.StatusCode != 200 {
		return nil, lrs_resp, nil
	}

	content, err := io.ReadAll(lrs_resp.Response.Body)

	if err != nil {
		return nil, nil, err
	}

	doc := activityprofile.ActivityDocument{
		Activity: params.Activity,
		Document: resources.Document{
			ID:      params.ProfileID,
			Content: content,
		},
	}

	if ts := lrs_resp.Response.Header.Get("last-modified"); len(ts) > 0 {
		t, err := time.Parse("Mon, 2 Jan 2006 15:04:05 GMT", ts)

		if err != nil {
			return nil, nil, err
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
func (lrs *RemoteLRS) SaveActivityProfile(profile *activityprofile.ActivityDocument) (*activityprofile.ActivityDocument, *Response, error) {

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
		return nil, nil, err
	}

	lrs_resp, err := lrs.sendRequest(req)

	if err != nil {
		return nil, nil, err
	}

	return profile, lrs_resp, nil
}

// DeleteActivityProfile is used to delete a an activty profile
func (lrs *RemoteLRS) DeleteActivityProfile(profile *activityprofile.ActivityDocument) (*Response, error) {

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
		return nil, err
	}

	lrs_resp, err := lrs.sendRequest(req)

	if err != nil {
		return nil, err
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
