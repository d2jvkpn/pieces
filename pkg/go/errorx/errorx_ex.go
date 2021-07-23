package errorx

func ToErrorX(err error) ErrorX {
	var (
		ok bool
		p  *ErrorX
	)

	if p, ok = err.(*ErrorX); !ok {
		return ErrUnexpected()
	}

	return *p
}

func (ex ErrorX) ToResData(mp map[string]interface{}) (rd ResData) {
	data := make(map[string]interface{}, len(mp))

	for k := range mp {
		data[k] = mp[k]
	}

	return ResData{
		Code:    ex.Code,
		Message: ex.Message,
		Data:    data,
	}
}
