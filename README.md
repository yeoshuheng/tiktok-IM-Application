# Instant Messaging Application

![Tests](https://github.com/TikTokTechImmersion/assignment_demo_2023/actions/workflows/test.yml/badge.svg)

This is a working backend developed as part of TikTok Tech Immersion 2023.

### Objectives
My objective for this program was to learn as much as possible. 
Hence, this application has been implemented for both MySQL and Redis.
Which database to be used can be configured within rpc-server/main.go, rpc-server/handlers.go and ./docker-compose.yml.

### Running the program.
The program can be set up by first running the following command on terminal within the project's main directory.
This should also be ran everytime you switch between the SQL and Redis implementation of the program.
```shell
    docker-compose up -d --build
```

Subsequently, the program can be ran with just this:
```shell
    docker-compose up
```

### Structure
The program consists of a HTTP Server that takes in HTTP requests, which are then passed onto our RPC Server, which will run the necessary logic to read / write to the database.
The program supports only POST and GET requests.

#### SEND
This uses endpoint *api/send*, the following is an example body for the POST request:
```json
    {
      "chat" : "party1:party2", // Note that both parties are kept seperated with ':'.
      "text" : "hi",
      "sender" : "party1"
      }
```

#### PULL
This uses endpoint *api/pull*, the following is an example body for the GET request:
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
For more information on Unix timestamps, read [here](https://unixtimestamp.com)
