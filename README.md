# The Goal Game Demonstration

This is an illustration of the game Alex Rogo, the protagonist in the novel 
["The Goal" by  Eliyahu M. Goldratt](https://en.wikipedia.org/wiki/The_Goal_(novel)), plays with scout troop
during a lunch break on the trail. The game illustrates a basic idea behind the Theory of Constraints:
- the bottleneck controls the output of the whole system; and
- all resources behind the bottleneck are always starved.

Here's a small visual explanation of the game and its significance: 

[![The Goal Game](https://i9.ytimg.com/vi/RPZG8r_poZg/mq3.jpg?sqp=CNy5vbMG-oaymwEmCMACELQB8quKqQMa8AEB-AH-CYAC0AWKAgwIABABGE8gZSgrMA8=&rs=AOn4CLAU5ZfnWtZ58tfXmWKPtYLALNQlNw)](https://youtu.be/RPZG8r_poZg)

The process described in the book may be difficult to follow. This app helps with visualizing the waves the inventory
moves through the plant. Unlike ALex Rogo who was limited to a piece pof paper, this small app draws possible 
results the game.

Building on the idea described in the book, this app provides modifications to this game that is tuned to the specifics of 
the software development projects.

## A small theory
The game shows the inefficiency of a balanced production pipeline (or, by extension, a project team). A perfectly balanced
production pipeline is where the capacity of the pipeline perfectly matches the work is given to it. 
The inefficiency arises from the very fact
that in the event a flow through the production line is upset out of balance, there isn't a possibility to
catch up.

In reality, a capacity of a work center on a production line fluctuate around mean capacity: sometimes a work center does a little bit
above mean, sometimes below. Same for software engineers in an agile project working on features in the sprint.
These little oscillations have a dramatic impact on the output.

In the novel, Alex Rogo theorizes that the production pipeline he's setting up has a mean output of 3.5 matches
per cycle. So after 20 runs they should have "processed" 70 matches. Running this simulation several times
shows that such production line would produce around 50 "processed matches" after 20 cycles.

This small container is a more visual replacement for the spreadsheet described in the book that Alex Rogo
is using to keep the track of the situation.

## Running this app in container
You can run this container in container indicating desired plots to be produced:
```shell
docker run dlspwd2/the_goal_game:latest -plot-achieved-output
```
```shell
podman run docker.io/lspwd2/the_goal_game:latest -plot-achieved-output
```
Below is the example output of such a run:
```
 70 ┤                                                        ╭──
 65 ┤                                                   ╭────╯
 61 ┤                                               ╭───╯
 56 ┤                                           ╭───╯
 52 ┤                                       ╭───╯
 47 ┤                                   ╭───╯                ╭──
 42 ┤                               ╭───╯                 ╭──╯
 38 ┤                           ╭───╯                  ╭──╯
 33 ┤                       ╭───╯                 ╭────╯
 29 ┤                   ╭───╯                 ╭───╯
 24 ┤               ╭───╯          ╭──────────╯
 19 ┤           ╭───╯         ╭────╯
 15 ┤       ╭───╯     ╭───────╯
 10 ┤   ╭───╯  ╭──────╯
  6 ┼───╭──────╯
  1 ┼───╯
                        Achieved 45 of 70 mean

                       ■ Expected   ■ Achieved
```
## Command Line Options

### Repeating options
- `-R` number of times to re-run the simulation (default to 1)

### Production line settings
- `-ps` production line size; default to `5`
- `-wrc-min` minimal capacity of a work center; default to `1`
- `-wrc-max` maximum capacity fo a work center; default to `6`
- `-c` number of cycles; default to `20` that are described in the book

### Match box behavior

- `-m`: feeding inventory option. The options are: `simple`, `epic-alternating`. Default to `simple`

In the book, Alex and scouts roll a die to determine the number of matches to draw from a box of matches.
This "match box" behaviour is modelled when `-m` option is set to `simple`. Here's the illustration:

| Cycle Number | Die Roll | Supplied |
| -------------| ---------| -------- |
| 1 | 2 | 2 individual matches |
| 2 | 5 | 5 individual matches |
| 3 | 1 | A single match |
| 4 | 2 | 2 individual matches |


The `epic-alternating`, as it's name implies, provides alternating "epic staff" for one cycle and set of individual
matches for another *regardless* of what the first work center in pipeline is capable of pulling (equal to the number
on the die).

There's the example using the pipeline described in the book with the mean capacity of the first work center of 3.5
and capability to draw determined by a throw of a die for the same results of a die roll.

| Cycle Number | Die Roll | Supplied |
| -------------| ---------| -------- |
| 1 | 2 | Epic with size 3 |
| 2 | 5 | 4 individual matches |
| 3 | 1 | Epic with size 3 |
| 4 | 2 | 4 individual matches |

This situation resembles closer a steady influx of work flowing into your average agile team from the stakeholders
with a complete disregard of how much work a team is able to take in and produce in the current iteration. This
situation is susceptible to producing bigger lags in delivery. 

> Don't be fooled by a coincidence. This is a game of chance. Probability can both favor and disfavour a particular
> simulation. The patter, though, remains.


### Output settings
- `-plot-achieved-output` plot achieved output vs mean output
- `-plot-lag` plot log of the output behind the mean output, the difference between two lines above
- `-plot-wrc-inventories` plot inventory size each work center has accumulated
- `-plot-wrc-starving` plot how much each work center starved (difference between how much a work center can
   do in this run vs inventory available)
- `-plot-starving` plot cumulative starving across all work centers
- `-G` plot all