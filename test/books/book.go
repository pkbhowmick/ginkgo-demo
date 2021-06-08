package books

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Pages  int    `json:"pages"`
}

func (b *Book) CategoryByLength() string {
	if b.Pages >= 300 {
		return "NOVEL"
	}
	return "SHORT STORY"
}
