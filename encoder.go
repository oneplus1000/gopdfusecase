package gopdfusecase

import (
	iconv "github.com/djimenez/iconv-go"
)

func Encoder_Utf8ToCp874(str string) string{
	str, _ = iconv.ConvertString( str, "utf-8", "cp874")
	return  str
}
