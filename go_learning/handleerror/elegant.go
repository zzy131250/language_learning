package handleerror

import (
	"io"
	"os"
	"syscall"
)

type errWriter struct {
	w   io.Writer
	err error
}

func (ew *errWriter) write(buf []byte) {
	if ew.err != nil {
		return
	}
	_, ew.err = ew.w.Write(buf)
}

// 使用实例变量存储错误，快速失败
func elegant() error {
	fd, _ := os.OpenFile("test", syscall.O_RDWR, 0666)
	p := []byte{'a', 'b'}
	ew := &errWriter{w: fd}
	ew.write(p)
	ew.write(p)
	ew.write(p)
	if ew.err != nil {
		return ew.err
	}
	return nil
}
