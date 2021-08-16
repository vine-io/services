// Code generated by vine. DO NOT EDIT.
package runtime

const Namespace = "go.vine"


const (
	ApiserverName = "go.vine.service.apiserver"
)

var (
	ApiserverId = "70518331-b5ad-48ef-bebf-1bb232bb69c8"
)

var (
	GitTag    string
	GitCommit string
	BuildDate string
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