package gopdfusecase

import (
	"net/http"
	"text/template"
)


type IndexHandler struct{


}

func (me *IndexHandler) ServeHTTP(w http.ResponseWriter,r *http.Request){
	t, _ := template.ParseFiles(Path_Www("index.html"))
	t.Execute(w,nil)
}
