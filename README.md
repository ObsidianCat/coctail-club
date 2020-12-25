# Cocktail club
Set of APi routes which allow to lockup cocktails by names and arrnage a cocktails collections which are saved as json files.
The application make use of data provided by [The Cocktail DB](https://www.thecocktaildb.com/api.php)
## Running
Make sure that you have Golang installed. Then, in the project root directory, run<br>
``$ go run .``

It is a shortcut for building and running application in one step. 
## Examples of usage
**Cocktail-club** micro-service provides two options to request for data.
- by name **/cocktail/name/:name** | http://localhost:8080/cocktail/name/mojito
- by ingredient **/cocktail/ingredient?{someParameter}**

Example: getting cocktail recipe and description, when you know the cocktail name and this cocktail is among **Cocktail-club** records<br>
http://localhost:8080/cocktail/name/mojito<br>
Example: getting list of all cocktail recipes and descriptions, containing specific ingredient<br>
http://localhost:8080/cocktail/ingredient?rum
## Linting
Run the following command, to lint all the files in all subdirectories.<br>
``$ golint ./...``
## Testing
Run the following command, to run all tests.<br>
``$ go test ./... -cover ``
