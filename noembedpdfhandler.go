package gopdfusecase

import (
	"net/http"

	"github.com/signintech/gopdf"
)

type NoembedPdfHandler struct {
}

func (me *NoembedPdfHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{Unit: "pt", PageSize: gopdf.Rect{W: 595.28, H: 841.89}}) //595.28, 841.89 = A4
	pdf.AddPage()
	pdf.SetFont("Loma", "B", 14)
	pdf.Cell(nil, Encoder_Utf8ToCp874("Hello world  = สวัสดี โลก in thai"))
	w.Header().Set("Content-type", "application/pdf")
	b := pdf.GetBytesPdf()
	w.Write(b)
}
