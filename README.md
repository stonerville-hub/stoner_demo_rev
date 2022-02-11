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

 >Docker is installed

 >Image is run from your local machine

# [Getting Started](#getting-started)

>run `startApp.sh` in root
- For test data, go to [Load Test Data](http://localhost:8080/loaddata).  Api keys will print to console. 

- If using default parms, goto [Home Page](http://localhost:8080/)


**`AeroSpike Config Parameters`**
- `env`  [**default** `env` := dev]
- `namespace` [**default** `namespace` := test]
- `set` [**default** `set` := users]

**`Examples`**

- startApp.sh

- startApp.sh dev test users