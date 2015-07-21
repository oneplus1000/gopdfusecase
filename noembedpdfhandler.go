package gopdfusecase

import (
	"fmt"
	"log"
	"net/http"

	"github.com/signintech/gopdf"
)

type NoembedPdfHandler struct {
}

func (me *NoembedPdfHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{Unit: "pt", PageSize: gopdf.Rect{W: 595.28, H: 841.89}}) //595.28, 841.89 = A4
	pdf.AddPage()
	err := pdf.AddTTFFont("Loma", "/home/oneplus/GOPATH/src/github.com/signintech/gopdf/fontmaker/ttf/Loma.ttf")
	if err != nil {
		log.Printf("ERROR:%s\n", err.Error())
	}
	err = pdf.SetFont("Loma", "B", 14)
	if err != nil {
		log.Printf("ERROR:%s\n", err.Error())
	}
	//log.Printf("ERROR")
	pdf.Cell(nil, "กข")
	pdf.Cell(nil, "คง")
	w.Header().Set("Content-type", "application/pdf")
	b := pdf.GetBytesPdf()
	fmt.Printf("byte=%s\n\n", string(b))
	w.Write(b)
}
