package xio

import (
	"errors"
	"io"

	"github.com/lukaproject/xerr"
)

const (
	defaultBlockSize = 1 << 10
)

var (
	ErrTransportMismatch = errors.New("writer reader transport mismatch occur")
)

func Transport(
	reader io.Reader,
	writer io.Writer,
) (err error) {
	return transportWithSize(reader, writer, defaultBlockSize)
}

func transportWithSize(
	reader io.Reader,
	writer io.Writer,
	blockSize int,
) (err error) {
	buf := make([]byte, blockSize)
	for {
		n, err := reader.Read(buf)
		if n > 0 {
			n1 := xerr.Must(writer.Write(buf[:n]))
			if n1 != n {
				return ErrTransportMismatch
			}
		}
		if err != nil {
			if err == io.EOF {
				err = nil
				break
			} else {
				return err
			}
		}
	}
	return
}
