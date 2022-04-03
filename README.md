# taxi_matching

## Install & Run
```
docker-compose build
docker-compose up
```

## APIS

### Authentication API (Port 9090)
Authentication API interacts with postgresql database which stores user information.  
Using this api, you can register a new user and signin with an existing user.  
When you sign in an existing user, the api will return a jwt token that will expire after 15 mins.  
You have to add this token to the request header with "Token" key in order to use other apis.  
  
Swagger Documentation: "localhost:9090/docs"

### Driver Location API (Port 8080, NOT EXPOSED)

Driver Location API is used to interact with mongodb that stores the drivers.  
It requires authentication.  
For authentication, Authentication API will be used.  
This api is not exposed to the users.  
Matching api will use this web service.  
  
Since this api is not exposed, the swagger document for this api will be reached from matching api.  
Swagger Documentation: "localhost:9191/driver-api-docs"

### Matching API (Port 9191)

The Matching API matches a suitable driver with the rider using Driver Location API. 
It allows the query to find the nearest driver around a given GeoJSON point.  
The API provides finding the nearest driver location with a given GeoJSON point and radius.  

It only accepts authenticated user requests (requests with Token in header).  
The "/find" endpoint that allows searching with a GeoJSON point and radiues to find a driver if there is any inside radius around location.  

Swagger Documentation: "localhost:9191/docs"
