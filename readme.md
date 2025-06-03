# Yesterday I...

A quick and easy way to log work for standup, whether that is so you can read it off during the meeting or if you can't remember what yyou did last week

# cli usage
The cli consists of two commands, `add` and `view`.
the flag `-help` can be used for each subcommand
# add
The `add` subcommand is used to add tasks. 
- `-t` is required and is used to store the task
- `-j` is optional and is used for a jira ticket or whatever other sort of category you wish.
## Examples
 - Adding a single task without a flag
    - *syntax:* ```yi -t '<task here>'```
    - *example:* ```yi -t 'implemented unit testing for the model dir'```
- Adding a single task with a jira flag
    - will assume no ticket without flag
    - *syntax:* ```yi -t '<task here>' -j '<jira here>'```
    - *syntax:* ```yi -t 'implemented testing' -j 'config-2159'```


## view
The `view` subcommand is used to view tasks. No flags(or only json flag) will print all tasks.
- `-json` will provide the output in json form. This can be used in combination with any of the other flags.
- `-start` will provide a start date. It can be used in tandem with `-end`. If no end date is specified it will show the range from `start` until now.
- `-end` will provide an end date. It requires the `start` flag to be used.
- `-date` will only print tasks of a specific date.

*Note: all dates must use the form 'xx/xx, example: 01/02.*

