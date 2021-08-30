package auth

const Namespace = "go.vine"

const (
	Name = "go.vine.service.auth"
)

var (
	Id = "bab996ea-f3b3-44dc-8268-3365ce90aa7c"
)

var (
	GitTag     string
	GitCommit  string
	BuildDate  string
	GetVersion = func() string {
		v := GitTag
		if GitCommit != "" {
			v += "-" + GitCommit
		}
		if BuildDate != "" {
			v += "-" + BuildDate
		}
		if v == "" {
			return "latest"
		}
		return v
	}
)
