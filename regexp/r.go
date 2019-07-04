package regexp

import "regexp"

var (
	Phone = regexp.MustCompile(`^\+?0?(86)?-?1\d{10}$`)
)
