## --- Day 15: Rambunctious Recitation ---

You catch the airport shuttle and try to book a new flight to your vacation island. Due to the storm, all direct flights have been cancelled, but a route is available to get around the storm. You take it.

While you wait for your flight, you decide to check in with the Elves back at the North Pole. They're playing a __memory game__ and are ever so excited to explain the rules!

In this game, the players take turns saying __numbers__. They begin by taking turns reading from a list of __starting numbers__ (your puzzle input). Then, each turn consists of considering the __most recently spoken number__:

*   If that was the __first__ time the number has been spoken, the current player says __`0`__.
*   Otherwise, the number had been spoken before; the current player announces __how many turns apart__ the number is from when it was previously spoken.

So, after the starting numbers, each turn results in that player speaking aloud either __`0`__ (if the last number is new) or an __age__ (if the last number is a repeat).

For example, suppose the starting numbers are `0,3,6`:

*   __Turn 1__: The `1`st number spoken is a starting number, __`0`__.
*   __Turn 2__: The `2`nd number spoken is a starting number, __`3`__.
*   __Turn 3__: The `3`rd number spoken is a starting number, __`6`__.
*   __Turn 4__: Now, consider the last number spoken, `6`. Since that was the first time the number had been spoken, the `4`th number spoken is __`0`__.
*   __Turn 5__: Next, again consider the last number spoken, `0`. Since it __had__ been spoken before, the next number to speak is the difference between the turn number when it was last spoken (the previous turn, `4`) and the turn number of the time it was most recently spoken before then (turn `1`). Thus, the `5`th number spoken is `4 - 1`, __`3`__.
*   __Turn 6__: The last number spoken, `3` had also been spoken before, most recently on turns `5` and `2`. So, the `6`th number spoken is `5 - 2`, __`3`__.
*   __Turn 7__: Since `3` was just spoken twice in a row, and the last two turns are `1` turn apart, the `7`th number spoken is __`1`__.
*   __Turn 8__: Since `1` is new, the `8`th number spoken is __`0`__.
*   __Turn 9__: `0` was last spoken on turns `8` and `4`, so the `9`th number spoken is the difference between them, __`4`__.
*   __Turn 10__: `4` is new, so the `10`th number spoken is __`0`__.

(The game ends when the Elves get sick of playing or dinner is ready, whichever comes first.)

Their question for you is: what will be the __`2020`th__ number spoken? In the example above, the `2020`th number spoken will be `436`.

Here are a few more examples:

*   Given the starting numbers `1,3,2`, the `2020`th number spoken is `1`.
*   Given the starting numbers `2,1,3`, the `2020`th number spoken is `10`.
*   Given the starting numbers `1,2,3`, the `2020`th number spoken is `27`.
*   Given the starting numbers `2,3,1`, the `2020`th number spoken is `78`.
*   Given the starting numbers `3,2,1`, the `2020`th number spoken is `438`.
*   Given the starting numbers `3,1,2`, the `2020`th number spoken is `1836`.

Given your starting numbers, __what will be the `2020`th number spoken?__

Your puzzle answer was `371`.

## --- Part Two ---

Impressed, the Elves issue you a challenge: determine the `30000000`th number spoken. For example, given the same starting numbers as above:

*   Given `0,3,6`, the `30000000`th number spoken is `175594`.
*   Given `1,3,2`, the `30000000`th number spoken is `2578`.
*   Given `2,1,3`, the `30000000`th number spoken is `3544142`.
*   Given `1,2,3`, the `30000000`th number spoken is `261214`.
*   Given `2,3,1`, the `30000000`th number spoken is `6895259`.
*   Given `3,2,1`, the `30000000`th number spoken is `18`.
*   Given `3,1,2`, the `30000000`th number spoken is `362`.

Given your starting numbers, __what will be the `30000000`th number spoken?__

Your puzzle answer was `352`.
