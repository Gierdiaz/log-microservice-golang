package scheduler

import "time"

// Every executa uma função a cada intervalo de tempo definido
func Every(internal time.Duration, task func()) {
	go func() {
		for {
			task()
			time.Sleep(internal)
		}
	}()
}