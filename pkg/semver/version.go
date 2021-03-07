package semver

// Version represents semver version
//
// For more info see https://semver.org/
type Version interface {
	// Major major version
	Major() int
	// Minor minor version
	Minor() int
	// Patch patch version
	Patch() int
}

// CheckCompatible returns a boolean value indicating if v2 is compatible again v1.
// The v2 will be compatible if accoplish this constraints:
// - Has same major version
// - If major version is 0 the version should be equals
// - If major version is not 0 the version v1 minor should be equals or lower than v2 minor
func CheckCompatible(v1, v2 Version) bool {
	if v1.Major() != v2.Major() {
		return false
	}

	if v1.Major() == 0 {
		return Compare(v1, v2) == 0
	}

	return v1.Minor() <= v2.Minor()
}

// Compare returns an integer comparing two versions according to semantic version precedence.
// The result will be 0 if v1 == v2, -1 if v1 < v2, or +1 if v1 > v2.
func Compare(v1, v2 Version) int {
	if res := compare(v1.Major(), v2.Major()); res != 0 {
		return res
	}

	if res := compare(v1.Minor(), v2.Minor()); res != 0 {
		return res
	}

	return compare(v1.Patch(), v2.Patch())
}

func compare(v1, v2 int) int {
	if v1 < v2 {
		return -1
	} else if v1 > v2 {
		return 1
	}
	return 0
}
