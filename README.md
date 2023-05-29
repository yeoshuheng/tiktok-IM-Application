# Instant Messaging Application

![Tests](https://github.com/TikTokTechImmersion/assignment_demo_2023/actions/workflows/test.yml/badge.svg)

This is a working backend developed as part of TikTok Tech Immersion 2023.

### Objectives
My objective for this program was to learn as much as possible. 
Hence, this application has been implemented for both MySQL and Redis.
Which database to be used can be configured within *./rpc-server/main.go*, *./rpc-server/handlers.go* and *./docker-compose.yml*.

### Running the program
The program can be set up by first running the following command on terminal within the project's main directory.
This should also be ran everytime you switch between the SQL and Redis implementation of the program.
```shell
    docker-compose up -d --build
```

Subsequently, the program can be ran with just this:
```shell
    docker-compose up
```

### Common Issues

A common issue I faced when setting this program up on MacOS was a permission error in the *.buildx* file.
This was resolved by running the following command which updates *./docker* permission settings from the terminal.
```shell
    sudo chown -R [user] ~./docker
```

Another issue I faced was running opening making port 3306 available for MySQL on MacOS. Before building the script you would have to clear port 3306 by going into *System Preferences/MySQL* and stopping the server currently running on port 3306. More information can be found [here](https://stackoverflow.com/questions/54575020/not-able-to-kill-mysql-process-with-kill-9-pid).

### Structure
The program consists of a HTTP Server that takes in HTTP requests, which are then passed onto our RPC Server, which will run the necessary logic to read / write to the database.
The program supports only POST and GET requests.

#### SEND
This uses endpoint `api/send`, the following is an example body for the POST request:
```json
    {
      "chat" : "party1:party2", // Note that both parties are kept seperated with ':'.
      "text" : "hi",
      "sender" : "party1"
      }
```

#### PULL
This uses endpoint `api/pull`, the following is an example body for the GET request:
```json
    {
        "chat" : "party1:party2",
        "cursor" : 0,
        "limit" : 10,
        "reverse" : true
      }
```

The expected output should look like this.
```json
    {
        "messages" : {
            
            "chat" : "party1:party2",
            "text" : "hi",
            "sender" : "party1",
            "time" : 1685209278  // Returns in Unix timestamp.

          } 
    }
```

If there are multiple POST requests called beforehand, and the limit of your PULL request is below that value, the expected output should look like this, when we call the following request:
```json
    {
        "chat" : "party1:party2",
        "cursor" : 0,
        "limit" : 2,
        "reverse" : true
      }

```

We should expect an output like this:
```json
    {
    "messages": [
        {
            "chat": "party1:party2",
            "text": "hi",
            "sender": "party1",
            "send_time": 1685333688
        },
        {
            "chat": "party2:party1",
            "text": "hi",
            "sender": "party2",
            "send_time": 1685333686
        }
    ],
    "has_more": true,
    "next_cursor": 2
}
```

*For more information on Unix timestamps, read [here](https://unixtimestamp.com).*
