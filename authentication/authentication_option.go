package authentication

// Option is used for passing options to the Authentication client.
type Option func(*Authentication)

// WithClientID configures the Client ID to be provided with requests if one is not provided.
func WithClientID(clientID string) Option {
	return func(a *Authentication) {
		a.clientID = clientID
	}
}

// WithClientSecret configures the Client ID to be provided with requests if one is not provided.
func WithClientSecret(clientSecret string) Option {
	return func(a *Authentication) {
		a.clientSecret = clientSecret
	}
}
