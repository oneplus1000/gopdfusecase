package gopdfusecase

import "net/http"

type Webserver struct{


}

func (me * Webserver) Start(){

	http.Handle("/",new(IndexHandler))
	http.Handle("/helloworldpdf",new(HelloworldPdfHandler))
	http.ListenAndServe(":8083",nil)
}





