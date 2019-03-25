package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
	"syscall"
	"time"

	"../pidusage"
)

type ProcessResult struct {
	Stdout      string
	Stderr      string
	SystemTime  time.Duration
	RealTime    time.Duration
	MemoryUsage int

	In          string
	ExecCmdText string
}

func (p *ProcessResult) DumpJSON() string {
	buf, _ := json.Marshal(p)
	return string(buf)
}

func LogProcessState(command *exec.Cmd) {
	fmt.Printf("[LOG] command.ProcessState.SysUsage() = %v\n", command.ProcessState.SysUsage())
	fmt.Printf("[LOG] command.ProcessState.SystemTime() = %v\n", command.ProcessState.SystemTime())
	fmt.Printf("[LOG] command.ProcessState.UserTime() = %v\n", command.ProcessState.UserTime())
}

func SetTimeoutExecCmdAndInput(cmd string, stdin string, timeout int) (p ProcessResult, err error) {
	p.In = stdin
	p.ExecCmdText = cmd

	if len(cmd) == 0 {
		err = fmt.Errorf("cannot run a empty command")
		return
	}
	var outbuf, errbuf bytes.Buffer
	args := strings.Fields(cmd)
	command := exec.Command(args[0], args[1:]...)
	command.Stdout = &outbuf
	command.Stderr = &errbuf
	command.Stdin = strings.NewReader(stdin)

	startTimer := time.Now()
	command.Start()
	pUsage, uerr := pidusage.GetStat(command.Process.Pid)
	if uerr == nil {
		p.MemoryUsage = int(pUsage.Memory)
	}

	if timeout > 0 {
		done := make(chan error)
		go func() { done <- command.Wait() }()

		after := time.After(time.Duration(timeout) * time.Millisecond)
		select {
		case <-after:
			command.Process.Signal(syscall.SIGINT)
			time.Sleep(time.Second)
			command.Process.Kill()
			err = fmt.Errorf("Timeout")
		case <-done:
		}
	} else {
		err = command.Wait()
	}

	p.Stderr = trimOutput(errbuf)
	p.Stdout = trimOutput(outbuf)
	p.RealTime = time.Since(startTimer)
	p.SystemTime = command.ProcessState.SystemTime()

	// fmt.Printf("[LOG] p.DumpJSON() = %v\n", p.DumpJSON())
	// fmt.Printf("[LOG] command.ProcessState.SysUsage().(*syscall.Rusage) = %v\n", command.ProcessState.SysUsage().(*syscall.Rusage))

	return p, err
}

func ExecCmd(cmd string) (p ProcessResult, err error) {
	return SetTimeoutExecCmdAndInput(cmd, "", -1)
}

func ExecCmdAndInput(cmd string, stdin string) (p ProcessResult, err error) {
	return SetTimeoutExecCmdAndInput(cmd, stdin, -1)
}
