package dhcpop

import (
	"fmt"
)

func validateSize(a []byte, s int) error {
	if len(a) != s {
		return &InvalidSizeError{
			Message: fmt.Sprintf("invalid size: expected %d bytes, but got %d bytes", s, len(a)),
		}
	}
	return nil
}

func validateSizeFactor(a []byte, s int) error {
	if len(a)%s > 0 {
		return &InvalidSizeError{
			Message: fmt.Sprintf("invalid size: expected mlutiplies of %d bytes, but got %d bytes", s, len(a)),
		}
	}
	return nil
}

func validateMinimumSize(a []byte, s int) error {
	if len(a) < s {
		return &InvalidSizeError{
			Message: fmt.Sprintf("invalid size: expected >= %d bytes, but got %d bytes", s, len(a)),
		}
	}
	return nil
}

type InvalidSizeError struct {
	Message string
}

func (err *InvalidSizeError) Error() string {
	return err.Message
}

type InvalidFormatError struct {
	Message string
}

func (err *InvalidFormatError) Error() string {
	return err.Message
}
