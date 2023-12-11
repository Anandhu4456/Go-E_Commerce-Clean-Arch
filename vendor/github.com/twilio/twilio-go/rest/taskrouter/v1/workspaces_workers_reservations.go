/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Taskrouter
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

//
func (c *ApiService) FetchWorkerReservation(WorkspaceSid string, WorkerSid string, Sid string) (*TaskrouterV1WorkerReservation, error) {
	path := "/v1/Workspaces/{WorkspaceSid}/Workers/{WorkerSid}/Reservations/{Sid}"
	path = strings.Replace(path, "{"+"WorkspaceSid"+"}", WorkspaceSid, -1)
	path = strings.Replace(path, "{"+"WorkerSid"+"}", WorkerSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TaskrouterV1WorkerReservation{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListWorkerReservation'
type ListWorkerReservationParams struct {
	// Returns the list of reservations for a worker with a specified ReservationStatus. Can be: `pending`, `accepted`, `rejected`, `timeout`, `canceled`, or `rescinded`.
	ReservationStatus *string `json:"ReservationStatus,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListWorkerReservationParams) SetReservationStatus(ReservationStatus string) *ListWorkerReservationParams {
	params.ReservationStatus = &ReservationStatus
	return params
}
func (params *ListWorkerReservationParams) SetPageSize(PageSize int) *ListWorkerReservationParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListWorkerReservationParams) SetLimit(Limit int) *ListWorkerReservationParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of WorkerReservation records from the API. Request is executed immediately.
func (c *ApiService) PageWorkerReservation(WorkspaceSid string, WorkerSid string, params *ListWorkerReservationParams, pageToken, pageNumber string) (*ListWorkerReservationResponse, error) {
	path := "/v1/Workspaces/{WorkspaceSid}/Workers/{WorkerSid}/Reservations"

	path = strings.Replace(path, "{"+"WorkspaceSid"+"}", WorkspaceSid, -1)
	path = strings.Replace(path, "{"+"WorkerSid"+"}", WorkerSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.ReservationStatus != nil {
		data.Set("ReservationStatus", *params.ReservationStatus)
	}
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

	ps := &ListWorkerReservationResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists WorkerReservation records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListWorkerReservation(WorkspaceSid string, WorkerSid string, params *ListWorkerReservationParams) ([]TaskrouterV1WorkerReservation, error) {
	response, errors := c.StreamWorkerReservation(WorkspaceSid, WorkerSid, params)

	records := make([]TaskrouterV1WorkerReservation, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams WorkerReservation records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamWorkerReservation(WorkspaceSid string, WorkerSid string, params *ListWorkerReservationParams) (chan TaskrouterV1WorkerReservation, chan error) {
	if params == nil {
		params = &ListWorkerReservationParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan TaskrouterV1WorkerReservation, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageWorkerReservation(WorkspaceSid, WorkerSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamWorkerReservation(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamWorkerReservation(response *ListWorkerReservationResponse, params *ListWorkerReservationParams, recordChannel chan TaskrouterV1WorkerReservation, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Reservations
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListWorkerReservationResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListWorkerReservationResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListWorkerReservationResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListWorkerReservationResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateWorkerReservation'
type UpdateWorkerReservationParams struct {
	// The If-Match HTTP request header
	IfMatch *string `json:"If-Match,omitempty"`
	//
	ReservationStatus *string `json:"ReservationStatus,omitempty"`
	// The new worker activity SID if rejecting a reservation.
	WorkerActivitySid *string `json:"WorkerActivitySid,omitempty"`
	// The assignment instruction for the reservation.
	Instruction *string `json:"Instruction,omitempty"`
	// The SID of the Activity resource to start after executing a Dequeue instruction.
	DequeuePostWorkActivitySid *string `json:"DequeuePostWorkActivitySid,omitempty"`
	// The caller ID of the call to the worker when executing a Dequeue instruction.
	DequeueFrom *string `json:"DequeueFrom,omitempty"`
	// Whether to record both legs of a call when executing a Dequeue instruction or which leg to record.
	DequeueRecord *string `json:"DequeueRecord,omitempty"`
	// The timeout for call when executing a Dequeue instruction.
	DequeueTimeout *int `json:"DequeueTimeout,omitempty"`
	// The contact URI of the worker when executing a Dequeue instruction. Can be the URI of the Twilio Client, the SIP URI for Programmable SIP, or the [E.164](https://www.twilio.com/docs/glossary/what-e164) formatted phone number, depending on the destination.
	DequeueTo *string `json:"DequeueTo,omitempty"`
	// The callback URL for completed call event when executing a Dequeue instruction.
	DequeueStatusCallbackUrl *string `json:"DequeueStatusCallbackUrl,omitempty"`
	// The Caller ID of the outbound call when executing a Call instruction.
	CallFrom *string `json:"CallFrom,omitempty"`
	// Whether to record both legs of a call when executing a Call instruction.
	CallRecord *string `json:"CallRecord,omitempty"`
	// The timeout for a call when executing a Call instruction.
	CallTimeout *int `json:"CallTimeout,omitempty"`
	// The contact URI of the worker when executing a Call instruction. Can be the URI of the Twilio Client, the SIP URI for Programmable SIP, or the [E.164](https://www.twilio.com/docs/glossary/what-e164) formatted phone number, depending on the destination.
	CallTo *string `json:"CallTo,omitempty"`
	// TwiML URI executed on answering the worker's leg as a result of the Call instruction.
	CallUrl *string `json:"CallUrl,omitempty"`
	// The URL to call for the completed call event when executing a Call instruction.
	CallStatusCallbackUrl *string `json:"CallStatusCallbackUrl,omitempty"`
	// Whether to accept a reservation when executing a Call instruction.
	CallAccept *bool `json:"CallAccept,omitempty"`
	// The Call SID of the call parked in the queue when executing a Redirect instruction.
	RedirectCallSid *string `json:"RedirectCallSid,omitempty"`
	// Whether the reservation should be accepted when executing a Redirect instruction.
	RedirectAccept *bool `json:"RedirectAccept,omitempty"`
	// TwiML URI to redirect the call to when executing the Redirect instruction.
	RedirectUrl *string `json:"RedirectUrl,omitempty"`
	// The Contact URI of the worker when executing a Conference instruction. Can be the URI of the Twilio Client, the SIP URI for Programmable SIP, or the [E.164](https://www.twilio.com/docs/glossary/what-e164) formatted phone number, depending on the destination.
	To *string `json:"To,omitempty"`
	// The caller ID of the call to the worker when executing a Conference instruction.
	From *string `json:"From,omitempty"`
	// The URL we should call using the `status_callback_method` to send status information to your application.
	StatusCallback *string `json:"StatusCallback,omitempty"`
	// The HTTP method we should use to call `status_callback`. Can be: `POST` or `GET` and the default is `POST`.
	StatusCallbackMethod *string `json:"StatusCallbackMethod,omitempty"`
	// The call progress events that we will send to `status_callback`. Can be: `initiated`, `ringing`, `answered`, or `completed`.
	StatusCallbackEvent *[]string `json:"StatusCallbackEvent,omitempty"`
	// The timeout for a call when executing a Conference instruction.
	Timeout *int `json:"Timeout,omitempty"`
	// Whether to record the participant and their conferences, including the time between conferences. Can be `true` or `false` and the default is `false`.
	Record *bool `json:"Record,omitempty"`
	// Whether the agent is muted in the conference. Defaults to `false`.
	Muted *bool `json:"Muted,omitempty"`
	// Whether to play a notification beep when the participant joins or when to play a beep. Can be: `true`, `false`, `onEnter`, or `onExit`. The default value is `true`.
	Beep *string `json:"Beep,omitempty"`
	// Whether to start the conference when the participant joins, if it has not already started. Can be: `true` or `false` and the default is `true`. If `false` and the conference has not started, the participant is muted and hears background music until another participant starts the conference.
	StartConferenceOnEnter *bool `json:"StartConferenceOnEnter,omitempty"`
	// Whether to end the conference when the agent leaves.
	EndConferenceOnExit *bool `json:"EndConferenceOnExit,omitempty"`
	// The URL we should call using the `wait_method` for the music to play while participants are waiting for the conference to start. The default value is the URL of our standard hold music. [Learn more about hold music](https://www.twilio.com/labs/twimlets/holdmusic).
	WaitUrl *string `json:"WaitUrl,omitempty"`
	// The HTTP method we should use to call `wait_url`. Can be `GET` or `POST` and the default is `POST`. When using a static audio file, this should be `GET` so that we can cache the file.
	WaitMethod *string `json:"WaitMethod,omitempty"`
	// Whether to allow an agent to hear the state of the outbound call, including ringing or disconnect messages. The default is `true`.
	EarlyMedia *bool `json:"EarlyMedia,omitempty"`
	// The maximum number of participants allowed in the conference. Can be a positive integer from `2` to `250`. The default value is `250`.
	MaxParticipants *int `json:"MaxParticipants,omitempty"`
	// The URL we should call using the `conference_status_callback_method` when the conference events in `conference_status_callback_event` occur. Only the value set by the first participant to join the conference is used. Subsequent `conference_status_callback` values are ignored.
	ConferenceStatusCallback *string `json:"ConferenceStatusCallback,omitempty"`
	// The HTTP method we should use to call `conference_status_callback`. Can be: `GET` or `POST` and defaults to `POST`.
	ConferenceStatusCallbackMethod *string `json:"ConferenceStatusCallbackMethod,omitempty"`
	// The conference status events that we will send to `conference_status_callback`. Can be: `start`, `end`, `join`, `leave`, `mute`, `hold`, `speaker`.
	ConferenceStatusCallbackEvent *[]string `json:"ConferenceStatusCallbackEvent,omitempty"`
	// Whether to record the conference the participant is joining or when to record the conference. Can be: `true`, `false`, `record-from-start`, and `do-not-record`. The default value is `false`.
	ConferenceRecord *string `json:"ConferenceRecord,omitempty"`
	// Whether to trim leading and trailing silence from your recorded conference audio files. Can be: `trim-silence` or `do-not-trim` and defaults to `trim-silence`.
	ConferenceTrim *string `json:"ConferenceTrim,omitempty"`
	// The recording channels for the final recording. Can be: `mono` or `dual` and the default is `mono`.
	RecordingChannels *string `json:"RecordingChannels,omitempty"`
	// The URL that we should call using the `recording_status_callback_method` when the recording status changes.
	RecordingStatusCallback *string `json:"RecordingStatusCallback,omitempty"`
	// The HTTP method we should use when we call `recording_status_callback`. Can be: `GET` or `POST` and defaults to `POST`.
	RecordingStatusCallbackMethod *string `json:"RecordingStatusCallbackMethod,omitempty"`
	// The URL we should call using the `conference_recording_status_callback_method` when the conference recording is available.
	ConferenceRecordingStatusCallback *string `json:"ConferenceRecordingStatusCallback,omitempty"`
	// The HTTP method we should use to call `conference_recording_status_callback`. Can be: `GET` or `POST` and defaults to `POST`.
	ConferenceRecordingStatusCallbackMethod *string `json:"ConferenceRecordingStatusCallbackMethod,omitempty"`
	// The [region](https://support.twilio.com/hc/en-us/articles/223132167-How-global-low-latency-routing-and-region-selection-work-for-conferences-and-Client-calls) where we should mix the recorded audio. Can be:`us1`, `ie1`, `de1`, `sg1`, `br1`, `au1`, or `jp1`.
	Region *string `json:"Region,omitempty"`
	// The SIP username used for authentication.
	SipAuthUsername *string `json:"SipAuthUsername,omitempty"`
	// The SIP password for authentication.
	SipAuthPassword *string `json:"SipAuthPassword,omitempty"`
	// The call progress events sent via webhooks as a result of a Dequeue instruction.
	DequeueStatusCallbackEvent *[]string `json:"DequeueStatusCallbackEvent,omitempty"`
	// The new worker activity SID after executing a Conference instruction.
	PostWorkActivitySid *string `json:"PostWorkActivitySid,omitempty"`
	// Whether to end the conference when the customer leaves.
	EndConferenceOnCustomerExit *bool `json:"EndConferenceOnCustomerExit,omitempty"`
	// Whether to play a notification beep when the customer joins.
	BeepOnCustomerEntrance *bool `json:"BeepOnCustomerEntrance,omitempty"`
}

func (params *UpdateWorkerReservationParams) SetIfMatch(IfMatch string) *UpdateWorkerReservationParams {
	params.IfMatch = &IfMatch
	return params
}
func (params *UpdateWorkerReservationParams) SetReservationStatus(ReservationStatus string) *UpdateWorkerReservationParams {
	params.ReservationStatus = &ReservationStatus
	return params
}
func (params *UpdateWorkerReservationParams) SetWorkerActivitySid(WorkerActivitySid string) *UpdateWorkerReservationParams {
	params.WorkerActivitySid = &WorkerActivitySid
	return params
}
func (params *UpdateWorkerReservationParams) SetInstruction(Instruction string) *UpdateWorkerReservationParams {
	params.Instruction = &Instruction
	return params
}
func (params *UpdateWorkerReservationParams) SetDequeuePostWorkActivitySid(DequeuePostWorkActivitySid string) *UpdateWorkerReservationParams {
	params.DequeuePostWorkActivitySid = &DequeuePostWorkActivitySid
	return params
}
func (params *UpdateWorkerReservationParams) SetDequeueFrom(DequeueFrom string) *UpdateWorkerReservationParams {
	params.DequeueFrom = &DequeueFrom
	return params
}
func (params *UpdateWorkerReservationParams) SetDequeueRecord(DequeueRecord string) *UpdateWorkerReservationParams {
	params.DequeueRecord = &DequeueRecord
	return params
}
func (params *UpdateWorkerReservationParams) SetDequeueTimeout(DequeueTimeout int) *UpdateWorkerReservationParams {
	params.DequeueTimeout = &DequeueTimeout
	return params
}
func (params *UpdateWorkerReservationParams) SetDequeueTo(DequeueTo string) *UpdateWorkerReservationParams {
	params.DequeueTo = &DequeueTo
	return params
}
func (params *UpdateWorkerReservationParams) SetDequeueStatusCallbackUrl(DequeueStatusCallbackUrl string) *UpdateWorkerReservationParams {
	params.DequeueStatusCallbackUrl = &DequeueStatusCallbackUrl
	return params
}
func (params *UpdateWorkerReservationParams) SetCallFrom(CallFrom string) *UpdateWorkerReservationParams {
	params.CallFrom = &CallFrom
	return params
}
func (params *UpdateWorkerReservationParams) SetCallRecord(CallRecord string) *UpdateWorkerReservationParams {
	params.CallRecord = &CallRecord
	return params
}
func (params *UpdateWorkerReservationParams) SetCallTimeout(CallTimeout int) *UpdateWorkerReservationParams {
	params.CallTimeout = &CallTimeout
	return params
}
func (params *UpdateWorkerReservationParams) SetCallTo(CallTo string) *UpdateWorkerReservationParams {
	params.CallTo = &CallTo
	return params
}
func (params *UpdateWorkerReservationParams) SetCallUrl(CallUrl string) *UpdateWorkerReservationParams {
	params.CallUrl = &CallUrl
	return params
}
func (params *UpdateWorkerReservationParams) SetCallStatusCallbackUrl(CallStatusCallbackUrl string) *UpdateWorkerReservationParams {
	params.CallStatusCallbackUrl = &CallStatusCallbackUrl
	return params
}
func (params *UpdateWorkerReservationParams) SetCallAccept(CallAccept bool) *UpdateWorkerReservationParams {
	params.CallAccept = &CallAccept
	return params
}
func (params *UpdateWorkerReservationParams) SetRedirectCallSid(RedirectCallSid string) *UpdateWorkerReservationParams {
	params.RedirectCallSid = &RedirectCallSid
	return params
}
func (params *UpdateWorkerReservationParams) SetRedirectAccept(RedirectAccept bool) *UpdateWorkerReservationParams {
	params.RedirectAccept = &RedirectAccept
	return params
}
func (params *UpdateWorkerReservationParams) SetRedirectUrl(RedirectUrl string) *UpdateWorkerReservationParams {
	params.RedirectUrl = &RedirectUrl
	return params
}
func (params *UpdateWorkerReservationParams) SetTo(To string) *UpdateWorkerReservationParams {
	params.To = &To
	return params
}
func (params *UpdateWorkerReservationParams) SetFrom(From string) *UpdateWorkerReservationParams {
	params.From = &From
	return params
}
func (params *UpdateWorkerReservationParams) SetStatusCallback(StatusCallback string) *UpdateWorkerReservationParams {
	params.StatusCallback = &StatusCallback
	return params
}
func (params *UpdateWorkerReservationParams) SetStatusCallbackMethod(StatusCallbackMethod string) *UpdateWorkerReservationParams {
	params.StatusCallbackMethod = &StatusCallbackMethod
	return params
}
func (params *UpdateWorkerReservationParams) SetStatusCallbackEvent(StatusCallbackEvent []string) *UpdateWorkerReservationParams {
	params.StatusCallbackEvent = &StatusCallbackEvent
	return params
}
func (params *UpdateWorkerReservationParams) SetTimeout(Timeout int) *UpdateWorkerReservationParams {
	params.Timeout = &Timeout
	return params
}
func (params *UpdateWorkerReservationParams) SetRecord(Record bool) *UpdateWorkerReservationParams {
	params.Record = &Record
	return params
}
func (params *UpdateWorkerReservationParams) SetMuted(Muted bool) *UpdateWorkerReservationParams {
	params.Muted = &Muted
	return params
}
func (params *UpdateWorkerReservationParams) SetBeep(Beep string) *UpdateWorkerReservationParams {
	params.Beep = &Beep
	return params
}
func (params *UpdateWorkerReservationParams) SetStartConferenceOnEnter(StartConferenceOnEnter bool) *UpdateWorkerReservationParams {
	params.StartConferenceOnEnter = &StartConferenceOnEnter
	return params
}
func (params *UpdateWorkerReservationParams) SetEndConferenceOnExit(EndConferenceOnExit bool) *UpdateWorkerReservationParams {
	params.EndConferenceOnExit = &EndConferenceOnExit
	return params
}
func (params *UpdateWorkerReservationParams) SetWaitUrl(WaitUrl string) *UpdateWorkerReservationParams {
	params.WaitUrl = &WaitUrl
	return params
}
func (params *UpdateWorkerReservationParams) SetWaitMethod(WaitMethod string) *UpdateWorkerReservationParams {
	params.WaitMethod = &WaitMethod
	return params
}
func (params *UpdateWorkerReservationParams) SetEarlyMedia(EarlyMedia bool) *UpdateWorkerReservationParams {
	params.EarlyMedia = &EarlyMedia
	return params
}
func (params *UpdateWorkerReservationParams) SetMaxParticipants(MaxParticipants int) *UpdateWorkerReservationParams {
	params.MaxParticipants = &MaxParticipants
	return params
}
func (params *UpdateWorkerReservationParams) SetConferenceStatusCallback(ConferenceStatusCallback string) *UpdateWorkerReservationParams {
	params.ConferenceStatusCallback = &ConferenceStatusCallback
	return params
}
func (params *UpdateWorkerReservationParams) SetConferenceStatusCallbackMethod(ConferenceStatusCallbackMethod string) *UpdateWorkerReservationParams {
	params.ConferenceStatusCallbackMethod = &ConferenceStatusCallbackMethod
	return params
}
func (params *UpdateWorkerReservationParams) SetConferenceStatusCallbackEvent(ConferenceStatusCallbackEvent []string) *UpdateWorkerReservationParams {
	params.ConferenceStatusCallbackEvent = &ConferenceStatusCallbackEvent
	return params
}
func (params *UpdateWorkerReservationParams) SetConferenceRecord(ConferenceRecord string) *UpdateWorkerReservationParams {
	params.ConferenceRecord = &ConferenceRecord
	return params
}
func (params *UpdateWorkerReservationParams) SetConferenceTrim(ConferenceTrim string) *UpdateWorkerReservationParams {
	params.ConferenceTrim = &ConferenceTrim
	return params
}
func (params *UpdateWorkerReservationParams) SetRecordingChannels(RecordingChannels string) *UpdateWorkerReservationParams {
	params.RecordingChannels = &RecordingChannels
	return params
}
func (params *UpdateWorkerReservationParams) SetRecordingStatusCallback(RecordingStatusCallback string) *UpdateWorkerReservationParams {
	params.RecordingStatusCallback = &RecordingStatusCallback
	return params
}
func (params *UpdateWorkerReservationParams) SetRecordingStatusCallbackMethod(RecordingStatusCallbackMethod string) *UpdateWorkerReservationParams {
	params.RecordingStatusCallbackMethod = &RecordingStatusCallbackMethod
	return params
}
func (params *UpdateWorkerReservationParams) SetConferenceRecordingStatusCallback(ConferenceRecordingStatusCallback string) *UpdateWorkerReservationParams {
	params.ConferenceRecordingStatusCallback = &ConferenceRecordingStatusCallback
	return params
}
func (params *UpdateWorkerReservationParams) SetConferenceRecordingStatusCallbackMethod(ConferenceRecordingStatusCallbackMethod string) *UpdateWorkerReservationParams {
	params.ConferenceRecordingStatusCallbackMethod = &ConferenceRecordingStatusCallbackMethod
	return params
}
func (params *UpdateWorkerReservationParams) SetRegion(Region string) *UpdateWorkerReservationParams {
	params.Region = &Region
	return params
}
func (params *UpdateWorkerReservationParams) SetSipAuthUsername(SipAuthUsername string) *UpdateWorkerReservationParams {
	params.SipAuthUsername = &SipAuthUsername
	return params
}
func (params *UpdateWorkerReservationParams) SetSipAuthPassword(SipAuthPassword string) *UpdateWorkerReservationParams {
	params.SipAuthPassword = &SipAuthPassword
	return params
}
func (params *UpdateWorkerReservationParams) SetDequeueStatusCallbackEvent(DequeueStatusCallbackEvent []string) *UpdateWorkerReservationParams {
	params.DequeueStatusCallbackEvent = &DequeueStatusCallbackEvent
	return params
}
func (params *UpdateWorkerReservationParams) SetPostWorkActivitySid(PostWorkActivitySid string) *UpdateWorkerReservationParams {
	params.PostWorkActivitySid = &PostWorkActivitySid
	return params
}
func (params *UpdateWorkerReservationParams) SetEndConferenceOnCustomerExit(EndConferenceOnCustomerExit bool) *UpdateWorkerReservationParams {
	params.EndConferenceOnCustomerExit = &EndConferenceOnCustomerExit
	return params
}
func (params *UpdateWorkerReservationParams) SetBeepOnCustomerEntrance(BeepOnCustomerEntrance bool) *UpdateWorkerReservationParams {
	params.BeepOnCustomerEntrance = &BeepOnCustomerEntrance
	return params
}

//
func (c *ApiService) UpdateWorkerReservation(WorkspaceSid string, WorkerSid string, Sid string, params *UpdateWorkerReservationParams) (*TaskrouterV1WorkerReservation, error) {
	path := "/v1/Workspaces/{WorkspaceSid}/Workers/{WorkerSid}/Reservations/{Sid}"
	path = strings.Replace(path, "{"+"WorkspaceSid"+"}", WorkspaceSid, -1)
	path = strings.Replace(path, "{"+"WorkerSid"+"}", WorkerSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.ReservationStatus != nil {
		data.Set("ReservationStatus", *params.ReservationStatus)
	}
	if params != nil && params.WorkerActivitySid != nil {
		data.Set("WorkerActivitySid", *params.WorkerActivitySid)
	}
	if params != nil && params.Instruction != nil {
		data.Set("Instruction", *params.Instruction)
	}
	if params != nil && params.DequeuePostWorkActivitySid != nil {
		data.Set("DequeuePostWorkActivitySid", *params.DequeuePostWorkActivitySid)
	}
	if params != nil && params.DequeueFrom != nil {
		data.Set("DequeueFrom", *params.DequeueFrom)
	}
	if params != nil && params.DequeueRecord != nil {
		data.Set("DequeueRecord", *params.DequeueRecord)
	}
	if params != nil && params.DequeueTimeout != nil {
		data.Set("DequeueTimeout", fmt.Sprint(*params.DequeueTimeout))
	}
	if params != nil && params.DequeueTo != nil {
		data.Set("DequeueTo", *params.DequeueTo)
	}
	if params != nil && params.DequeueStatusCallbackUrl != nil {
		data.Set("DequeueStatusCallbackUrl", *params.DequeueStatusCallbackUrl)
	}
	if params != nil && params.CallFrom != nil {
		data.Set("CallFrom", *params.CallFrom)
	}
	if params != nil && params.CallRecord != nil {
		data.Set("CallRecord", *params.CallRecord)
	}
	if params != nil && params.CallTimeout != nil {
		data.Set("CallTimeout", fmt.Sprint(*params.CallTimeout))
	}
	if params != nil && params.CallTo != nil {
		data.Set("CallTo", *params.CallTo)
	}
	if params != nil && params.CallUrl != nil {
		data.Set("CallUrl", *params.CallUrl)
	}
	if params != nil && params.CallStatusCallbackUrl != nil {
		data.Set("CallStatusCallbackUrl", *params.CallStatusCallbackUrl)
	}
	if params != nil && params.CallAccept != nil {
		data.Set("CallAccept", fmt.Sprint(*params.CallAccept))
	}
	if params != nil && params.RedirectCallSid != nil {
		data.Set("RedirectCallSid", *params.RedirectCallSid)
	}
	if params != nil && params.RedirectAccept != nil {
		data.Set("RedirectAccept", fmt.Sprint(*params.RedirectAccept))
	}
	if params != nil && params.RedirectUrl != nil {
		data.Set("RedirectUrl", *params.RedirectUrl)
	}
	if params != nil && params.To != nil {
		data.Set("To", *params.To)
	}
	if params != nil && params.From != nil {
		data.Set("From", *params.From)
	}
	if params != nil && params.StatusCallback != nil {
		data.Set("StatusCallback", *params.StatusCallback)
	}
	if params != nil && params.StatusCallbackMethod != nil {
		data.Set("StatusCallbackMethod", *params.StatusCallbackMethod)
	}
	if params != nil && params.StatusCallbackEvent != nil {
		for _, item := range *params.StatusCallbackEvent {
			data.Add("StatusCallbackEvent", item)
		}
	}
	if params != nil && params.Timeout != nil {
		data.Set("Timeout", fmt.Sprint(*params.Timeout))
	}
	if params != nil && params.Record != nil {
		data.Set("Record", fmt.Sprint(*params.Record))
	}
	if params != nil && params.Muted != nil {
		data.Set("Muted", fmt.Sprint(*params.Muted))
	}
	if params != nil && params.Beep != nil {
		data.Set("Beep", *params.Beep)
	}
	if params != nil && params.StartConferenceOnEnter != nil {
		data.Set("StartConferenceOnEnter", fmt.Sprint(*params.StartConferenceOnEnter))
	}
	if params != nil && params.EndConferenceOnExit != nil {
		data.Set("EndConferenceOnExit", fmt.Sprint(*params.EndConferenceOnExit))
	}
	if params != nil && params.WaitUrl != nil {
		data.Set("WaitUrl", *params.WaitUrl)
	}
	if params != nil && params.WaitMethod != nil {
		data.Set("WaitMethod", *params.WaitMethod)
	}
	if params != nil && params.EarlyMedia != nil {
		data.Set("EarlyMedia", fmt.Sprint(*params.EarlyMedia))
	}
	if params != nil && params.MaxParticipants != nil {
		data.Set("MaxParticipants", fmt.Sprint(*params.MaxParticipants))
	}
	if params != nil && params.ConferenceStatusCallback != nil {
		data.Set("ConferenceStatusCallback", *params.ConferenceStatusCallback)
	}
	if params != nil && params.ConferenceStatusCallbackMethod != nil {
		data.Set("ConferenceStatusCallbackMethod", *params.ConferenceStatusCallbackMethod)
	}
	if params != nil && params.ConferenceStatusCallbackEvent != nil {
		for _, item := range *params.ConferenceStatusCallbackEvent {
			data.Add("ConferenceStatusCallbackEvent", item)
		}
	}
	if params != nil && params.ConferenceRecord != nil {
		data.Set("ConferenceRecord", *params.ConferenceRecord)
	}
	if params != nil && params.ConferenceTrim != nil {
		data.Set("ConferenceTrim", *params.ConferenceTrim)
	}
	if params != nil && params.RecordingChannels != nil {
		data.Set("RecordingChannels", *params.RecordingChannels)
	}
	if params != nil && params.RecordingStatusCallback != nil {
		data.Set("RecordingStatusCallback", *params.RecordingStatusCallback)
	}
	if params != nil && params.RecordingStatusCallbackMethod != nil {
		data.Set("RecordingStatusCallbackMethod", *params.RecordingStatusCallbackMethod)
	}
	if params != nil && params.ConferenceRecordingStatusCallback != nil {
		data.Set("ConferenceRecordingStatusCallback", *params.ConferenceRecordingStatusCallback)
	}
	if params != nil && params.ConferenceRecordingStatusCallbackMethod != nil {
		data.Set("ConferenceRecordingStatusCallbackMethod", *params.ConferenceRecordingStatusCallbackMethod)
	}
	if params != nil && params.Region != nil {
		data.Set("Region", *params.Region)
	}
	if params != nil && params.SipAuthUsername != nil {
		data.Set("SipAuthUsername", *params.SipAuthUsername)
	}
	if params != nil && params.SipAuthPassword != nil {
		data.Set("SipAuthPassword", *params.SipAuthPassword)
	}
	if params != nil && params.DequeueStatusCallbackEvent != nil {
		for _, item := range *params.DequeueStatusCallbackEvent {
			data.Add("DequeueStatusCallbackEvent", item)
		}
	}
	if params != nil && params.PostWorkActivitySid != nil {
		data.Set("PostWorkActivitySid", *params.PostWorkActivitySid)
	}
	if params != nil && params.EndConferenceOnCustomerExit != nil {
		data.Set("EndConferenceOnCustomerExit", fmt.Sprint(*params.EndConferenceOnCustomerExit))
	}
	if params != nil && params.BeepOnCustomerEntrance != nil {
		data.Set("BeepOnCustomerEntrance", fmt.Sprint(*params.BeepOnCustomerEntrance))
	}

	if params != nil && params.IfMatch != nil {
		headers["If-Match"] = *params.IfMatch
	}
	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TaskrouterV1WorkerReservation{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
