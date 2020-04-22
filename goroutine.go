package utils

import "sync"

func WaitGroupFunc(fs ...func())  {
	wg := &sync.WaitGroup{}
	for _, f := range fs {
		wg.Add(1)
		go func(fc func()) {
			defer func() {
				recover()
				wg.Done()
			}()
			fc()
		}(f)
	}
	wg.Wait()
}