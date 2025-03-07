# LearningGo
A repository to learn Go, concurrency and to practice applications of Go in a backend context.
## QuizGame
A simple game extracted from gophersises.com course.

## Calculator
A Calculator in Go, to learn structs and the net.Conn usage.

## TextChat
This proyect consists in a server that recieves connections and process the connections using a SQLite database.

### Server
*Where everything happens.* The steps are:
1. The server recieves the connections from the clients and stores them in a channel that works as a queue
2. The distributor creates an specific number of handlers to handle the connection in their own queues.
3. The distributor will check and assing the client to the freest handler.
4. The handler will recieve the connection from its channel and manage it until the connection is closed.

#### The database
In this proyect, the database is in SQLite, but it would be better if it was on a NoSQL databse. Anyways, it contains three major tables:
1. User
2. Conversation
3. Message
This tables connects via other tables like UserConversation table and UserMessage table to support group chats.

*To read this database SQL use DB Browser for SQLite.*
TODO: Add the scripts in its own folder.

### Client
TODO
