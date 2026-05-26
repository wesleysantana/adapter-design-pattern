package generator

type PDFGeneratorTarget interface {
	Create(title string, body string) (string, error)
}
