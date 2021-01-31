package misc

type ErrContainer struct {
	value  interface{}
	errmsg string
	level  int
}

func NewErrContainer(value interface{}, errmsg string) (ec *ErrContainer) {
	return &ErrContainer{value: value, errmsg: errmsg}
}

func (ec *ErrContainer) SetLevel(level int) *ErrContainer {
	ec.level = level
	return ec
}

func (ec *ErrContainer) Error() string {
	return ec.errmsg
}

func (ec *ErrContainer) Value() (value interface{}) {
	return ec.value
}

func (ec *ErrContainer) Level() (level int) {
	return ec.level
}
