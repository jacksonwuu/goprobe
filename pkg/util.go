package pkg

import "time"

// Period call f periodic.
func Period(f func(), period time.Duration, stopCh <-chan struct{}, doAtStart bool) {
	if doAtStart {
		f()
	}
	ticker := time.NewTicker(period)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			f()
		case <-stopCh:
			return
		}
	}
}

// Retry try to call callable until it the callable return nil.
func Retry(retryMax int, sleep time.Duration, callable func(index int) error) error {
	var err error
	for i := 1; i <= retryMax; i++ {
		err = callable(i)
		if err == nil {
			return nil
		}
		if i == retryMax {
			return err
		}
		time.Sleep(sleep)
	}
	return err
}
