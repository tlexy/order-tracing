package uerrors

import "strconv"

type VideoTranslateErr struct {
	Msg     string
	RetCode int
}

func (e *VideoTranslateErr) Error() string {
	return e.Msg + " | ret_code: " + strconv.FormatInt(int64(e.RetCode), 10)
}
func NewVideoTranslateErr(msg string, retCode int) *VideoTranslateErr {
	return &VideoTranslateErr{
		Msg:     msg,
		RetCode: retCode,
	}
}
