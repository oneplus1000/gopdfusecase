package gopdfusecase

import (
	"net/http"

	"github.com/signintech/gopdf"
)

type CidHandler struct {
}

func (me *CidHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	gopdf.Ttf2CID("ก", "/home/oneplus/GOPATH/src/github.com/signintech/gopdf/fontmaker/ttf/Loma.ttf")
	w.Write([]byte("ok"))
}
