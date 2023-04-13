package visitor

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestVisitor(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Visitor Suite")
}
