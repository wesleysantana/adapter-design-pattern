package generator

type PDFGeneratorInterface interface {
	Create(title string, body string) (string, error)
}
