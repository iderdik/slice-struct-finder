package finder_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestFinder(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Finder Suite")
}
