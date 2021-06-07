package misc

type PX struct {
	n     int
	errch chan error
}

func NewPx() (px *PX) {
	return &PX{
		n:     0,
		errch: make(chan error),
	}
}

func (px *PX) Run(fn func() error) {
	px.n++
	go func() {
		px.errch <- fn()
	}()
}

func (px *PX) Wait() (m int) {
	for i := 0; i < px.n; i++ {
		if err := <-px.errch; err != nil {
			m++
		}
	}

	return m
}
