package handleerror

import (
	"os"
	"syscall"
)

// 使用局部变量存储错误，快速失败
func useFunc() error {
	fd, _ := os.OpenFile("test", syscall.O_RDWR, 0666)
	p := []byte{'a', 'b'}
	var err error
	write := func(buf []byte) {
		if err != nil {
			return
		}
		_, err = fd.Write(buf)
	}
	write(p)
	write(p)
	write(p)
	if err != nil {
		return err
	}
	return nil
}
