// +build integration

package integration_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCloudMtaBuildTool(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "MtaBuildTool Integration Suite")
}
