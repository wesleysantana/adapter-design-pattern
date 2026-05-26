package salesreportgenerator

import pdfg "github.com/wesleysantana/adapter-design-pattern/internal/pdf/generator"

type Content struct {
	Title   string
	Content string
}

type SalesReportGenerator struct {
	pdfGenerator pdfg.PDFGeneratorTarget
}

// Construtor que recebe QUALQUER adaptador de PDF
func New(pdfGenerator pdfg.PDFGeneratorTarget) *SalesReportGenerator {
	return &SalesReportGenerator{pdfGenerator: pdfGenerator}
}

func (s *SalesReportGenerator) Generate() (string, error) {
	content := Content{
		Title:   "Sales Report",
		Content: "This is the sales report content.",
	}

	fileName, err := s.pdfGenerator.Create(content.Title, content.Content)
	if err != nil {
		return "", err
	}

	return fileName, nil
}
