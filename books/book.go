package books

type Book struct {
	Title  string
	Author string
	Pages  string
}

func (b *Book) CategoryByLength() string {
	if b.Pages > 300 {
		return "NOVEL"
	} else {
		return "SHORT STORY"
	}
}
