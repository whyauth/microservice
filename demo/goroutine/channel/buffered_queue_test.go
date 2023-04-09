package channel

import "testing"

var c1 chan string

var c2 chan string

func init() {
	c1 = make(chan string, 100)

	go func() {
		for {
			<-c1
		}
	}()

	c2 = make(chan string, 100)
	go func() {
		for {
			c2 <- "hello"
		}
	}()


}

func send (msg string) {
	c1 <- msg
}

func recv () {
	<- c2
}

func BenchmarkBufferedChan1To1SendCap10(b *testing.B)  {
	for n := 0; n < b.N; n++ {
		send("hello")
	}
}

func BenchmarkBufferedChan1To1RecvCap10(b *testing.B) {
	for n :=0; n < b.N; n++ {
		recv()
	}
}


