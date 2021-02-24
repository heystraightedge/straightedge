package types

import (
	"regexp"
)

// IsAlphaNumeric checks if string is in the alphabet and numeric
var IsAlphaNumeric = regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString
