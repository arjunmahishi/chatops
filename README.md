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

- Clone repository
- Install dependencies
    ```
    $ go get -d
    ```
- Build the binary
    ```
    $ go build -o $GOPATH/bin/chatops
    ```
- Create a `commands.json` file with the following format
    ```
    {
        "commands": [
            {
                "name": "command11",
                "command": "echo 'hello world from command 1'",
                "regex": "^test-command-1$"
            },
            {
                "name": "command-2",
                "command": "echo 'hello world from command 2'",
                "regex": "^test-command-2$"
            }
        ]
    }
    ```
- Each command can be of the following variation
    ```
    // Command that runs on a remote machine 
    {
        "name": "command-name",
        "hostname": "<ip address>",
        "command": "<shell command>",
        "regex": "^The text that triggers this command$",
        "example": "<example usage>"
    }

    // Command that runs locally
    {
        "name": "command-name",
        "command": "<shell command>",
        "regex": "^The text that triggers this command$",
        "example": "<example usage>"
    }

    // Command that takes an arguement
    {
        "name": "command-name",
        "command": "command {{arg1}} {{arg2}}",
        "regex": "^text with (?P<arg1>\\S+) (?P<arg2>\\S+)$",
        "example": "text with 1 2"
    }
    ```
- Create a `config.json` file with the following details
    ```
    {
        "BotName": "<bot name>",
        "HangoutsToken": "<hangouts chat varification token>",
        "DialogFlowAccessToken": "<dialogflow access token>",
        "CommandsPath": "<path to commands.json>",
        "ServiceAccountCredsPath": "<path to google service account credentials file>"
    }
    ```
- Run the binary file 
    ```
    $ chatops -config config.json
    ```

    The server should start. The output would look something like this
    ```
    $ chatops -config config.json                                                                                
    2019/05/09 23:03:19 Syncing commands list
    2019/05/09 23:03:19 Total commands 1
    ⇨ http server started on [::]:1323
    ```
- Now, go to google developer console and configure the hangouts chat API to hit this server. 


## Contributing
No guidelines yet. Just make a PR :)