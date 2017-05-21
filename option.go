package nexmo

import "net/url"
import "strconv"

// Option applies an optional parameter to a request
type Option func(params url.Values)

// optParams gets url parameters from a series of options
func optParams(opts ...Option) url.Values {
	v := make(url.Values)
	for _, o := range opts {
		o(v)
	}
	return v
}

// Country sets the country code of a request.
// Its in ISO 3166-1 alpha-2 format
func Country(countryCode string) Option {
	return func(params url.Values) {
		params.Set("country", countryCode)
	}
}

// CheckName sets cname
// It charges an additional ~1 cent
func CheckName(do bool) Option {
	return func(params url.Values) {
		params.Set("cnam", strconv.FormatBool(do))
	}
}

// IP sets the ip
func IP(ip string) Option {
	return func(params url.Values) {
		params.Set("ip", ip)
	}
}

// Callback sets a notifcation url.
// See https://docs.nexmo.com/messaging/setup-callbacks and https://docs.nexmo.com/messaging/sms-api/api-reference#delivery_receipt
func Callback(callback string) Option {
	return func(params url.Values) {
		params.Set("callback", callback)
	}
}
