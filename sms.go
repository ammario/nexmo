package nexmo

// SMSResp is returned by the SMS api
type SMSResp struct {
	MessageCount string `json:"message-count"`
	Messages     []struct {
		Status           string `json:"status"`
		MessageID        string `json:"message-id"`
		To               string `json:"to"`
		ClientRef        string `json:"client-ref"`
		RemainingBalance string `json:"remaining-balance"`
		MessagePrice     string `json:"message-price"`
		Network          string `json:"network"`
		ErrorText        string `json:"error-text"`
	} `json:"messages"`
}

// SMS sends an SMS mesage.
// Valid options are Callback.
func (c *Client) SMS(from string, to string, text string, opts ...Option) (resp *SMSResp, err error) {
	params := optParams(opts...)
	params.Set("from", from)
	params.Set("to", to)
	params.Set("text", text)
	resp = new(SMSResp)
	return resp, c.request("GET", "/sms/json", params, resp)
}
