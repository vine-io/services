package apiserver

const Namespace = "go.vine"

const (
	Name = "go.vine.service.apiserver"
)

var (
	Id = "744f4422-694b-4ebf-a22c-5fed009c70fc"
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
