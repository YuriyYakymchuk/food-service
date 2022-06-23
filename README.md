# food-service

REST API service written in Go language for learning purposes

The service exposes the following API:
### Food
* Get all food orders: **/api/food** GET
* Get food orders for the user: **/api/food/{userId}** GET
* Create a food order: **/api/food** POST
* Update a food order: **/api/food** PUT
* Delete a food order: **/api/food** DELETE

### User
* Create a user: **/api/user** POST

### Greeting
* Get greeting **/api/hello/{name}** GET
