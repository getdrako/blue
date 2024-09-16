package build

import (
	"fmt"
	"strings"
)

const (
	VERSION_TYPE_CUSTOM       = "custom"
	VERSION_TYPE_EXPERIMENTAL = "experimental"
	VERSION_TYPE_UNSTABLE     = "unstable"
	VERSION_TYPE_UNTESTED     = "untested"
	VERSION_TYPE_RELEASE      = "release"
)

var (
	VersionType      = VERSION_TYPE_CUSTOM
	DRAKO_DEBUG_BOOL = "true"

	DRAKO_VERSION_MAJOR = "0"
	DRAKO_VERSION_MINOR = "0"
	DRAKO_VERSION_PATCH = "0"
	Tag                 = fmt.Sprintf("%s.%s.%s", DRAKO_VERSION_MAJOR, DRAKO_VERSION_MINOR, DRAKO_VERSION_PATCH)
	Commit              = "-"
	Datetime            = "0000-00-00T00:00:00"
)

func init() {
	if VersionType != VERSION_TYPE_RELEASE {
		Tag = fmt.Sprintf("%s.%s.%s-%s", DRAKO_VERSION_MAJOR, DRAKO_VERSION_MINOR, DRAKO_VERSION_PATCH, VersionType)
	}

	// Convert string DRAKO_DEBUG to bool
	if strings.ToLower(DRAKO_DEBUG_BOOL) == "true" {
		DRAKO_DEBUG = true
	} else {
		DRAKO_DEBUG = false
	}

}

var DRAKO_DEBUG bool
