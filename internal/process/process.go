package process

import (
	"fmt"
	"os"
	"strconv"
	"syscall"
)

type ProcessPriority int

const (
	PROCESS_PRIORITY_HIGH   ProcessPriority = 0
	PROCESS_PRIORITY_MEDIUM ProcessPriority = 10
	PROCESS_PRIORITY_LOW    ProcessPriority = 20
)

type SystemProcess struct {
	Pid  int
	Path string
}

func GetSystemProcess() (*SystemProcess, error) {
	exePath, err := os.Executable()
	if err != nil {
		return nil, err
	}

	return &SystemProcess{
		Pid:  os.Getpid(),
		Path: exePath,
	}, nil
}

func (sp *SystemProcess) GetProcessPriority() (int, error) {
	prio, err := syscall.Getpriority(syscall.PRIO_PROCESS, sp.Pid)
	if err != nil {
		return -1, err
	}

	return prio, nil
}

func (sp *SystemProcess) SetProcessPriority(priority ProcessPriority) error {
	if err := syscall.Setpriority(syscall.PRIO_PROCESS, sp.Pid, int(priority)); err != nil {
		return err
	}

	return nil
}

func (sp *SystemProcess) CreatePidFile() error {
	pid := strconv.Itoa(sp.Pid)
	path := fmt.Sprintf("%v.pid", sp.Path)

	if err := os.WriteFile(path, []byte(pid), 0644); err != nil {
		return err
	}

	return nil
}
