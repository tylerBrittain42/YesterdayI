# Yesterday I...

A quick and easy way to log work for standup, whether that is so you can read it off during the meeting or if you can't remember what yyou did last week

# Ideas
## Adding data
- should have long form to log whole day
- should have short form to log single item
- should include date of logging(can allow override to date minus 1 or custom)
## Should be able to view data
- views should include:
    - today
    - yesterday
    - yesterday + today
    - all days
## edit
- use view window and multiple choice to select option to override
## delete 
- use view window and multiple choice to select option to override

# Other libraries
- mostly standard lib
- consider charm when we get to edit and delete

# User input
- use yi as command
- tui for now think enter blank line to stop
- cli: also single line with args
## cli input
- Adding a single task(explicit)
    - *syntax:* ```yi -t <task here>```
    - *example:* ```yi -t implemented unit testing for the model dir```
- Assigning a jira ticket
    - will assume no ticket without flag
    - *sytanx:* ```yi -j a" -t "created a database"```
    - *example:* ```yi -j "CONFIG-1234" -t "created a database"```

## cli output
Just something simple for now
- no args should output everything



## to be done later
- Adding multiple tasks
    - opt in with flag
    - *sytanx:* ```yi -m "<task 1>" "<task 2>"```
    - *example:* ```yi -m "created a database" "updated docs"```
- Adding a single task(default)
    - default usage
    - *syntax:* ```yi "<task here>"```
    - *example:* ```yi "implemented unit testing for the model dir"```

    
# links to check out
- advice on parsing and testing parsing https://eli.thegreenplace.net/2020/testing-flag-parsing-in-go-programs/