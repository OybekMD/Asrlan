package certificate

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/jung-kurt/gofpdf"
	"github.com/skip2/go-qrcode"
)

func generateQRCode(content, filename string) error {
	// Generate QR code and encode it to a PNG file
	err := qrcode.WriteFile(content, qrcode.Medium, 256, filename)
	if err != nil {
		return err
	}
	return nil
}

func GenerateCertificate(name, level, language string) (string, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	outputPDF := "./media/certificate/" + id.String() + ".pdf"
	qrContent := fmt.Sprintf("https://asrlan.jprq.app/media/certificate/%s.pdf", id.String())

	// Generate QR Code and save it to a file
	qrImagePath := "qrcode.png"
	err = generateQRCode(qrContent, qrImagePath)
	if err != nil {
		log.Println("Error generating QR Code: %v", err)
		return "", err
	}
	defer os.Remove(qrImagePath) // Clean up QR code image file after use

	// Create PDF certificate with background template
	templatePath := "./api/helper/certificate/template.png"
	err = createCertificate(outputPDF, name, level, language, qrImagePath, templatePath)
	if err != nil {
		log.Println("Error creating certificate: %v", err)
		return "", err
	}

	fmt.Println("Certificate created successfully!")
	return qrContent, nil
}

func createCertificate(outputPath, name, level, language, qrImagePath, templatePath string) error {
	// Create a new PDF document
	pdf := gofpdf.New("L", "mm", "A4", "")
	pdf.SetMargins(10, 10, 10)
	pdf.AddPage()

	// Load template PDF as background image
	templateWidth, templateHeight := 297.0, 210.0 // Adjust based on your template PDF size
	pdf.ImageOptions(templatePath, 0, 0, templateWidth, templateHeight, false, gofpdf.ImageOptions{}, 0, "")

	// Set font for certificate title
	pdf.SetFont("Arial", "B", 36)

	// Add certificate title (adjust position as needed)
	pdf.SetTextColor(29, 87, 252)
	pdf.CellFormat(0, 135, name, "", 1, "C", false, 0, "")

	// Set font for recipient's name
	pdf.SetFont("Arial", "", 32)

	// Add recipient's name (adjust position as needed)
	pdf.SetTextColor(94, 178, 255)
	pdf.CellFormat(0, -105, "has successfully completed the", "", 1, "C", false, 0, "")
	pdf.SetFont("Arial", "B", 32)
	pdf.CellFormat(0, 130, fmt.Sprintf("%s %s", language, level), "", 1, "C", false, 0, "")
	pdf.SetFont("Arial", "", 32)
	pdf.CellFormat(0, -105, "level course conducted by Asrlan.com", "", 1, "C", false, 0, "")

	pdf.SetTextColor(0, 0, 0)
	pdf.SetFont("Arial", "", 16)
	currentTime := formatCurrentDate()

	// Add QR Code image (adjust position and size as needed)
	pdf.Image(qrImagePath, 213, 154, 45, 45, false, "", 0, "")
	pdf.CellFormat(0, 124, fmt.Sprintf("This certificate is awarded on %s", currentTime), "", 1, "C", false, 0, "")

	// Save the output PDF
	return pdf.OutputFileAndClose(outputPath)
}

func formatCurrentDate() string {
	// Get the current time
	currentTime := time.Now()

	// Format the current date in the desired layout
	formattedDate := currentTime.Format("02.01.2006") // "02" for day, "01" for month, "2006" for year
	return formattedDate
}
