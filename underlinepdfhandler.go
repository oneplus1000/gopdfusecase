package gopdfusecase

import (
	"net/http"
	"github.com/signintech/gopdf"
	"github.com/signintech/gopdf/fonts"
)

type UnderlinePdfHandler struct{

}


func (me *UnderlinePdfHandler) ServeHTTP(w http.ResponseWriter,r *http.Request){
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{Unit: "pt", PageSize: gopdf.Rect{W: 595.28, H: 841.89}}) //595.28, 841.89 = A4
	pdf.AddFont("THSarabunPSK",new(fonts.THSarabun),Path_ResFont("THSarabun.z"))
	pdf.AddFont("Loma",new(fonts.Loma),Path_ResFont("Loma.z"))
	pdf.AddPage()
	pdf.SetFont("THSarabunPSK","U",44)
	pdf.Cell(nil,   Encoder_Utf8ToCp874("Hello world  = สวัสดี โลก in thai"))
	pdf.Br(44)
	pdf.SetFont("THSarabunPSK","U",14)
	pdf.Cell(nil,   Encoder_Utf8ToCp874("Hello world  = สวัสดี โลก in thai"))
	pdf.Cell(nil,   Encoder_Utf8ToCp874("เดี๋ยวมันก็ผ่านพ้น เดี๋ยวไม่นานก็จาง เดี๋ยวมันก็มีทางไป"))
	pdf.Br(14)
	pdf.Cell(nil,   Encoder_Utf8ToCp874("ทดสอบๆๆ การบ้านกั่ง"))
	w.Header().Set("Content-type", "application/pdf")
	w.Write(pdf.GetBytesPdf())
}
