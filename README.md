# Cocktail club
## Running
Make sure that you have Golang installed. Then, in the project root directory, run<br>
``$ go run .``
It is a shortcut for building and running application in one step. 
## Examples of usage
**Cocktail-club** micro-service provides two endpoints. 
### By name
/cocktail/name/:name <br>
To get a cocktail recipe and description, if you know the cocktail name and this cocktail is among Cocktail-club records  
Example: http://localhost:8080/cocktail/name/mojito
### By ingredient
/cocktail/ingredient?{someParameter} <br>
To get all coctails which contain specific ingredient
Example: http://localhost:8080/cocktail/ingredient?rum
## Linting
Run the following command, to lint all the files in all subdirectories.<br>
``$ golint ./...``
## Testing
Run the following command, to run all tests.<br>
``$ go test ./... -cover ``
