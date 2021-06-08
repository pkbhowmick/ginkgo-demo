package books_test

import (
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBooks(t *testing.T) {
	RegisterFailHandler(Fail)            // This is the sole connection point between Ginkgo and Gomega
	RunSpecs(t, "Books Suite")
}

var _ = BeforeSuite(func() {
	fmt.Println("Some initialization before the entire test")
})

var _ = AfterSuite(func() {
	fmt.Println("Clean up after the entire test")
})
