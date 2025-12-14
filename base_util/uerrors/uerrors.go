package uerrors

type VideoTranslateErr struct {
	Msg     string
	RetCode int
}

func (e *VideoTranslateErr) Error() string {
	return e.Msg
}
func NewVideoTranslateErr(msg string, retCode int) *VideoTranslateErr {
	return &VideoTranslateErr{
		Msg:     msg,
		RetCode: retCode,
	}
}
