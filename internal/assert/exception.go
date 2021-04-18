package assert

import (
	"errors"
	"fmt"
	"strings"
)

func exception(message string, expected, actual interface{}) error {
	m := fmt.Sprintf("%s failed: expected %T(%#v), but got %T(%#v)", message, expected, expected, actual, actual)
	m = strings.TrimSpace(m)
	return errors.New(m)
}
