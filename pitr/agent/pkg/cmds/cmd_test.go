package cmds

import "testing"

const (
	sh = "/bin/sh"
)

func TestCommand(t *testing.T) {
	if err := command(sh, "ping www.baidu.com"); err != nil {
		t.Fatal(err)
	}
}
