package service

import (
	l4g "code.google.com/p/log4go"
	"syscall"
)

func OnOsCreateError(err error) (bool, error) {
	if err != nil {
		l4g.Error(err)
		if err == syscall.ERROR_ACCESS_DENIED {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func IsFileNotFoundError(err error) bool {
	if err != nil {
		if err == syscall.ERROR_FILE_NOT_FOUND {
			return true
		}
	}
	return false
}