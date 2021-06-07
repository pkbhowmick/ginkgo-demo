package books_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pkbhowmick/ginkgo-demo/test/books"
)

var _ = Describe("Book", func() {
	var (
		longBook books.Book
		shortBook books.Book
	)

	BeforeEach(func() {
		longBook = books.Book{
			Title: "A",
			Author: "Victor",
			Pages: 2783,
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
	})
})

