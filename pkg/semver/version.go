package semver

// Version represents version based in semver
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
