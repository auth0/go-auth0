package management

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

var logTypeName = map[string]string{
	"s":         "Success Login",
	"ssa":       "Success Silent Auth",
	"fsa":       "Failed Silent Auth",
	"seacft":    "Success Exchange (Authorization Code for Access Token)",
	"feacft":    "Failed Exchange (Authorization Code for Access Token)",
	"seccft":    "Success Exchange (Client Credentials for Access Token)",
	"feccft":    "Failed Exchange (Client Credentials for Access Token)",
	"sepft":     "Success Exchange (Password for Access Token)",
	"fepft":     "Failed Exchange (Password for Access Token)",
	"f":         "Failed Login",
	"w":         "Warnings During Login",
	"du":        "Deleted User",
	"fu":        "Failed Login (invalid email/username)",
	"fp":        "Failed Login (wrong password)",
	"fc":        "Failed by Connector",
	"fco":       "Failed by CORS",
	"con":       "Connector Online",
	"coff":      "Connector Offline",
	"fcpro":     "Failed Connector Provisioning",
	"ss":        "Success Signup",
	"fs":        "Failed Signup",
	"cs":        "Code Sent",
	"cls":       "Code/Link Sent",
	"sv":        "Success Verification Email",
	"fv":        "Failed Verification Email",
	"scp":       "Success Change Password",
	"fcp":       "Failed Change Password",
	"sce":       "Success Change Email",
	"fce":       "Failed Change Email",
	"scu":       "Success Change Username",
	"fcu":       "Failed Change Username",
	"scpn":      "Success Change Phone Number",
	"fcpn":      "Failed Change Phone Number",
	"svr":       "Success Verification Email Request",
	"fvr":       "Failed Verification Email Request",
	"scpr":      "Success Change Password Request",
	"fcpr":      "Failed Change Password Request",
	"fn":        "Failed Sending Notification",
	"sapi":      "API Operation",
	"fapi":      "Failed API Operation",
	"limit_wc":  "Blocked Account",
	"limit_mu":  "Blocked IP Address",
	"limit_ui":  "Too Many Calls to /userinfo",
	"api_limit": "Rate Limit On API",
	"sdu":       "Successful User Deletion",
	"fdu":       "Failed User Deletion",
	"slo":       "Success Logout",
	"flo":       "Failed Logout",
	"sd":        "Success Delegation",
	"fd":        "Failed Delegation",
	"fcoa":      "Failed Cross Origin Authentication",
	"scoa":      "Success Cross Origin Authentication",
}

// Log for analyzing business needs.
//
// See: https://auth0.com/docs/deploy-monitor/logs
type Log struct {
	ID *string `json:"_id"`

	// Unique ID of the log event.
	LogID *string `json:"log_id"`

	// The date when the log event was created.
	Date *time.Time `json:"date"`

	// The log event type.
	Type *string `json:"type"`

	// The log event description.
	Description *string `json:"description"`

	// Name of the connection the log event relates to.
	Connection *string `json:"connection"`

	// ID of the connection the log event relates to.
	ConnectionID *string `json:"connection_id"`

	// ID of the organization the log event relates to.
	OrganizationID *string `json:"organization_id"`

	// Name of the organization the log event relates to.
	OrganizationName *string `json:"organization_name"`

	// The ID of the client (application).
	ClientID *string `json:"client_id"`

	// The name of the client (application).
	ClientName *string `json:"client_name"`

	// The IP address of the log event source.
	IP *string `json:"ip"`

	// Hostname the log event applies to.
	Hostname *string `json:"hostname"`

	// Additional useful details about this event (structure is dependent upon event type).
	Details map[string]interface{} `json:"details"`

	// ID of the user involved in the log event.
	UserID *string `json:"user_id"`

	// Name of the user involved in the log event.
	UserName *string `json:"user_name"`

	// User agent string from the client device that caused the event.
	UserAgent *string `json:"user_agent"`

	// API audience the event applies to.
	Audience *string `json:"audience"`

	// Scope permissions applied to the event.
	Scope *string `json:"-"`

	// Name of the strategy involved in the event.
	Strategy *string `json:"strategy"`

	// Type of strategy involved in the event.
	StrategyType *string `json:"strategy_type"`

	// Whether the client was a mobile device (true) or desktop/laptop/server (false).
	IsMobile *bool `json:"isMobile"`

	// Information about the location that triggered this event based on the `ip`.
	LocationInfo map[string]interface{} `json:"location_info"`
}

// TypeName returns the type name of an Event Log.
func (l *Log) TypeName() string {
	if l.Type == nil {
		return ""
	}
	if name, ok := logTypeName[*l.Type]; ok {
		return name
	}
	return ""
}

// UnmarshalJSON is a custom deserializer for the Log type.
//
// It is required due to differences in the scope field which cane be a string or array of strings.
func (l *Log) UnmarshalJSON(data []byte) error {
	type log Log
	type logWrapper struct {
		*log
		RawScope interface{} `json:"scope"`
	}

	alias := &logWrapper{(*log)(l), nil}

	err := json.Unmarshal(data, alias)

	if err != nil {
		return err
	}

	if alias.RawScope != nil {
		switch rawScope := alias.RawScope.(type) {
		case []interface{}:
			scopes := make([]string, len(rawScope))
			for i, v := range rawScope {
				scopes[i] = v.(string)
			}
			scope := strings.Join(scopes, " ")
			l.Scope = &scope
		case string:
			l.Scope = &rawScope
		default:
			return fmt.Errorf("unexpected type for field scope: %T", alias.RawScope)
		}
	}

	return nil
}

// LogManager manages Auth0 Log resources.
type LogManager manager

// Read the data related to the log entry identified by id. This returns a
// single log entry representation as specified in the schema.
//
// See: https://auth0.com/docs/api/management/v2#!/Logs/get_logs_by_id
func (m *LogManager) Read(ctx context.Context, id string, opts ...RequestOption) (l *Log, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("logs", id), &l, opts...)
	return
}

// List all log entries that match the specified search criteria (or lists all
// log entries if no criteria are used). Set custom search criteria using the
// `q` parameter, or search from a specific log id ("search from checkpoint").
//
// For more information on all possible event types, their respective acronyms
// and descriptions, Log Data Event Listing.
//
// See: https://auth0.com/docs/api/management/v2#!/Logs/get_logs
func (m *LogManager) List(ctx context.Context, opts ...RequestOption) (l []*Log, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("logs"), &l, opts...)
	return
}

// Search is an alias for List.
func (m *LogManager) Search(ctx context.Context, opts ...RequestOption) ([]*Log, error) {
	return m.List(ctx, opts...)
}
