# Test_Go_EC

## Fork this repo and show us your development progress via a Pull Request

## Implement the function executeParallel so that:

  1. Functions received as the second argument should be executed in parallel.
  
  2. As soon as a function finishes, its result should be written to the channel ch.
  
  3. After all functions finish, channel ch should be closed, and executeParallel should finish its execution.
  
 For example, executing main function should produce the following console output:

Result: 10000000
Result: 200000000

```golang
package main
import "fmt"
func executeParallel(ch chan<- int, functions ...func() int) {
}
func exampleFunction(counter int) int {
    sum := 0
    for i := 0; i < counter; i++ {
        sum += 1
    }
    return sum
}
func main() {
    expensiveFunction := func() int { return exampleFunction(200000000) }  
    cheapFunction := func() int { return exampleFunction(10000000) }
    ch := make(chan int)
    go executeParallel(ch, expensiveFunction, cheapFunction)
    for result := range ch {
        fmt.Printf("Result: %d\n", result)
    }
}
```


## Fill Channel

Implement the function fillChannel so that:
1.  It creates a channel of integers with a size that is the same as the number of functions received.
2.  Functions received as an argument should be executed in parallel.
3.  As soon as a function finishes, its result should be written to the channel.
4.  fillChannel should return the channel immediately, without waiting for functions to end.

For example, executing main function should produce the following console output:

```golang
Result: 10000000
Result: 200000000
```

```golang
package main
import "fmt"
func fillChannel(functions ...func() int) chan int {
    return nil
}
func exampleFunction(counter int) int {
    sum := 0
    for i := 0; i < counter; i++ {
        sum += 1
    }
    return sum
}
func main() {
    expensiveFunction := func() int { return exampleFunction(200000000) }  
    cheapFunction := func() int { return exampleFunction(10000000) }
    ch := fillChannel(expensiveFunction, cheapFunction)
    if ch != nil {
        for i := 0; i < 2; i++ {
            fmt.Printf("Result: %d\n", <-ch)
        }
    }
}
```


## Hobbies

Implement the _findAllHobbyists_ function that takes a hobby, and a map consisting of people's names mapped to their respective hobbies. The function should return a slice containing the names of the people (in any order) who enjoy the hobby.

For example, the following code should display the name 'Chad'.

```golang
hobbies := map[string][]string {
    "Steve": []string{"Fashion", "Piano", "Reading"},
    "Patty": []string{"Drama", "Magic", "Pets"},
    "Chad": []string{"Puzzles", "Pets", "Yoga"},
}
    
fmt.Println(findAllHobbyists("Yoga", hobbies))
```

```golang
package main
import "fmt"
func findAllHobbyists(hobby string, hobbies map[string][]string) []string {
    var hobbyists []string
    return hobbyists
}
func main() {
    hobbies := map[string][]string {
        "Steve": []string{"Fashion", "Piano", "Reading"},
        "Patty": []string{"Drama", "Magic", "Pets"},
        "Chad": []string{"Puzzles", "Pets", "Yoga"},
    }
    fmt.Println(findAllHobbyists("Yoga", hobbies))
}
```


## Max Sum

Implement the _findMaxSum_ function that efficiently, with respect to time used, returns the largest sum of any two elements in the given slice of positive numbers.

For example, the largest sum of the slice [5, 9, 7, 11] is the sum of the elements 9 and 11, which is 20.


```golang
package main
import "fmt"

func findMaxSum(numbers []int) int {
    return 0
}
    
func main() {
    list := []int { 5, 9, 7, 11 }
    fmt.Println(findMaxSum(list))
}
```


## Fire Dragon

Scientists have discovered a species of fire-breathing dragons. DNA analysis of the dragon reveals that it is a reptile evolved from a common ancestor of crocodile, hundreds of millions of years ago. Even though they're related, the different reptile species cannot cross-breed.

Researchers would like to develop a lifecycle model of this rare species, in order to better study them. Complete the implementation below so that:
1.  The FireDragon species implements the Reptile interface.
2.	When a ReptileEgg hatches, a new reptile will be created of the same species that laid the egg.
3.	nil is returned if a ReptileEgg tries to hatch more than once.

```golang
package main
//import "fmt"

type Reptile interface {
    Lay() ReptileEgg
}

type ReptileCreator func() Reptile

type ReptileEgg struct {
    CreateReptile ReptileCreator        
}

func(egg *ReptileEgg) Hatch() Reptile {
    return nil
}

type FireDragon struct {
}
```

