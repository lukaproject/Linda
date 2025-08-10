package procs

import (
	"os"
	"os/exec"
)

type Process struct {
	Pid int

	OSProc *os.Process `json:"-"`
}

// Record 把command的process记录下来
func (p *Process) Record(cmd *exec.Cmd) {
	p.Pid = cmd.Process.Pid
	p.OSProc = cmd.Process
}
