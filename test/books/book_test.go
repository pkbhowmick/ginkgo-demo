package books_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pkbhowmick/ginkgo-demo/test/books"
)

func doSomething() bool {
	return true
}

var _ = Describe("Book", func() {
	var (
		longBook  books.Book
		shortBook books.Book
	)

	BeforeEach(func() {
		longBook = books.Book{
			Title:  "A",
			Author: "Victor",
			Pages:  2783,
		}

		shortBook = books.Book{
			Title:  "B",
			Author: "Mike",
			Pages:  24,
		}
	})

	Describe("Categorized book length", func() {
		Context("With more than 300 pages", func() {
			It("should be a novel", func() {
				Expect(longBook.CategoryByLength()).To(Equal("NOVEL"))
			})
		})

		Context("With fewer than 300 pages", func() {
			It("should be a short story", func() {
				Expect(shortBook.CategoryByLength()).To(Equal("SHORT STORY"))
			})
		})

		Context("Test failure in Go routine", func() {
			It("panic in a go routine", func() {
				done := make(chan interface{})
				go func() {
					defer GinkgoRecover()

					Expect(doSomething()).Should(BeTrue())

					close(done)
				}()
				Eventually(done, 3).Should(BeClosed())
			})
		})

		Context("Test failure and recovery", func() {
			It("panic and recover by ginkgo", func() {
				defer GinkgoRecover()
				Fail("for some reason")
			})
		})
	})
})


