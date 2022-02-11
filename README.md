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

 1.  _Docker is installed_

 2.  _Image is run from your local machine_

# [Getting Started](#getting-started)

There are a few options for running the application.
1. If Aerospike is up and running 172.0.0.1:3000, run the this from command line: *go run . dev 172.0.0.1 3000 test user*
2. Start up script. Run the `sh startApp.sh` in root
3. *WIP* Run *docker-compose up* in root
 
 Once application is up, API endpoints are functioning.  Here are your options:
- For test data, click [Load Test Data](http://localhost:8080/loaddata).  Api keys will print to console. 

- For validating the API, use this link `http://localhost:8080/user/{id}?api_key={api_key}`

**AeroSpike Config Parameters**
- `env`  [**default** `env` := dev]
- `namespace` [**default** `namespace` := test]
- `set` [**default** `set` := users]

**Examples**

1. sh startApp.sh *(default parms will be used)*

2. sh startApp.sh dev test users


## `Code Challenge Summary`
**`Existing functionality`**
- loading test data from /loaddata API.  Api_keys are output to console
- Read record from aerospike using api_key
- Build a new image of app using Dockerfile
- If you are reading this, you know there is a ReadMe
- Docker compose builds and start app image.  This would eliminating the `startApp.sh`.


**`Functionality that needs additional Work`**
- docker-compose connection to aerospike database from application inside container
- API `/user/{id}` endpoint needs to return json format.
- API path parm is not being used but is required.
- All output is on the console.