package notifications

func mapNotificationNameToSettings(name string) string {
	switch name {
	case "client.token_introspection":
		return "notifications.client-application-email-notifications.token-introspection"
	case "client.userinfo_introspection":
		return "notifications.client-application-email-notifications.userinfo-introspection"
	case "user.authorization_granted":
		return "notifications.system-email-notifications.client-authorization"
	case "session.created":
		return "notifications.system-email-notifications.authentication"
	}

	return ""
}
