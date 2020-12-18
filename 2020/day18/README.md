## --- Day 18: Operation Order ---

As you look out the window and notice a heavily-forested continent slowly appear over the horizon, you are interrupted by the child sitting next to you. They're curious if you could help them with their math homework.

Unfortunately, it seems like this "math" [follows different rules](https://www.youtube.com/watch?v=3QtRK7Y2pPU&t=15) than you remember.

The homework (your puzzle input) consists of a series of expressions that consist of addition (`+`), multiplication (`*`), and parentheses (`(...)`). Just like normal math, parentheses indicate that the expression inside must be evaluated before it can be used by the surrounding expression. Addition still finds the sum of the numbers on both sides of the operator, and multiplication still finds the product.

However, the rules of __operator precedence__ have changed. Rather than evaluating multiplication before addition, the operators have the __same precedence__, and are evaluated left-to-right regardless of the order in which they appear.

For example, the steps to evaluate the expression `1 + 2 * 3 + 4 * 5 + 6` are as follows:

    1 + 2 * 3 + 4 * 5 + 6
      3   * 3 + 4 * 5 + 6
          9   + 4 * 5 + 6
             13   * 5 + 6
                 65   + 6
                     71

Parentheses can override this order; for example, here is what happens if parentheses are added to form `1 + (2 * 3) + (4 * (5 + 6))`:

    1 + (2 * 3) + (4 * (5 + 6))
    1 +    6    + (4 * (5 + 6))
         7      + (4 * (5 + 6))
         7      + (4 *   11   )
         7      +     44
                51

Here are a few more examples:

*   `2 * 3 + (4 * 5)` becomes __`26`__.
*   `5 + (8 * 3 + 9 + 3 * 4 * 3)` becomes __`437`__.
*   `5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))` becomes __`12240`__.
*   `((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2` becomes __`13632`__.

Before you can help with the homework, you need to understand it yourself. __Evaluate the expression on each line of the homework; what is the sum of the resulting values?__

Your puzzle answer was `86311597203806`.

## --- Part Two ---

You manage to answer the child's questions and they finish part 1 of their homework, but get stuck when they reach the next section: __advanced__ math.

Now, addition and multiplication have __different__ precedence levels, but they're not the ones you're familiar with. Instead, addition is evaluated __before__ multiplication.

For example, the steps to evaluate the expression `1 + 2 * 3 + 4 * 5 + 6` are now as follows:

    1 + 2 * 3 + 4 * 5 + 6
      3   * 3 + 4 * 5 + 6
      3   *   7   * 5 + 6
      3   *   7   *  11
         21       *  11
             231

Here are the other examples from above:

*   `1 + (2 * 3) + (4 * (5 + 6))` still becomes __`51`__.
*   `2 * 3 + (4 * 5)` becomes __`46`__.
*   `5 + (8 * 3 + 9 + 3 * 4 * 3)` becomes __`1445`__.
*   `5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))` becomes __`669060`__.
*   `((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2` becomes __`23340`__.

__What do you get if you add up the results of evaluating the homework problems using these new rules?__

Your puzzle answer was `276894767062189`.
