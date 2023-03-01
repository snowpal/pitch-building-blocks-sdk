package golang

const Host = "https://gateway-dev.snowpal.com/"
const XApiKey = "wf8sHELzWp9MGizZME5Zsjk4IntZS0e8mdYMYjjg"

const (
	UrlSignUp   = Host + "app/users/sign-up"
	UrlSignIn   = Host + "app/users/sign-in"
	UrlActivate = Host + "app/user-verified/%s"
	UrlGeyKeys  = Host + "keys"
)

const (
	UrlGetAttributes       = Host + "app/resource/attributes"
	UrlUpdateKeyAttributes = Host + "keys/%s/attributes"
)
