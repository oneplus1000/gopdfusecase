package gopdfusecase

import (
	"net/http"
	"github.com/signintech/gopdf"
	//"github.com/signintech/gopdf/fonts"
)

type ImgsPdfHandler struct{


}

func (me *ImgsPdfHandler) ServeHTTP(w http.ResponseWriter,r *http.Request){
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{Unit: "pt", PageSize: gopdf.Rect{W: 595.28, H: 841.89}}) //595.28, 841.89 = A4
	//pdf.AddFont("THSarabunPSK",new(fonts.THSarabun),Path_ResFont("THSarabun.z"))
	pdf.AddPage()

	x := 10
	y := 5

	for c := 0; c < x; c++ {
		pdf.Image("../www/imgs/im1.jpg", float64(10+(c*60)), pdf.GetY()+25, nil)
	}
	//pdf.Image("../www/imgs/im1.jpg", 10,10, nil)
	//pdf.SetFont("THSarabunPSK","B",14)
	//pdf.Cell(nil,   Encoder_Utf8ToCp874("Hello world  = สวัสดี โลก in thai"))
	//pdf.Image("../www/imgs/im2.jpg", 100,100, nil)

	for d := 0; d < y; d++ {
		pdf.Image("../www/imgs/im2.jpg", float64(10+(d*50)), pdf.GetY()+65, nil)
	}

	w.Header().Set("Content-type", "application/pdf")
	w.Write(pdf.GetBytesPdf())
}


