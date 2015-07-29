package gopdfusecase

import (
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

	var err error
	//err = pdf.AddTTFFont("HDZB_5", "/home/oneplus/GOPATH/src/github.com/signintech/gopdf/fontmaker/ttf/aa.ttf")
	err = pdf.AddTTFFont("HDZB_5", "/home/oneplus/GOPATH/src/github.com/signintech/gopdf/fontmaker/ttf/wts11.ttf")
	if err != nil {
		log.Printf("ERROR:%s\n", err.Error())
		return
	}

	err = pdf.AddTTFFont("Loma", "/home/oneplus/GOPATH/src/github.com/signintech/gopdf/fontmaker/ttf/Loma.ttf")
	if err != nil {
		log.Printf("ERROR:%s\n", err.Error())
		return
	}

	//BokutachinoGothic2Regular.ttf
	err = pdf.AddTTFFont("BokutachinoGothic2Regular", "/home/oneplus/GOPATH/src/github.com/signintech/gopdf/fontmaker/ttf/BokutachinoGothic2Regular.ttf")
	if err != nil {
		log.Printf("ERROR:%s\n", err.Error())
		return
	}

	//Thai
	err = pdf.SetFont("Loma", "B", 14)
	if err != nil {
		log.Printf("ERROR:%s\n", err.Error())
		return
	}
	pdf.Cell(nil, "Hello")
	pdf.Br(20)
	pdf.Cell(nil, "สวัสดี")
	pdf.Br(20)

	//China
	err = pdf.SetFont("HDZB_5", "B", 14)
	if err != nil {
		log.Printf("ERROR:%s\n", err.Error())
		return
	}
	pdf.Cell(nil, "您好")
	pdf.Br(20)

	//Japan
	err = pdf.SetFont("BokutachinoGothic2Regular", "B", 14)
	if err != nil {
		log.Printf("ERROR:%s\n", err.Error())
		return
	}
	pdf.Cell(nil, "こんにちは")
	pdf.Br(20)

	w.Header().Set("Content-type", "application/pdf")
	b := pdf.GetBytesPdf()
	//fmt.Printf("byte=%s\n\n", string(b))
	w.Write(b)
}
