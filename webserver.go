package gopdfusecase

import (
	"net/http"
)

type Webserver struct{


}

func (me * Webserver) Start(){

	http.Handle("/",new(IndexHandler))
	http.Handle("/helloworldpdf",new(HelloworldPdfHandler))
	http.Handle("/underlinepdf",new(UnderlinePdfHandler))
	http.Handle("/imgspdf",new(ImgsPdfHandler))
	http.ListenAndServe(":8080",nil)
	fmt.Printf("http://localhost:8080")
}





