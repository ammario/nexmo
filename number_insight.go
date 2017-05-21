package nexmo

// NumberInsightBasicResp is returned from the number insight basic API
type NumberInsightBasicResp struct {
	Status                    Status `json:"status"`
	StatusMessage             string `json:"status_message"`
	RequestID                 string `json:"request_id"`
	InternationalFormatNumber string `json:"international_format_number"`
	NationalFormatNumber      string `json:"national_format_number"`
	CountryCode               string `json:"country_code"`
	CountryCodeIso3           string `json:"country_code_iso3"`
	CountryName               string `json:"country_name"`
	CountryPrefix             string `json:"country_prefix"`
}

// NumberInsightBasic sends a request to Nexmo's number insight basic API.
// See https://docs.nexmo.com/number-insight/basic/api-reference
// It accepts options Country and CheckName.
func (c *Client) NumberInsightBasic(number string, opts ...Option) (resp *NumberInsightBasicResp, err error) {
	params := optParams(opts...)
	params.Set("number", number)
	resp = new(NumberInsightBasicResp)
	return resp, c.request("GET", "/ni/basic/json", params, resp)
}

// NumberInsightStandardResp is returned from the number insight standard API
type NumberInsightStandardResp struct {
	Status                    Status `json:"status"`
	StatusMessage             string `json:"status_message"`
	RequestID                 string `json:"request_id"`
	InternationalFormatNumber string `json:"international_format_number"`
	NationalFormatNumber      string `json:"national_format_number"`
	CountryCode               string `json:"country_code"`
	CountryCodeIso3           string `json:"country_code_iso3"`
	CountryName               string `json:"country_name"`
	CountryPrefix             string `json:"country_prefix"`
	RequestPrice              string `json:"request_price"`
	RemainingBalance          string `json:"remaining_balance"`
	CurrentCarrier            struct {
		NetworkCode string `json:"network_code"`
		Name        string `json:"name"`
		Country     string `json:"country"`
		NetworkType string `json:"network_type"`
	} `json:"current_carrier"`
	OriginalCarrier struct {
		NetworkCode string `json:"network_code"`
		Name        string `json:"name"`
		Country     string `json:"country"`
		NetworkType string `json:"network_type"`
	} `json:"original_carrier"`
	Ported     string `json:"ported"`
	CallerName string `json:"caller_name"`
	LastName   string `json:"last_name"`
	FirstName  string `json:"first_name"`
	CallerType string `json:"caller_type"`
}

// NumberInsightStandard sends a request to Nexmo's number insight standard API.
// See https://docs.nexmo.com/number-insight/standard/api-reference
// It accepts options Country and CheckName.
func (c *Client) NumberInsightStandard(number string, opts ...Option) (resp *NumberInsightStandardResp, err error) {
	params := optParams(opts...)
	params.Set("number", number)
	resp = new(NumberInsightStandardResp)
	return resp, c.request("GET", "/ni/standard/json", params, resp)
}

// NumberInsightAdvancedResp is returned from the number insight advanced API
type NumberInsightAdvancedResp struct {
	Status                    int    `json:"status"`
	StatusMessage             string `json:"status_message"`
	LookupOutcome             int    `json:"lookup_outcome"`
	LookupOutcomeMessage      string `json:"lookup_outcome_message"`
	RequestID                 string `json:"request_id"`
	InternationalFormatNumber string `json:"international_format_number"`
	NationalFormatNumber      string `json:"national_format_number"`
	CountryCode               string `json:"country_code"`
	CountryCodeIso3           string `json:"country_code_iso3"`
	CountryName               string `json:"country_name"`
	CountryPrefix             string `json:"country_prefix"`
	RequestPrice              string `json:"request_price"`
	RemainingBalance          string `json:"remaining_balance"`
	CurrentCarrier            struct {
		NetworkCode string `json:"network_code"`
		Name        string `json:"name"`
		Country     string `json:"country"`
		NetworkType string `json:"network_type"`
	} `json:"current_carrier"`
	OriginalCarrier struct {
		NetworkCode string `json:"network_code"`
		Name        string `json:"name"`
		Country     string `json:"country"`
		NetworkType string `json:"network_type"`
	} `json:"original_carrier"`
	ValidNumber string `json:"valid_number"`
	Reachable   string `json:"reachable"`
	Ported      string `json:"ported"`
	Roaming     struct {
		Status string `json:"status"`
	} `json:"roaming"`
	CallerName string `json:"caller_name"`
	LastName   string `json:"last_name"`
	FirstName  string `json:"first_name"`
	CallerType string `json:"caller_type"`
}

// NumberInsightAdvanced sends a request to Nexmo's number insight advanced API.
// See https://docs.nexmo.com/number-insight/advanced/api-reference
// It accepts options Country, CheckName and IP.
func (c *Client) NumberInsightAdvanced(number string, opts ...Option) (resp *NumberInsightAdvancedResp, err error) {
	params := optParams(opts...)
	params.Set("number", number)
	resp = new(NumberInsightAdvancedResp)
	return resp, c.request("GET", "/ni/advanced/json", params, resp)
}
