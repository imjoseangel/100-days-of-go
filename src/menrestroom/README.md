# The Urinal Algorithm. How I improved my Python and Golang skills with this funny Algorithm

The first time I read about the Urinal Algorithm was in a [Reddit post](https://www.reddit.com/r/learnpython/comments/b7kq94/men_restroom_algorithm/). In fact, the author called it the **Men Restroom Algorithm** so I named my _Urinal Based Algorithm_ as _MenRestRoom_. I started to develop my own algorithm based on that post and although it is unfinished, I use it to improve my programming skills.

At the time of writing this article, I realized that the Algorithm with its original name, had some post in [StackOverflow](https://stackoverflow.com/questions/32645046/urinal-algorithm-a-simple-optimization) even with an [Unisex](https://stackoverflow.com/questions/39661826/unisex-bathroom-algorithm-with-priority) implementation with Priority. It also has a cool document from 2010 called the [Urinal Problem](https://people.scs.carleton.ca/~kranakis/Papers/urinal.pdf) which I promise I will review in the future to find ideas to improve mine.

## The Algorithm

The concept is simple: It is a well-researched fact that men in a restroom generally prefer to maximize their distance from already occupied stalls, by occupying the middle of the longest sequence of unoccupied places. For example, consider the situation where ten stalls are empty.

The first visitor will occupy a middle position:

🚽 🚽 🚽 🚽 🚶 🚽 🚽 🚽 🚽 🚽 🚪

In my variant, there is a door, so the left side is clearly the farther stalls from the door. Also I have improved it with emojis. Cool, isn't it?.

The next visitor will use the middle or the farther stall from empty area at the left. Always keeping at least one stall empty between visitors.

🚶 🚽 🚽 🚽 🚶 🚽 🚽 🚽 🚽 🚽 🚪

or

🚽 🚽 🚶 🚽 🚶 🚽 🚽 🚽 🚽 🚽 🚪

The initial sequence can be like:

🚽 🚽 🚽 🚽 🚶 🚽 🚽 🚽 🚽 🚽 🚪

🚶 🚽 🚽 🚽 🚶 🚽 🚽 🚽 🚽 🚽 🚪

🚶 🚽 🚶 🚽 🚶 🚽 🚽 🚽 🚽 🚽 🚪

🚶 🚽 🚶 🚽 🚶 🚽 🚽 🚽 🚶 🚽 🚪

🚶 🚽 🚶 🚽 🚶 🚽 🚶 🚽 🚶 🚽 🚪

Once every other stall is occupied, we can randomly use the rest of the stalls until finish:

🚶 🚽 🚶 🚽 🚶 🚽 🚶 🚶 🚶 🚽 🚪

🚶 🚶 🚶 🚽 🚶 🚽 🚶 🚶 🚶 🚽 🚪

🚶 🚶 🚶 🚽 🚶 🚶 🚶 🚶 🚶 🚽 🚪

🚶 🚶 🚶 🚽 🚶 🚶 🚶 🚶 🚶 🚶 🚪

🚶 🚶 🚶 🚶 🚶 🚶 🚶 🚶 🚶 🚶 🚪

## More Fun added

Well, this is how the stalls become occupied but in real life visitors use them for a limited time. Also, the time of peeing is different from visitor to visitor. In the other hand, we assume that the time to take a stall is fixed (By default *2 seconds*). The code needs to manage both *take* and *leave* functionality at the same time.

As said, the code needs to be improved. When some consecutive stalls gets empty, the algorithm continues to use a random value instead of maximizing the distance. I promise I will work on this option in the future.

>**Note:** I will work on adding some code tests for Python too. I have added some interesting ones for the Golang functions.

## Implementation

I'm going to explain What and How I learned implementing this algorithm in Python and how I used it to learn some Golang techniques.

### Variables

Python and Golang have different variable scopes (Local and Global) but the way to define and use them is harder in Golang. Constant and variables types need to be properly defined and due to the fact that it is a statically typed language, their variables cannot be changed at runtime as with Python.

Additionally, Go is not a pure object oriented programming language. The way to manage functions and how to call them, changes significantly if you are used to program in Python.

Let's define our constants and variables:

#### Constants

| Description                                 | Constant      | Type          |
|---------------------------------------------|---------------|---------------|
| Maximum time a visitor is occupying a stall | maxtimepeeing | Int           |
| Minimum time a visitor is occupying a stall | mintimepeeing | Int           |
| Number of stalls                            | stalls        | Int           |
| Emoji for door                              | emoDoor       | string        |
| Emoji for empty stall                       | emoEmpty      | string        |
| Emoji for taken stall                       | emoTaken      | string        |
| Stall occupancy frequency                   | stallfreq     | time.Duration |

#### Vars

| Description                         | Variable   | Type             |
|-------------------------------------|------------|------------------|
| Untaken stalls                      | untaken    | List of integers |
| Taken stalls                        | taken      | List of integers |
| time a visitor is occupying a stall | timePeeing | time.Duration    |
| Stall occupied on every iteration   | stall      | Int              |
| Left side of the stalls             | left       | List of integers |
| Right side of the stalls            | right      | List of integers |
| Shows stall status on Screen        | stallPrint | List of strings  |

### The code explained

#### Generate a list of integers

First of all, we need to initialize all our variables with the startup values. We need a list of *n untaken* stalls.

**Python:**

```py
untaken = list(range(0, stalls))
```

If `stalls = 4`, it creates the list: [1, 2, 3, 4]

**Go:**

In go, there is no a built-in range function, so we need to create it first:

```go
type Range struct {
    MinList int
    MaxList int
}

// RangeArray ...
// Created this way to use struct with a function inside a module
func (arrayrange Range) RangeArray() []int {

    result := make([]int, arrayrange.MaxList-arrayrange.MinList+1)
    for Item := range result {
        result[Item] = arrayrange.MinList + Item
    }

    return result
}

untaken = functions.Range{
    MinList: 1,
    MaxList: stalls,
}.RangeArray()
```

I created this way to learn about *structs* and Object Oriented Programming (OOP) within Go.

#### Generate a random number from a given range

Let's go with the `timepeeing` var. This var will be created at init time (Another important topic that we will review later) and reset with a new value when calling the leaveStall function. It contains a random integer between `mintimepeeing` and `maxtimepeeing`.

**Python:**

```py
timepeeing = random.randint(mintimepeeing, maxtimepeeing)
```

**Golang:**

To generate a random value between two values in go, use this:

```go
rand.Seed(time.Now().UnixNano())
timePeeing = time.Duration(rand.Intn(maxtimepeeing-mintimepeeing+1) + mintimepeeing)
```

#### Calculate the stall to be taken on every iteration

In the first iteration, the taken stall is the one in the middle of the untaken row. It is calculated with the sum of the integers of the array and divided by its length. The result is rounded to ceil.

**Python:**

```python
new_stall = round(sum(untaken) / len(untaken) + .5)
```

**Golang:**

There is no built-in implementation for the sum of the items of a list. I had to implement too.

```go
// Array ...
type Array []int

// SumArray ...
// Created this way to use struct with a function inside a module
func (array Array) SumArray() int {
    result := 0
    for _, numb := range array {
        result += numb
    }
    return result
}

newStall = int(math.Ceil(float64(functions.Array(untaken).SumArray()) / float64(len(untaken))))
```

And again, I created this way to learn about OOP within Go. Cool, isn't it?

#### The Left and Right arrays with alternative stalls

The idea is divide the stalls in two (Left and Right) given the already occupied newStall. Having this untaken list:

`[0, 1, 2, 3, 4, 5, 6, 7, 8, 9]`

If `newStall` is *5* in this scenario, the *left* list is `[0, 1, 2, 3, 4]` and the right list `[6, 7, 8, 9]`. Integers *4* and *6* are next to `newStall` and must be empty. Following the same rule, we will build a list of potential stalls to be taken:

For the left side:

`[1, 3]`

For the right side:

`[7, 9]`

Let's do it. In Python, is quite easy:

**Python:**

```python
left = untaken[1:new_stall:2]
right = untaken[new_stall + 2::2]
```

Again, in golang I had to create a new function:

**Golang:**

```go
// SliceArray ...
func SliceArray(array []int, start int) []int {

    result := []int{}
    for i := start; i < len(array); i += 2 {
        result = append(result, array[i])
    }
    return result
}

if stalls%2 == 0 {
    left = functions.SliceArray(untaken[0:newStall-1], 1)
} else {
    left = functions.SliceArray(untaken[0:newStall-1], 0)
}
right = functions.SliceArray(untaken[newStall:], 1)
```

With this function, we need to check if the `stalls` integer is even or odd to create the `left` list from `untaken` list. The start integer changes depending on the length of the list, something I need to improve in future releases.

#### Printing on screen

This is the easier one for the first stage. We need to prepare the empty stalls with the door and show it on screen:

**Python:**

```python
stall_print = list(emo_empty * stalls + emo_door)
```

**Golang:**

```go
stallPrint = strings.SplitN(strings.Repeat(emoEmpty, stalls)+emoDoor, "", stalls+1)
```

#### Init Functions and Methods

Before carry on with the main logic, I think it is interesting to refer about the way Python and Golang manage the Init functions.

The `__init__` in **Python** is one of the main pieces when talking about OOP. This method initializes the object state. It contains the values and instructions executed at the time of the Object Creation.

The `init()` function in **Golang** are defined in the package block. It contains the values and instructions executed at the time of the Package call although it runs only once even if the package is imported many times.

#### Variable name convention

I think it is also nice to realize that the name convention changes among programming languages. I recommend to search for the different naming conventions for Python and Golang on the Internet.

For instance:

`stall_print` for Python
`stallPrint` for Go

#### The Take Stall logic

As explained in the ![The Algorithm](#the-algorithm) section, we need to move items between the `untaken` and `taken` lists and vice versa.

When the process starts, it puts the middle of the stalls in the `taken` and paints the new scenario. After that, it checks if the `left` list has elements to use them. Then the `right` list and finally the ones unoccupied.

The code:

**Python:**

```python
        if untaken:
            if not taken:
                new_stall = new_stall
            else:
                if left:
                    new_stall = random.choice(left)
                    left.remove(new_stall)
                elif right:
                    new_stall = random.choice(right)
                    right.remove(new_stall)
                else:
                    new_stall = random.choice(untaken)

            untaken.remove(new_stall)
            taken.append(new_stall)
```

**Golang:**

To get the position of the stall, I created the IndexArray function, also take a look to the append built-in function. The [documentation](https://golang.org/pkg/builtin/#append) of the built-in package describes append.

`func append(s []T, vs ...T) []T`

```go
// IndexArray ...
func IndexArray(array []int, item int) int {
    for result := range array {
        if array[result] == item {
            return result
        }
    }
    return -1
}

if len(untaken) > 0 {
    if len(taken) == 0 {
        stall = newStall
    } else {
        if len(left) > 0 {
            randomIndex := rand.Intn(len(left))
            stall = left[randomIndex]
            left = append(left[:randomIndex], left[randomIndex+1:]...)
        } else if len(right) > 0 {
            randomIndex := rand.Intn(len(right))
            stall = right[randomIndex]
            right = append(right[:randomIndex], right[randomIndex+1:]...)

        } else {
            randomIndex := rand.Intn(len(untaken))
            stall = untaken[randomIndex]
        }
    }
    stallIndex := functions.IndexArray(untaken, stall)
    untaken = append(untaken[:stallIndex], untaken[stallIndex+1:]...)
}

```

##### Create the stall for printing

Once the logic
