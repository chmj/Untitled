package main

/*
Concurrency Synchronization
Concurrent computations may share resources, generally memory resource.
The following are some circumstances that may occur during concurrent computing:

    In the same period that one computation is writing data to a memory segment,
	another computation is reading data from the same memory segment.
	Then the integrity of the data read by the other computation might be not preserved.
    In the same period that one computation is writing data to a memory segment,
	another computation is also writing data to the same memory segment.
	Then the integrity of the data stored at the memory segment might be not preserved.

These circumstances are called data races.
One of the duties in concurrent programming is to control resource sharing among concurrent computations,
so that data races will never happen.
The ways to implement this duty are called concurrency synchronizations,
or data synchronizations, which will be introduced one by one in later Go 101 articles.
Other duties in concurrent programming include

    determine how many computations are needed.
    determine when to start, block, unblock and end a computation.
    determine how to distribute workload among concurrent computations.

The program shown in the last section is not perfect.
The two new goroutines are intended to print ten greetings each.
However, the main goroutine will exit in two seconds, so many greetings don't have a chance to get printed.
How to let the main goroutine know when the two new goroutines have both finished their tasks?
We must use something called concurrency synchronization techniques.

Go supports several concurrency synchronization techniques.
Among them, the channel technique is the most unique and popularly used one.
However, for simplicity, here we will use another technique, the WaitGroup type in the sync standard package,
to synchronize the executions between the two new goroutines and the main goroutine.
The WaitGroup type has three methods (special functions, will be explained later):
Add, Done and Wait. This type will be explained in detail later in another article. Here we can simply think

    the Add method is used to register the number of new tasks.
    the Done method is used to notify that a task is finished.
    and the Wait method makes the caller goroutine become blocking until all registered tasks are finished.

Example:
*/
import (
	"log"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func SayGreetings(greeting string, times int) {
	for i := 0; i < times; i++ {
		log.Println(greeting)
		d := time.Second * time.Duration(rand.Intn(5)) / 2
		time.Sleep(d)
	}
	// Notify a task is finished.
	wg.Done() // <=> wg.Add(-1)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)
	wg.Add(2) // register two tasks.
	go SayGreetings("hi!", 10)
	go SayGreetings("hello!", 10)
	wg.Wait() // block until all tasks are finished.
}
