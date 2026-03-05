package utils

type IDGenerator <-chan uint

func NewIDGenerator() IDGenerator {
	c := make(chan uint)
	go func() {
		var i uint
		for {
			c <- i
			i++
		}
	}()
	return c
}
