package caller

import (
	"fmt"
)

// Wrap error with caller name, e.g., "caller.testWrapObj.Fn: some err"
func Wrap(err error) error {
	if err == nil {
		return nil
	}

	name := getName()

	return fmt.Errorf("%s: %v", name, err)
}

// Wrap error with current package name, e.g., "caller: some err"
func WrapPackage(err error) error {
	if err == nil {
		return nil
	}

	return fmt.Errorf("%s: %w", getPackage(), err)
}
