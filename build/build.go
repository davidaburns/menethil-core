package build

import "fmt"

var (
	name = "menethil-core"
	version = "0.0.1"
	buildNumber = "0"
	buildTime = ""
	buildType = ""
)

const (
	BuildTypeDev string = "dev"
	BuildTypeProd string = "prod"
)

type BuildInfo struct {
	Name string
	Version string
	BuildNumber string
	BuildTime string
	BuildType string
	BuildCommit string
}

func NewBuildInfo() *BuildInfo {
	return &BuildInfo {
		Name: name,
		Version: version,
		BuildNumber: buildNumber,
		BuildTime: buildTime,
		BuildType: buildType,
	}
}

func (build *BuildInfo) String() string {
	return fmt.Sprintf("%s-%s.%s:b%s", build.Name, build.Version, build.BuildType, build.BuildNumber)
}
