package nexmo

//go:generate stringer -type=Status

// Status is a special status code returned in the body of a nexmo response
type Status int

// Enum
const (
	StatusSuccess              Status = 0
	StatusBusy                 Status = 1
	StatusInvalid              Status = 3
	StatusInvalidCredentials   Status = 4
	StatusInternalError        Status = 5
	StatusPartnerQuotaExceeded Status = 9
	StatusFacilityNotAllowed   Status = 19
	StatusUnparseable          Status = 999
)
