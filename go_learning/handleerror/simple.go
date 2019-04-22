package handleerror

import (
	"os"
	"syscall"
)

// 每次操作都要检查错误
func simple() error {
	fd, _ := os.OpenFile("test", syscall.O_RDWR, 0666)
	p := []byte{'a', 'b'}
	_, err := fd.Write(p)
	if err != nil {
		return err
	}
	_, err = fd.Write(p)
	if err != nil {
		return err
	}
	_, err = fd.Write(p)
	if err != nil {
		return err
	}
}
