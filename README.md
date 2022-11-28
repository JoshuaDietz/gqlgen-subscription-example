# gqlgen-subscription-example
The subscription example from the gqlgen documentation.

The following repository demonstrates a bug in query verification of gqlgen.

Start the server with `go run main.go`

When sending the correct graphql query

    subscription {
        currentTime {
            unixTime
            timeStamp
        }
    }

everything works fine. When instead sending the syntactically incorrect query (note the opening and closing brackets after `subscription`)

    subscription() {
        currentTime {
            unixTime
            timeStamp
        }
    }

The server does:
- not report an error
- still answering with the reply to the query
- after that closing the channel without any comment leaving the user in the dark why the updates to the subscription do not work