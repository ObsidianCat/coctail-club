# Cocktail club
Set of API routes which allow two functions. The first is search for cocktails by name, ingredient or id.<br> 
The second is creating collection of cocktails anf saving it into a json file. 
Service makes use of cocktails data provided by [The Cocktail DB](https://www.thecocktaildb.com/api.php)
## Running
Make sure that you have Golang installed. Then, in the project root directory, run<br>
``$ go run .``
It is a shortcut for building and running application in one step. 
## usage
**Cocktails** 
- **by name GET /cocktail/name/:name** | http://localhost:8080/cocktails/name/margarita
- **by ingredient GET /cocktails/ingredient/:name** | http://localhost:8080/cocktails/ingredient/rum
- **by Id GET /cocktails/id/:id** | http://localhost:8080/cocktails/id/13192

**Collection**
- **Get cocktails in collection GET /collection** | http://localhost:8080/collection
- **Save collection into a file POST /collection** | http://localhost:8080/collection
- **Add cocktail to collection PUT /collection/add/id/:id** | http://localhost:8080/collection/add/id/13192

## Linting
Run the following command, to lint all the files in all subdirectories.<br>
``$ golint ./...``
## Testing
Run the following command, to run all tests.<br>
``$ go test ./... -cover ``
