## --- Day 12: Rain Risk ---

Your ferry made decent progress toward the island, but the storm came in faster than anyone expected. The ferry needs to take __evasive actions__!

Unfortunately, the ship's navigation computer seems to be malfunctioning; rather than giving a route directly to safety, it produced extremely circuitous instructions. When the captain uses the [PA system](https://en.wikipedia.org/wiki/Public_address_system) to ask if anyone can help, you quickly volunteer.

The navigation instructions (your puzzle input) consists of a sequence of single-character __actions__ paired with integer input __values__. After staring at them for a few minutes, you work out what they probably mean:

*   Action __`N`__ means to move __north__ by the given value.
*   Action __`S`__ means to move __south__ by the given value.
*   Action __`E`__ means to move __east__ by the given value.
*   Action __`W`__ means to move __west__ by the given value.
*   Action __`L`__ means to turn __left__ the given number of degrees.
*   Action __`R`__ means to turn __right__ the given number of degrees.
*   Action __`F`__ means to move __forward__ by the given value in the direction the ship is currently facing.

The ship starts by facing __east__. Only the `L` and `R` actions change the direction the ship is facing. (That is, if the ship is facing east and the next instruction is `N10`, the ship would move north 10 units, but would still move east if the following action were `F`.)

For example:

    F10
    N3
    F7
    R90
    F11

These instructions would be handled as follows:

*   `F10` would move the ship 10 units east (because the ship starts by facing east) to __east 10, north 0__.
*   `N3` would move the ship 3 units north to __east 10, north 3__.
*   `F7` would move the ship another 7 units east (because the ship is still facing east) to __east 17, north 3__.
*   `R90` would cause the ship to turn right by 90 degrees and face __south__; it remains at __east 17, north 3__.
*   `F11` would move the ship 11 units south to __east 17, south 8__.

At the end of these instructions, the ship's [Manhattan distance](https://en.wikipedia.org/wiki/Manhattan_distance) (sum of the absolute values of its east/west position and its north/south position) from its starting position is `17 + 8` = __`25`__.

Figure out where the navigation instructions lead. __What is the Manhattan distance between that location and the ship's starting position?__

Your puzzle answer was `923`.

## --- Part Two ---

Before you can give the destination to the captain, you realize that the actual action meanings were printed on the back of the instructions the whole time.

Almost all of the actions indicate how to move a __waypoint__ which is relative to the ship's position:

*   Action __`N`__ means to move the waypoint __north__ by the given value.
*   Action __`S`__ means to move the waypoint __south__ by the given value.
*   Action __`E`__ means to move the waypoint __east__ by the given value.
*   Action __`W`__ means to move the waypoint __west__ by the given value.
*   Action __`L`__ means to rotate the waypoint around the ship __left__ (__counter-clockwise__) the given number of degrees.
*   Action __`R`__ means to rotate the waypoint around the ship __right__ (__clockwise__) the given number of degrees.
*   Action __`F`__ means to move __forward__ to the waypoint a number of times equal to the given value.

The waypoint starts __10 units east and 1 unit north__ relative to the ship. The waypoint is relative to the ship; that is, if the ship moves, the waypoint moves with it.

For example, using the same instructions as above:

*   `F10` moves the ship to the waypoint 10 times (a total of __100 units east and 10 units north__), leaving the ship at __east 100, north 10__. The waypoint stays 10 units east and 1 unit north of the ship.
*   `N3` moves the waypoint 3 units north to __10 units east and 4 units north of the ship__. The ship remains at __east 100, north 10__.
*   `F7` moves the ship to the waypoint 7 times (a total of __70 units east and 28 units north__), leaving the ship at __east 170, north 38__. The waypoint stays 10 units east and 4 units north of the ship.
*   `R90` rotates the waypoint around the ship clockwise 90 degrees, moving it to __4 units east and 10 units south of the ship__. The ship remains at __east 170, north 38__.
*   `F11` moves the ship to the waypoint 11 times (a total of __44 units east and 110 units south__), leaving the ship at __east 214, south 72__. The waypoint stays 4 units east and 10 units south of the ship.

After these operations, the ship's Manhattan distance from its starting position is `214 + 72` = __`286`__.

Figure out where the navigation instructions actually lead. __What is the Manhattan distance between that location and the ship's starting position?__

Your puzzle answer was `24769`.
