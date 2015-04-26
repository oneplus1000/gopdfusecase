package gopdfusecase

import (
	"fmt"
	"net/http"
)

type Webserver struct {
}

func (me *Webserver) Start() {

	http.Handle("/", new(IndexHandler))
	http.Handle("/helloworldpdf", new(HelloworldPdfHandler))
	http.Handle("/underlinepdf", new(UnderlinePdfHandler))
	http.Handle("/imgspdf", new(ImgsPdfHandler))
	http.Handle("/noembedpdf", new(NoembedPdfHandler))
	http.Handle("/cid", new(CidHandler))
	fmt.Printf("http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}
