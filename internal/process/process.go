package process

import (
	"fmt"
	"os"
	"strconv"
)

func CreatePIDFile() error {
	exePath, err := os.Executable()
	fmt.Println(exePath)
	if err != nil {
		return err
	}

	pid := strconv.Itoa(os.Getpid())
	pidPath := fmt.Sprintf("%v.pid", exePath)
	if err := os.WriteFile(pidPath, []byte(pid), 0644); err != nil {
		return err
	}

	return nil
}
