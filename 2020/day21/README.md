## --- Day 21: Allergen Assessment ---

You reach the train's last stop and the closest you can get to your vacation island without getting wet. There aren't even any boats here, but nothing can stop you now: you build a raft. You just need a few days' worth of food for your journey.

You don't speak the local language, so you can't read any ingredients lists. However, sometimes, allergens are listed in a language you __do__ understand. You should be able to use this information to determine which ingredient contains which allergen and work out which foods are safe to take with you on your trip.

You start by compiling a list of foods (your puzzle input), one food per line. Each line includes that food's __ingredients list__ followed by some or all of the allergens the food contains.

Each allergen is found in exactly one ingredient. Each ingredient contains zero or one allergen. __Allergens aren't always marked__; when they're listed (as in `(contains nuts, shellfish)` after an ingredients list), the ingredient that contains each listed allergen will be __somewhere in the corresponding ingredients list__. However, even if an allergen isn't listed, the ingredient that contains that allergen could still be present: maybe they forgot to label it, or maybe it was labeled in a language you don't know.

For example, consider the following list of foods:

    mxmxvkd kfcds sqjhc nhms (contains dairy, fish)
    trh fvjkl sbzzf mxmxvkd (contains dairy)
    sqjhc fvjkl (contains soy)
    sqjhc mxmxvkd sbzzf (contains fish)

The first food in the list has four ingredients (written in a language you don't understand): `mxmxvkd`, `kfcds`, `sqjhc`, and `nhms`. While the food might contain other allergens, a few allergens the food definitely contains are listed afterward: `dairy` and `fish`.

The first step is to determine which ingredients __can't possibly__ contain any of the allergens in any food in your list. In the above example, none of the ingredients `kfcds`, `nhms`, `sbzzf`, or `trh` can contain an allergen. Counting the number of times any of these ingredients appear in any ingredients list produces __`5`__: they all appear once each except `sbzzf`, which appears twice.

Determine which ingredients cannot possibly contain any of the allergens in your list. __How many times do any of those ingredients appear?__

Your puzzle answer was `2380`.

## --- Part Two ---

Now that you've isolated the inert ingredients, you should have enough information to figure out which ingredient contains which allergen.

In the above example:

*   `mxmxvkd` contains `dairy`.
*   `sqjhc` contains `fish`.
*   `fvjkl` contains `soy`.

Arrange the ingredients __alphabetically by their allergen__ and separate them by commas to produce your __canonical dangerous ingredient list__. (There should __not be any spaces__ in your canonical dangerous ingredient list.) In the above example, this would be __`mxmxvkd,sqjhc,fvjkl`__.

Time to stock your raft with supplies. __What is your canonical dangerous ingredient list?__

Your puzzle answer was `ktpbgdn,pnpfjb,ndfb,rdhljms,xzfj,bfgcms,fkcmf,hdqkqhh`.