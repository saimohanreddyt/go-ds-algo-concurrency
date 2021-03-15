//concurrency is not same as parallel programming or parallelism
//concurrency is about design
/*
Design your program as a collection of independent processes
Design these processes to eventually run in parallel
Design your code so that the outcomes is always the same
*/

//Cocurrency in detail
/*group code (and data) by identifying independent tasks
no race conditions
no deadlocks
more workers = faster execution
*/

//Communicating Sequential Processes(CSP)
/*Each process is built for sequential execution
Data is communicated between processes via channels
No shared stste!

Scale by adding more of the same
*/

//Go's concurrency toolset
/*
go routines
channels
select
sync package
*/

//Channels
/*
Think of a bucket chain
3 components:sender, buffer, receiver
The buffer is optional
unbuffered := make(chan int)
1)blocks
a := <- unbuffered
2)blocks
unbuffered <- 1
3)synchronises
go fun() { <-unbuffered}()
unbuffered <-1
unbuffered := make(chan int, 1)
4)still blocks
a := <-buffered
5)fine
buffered <- 1
6)blocks(buffer full)
buffer <- 2
*/

//Blocking breaks concurrency
/*Remember ?
  . no deadlocks
  . more workers = faster execution
Blocking can lead to deadlocks
Blocking can prevent scaling
*/

//Closing channels
/*
Close sends a special ,,closed". Yay! nonthing to do
if you try to send more:panic!
c := make(chan int)
close(c)
fmt.Println(<-c)//receive ans print
*/

//Select
/*
Like a switch statement on channel operations
The order of case, too
The first non-blocking case is chosen(send and/ or recieve)
*/
package main

import "fmt"

func main() {
	fmt.Println("Concurrency")
}
