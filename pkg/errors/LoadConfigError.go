package errors

import "github.com/pkg/errors"

const LoadConfigErrorString = "failed to load config: environment string is empty"

func LoadConfigError() error {
	return errors.New(LoadConfigErrorString)
}
