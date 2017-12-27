package oauth2

type Auth interface {
	Authenticated() error
}

type Authentication struct {
	ID    string
	Token string
}

func Authenticate(auth Auth) {
	auth.Authenticated()
}
