#
# [Code Challenge - Microservice]()

## [Description](#description)
We need a new microservice to return the user data stored in Aerospike. The data resides within the "test" namespace, "users" set, and each record has a primary key integer corresponding to the user's id in RDS. Each record contains the following bins:
* api_key    
* first_name
* last_name
* company

Please create a GoLang HTTP REST service with a GET `/user/<id>` endpoint. The request will accept a `?api_key=<string>` GET param and will validate the value against the Aerospike user record before returning the user record as a JSON response.

This is your time to shine. Let's see all the best practices in play. A docker-compose including Aerospike, test record loading, README, etc... Publish the repository to GitHub and pass along the URL.
## [Prerequisites](#prerequisites)

 1. _Docker is installed_
 2. _Docker compose is installed_
 3. _Image is run from your local machine_  
 

## [Getting Started](#getting-started)

1. Start application by running `sh appUp.sh`
2. Go to one of these endpoints:
    - [Home](http://localhost:8080/) A list of api_keys for testing `Get By ID` endpoint will be displayed on this page.  If no records are displayed, click [here](http://localhost:8080/loadNewData) to generate 10 new records.
    - [Load Test Data](http://localhost:8080/loadNewData) Generates 10 new records
    - [Get by User ID](http://localhost:8080/user/2?api_key=249fd4fbff52414aa81a670d696bc2c9) Returns the json for the api_key
3. Stop application by running `sh appDown.sh`

# [Code Challenge Summary](#code-challenge-summary)
**`Existing functionality`**
- Loading test data
- Read record from aerospike using api_key
- Docker compose builds image, starts aerospike, and start/stop application
- If you are reading this, there is a ReadMe
- .env file has Aerospike configs

**`Functionality that needs additional Work`**
- API id path parm is not being utilized but is required.