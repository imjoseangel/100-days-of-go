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

Let's define our variables:

#### Constants

| Description                                 | Variable      | Type          |
|---------------------------------------------|---------------|---------------|
| Maximum Time a visitor is occupying a stall | maxtimepeeing | Int           |
| Minimum Time a visitor is occupying a stall | mintimepeeing | Int           |
| Number of Stalls                            | stalls        | Int           |
| Emoji for Door                              | emoDoor       | string        |
| Emoji for Empty stall                       | emoEmpty      | string        |
| Emoji for Taken stall                       | emoTaken      | string        |
| Stall Occupancy Frequency                   | stallfreq     | time.Duration |
