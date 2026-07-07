package lc278

var BAD = 0

func isBadVersion(version int) bool {
	if version >= BAD {
		return true
	}
	return false
}
