# Review bot

## Architecture
- The backend is made entirely in golang
- The frontend is made with HTMX for making ajax requests and HTML + CSS(tailwind) for the UI

To run you can use the following commands in the root folder of this repository

``go run cmd/main.go``

``air``

## How the bot works

The bot works by starting up a conversation after receiving an event
(for the purpose of this demonstration, only a http request)

The bot manages a state machine, using a field in the database to keep track of which state the conversation is in.

![Bot Design System](/images/design_system.png)

![Bot Database](/images/tables.png)

The conversation states and transitions are described in the following image:

![Bot Flow](/images/flow_design.png)


The sentiment analysis is done using the vader lib