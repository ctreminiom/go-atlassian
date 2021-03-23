package sm

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type RequestFeedbackService struct{ client *Client }

func (r *RequestFeedbackService) Get(ctx context.Context, requestIDOrKey string) (result *CustomerFeedbackScheme, response *Response, err error) {

	if len(requestIDOrKey) == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valid requestIDOrKey value")
	}

	var endpoint = fmt.Sprintf("rest/servicedeskapi/request/%v/feedback", requestIDOrKey)

	request, err := r.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("X-ExperimentalApi", "opt-in")

	response, err = r.client.Do(request)
	if err != nil {
		return
	}

	result = new(CustomerFeedbackScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

func (r *RequestFeedbackService) Post(ctx context.Context, requestIDOrKey string, rating int, comment string) (result *CustomerFeedbackScheme, response *Response, err error) {

	if len(requestIDOrKey) == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valid requestIDOrKey value")
	}

	payload := struct {
		Rating  int `json:"rating"`
		Comment struct {
			Body string `json:"body,omitempty"`
		} `json:"comment,omitempty"`
		Type string `json:"type"`
	}{
		Rating: rating,
		Comment: struct {
			Body string `json:"body,omitempty"`
		}{
			Body: comment,
		},
		Type: "csat",
	}

	var endpoint = fmt.Sprintf("rest/servicedeskapi/request/%v/feedback", requestIDOrKey)

	request, err := r.client.newRequest(ctx, http.MethodPost, endpoint, &payload)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-ExperimentalApi", "opt-in")

	response, err = r.client.Do(request)
	if err != nil {
		return
	}

	result = new(CustomerFeedbackScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

func (r *RequestFeedbackService) Delete(ctx context.Context, requestIDOrKey string) (response *Response, err error) {

	if len(requestIDOrKey) == 0 {
		return nil, fmt.Errorf("error, please provide a valid requestIDOrKey value")
	}

	var endpoint = fmt.Sprintf("rest/servicedeskapi/request/%v/feedback", requestIDOrKey)

	request, err := r.client.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("X-ExperimentalApi", "opt-in")

	response, err = r.client.Do(request)
	if err != nil {
		return
	}

	return
}

type CustomerFeedbackScheme struct {
	Type    string `json:"type"`
	Rating  int    `json:"rating"`
	Comment struct {
		Body string `json:"body"`
	} `json:"comment"`
}
