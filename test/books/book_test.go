package books_test

import (
	"encoding/json"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pkbhowmick/ginkgo-demo/test/books"
)

func doSomething() bool {
	return true
}

func NewBookFromJson(data []byte) (books.Book, error) {
	var book books.Book
	err := json.Unmarshal(data, &book)
	if err != nil {
		return books.Book{}, err
	}
	return book, nil
}

var _ = Describe("Book", func() {
	var (
		longBook  books.Book
		shortBook books.Book
		book      books.Book
		err       error
		json      []byte
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

		json = []byte(`{
            "title":"Les Miserables",
			"author":"Victor Hugo",
			"pages":2783
            }`)
	})

	// JustBeforeEach blocks are guaranteed to be run
	// after all the BeforeEach blocks have run
	// and just before the It block has run.
	JustBeforeEach(func() {
		book, err = NewBookFromJson(json)
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

	Describe("Test failure and recovery", func() {
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

		//Context("Test failure and recovery", func() {
		//	It("panic and recover by ginkgo", func() {
		//		defer GinkgoRecover()
		//		Fail("for some reason")
		//	})
		//})
	})

	// Some important lines from doc:
	// You use Describe blocks to describe the individual behaviors of your code
	// and Context blocks to exercise those behaviors under different circumstances.
	// In this example we Describe loading a book from JSON and specify two Contexts:
	// when the JSON parses succesfully and when the JSON fails to parse.
	Describe("loading from json", func() {
		Context("when the json parses successfully", func() {
			It("should not error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
			It("should populate the fields correctly", func() {
				Expect(book.Title).To(Equal("Les Miserables"))
				Expect(book.Author).To(Equal("Victor Hugo"))
				Expect(book.Pages).To(Equal(2783))
			})
		})
		Context("when the json fails to parse", func() {
			BeforeEach(func() {
				json = []byte(`{
					"title":"Les Miserables",
					"author":"Victor Hugo",
					"pages":2783invalid
				}`)
			})
			It("should return the zero value for the book", func() {
				Expect(book).To(BeZero())
			})

			It("should error", func() {
				Expect(err).To(HaveOccurred())
			})
		})
	})
})
