package cmds

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
)

func command(name string, args ...string) error {
	c := "-c"
	args = append([]string{c}, args...)

	cmd := exec.Command(name, args...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("can not obtain stdout pipe for command[args=%+v]:%s", args, err)
	}
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("the command is err[args=%+v]:%s", args, err)
	}

	reader := bufio.NewReader(stdout)

	index := 1
	for {
		line, err := reader.ReadString('\n')
		if io.EOF == err {
			break
		} else if err != nil {
			return fmt.Errorf("read string is err[args=%+v]:%s", args, err)
		}

		fmt.Print(index, "\t", line)
		index++
	}

	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("cmd wait is err[args=%+v]:%s", args, err)
	}

	return nil
}
