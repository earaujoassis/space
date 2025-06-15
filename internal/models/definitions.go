package models

const (
	// PublicClient client type
	PublicClient string = "public"
	// ConfidentialClient client type
	ConfidentialClient string = "confidential"
)

const (
	// NotSpecialAction action description, used for ephemeral actions with no special meaning
	NotSpecialAction string = "not_special"
	// UpdateUserAction action description, user for ephemeral actions updating user data
	UpdateUserAction string = "update_user"
)

const (
	// AccessToken token type
	AccessToken string = "access_token"
	// RefreshToken token type
	RefreshToken string = "refresh_token"
	// GrantToken token type
	GrantToken string = "grant_token"
	// IDToken token type
	IDToken string = "id_token"

	// PublicScope session scope
	// This is used by public clients (they can't read or write user data)
	PublicScope string = "public"
	// ReadScope session scope
	// This is used by confidential clients (they can only read user data)
	ReadScope string = "read"
	// WriteScope session scope
	// No client is allowed to hold this scope (they can't write user data)
	WriteScope string = "write"
	// OpenIDScope session scope
	// This is used for OpenID Connect and confidential clients
	OpenIDScope string = "openid"
	// ProfileScope session scope
	// This is used for OpenID Connect and confidential clients
	ProfileScope string = "openid"
)

const (
	// PublicService service type
	PublicService string = "public"
	// AttachedService service type
	AttachedService string = "attached"
)
