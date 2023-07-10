## Goroutines [Link](Goroutines/goroutines.go)
- Lighthweight thread managed by the GO runtime
  ```
  go (action)
  ```
- It starts a new thread with (action)
- The evaluation of the action happens in the current thread but the execution is on the new thread
- This threads run in the same address space, so the access to shared memory must be sync

## Channels [Link](Channels/channels.go)
- Channels are a way to sync goroutines
- You can send a function to a channel, assigning the result of that func to the channel
- Latter on you can use the value that the channel will eventually retain and assign that value from the channel to something
 ```
 ch := make(chan int)
 ch <- v
 v:= <-ch 
 ```
- Sends and receives a block until the other side is ready

## Buffered Channels [Link](BufferedChannels/buffered-channels.go)
- This is channel only receives values when it has space for it
- In case we put more than the number of goroutines we specified, it will give a error
- Case we call more than what we actually have, it will give a error, it is like a lock that only accepts the number of goroutines it has as limit
- It works like a stack
```
ch := make(chan int,1)
ch <- 1 //OK
```
```
ch := make(cha int,1)
ch <- 1
ch <- 2//NOK
```
```
ch := make(cha int,1)
ch <- 1
ch <- 2//NOK
```
- Note that when you assign the value, you are removing one goroutine from the buffer
```
ch := make(chan int, 2)
	ch <- 1
	fmt.Println(<-ch)
	ch <- 2
	ch <- 3 //OK
```

## Range and Close [Link](Range-and-close/range-and-close.go)
- A sender can close a channel to indicate that no more values will be sent
- Receivers can test whether the channel as been closed or not with a aditional variavel
```
v,ok := <-ch
//ok can either be false or true, case true then it is open, case false it is closed
```
- Only the sender should close a channel
- Sending on a closed channel will cause a panic
- Channel is not like a file, it does not need to be closed but closing it is a way to stop a action that is running such as a lop

## Select [Link](Select/select.go) [Link2](Select/select2.go)
- select statement lets a goroutine wait on multiple communication operations
- A select blocks until one of its cases can run, then it executes that case
- It is like a endpoint, where according to the receival it runs on behalf of that thread
  
## Default Select  [Link](DefaultSelection/default-selection.go)
- In case theres no receival of the channels we can use the default to do some action

##  Sync.Mutex [Link](SyncMutex/mutex-counter.go)
- Concept to avoid that multiple threads access the same variable
- It locks and unlocks 
- The concepts its called mutual exclusion
- Defer is used to after all the other methods run unlock the variable