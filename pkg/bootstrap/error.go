package bootstrap

import "fmt"

// CustomError ...
func CustomError(layer string, method string, err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%s.%s:%s", layer, method, err.Error())
}
