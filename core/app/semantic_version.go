package app

import (
	"fmt"
	"strconv"
	"strings"
)

// ParseVersion returns a SemanticVersion parsed from a string in the format "X.X.X" where X are integers
func ParseVersion(str string) (SemanticVersion, error) {
	ver := SemanticVersion{}
	err := ver.Parse(str)
	return ver, err
}

// SemanticVersion represents the versioning of software using Semantic Versioning 2.0.0
// See https://semver.org/ for more information regarding semantic versioning
type SemanticVersion struct {
	MajorRelease int
	MinorRelease int
	Patch        int
}

// String implements the string representation of a semantic version
func (version SemanticVersion) String() string {
	return fmt.Sprintf("%v.%v.%v", version.MajorRelease, version.MinorRelease, version.Patch)
}

// Parse parses the values for the version from a string in the format of "X.X.X" where X are integers
func (version SemanticVersion) Parse(str string) error {
	substrs := strings.Split(str, ".")
	var err error
	version.MajorRelease, err = strconv.Atoi(substrs[0])
	if err != nil {
		return fmt.Errorf("failed to parse MajorRelease from string \"%s\".\n Atoi Error: %s", substrs[0], err.Error())
	}
	version.MinorRelease, err = strconv.Atoi(substrs[1])
	if err != nil {
		return fmt.Errorf("failed to parse MinorRelease from string \"%s\".\n Atoi Error: %s", substrs[1], err.Error())
	}
	version.Patch, err = strconv.Atoi(substrs[2])
	if err != nil {
		return fmt.Errorf("failed to parse Patch from string \"%s\".\n Atoi Error: %s", substrs[2], err.Error())
	}
	return nil
}
