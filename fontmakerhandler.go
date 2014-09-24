package gopdfusecase

import (
	"github.com/signintech/gopdf"
	"github.com/signintech/gopdf/fonts"
	"net/http"
)

type FontMakerHandler struct {
}

func (me *FontMakerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{Unit: "pt", PageSize: gopdf.Rect{W: 595.28, H: 841.89}}) //595.28, 841.89 = A4
	pdf.AddFont("THSarabunPSK", new(fonts.THSarabun), Path_ResFont("THSarabun.z"))
	pdf.AddFont("Loma", new(fonts.Loma), Path_ResFont("Loma.z"))
	pdf.AddPage()
	pdf.SetFont("THSarabunPSK", "B", 14)
	pdf.Cell(nil, Encoder_Utf8ToCp874("Hello world  = สวัสดี โลก in thai"))
	w.Header().Set("Content-type", "application/pdf")
	w.Write(pdf.GetBytesPdf())
}
