package gopdfusecase

import (
	"net/http"

	"github.com/signintech/gopdf"
)

//test cid
type CidHandler struct {
}

//x
func (c *CidHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var subsetObj gopdf.SubsetFontObj
	subsetObj.AddChars("ก")
	subsetObj.Build()
}
