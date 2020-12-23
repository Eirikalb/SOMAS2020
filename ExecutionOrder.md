
# Execution Order

Sequence of function calls from start to end of game.

## Initialisations

| Filename | Function | Description |
| ---- | ---- | ---- |
| main.go |  main | Begins game |
| internal/server/server.go | EntryPoint| Creates deep copy of list of gamestates. Then, while game is not over, calls runTurn and keeps track of the states during the game. |
|internal/server/turn.go|gameOver| Checks at least one client is alive and we haven't reached maximum number of turns or seasons. |
|internal/server/turn.go |runTurn| Gets a start of turn update, runs the organisations and runs the end of turn procedures. |
|internal/server/turn.go|startOfTurnUpdate| Sends update of gameState to alive clients. |
|internal/common/baseclient/baseclient.go|Client.StartOfTurnUpdate| Where the client receives the updated gameState from the server. |
|internal/server/turn.go|runOrgs| Runs IIGO, IIFO, IITO. |
|internal/server/iigo.go|      |
|      |      |