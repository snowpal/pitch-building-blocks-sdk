package response

type Version struct {
	Status              string  `json:"status"`
	Version             string  `json:"version"`
	InternalVersion     string  `json:"internalVersion"`
	DeploymentIteration *string `json:"deploymentIteration"`
}
