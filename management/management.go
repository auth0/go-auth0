package management

//go:generate go run gen-methods.go

import (
	"context"
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

	// Email manages Auth0 Email Providers.
	// Deprecated: Use EmailProvider instead.
	Email *EmailManager

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
	ctx             context.Context
	tokenSource     oauth2.TokenSource
	http            *http.Client
	auth0ClientInfo *client.Auth0ClientInfo
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
		ctx:             context.Background(),
		http:            http.DefaultClient,
		auth0ClientInfo: client.DefaultAuth0ClientInfo,
	}

	for _, option := range options {
		option(m)
	}

	m.http = client.Wrap(
		m.http,
		m.tokenSource,
		client.WithDebug(m.debug),
		client.WithUserAgent(m.userAgent),
		client.WithRateLimit(),
		client.WithAuth0ClientInfo(m.auth0ClientInfo),
	)

	m.Client = newClientManager(m)
	m.ClientGrant = newClientGrantManager(m)
	m.Connection = newConnectionManager(m)
	m.CustomDomain = newCustomDomainManager(m)
	m.Grant = newGrantManager(m)
	m.LogStream = newLogStreamManager(m)
	m.Log = newLogManager(m)
	m.ResourceServer = newResourceServerManager(m)
	m.Role = newRoleManager(m)
	m.Rule = newRuleManager(m)
	m.Hook = newHookManager(m)
	m.RuleConfig = newRuleConfigManager(m)
	m.EmailTemplate = newEmailTemplateManager(m)
	m.Email = newEmailManager(m)
	m.User = newUserManager(m)
	m.Job = newJobManager(m)
	m.Tenant = newTenantManager(m)
	m.Ticket = newTicketManager(m)
	m.Stat = newStatManager(m)
	m.Branding = newBrandingManager(m)
	m.Guardian = newGuardianManager(m)
	m.Prompt = newPromptManager(m)
	m.Blacklist = newBlacklistManager(m)
	m.SigningKey = newSigningKeyManager(m)
	m.Anomaly = newAnomalyManager(m)
	m.Action = newActionManager(m)
	m.Organization = newOrganizationManager(m)
	m.AttackProtection = newAttackProtectionManager(m)
	m.BrandingTheme = newBrandingThemeManager(m)
	m.EmailProvider = newEmailProviderManager(m)

	return m, nil
}
