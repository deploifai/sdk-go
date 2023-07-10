package host

type endpoint struct {
	GraphQL string
	Rest    rest
}

type rest struct {
	Base string
	Auth auth
}

type auth struct {
	Login string
}

const base = "https://api.deploif.ai"

var Endpoint = endpoint{
	GraphQL: base + "/graphql",

	Rest: rest{
		Base: base,

		Auth: auth{
			Login: "/auth/login/cli",
		},
	},
}
