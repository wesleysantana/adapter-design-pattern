package main

import (
	"fmt"

	"github.com/wesleysantana/adapter-design-pattern/internal/pdf/generator"
	report "github.com/wesleysantana/adapter-design-pattern/internal/sale/report"
)

func main() {
	pdfAdapter := generator.New("tmp")

	// Injeta o adaptador no gerador de relatórios
	reportService := report.New(pdfAdapter)

	// Executa
	fileName, err := reportService.Generate()
	if err != nil {
		fmt.Println("Erro ao gerar relatório:", err)
		return
	}

	fmt.Println("PDF gerado com sucesso em:", fileName)
}
