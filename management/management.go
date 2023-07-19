package management

//go:generate go run gen-methods.go

import (
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/oauth2"

	"github.com/auth0/go-auth0/internal/client"
)

// Management is an Auth0 management client used to interact with the Auth0
// Management API v2.
type Management struct {
	// Client manages Auth0 Client (also known as Application) resources.
	Client *ClientManager

	// ClientGrant manages Auth0 ClientGrant resources.
	ClientGrant *ClientGrantManager

	// ResourceServer manages Auth0 Resource Server (also known as API)
	// resources.
	ResourceServer *ResourceServerManager

	// Connection manages Auth0 Connection resources.
	Connection *ConnectionManager

	// CustomDomain manages Auth0 Custom Domains.
	CustomDomain *CustomDomainManager

	// Grant manages Auth0 Grants.
	Grant *GrantManager

	// Log reads Auth0 Logs.
	Log *LogManager

	// LogStream reads Auth0 Logs.
	LogStream *LogStreamManager

	// RoleManager manages Auth0 Roles.
	Role *RoleManager

	// RuleManager manages Auth0 Rules.
	Rule *RuleManager

	// HookManager manages Auth0 Hooks
	Hook *HookManager

	// RuleManager manages Auth0 Rule Configurations.
	RuleConfig *RuleConfigManager

	// EmailTemplate manages Auth0 Email Templates.
	EmailTemplate *EmailTemplateManager

	// User manages Auth0 User resources.
	User *UserManager

	// Job manages Auth0 jobs.
	Job *JobManager

	// Tenant manages your Auth0 Tenant.
	Tenant *TenantManager

	// Ticket creates verify email or change password tickets.
	Ticket *TicketManager

	// Stat is used to retrieve usage statistics.
	Stat *StatManager

	// Branding settings such as company logo or primary color.
	Branding *BrandingManager

	// Guardian manages your Auth0 Guardian settings
	Guardian *GuardianManager

	// Prompt manages your prompt settings.
	Prompt *PromptManager

	// Blacklist manages the auth0 blacklists
	Blacklist *BlacklistManager

	// SigningKey manages Auth0 Application Signing Keys.
	SigningKey *SigningKeyManager

	// Anomaly manages the IP blocks
	Anomaly *AnomalyManager

	// Actions manages Actions extensibility
	Action *ActionManager

	// Organization manages Auth0 Organizations.
	Organization *OrganizationManager

	// AttackProtection manages Auth0 Attack Protection.
	AttackProtection *AttackProtectionManager

	// BrandingTheme manages Auth0 Branding Themes.
	BrandingTheme *BrandingThemeManager

	// EmailProvider manages Auth0 Email Providers.
	EmailProvider *EmailProviderManager

	url             *url.URL
	basePath        string
	userAgent       string
	debug           bool
	tokenSource     oauth2.TokenSource
	http            *http.Client
	auth0ClientInfo *client.Auth0ClientInfo
	common          manager
	retryStrategy   client.RetryOptions
}

type manager struct {
	management *Management
}

// New creates a new Auth0 Management client by authenticating using the
// supplied client id and secret.
func New(domain string, options ...Option) (*Management, error) {
	// Ignore the scheme if it was defined in the domain variable, then prefix
	// with https as it's the only scheme supported by the Auth0 API.
	if i := strings.Index(domain, "//"); i != -1 {
		domain = domain[i+2:]
	}
	domain = "https://" + domain

	u, err := url.Parse(domain)
	if err != nil {
		return nil, err
	}

	m := &Management{
		url:             u,
		basePath:        "api/v2",
		userAgent:       client.UserAgent,
		debug:           false,
		http:            http.DefaultClient,
		auth0ClientInfo: client.DefaultAuth0ClientInfo,
		retryStrategy:   client.DefaultRetryOptions,
	}

	for _, option := range options {
		option(m)
	}

	m.http = client.WrapWithTokenSource(
		m.http,
		m.tokenSource,
		client.WithDebug(m.debug),
		client.WithUserAgent(m.userAgent),
		client.WithAuth0ClientInfo(m.auth0ClientInfo),
		client.WithRetries(m.retryStrategy),
	)

	m.common.management = m

	m.Action = (*ActionManager)(&m.common)
	m.Anomaly = (*AnomalyManager)(&m.common)
	m.AttackProtection = (*AttackProtectionManager)(&m.common)
	m.Blacklist = (*BlacklistManager)(&m.common)
	m.Branding = (*BrandingManager)(&m.common)
	m.BrandingTheme = (*BrandingThemeManager)(&m.common)
	m.Client = (*ClientManager)(&m.common)
	m.ClientGrant = (*ClientGrantManager)(&m.common)
	m.Connection = (*ConnectionManager)(&m.common)
	m.CustomDomain = (*CustomDomainManager)(&m.common)
	m.EmailProvider = (*EmailProviderManager)(&m.common)
	m.EmailTemplate = (*EmailTemplateManager)(&m.common)
	m.Grant = (*GrantManager)(&m.common)
	m.Guardian = &GuardianManager{
		Enrollment: (*EnrollmentManager)(&m.common),
		MultiFactor: &MultiFactorManager{
			manager:          m.common,
			DUO:              (*MultiFactorDUO)(&m.common),
			Email:            (*MultiFactorEmail)(&m.common),
			OTP:              (*MultiFactorOTP)(&m.common),
			Phone:            (*MultiFactorPhone)(&m.common),
			Push:             (*MultiFactorPush)(&m.common),
			RecoveryCode:     (*MultiFactorRecoveryCode)(&m.common),
			SMS:              (*MultiFactorSMS)(&m.common),
			WebAuthnPlatform: (*MultiFactorWebAuthnPlatform)(&m.common),
			WebAuthnRoaming:  (*MultiFactorWebAuthnRoaming)(&m.common),
		},
	}
	m.Hook = (*HookManager)(&m.common)
	m.Job = (*JobManager)(&m.common)
	m.Log = (*LogManager)(&m.common)
	m.LogStream = (*LogStreamManager)(&m.common)
	m.Organization = (*OrganizationManager)(&m.common)
	m.Prompt = (*PromptManager)(&m.common)
	m.ResourceServer = (*ResourceServerManager)(&m.common)
	m.Role = (*RoleManager)(&m.common)
	m.Rule = (*RuleManager)(&m.common)
	m.RuleConfig = (*RuleConfigManager)(&m.common)
	m.SigningKey = (*SigningKeyManager)(&m.common)
	m.Stat = (*StatManager)(&m.common)
	m.Tenant = (*TenantManager)(&m.common)
	m.Ticket = (*TicketManager)(&m.common)
	m.User = (*UserManager)(&m.common)

	return m, nil
}
