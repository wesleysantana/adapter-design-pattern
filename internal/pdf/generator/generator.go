package generator

import (
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/jung-kurt/gofpdf"
)

type GofpdfAdapter struct {
	rootPath string
}

func New(rootPath string) PDFGeneratorInterface {
	return &GofpdfAdapter{rootPath: rootPath}
}

// Modificamos a assinatura ou o comportamento para receber o texto/dados diretamente
func (g *GofpdfAdapter) Create(title string, body string) (string, error) {
	// Cria um novo PDF: Orientação Retrato ("P"), unidade em milímetros ("mm"), tamanho A4
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	// Configura o Título
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, title)
	pdf.Ln(12) // Quebra de linha

	// Configura o Corpo do texto
	pdf.SetFont("Arial", "", 12)
	// MultiCell é excelente para parágrafos longos, pois quebra a linha automaticamente
	pdf.MultiCell(0, 6, body, "", "", false)

	// Garante que a pasta de destino exista
	if err := os.MkdirAll(g.rootPath, os.ModePerm); err != nil {
		return "", err
	}

	// Gera um nome único para o PDF
	fileName := filepath.Join(g.rootPath, uuid.New().String()+".pdf")

	// Salva o arquivo no disco
	err := pdf.OutputFileAndClose(fileName)
	if err != nil {
		return "", err
	}

	return fileName, nil
}
