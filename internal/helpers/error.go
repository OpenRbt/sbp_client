package helpers

import "fmt"

func CustomError(layer string, method string, err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%s.%s:%w", layer, method, err)
}
