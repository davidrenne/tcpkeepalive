package tcpkeepalive

import (
	"errors"
)

func setIdle(fd int, secs int) error {
	return errors.New("not supported")
}

func setCount(fd int, n int) error {
	return errors.New("not supported")
}

func setInterval(fd int, secs int) error {
	return errors.New("not supported")
}

func setNonblock(fd int) error {
	return errors.New("not supported")
}
