# Chatops
Chatops is a chatbot which can integrate with google hangouts chat and perform a lot of internal ops related tasks. It currently does not interpret natural langauge. So, the interaction happens through a limited set of *commands/statements*

## What it can do
It is capable of executing commands/scripts based on a user's message. You can use `help` to get all the available commands and how to use them.

## How it works

- When a user messages the bot on hangouts, hangouts will make a post requests to chatops service with the message and other details. 
- Chatops will try to match the message with the defined set of regex strings. 
- Each regex string is tied to a command/script. 
- If the regex matches the message the user sent, the tied command/script will be executed. 
- The output of the command/script will be sent to the same user (async).
- These regex strings are stored in a file called `commands.json`. This file will be read only once, when the service starts.

## Get started


## Adding a command


```
// Command that runs on a remote machine 
{
    "name": "command-name",
    "hostname": "0.0.0.0",
    "command": "command arg1 arg2 ...",
    "regex": "^The text that triggers this command$",
    "example": "command 1 2 3"
}

// Command that runs locally
{
    "name": "command-name",
    "command": "command arg1 arg2 ...",
    "regex": "^the text that triggers this command$",
    "example": ""
}

// Command that takes an arguement
{
    "name": "command-name",
    "command": "command {{arg1}} {{arg2}}",
    "regex": "^text with (?P<arg1>\\S+) (?P<arg2>\\S+)$",
    "example": "text with 1 2"
}
```