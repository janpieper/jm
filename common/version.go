package common

import (
	"fmt"
	"runtime"
)

var (
	version      = "0.0.0"
	buildDate    = "1970-01-01T00:00:00Z"
	gitCommit    = ""
	gitTag       = ""
	gitTreeState = ""
)

type Version struct {
	Version      string
	BuildDate    string
	GitCommit    string
	GitTag       string
	GitTreeState string
	GoVersion    string
	Compiler     string
	Platform     string
}

func (v Version) String() string {
	return v.Version
}

func GetVersion() Version {
	var versionStr string

	if gitCommit != "" && gitTag != "" && gitTreeState == "clean" {
		versionStr = gitTag
	} else {
		versionStr = "v" + version
		if len(gitCommit) >= 7 {
			versionStr += "+" + gitCommit[0:7]
			if gitTreeState != "clean" {
				versionStr += ".dirty"
			}
		} else {
			versionStr += "+unknown"
		}
	}

	return Version{
		Version:      versionStr,
		BuildDate:    buildDate,
		GitCommit:    gitCommit,
		GitTag:       gitTag,
		GitTreeState: gitTreeState,
		GoVersion:    runtime.Version(),
		Compiler:     runtime.Compiler,
		Platform:     fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}
