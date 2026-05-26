package salesreportgenerator

import pdfg "github.com/wesleysantana/adapter-design-pattern/internal/pdf/generator"

type Content struct {
	Title   string
	Content string
}

func Generate() (string, error) {
	fileName, err := createPDF()
	if err != nil {
		return "", err
	}

	return fileName, nil
}

func createPDF() (string, error) {
	content := Content{
		Title:   "Sales Report",
		Content: "This is the sales report content.",
	}

	wk := pdfg.New("tmp")

	fileName, err := wk.Create(content.Title, content.Content)
	if err != nil {
		return "", err
	}

	return fileName, nil
}
