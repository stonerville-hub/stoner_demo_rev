># Code Challenge - Microservice

## [Description](#description)
We need a new microservice to return the user data stored in Aerospike. The data resides within the "test" namespace, "users" set, and each record has a primary key integer corresponding to the user's id in RDS. Each record contains the following bins:
* api_key    
* first_name
* last_name
* company

Please create a GoLang HTTP REST service with a GET `/user/<id>` endpoint. The request will accept a `?api_key=<string>` GET param and will validate the value against the Aerospike user record before returning the user record as a JSON response.

This is your time to shine. Let's see all the best practices in play. A docker-compose including Aerospike, test record loading, README, etc... Publish the repository to GitHub and pass along the URL.
# [Prerequisites](#prerequisites)

 1. _Docker is installed_
 2. _Docker compose is installed_
 3. _Image is run from your local machine_  
 

# [Getting Started](#getting-started)

1. run  `docker-compose up -d --force-recreate`
2. Go to one of the following endpoints:
    - For test data, click [Load Test Data](http://localhost:8080/loaddata).  Api keys will print to console. 
    - For validating the API, use this link `http://localhost:8080/user/{id}?api_key={api_key}`


# [Code Challenge Summary](#code-challenge-summary)
**`Existing functionality`**
- loading test data from /loaddata API.  Api_keys are output to console
- Read record from aerospike using api_key
- Build a new image of app using Dockerfile
- Docker compose builds image, starts aerospike, and starts application
- If you are reading this, you know there is a ReadMe


**`Functionality that needs additional Work`**
- API `/user/{id}` endpoint needs to return json format.
- API path parm is not being used but is required.
- All output is on the console.