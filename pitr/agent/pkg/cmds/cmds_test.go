package cmds

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCmds(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Cmds suits")
}
