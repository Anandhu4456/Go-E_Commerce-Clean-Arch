/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Chat
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'CreateUser'
type CreateUserParams struct {
	// The `identity` value that uniquely identifies the new resource's [User](https://www.twilio.com/docs/api/chat/rest/v1/user) within the [Service](https://www.twilio.com/docs/api/chat/rest/v1/service). This value is often a username or email address. See the Identity documentation for more details.
	Identity *string `json:"Identity,omitempty"`
	// The SID of the [Role](https://www.twilio.com/docs/api/chat/rest/roles) assigned to the new User.
	RoleSid *string `json:"RoleSid,omitempty"`
	// A valid JSON string that contains application-specific data.
	Attributes *string `json:"Attributes,omitempty"`
	// A descriptive string that you create to describe the new resource. This value is often used for display purposes.
	FriendlyName *string `json:"FriendlyName,omitempty"`
}

func (params *CreateUserParams) SetIdentity(Identity string) *CreateUserParams {
	params.Identity = &Identity
	return params
}
func (params *CreateUserParams) SetRoleSid(RoleSid string) *CreateUserParams {
	params.RoleSid = &RoleSid
	return params
}
func (params *CreateUserParams) SetAttributes(Attributes string) *CreateUserParams {
	params.Attributes = &Attributes
	return params
}
func (params *CreateUserParams) SetFriendlyName(FriendlyName string) *CreateUserParams {
	params.FriendlyName = &FriendlyName
	return params
}

//
func (c *ApiService) CreateUser(ServiceSid string, params *CreateUserParams) (*ChatV1User, error) {
	path := "/v1/Services/{ServiceSid}/Users"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Identity != nil {
		data.Set("Identity", *params.Identity)
	}
	if params != nil && params.RoleSid != nil {
		data.Set("RoleSid", *params.RoleSid)
	}
	if params != nil && params.Attributes != nil {
		data.Set("Attributes", *params.Attributes)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ChatV1User{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

//
func (c *ApiService) DeleteUser(ServiceSid string, Sid string) error {
	path := "/v1/Services/{ServiceSid}/Users/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

//
func (c *ApiService) FetchUser(ServiceSid string, Sid string) (*ChatV1User, error) {
	path := "/v1/Services/{ServiceSid}/Users/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ChatV1User{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListUser'
type ListUserParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListUserParams) SetPageSize(PageSize int) *ListUserParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListUserParams) SetLimit(Limit int) *ListUserParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of User records from the API. Request is executed immediately.
func (c *ApiService) PageUser(ServiceSid string, params *ListUserParams, pageToken, pageNumber string) (*ListUserResponse, error) {
	path := "/v1/Services/{ServiceSid}/Users"

	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	if pageToken != "" {
		data.Set("PageToken", pageToken)
	}
	if pageNumber != "" {
		data.Set("Page", pageNumber)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListUserResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists User records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListUser(ServiceSid string, params *ListUserParams) ([]ChatV1User, error) {
	response, errors := c.StreamUser(ServiceSid, params)

	records := make([]ChatV1User, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams User records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamUser(ServiceSid string, params *ListUserParams) (chan ChatV1User, chan error) {
	if params == nil {
		params = &ListUserParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan ChatV1User, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageUser(ServiceSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamUser(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamUser(response *ListUserResponse, params *ListUserParams, recordChannel chan ChatV1User, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Users
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListUserResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListUserResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListUserResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListUserResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateUser'
type UpdateUserParams struct {
	// The SID of the [Role](https://www.twilio.com/docs/api/chat/rest/roles) assigned to this user.
	RoleSid *string `json:"RoleSid,omitempty"`
	// A valid JSON string that contains application-specific data.
	Attributes *string `json:"Attributes,omitempty"`
	// A descriptive string that you create to describe the resource. It is often used for display purposes.
	FriendlyName *string `json:"FriendlyName,omitempty"`
}

func (params *UpdateUserParams) SetRoleSid(RoleSid string) *UpdateUserParams {
	params.RoleSid = &RoleSid
	return params
}
func (params *UpdateUserParams) SetAttributes(Attributes string) *UpdateUserParams {
	params.Attributes = &Attributes
	return params
}
func (params *UpdateUserParams) SetFriendlyName(FriendlyName string) *UpdateUserParams {
	params.FriendlyName = &FriendlyName
	return params
}

//
func (c *ApiService) UpdateUser(ServiceSid string, Sid string, params *UpdateUserParams) (*ChatV1User, error) {
	path := "/v1/Services/{ServiceSid}/Users/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.RoleSid != nil {
		data.Set("RoleSid", *params.RoleSid)
	}
	if params != nil && params.Attributes != nil {
		data.Set("Attributes", *params.Attributes)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ChatV1User{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}