# dice-pouch

A dice rolling tool for all kinds of tabletop RPG players.

## Getting started

Run the docker image and map the port of your choice on your host, eg:
```
docker run -d -p 80:9001 ge0rr/dice-pouch
```
You can now check that the service is running at `127.0.0.1:80`.

### Roll the dice
Rolling a new roll is easy, for instance:
```
http://127.0.0.1/roll?r1=d20+5
```
The above requested a roll of one d20 (20 sided die), plus a modifier of 5. 
We get the results in json, like so:
```json
[[{"sum":17},{"d20":12},{"modifier":5}]]
```
Here we see the results of any dice rolled and any modifiers.
Sum is the final result of the roll, here it's `17`, where all of the above are added together.

#### Add negative effects
Here is a different example:
```
http://127.0.0.1/roll?r1=2d6-1-d4
```
We rolled two d6 with a modifier of -1 and also deduct from the result one d4. The final sum is `4` as shown on the json:
```json
[[{"sum":4},{"d6":2},{"d6":5},{"modifier":-1},{"d4":-2}]]
```

#### Multiple rolls
You could also have multiple rolls:
```
127.0.0.1/roll?r1=d20+5&r2=2d6+3+d4&r3=3d12-2d4
```
Here we have three rolls:
* the first roll `r1` is a d20 with a modifier of 5
* the second roll `r2` is two d6 with a modifier of 3 plus a d4
* the third roll `r3` is three d12 minus two d4

## Suggested use
Let's say you play DnD and you are usually attacking with your two shortswords (two weapon fighting).
You can quickly get the result of your attack and bonus attack with a roll like that:
```
127.0.0.1/roll?r1=d20+5&r2=d6+3&r3=d20&r4=d6+3
```

Here `r1` and `r2` are the attack roll and damage roll for the first hand. Similarly, `r3` and `r4` are the attack and damage rolls for the second hand.
I usually bookmark the roll so that it is one click away. A second bookmark might go to my ranged attack roll.

### Suggestion for DMs
DMs might benefit from this by organizing a number of attacks from different enemies as bookmarks. 
This eliminates the added complexity of looking up attacks during battle, and helps maintain momentum in combat.

What I usually do when DMing is create a bookmark folder for each planned encounter. Inside each encounter folder I have a list of `dice-pouch/roll` bookmarks, each bookmark corresponding to rolls of one enemy.
