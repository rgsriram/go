package main

import (
	"fmt"
)

func increment(i *int) {
	*i++;
}

func increment2(i * int) {
	*i++;
}

func printNum() int {
	var i int
	go increment(&i)
	go increment2(&i)
	return i
}

func main() {
	fmt.Println(printNum())
}

/*

Race Condition:
---------------

When multiple threads trying to access same data at the same time which cause data to be changed or results into unexpected outcome, then race condition occurs

Example:
-------

In this example, I have created variable 'i', and I am calling 2 function with reference of 'i' to incrementing it.

Without go rountines
---------------------
Ideally, without go routines this program should print 2

➜  Programs git:(master) ✗ go run race_example.go
2


With go rountines
-----------------
Execution flow of the above program is,

1. Initialize 'i' (by default 0)
2. Execute go routine 1
3. Execute go routine 2
4. Read value of 'i'

Here we don't which operation will complete first and it output is depends on that.
As you can imagine, its horrible having to test and use code which acts differently every single time you call it, and this is why data races pose such a huge problem.  


Detect race condition occured:
------------------------------
We can run a program with '-race' to detect any race condiiton happened in the program. A you see below there are 2 race conditions happened.

➜  Programs git:(master) ✗ go run -race race_example.go
==================
WARNING: DATA RACE
Read at 0x00c4200a0008 by main goroutine:
  main.printNum()
      /Users/sriramganesh/Study/Go/Exercise/Programs/race_example.go:19 +0xaa
  main.main()
      /Users/sriramganesh/Study/Go/Exercise/Programs/race_example.go:23 +0x33

Previous write at 0x00c4200a0008 by goroutine 6:
  main.increment()
      /Users/sriramganesh/Study/Go/Exercise/Programs/race_example.go:8 +0x4e

Goroutine 6 (finished) created at:
  main.printNum()
      /Users/sriramganesh/Study/Go/Exercise/Programs/race_example.go:17 +0x7a
  main.main()
      /Users/sriramganesh/Study/Go/Exercise/Programs/race_example.go:23 +0x33
==================
==================
WARNING: DATA RACE
Read at 0x00c4200a0008 by goroutine 7:
  main.increment2()
      /Users/sriramganesh/Study/Go/Exercise/Programs/race_example.go:12 +0x38

Previous write at 0x00c4200a0008 by goroutine 6:
  main.increment()
      /Users/sriramganesh/Study/Go/Exercise/Programs/race_example.go:8 +0x4e

Goroutine 7 (running) created at:
  main.printNum()
      /Users/sriramganesh/Study/Go/Exercise/Programs/race_example.go:18 +0x9c
  main.main()
      /Users/sriramganesh/Study/Go/Exercise/Programs/race_example.go:23 +0x33

Goroutine 6 (finished) created at:
  main.printNum()
      /Users/sriramganesh/Study/Go/Exercise/Programs/race_example.go:17 +0x7a
  main.main()
      /Users/sriramganesh/Study/Go/Exercise/Programs/race_example.go:23 +0x33
==================
1
Found 2 data race(s)
exit status 66



*/