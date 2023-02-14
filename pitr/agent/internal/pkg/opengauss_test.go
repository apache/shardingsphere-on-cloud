package pkg

import (
	"fmt"
	"testing"
)

func TestOpenGauss_AsyncBackup(t *testing.T) {
	og := &openGauss{
		shell: "/bin/sh",
	}
	backupID, err := og.AsyncBackup(
		"/home/omm/data",
		"ins-default-0",
		"full",
		"/data/opengauss/3.1.1/data/single_node/",
	)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(backupID)
}
