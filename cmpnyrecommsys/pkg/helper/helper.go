package helper

import "fmt"

// StringToUint Converts to uint if valid input else returns error
func StringToUint(s string) (uint, error) {
	var u uint
	if _, err := fmt.Sscanf(s, "%d", &u); err != nil {
		return 0, err
	}
	return u, nil
}
