package rover

import (
	"fmt"
	// "log"
	"sync"
	"time"
)

type UpdateTrigger struct {
	m        sync.Mutex    // locker
	d        time.Duration // update time interval
	l        time.Time     // last updated at
	mustSync bool          // block other UpdateTrigger when executing fn
	fn       func() error  // function to execute
	err      error         // last updated error
}

func NewUpdateTrigger(fn func() error, d time.Duration, mustSync, atOnce bool) (
	ut *UpdateTrigger, err error) {

	if fn == nil || d <= 0 {
		return nil, fmt.Errorf("invalid input")
	}

	ut = &UpdateTrigger{d: d, fn: fn}
	if atOnce {
		ut.err, ut.l = fn(), time.Now()
	}

	return
}

func (ut *UpdateTrigger) Update() (err error) {
	if !ut.Expired() {
		return
	}

	ut.m.Lock()
	if !ut.Expired() { // check again, may other has updated
		// if !lastUpdateIsOK || (lastUpdateIsOK && ut.err == nil) {
		//	return nil
		// }
		ut.m.Unlock()
		return
	}

	if !ut.mustSync { // free others' update trigger
		ut.l = time.Now()
		ut.m.Unlock()
	}

	ut.err = ut.fn()
	if ut.mustSync {
		ut.l = time.Now()
		ut.m.Unlock()
	}
	return ut.err
}

func (ut *UpdateTrigger) Expired() bool {
	return ut.l.Add(ut.d).After(time.Now())
}

func (ut *UpdateTrigger) LastOne() (time.Time, error) {
	return ut.l, ut.err
}

func (ut *UpdateTrigger) ResetInterval(d time.Duration) error {
	if d <= 0 {
		return fmt.Errorf("invalid input")
	}

	ut.m.Lock()
	ut.d = d
	ut.m.Unlock()
	return nil
}
