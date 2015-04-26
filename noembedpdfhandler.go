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
	err := pdf.AddFontUnicode("Loma", "/home/oneplus/GOPATH/src/github.com/signintech/gopdf/fontmaker/ttf/Loma.ttf")
	//err := pdf.AddFontUnicode("Refine", "/home/oneplus/GOPATH/src/github.com/signintech/gopdf/fontmaker/ttf/RefinementSuperSongStyleGBKFontSimplifiedChinese.TTF")
	if err != nil {
		log.Printf("ERROR:%s\n", err.Error())
	}
	pdf.SetFont("Loma", "B", 14)
	pdf.Cell(nil, "ก")
	w.Header().Set("Content-type", "application/pdf")
	b := pdf.GetBytesPdf()
	w.Write(b)
}
