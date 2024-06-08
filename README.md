# The Goal Game Demonstration

This is an illustration of the game Alex Rogo, the protagonist in the novel 
["The Goal" by  Eliyahu M. Goldratt](https://en.wikipedia.org/wiki/The_Goal_(novel)), plays with scout troop
during a lunch break on the trail. The game illustrates a basic idea behind the Theory of Constraints:
- the bottleneck controls the output of the whole system; and
- all resources behind the bottleneck are always starved.

This simple app draws possible results of such a game, showing a great variability of possible output.

## A small theory
The game shows the inefficiency of a balanced production pipeline (or a project team). A perfectly balanced
production pipeline is where the capacity of the pipeline perfectly matches the work is given to it. 
The inefficiency arises from the fact
that the team is balanced. In the event a flow through the production line is upset, there isn't a possibility to
catch up.

In reality, a capacity of a work center will fluctuate around mean capacity: sometimes a work center does a little bit
above mean, sometimes below. These little oscillations have a dramatic impact on the output.

In the novel, Alex Rogo theorizes that the production pipeline he's setting up has a mean output of 3.5 matches
per cycle. So after 20 runs they should have "processed" 70 matches. Running this simulation several times
shows that such production line would produce around 50 "processed matches" after 20 cycles.

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

### Production line settings
- `-ps` production line size; default to `5`
- `-wrc-min` minimal capacity of a work center; default to `1`
- `-wrc-max` maximum capacity fo a work center; default to `6`
- `-r` number of runs; default to `20`

### Starting inventory
- `-i`: feeding inventory option. The options are: `simple`, `epic-alternating`

### Output settings
- `-plot-achieved-output` plot achieved output vs mean output
- `-plot-lag` plot log of the output behind the mean output
- `-plot-wrc-inventories` plot inventory size each work center has accumulated
- `-plot-wrc-starving` plot how much each work center starved (difference between how much a work center can
   do in this run vs inventory available)
- `-plot-starving` plot cumulative starving across all work centers
- `-G` plot all