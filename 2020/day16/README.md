## --- Day 16: Ticket Translation ---

As you're walking to yet another connecting flight, you realize that one of the legs of your re-routed trip coming up is on a high-speed train. However, the train ticket you were given is in a language you don't understand. You should probably figure out what it says before you get to the train station after the next flight.

Unfortunately, you can't actually __read__ the words on the ticket. You can, however, read the numbers, and so you figure out __the fields these tickets must have__ and __the valid ranges__ for values in those fields.

You collect the __rules for ticket fields__, the __numbers on your ticket__, and the __numbers on other nearby tickets__ for the same train service (via the airport security cameras) together into a single document you can reference (your puzzle input).

The __rules for ticket fields__ specify a list of fields that exist __somewhere__ on the ticket and the __valid ranges of values__ for each field. For example, a rule like `class: 1-3 or 5-7` means that one of the fields in every ticket is named `class` and can be any value in the ranges `1-3` or `5-7` (inclusive, such that `3` and `5` are both valid in this field, but `4` is not).

Each ticket is represented by a single line of comma-separated values. The values are the numbers on the ticket in the order they appear; every ticket has the same format. For example, consider this ticket:

    .--------------------------------------------------------.
    | ????: 101    ?????: 102   ??????????: 103     ???: 104 |
    |                                                        |
    | ??: 301  ??: 302             ???????: 303      ??????? |
    | ??: 401  ??: 402           ???? ????: 403    ????????? |
    '--------------------------------------------------------'

Here, `?` represents text in a language you don't understand. This ticket might be represented as `101,102,103,104,301,302,303,401,402,403`; of course, the actual train tickets you're looking at are __much__ more complicated. In any case, you've extracted just the numbers in such a way that the first number is always the same specific field, the second number is always a different specific field, and so on - you just don't know what each position actually means!

Start by determining which tickets are __completely invalid__; these are tickets that contain values which __aren't valid for any field__. Ignore __your ticket__ for now.

For example, suppose you have the following notes:

    class: 1-3 or 5-7
    row: 6-11 or 33-44
    seat: 13-40 or 45-50

    your ticket:
    7,1,14

    nearby tickets:
    7,3,47
    40,4,50
    55,2,20
    38,6,12

It doesn't matter which position corresponds to which field; you can identify invalid __nearby tickets__ by considering only whether tickets contain __values that are not valid for any field__. In this example, the values on the first __nearby ticket__ are all valid for at least one field. This is not true of the other three __nearby tickets__: the values `4`, `55`, and `12` are are not valid for any field. Adding together all of the invalid values produces your __ticket scanning error rate__: `4 + 55 + 12` = __`71`__.

Consider the validity of the __nearby tickets__ you scanned. __What is your ticket scanning error rate?__

Your puzzle answer was `29851`.

## --- Part Two ---

Now that you've identified which tickets contain invalid values, __discard those tickets entirely__. Use the remaining valid tickets to determine which field is which.

Using the valid ranges for each field, determine what order the fields appear on the tickets. The order is consistent between all tickets: if `seat` is the third field, it is the third field on every ticket, including __your ticket__.

For example, suppose you have the following notes:

    class: 0-1 or 4-19
    row: 0-5 or 8-19
    seat: 0-13 or 16-19

    your ticket:
    11,12,13

    nearby tickets:
    3,9,18
    15,1,5
    5,14,9

Based on the __nearby tickets__ in the above example, the first position must be `row`, the second position must be `class`, and the third position must be `seat`; you can conclude that in __your ticket__, `class` is `12`, `row` is `11`, and `seat` is `13`.

Once you work out which field is which, look for the six fields on __your ticket__ that start with the word `departure`. __What do you get if you multiply those six values together?__

Your puzzle answer was `3029180675981`.
