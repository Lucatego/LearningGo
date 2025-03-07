# LearningGo
A repository to learn Go, concurrency and to practice applications of Go in a backend context.
## QuizGame
A simple game extracted from gophersises.com course.

## Calculator
A Calculator in Go, to learn structs and the net.Conn usage.

## TextChat
This proyect consists in a server that recieves connections and process the connections using a SQLite database.

### Server
Where everything happens.
Steps:
1. The server recieves the connections from the clients and stores them in a channel that works as a queue
2. The distributor creates an specific number of handlers to handle the connection in their own queues.
3. The distributor will check and assing the client to the freest handler.
4. The handler will recieve the connection from its channel and manage it until the connection is closed.

### Client
TODO