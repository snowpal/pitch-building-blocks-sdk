package golang

const XApiKey = "wf8sHELzWp9MGizZME5Zsjk4IntZS0e8mdYMYjjg"

const (
	HostDev = "https://gateway-dev.snowpal.com/"
)

const (
	RouteSignUp   = "app/users/sign-up"
	RouteSignIn   = "app/users/sign-in"
	RouteActivate = "app/user-verified/%s"
	RouteGeyKeys  = "keys"
)

const (
	RouteGetAttributes            = "app/resource/attributes"
	RouteUpdateKeyAttributes      = "keys/%s/attributes"
	RouteUpdateBlockAttributes    = "blocks/%s/attributes"
	RouteUpdatePodAttributes      = "pods/%s/attributes"
	RouteUpdateBlockPodAttributes = "block-pods/%s/attributes"
)

const (
	RouteGetBlocks             = "keys/%s/blocks?filter=%s&batchIndex=%s&aclWriteOrHigher=%s"
	RouteAddBlock              = "keys/%s/blocks"
	RouteGetBlocksLinkedToPods = "pods/%s/linked-to/blocks?keyId=%s"
)
