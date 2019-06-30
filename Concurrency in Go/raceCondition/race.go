package raceCondition

import (
	"fmt"
	"sync"
)

var Wait sync.WaitGroup
var Counter = 0

func main() {
	for routine := 1; routine <= 2; routine++ {
		Wait.Add(1)
		go RoutineOne(routine)
	}

	Wait.Wait()
	fmt.Printf("Final Counter: %d\n", Counter)
}

func RoutineOne(id int) {
	for count := 0; count < 2; count++ {
		value := Counter
		//time.Sleep(1 * time.Nanosecond) // Uncomment this line for second test
		value++
		Counter = value
	}
	Wait.Done()
}

/* Explanation
------------------------------------------------------------------------------------------------------
Race Condition
A data race happens when two goroutines access the same variable concurrently, and at least one of the access is a write instruction.

Detecting a race
Use the command
	go run -race routineOne.go to generate the race report
	uncomment line 28
	go run -race routineOne.go to generate the race report


Conclusion
•	Result one and two generates 2 different Final Counter after adding wait period
•	Adding the pause caused the result to fail (should be 4)
•	-race parameter generates a report to find data races.
•	The report describes when the write value happened

The example provided commenting/uncommenting line 28 shows the difference of context switching

Without pause
-------------
Without the pause, the final Counter value = 4
This is because, before the second routine is executed (Line 13), routine 1 loop has already completed
writing the value to Counter.

Hence;
Loop 1 -> Increment -> Context Switch -> Loop 2

With pause
----------
With the pause, the context switch happens before routine 1 loop is finished.
Therefore resulting to final counter value = 2

Hence
Loop 1 -> Context Switch -> Loop 2
Resulting the second loop to start count from 0, instead of 2
*/
