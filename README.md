# assignment-2

## Description
This is an assignment for the course [IDATG2005](https://www.ntnu.edu/studies/courses/PROG2005#tab=omEmnet) at 
[NTNU](https://www.ntnu.edu/). The assignment is to create a REST API that provides data about the percentage of
renewables in the energy mix for different countries. The data for the share of renewable energy can be found 
[here](https://drive.google.com/file/d/18G470pU2NRniDfAYJ27XgHyrWOThP__p/view), and the API for the neighbouring 
countries can be found [here](https://restcountries.com/).

## Graphical user interface
As an addition to the requirements a simple graphical user interface has been created to navigate the API and search
for countries and view this readme document. The GUI is available at the root path of the API. If you are running it 
locally this would be [localhost:8080](http://localhost:8080/) and if you run it though the deployed docker container
it can be found at this [ip-address](http://10.212.170.218:8080/).

If you do not wish to use the gui and instead enter the URL manually you can find examples of the different endpoints
further down in this document.

The GUI can be navigated by clicking on the different links. When you click either the "Current percentage of 
renewables" or the "Historical percentages of renewables" links you will be presented with input fields where you can
enter the relevant parameters for the endpoint. The input fields will be different depending on which endpoint you
choose. After entering the parameters you can click the "Search" button and the response will be displayed below the
input fields. If you wish to search for another country or countries you can simply enter the new parameters and click
the "Search" button again. The response will be updated with the new response. If you click the same link again the 
input fields will be hidden again.

If you click the "Readme" link you will be taken to this document.


## - Tree view of the file structure
```
Assignment-2
|   .gitignore
|   assignment-2-key.json
|   docker-compose.yml
|   Dockerfile
|   go.mod
|   go.sum
|   README.md
|
|
|
+---cmd
|       main.go
|
+---constants
|       commonstrings.go
|       internalPaths.go
|       paths.go
|
+---database
|       firestoreWebhooksDB.go
|       firestoreWebhooksDB_test.go
|       initFirestore.go
|
+---handlers
|   +---defaultHandler
|   |       defaultHandler.go
|   |       defaultHandler_test.go
|   |
|   +---notificationHandler
|   |       notificationHandler.go
|   |       notificationHandler_test.go
|   |
|   +---readmeHandler
|   |       readmeHandler.go
|   |
|   +---renewableHandlers
|   |   +---currentHandler
|   |   |       currentHandler.go
|   |   |       currentHandler_test.go
|   |   |
|   |   +---historicalHandler
|   |   |       historicalHandler.go
|   |   |       historicalHandler_test.go
|   |   |
|   |   \---renewableUtils
|   |           renewableUtils.go
|   |
|   \---statusHandler
|           statusHandler.go
|           statusHandler_test.go
|
+---json_coder
|       decoder.go
|       encoder.go
|
+---res
|       renewable-share-energy.csv
|
+---responses
|       countries.json
|       countries_formatted.json
|       formatJSON.go
|       renewable-share-energy.csv
|
+---static
|   +---css
|   |       default.css
|   |
|   \---html
|           defEndpoint.html
|           inputFields.js
|           responseContainer.js
|
+---structs
|       structs.go
|
+---utils
|   |   fileHandler.go
|   |   showHTML.go
|   |   time.go
|   |   urlToParams.go
|   |
|   \---hashing-utility
|           hash_secret.go
|           webhookHashing.go
|
\---webhooks
        webhooks.go

```

## Renewables
### Description
Renewables is a directory containing the historical percentage of renewables endpoint and the current percentage of
renewables endpoint. 

Both endpoints will return a JSON object containing the percentage of renewables in the energy mix for the selected 
country or countries.
- Historical percentage of renewables endpoint will return the percentage of renewables for
each year from the begin year to the end year.
-  Current percentage of renewables endpoint will return the percentage
of renewables for the latest year for which data is available.

##### Root paths
The endpoints have the following root paths:
- Historical percentage of renewables: `/energy/v1/renewables/history`
- Current percentage of renewables: `/energy/v1/renewables/current`

##### placeholders for query parameters
This document have the following conventions for placeholders:
- `{value?}` - optional value
- `{?parameter?}` - optional parameter 

#### Full endpoint paths
The endpoints have the following full paths:
- Current percentage of renewables: `/energy/v1/renewables/current/{country?}{?sortByValue=bool?}`
- Historical percentage of renewables: `/energy/v1/renewables/history/{country?}{?begin=year&end=year?}{?sortByValue=bool?}`

##### Alternative to query parameters
The endpoints also have an alternative form where the parameters are placed in the path instead of as query parameters.
The alternative form is as follows:
- Current percentage of renewables: `/energy/v1/renewables/current/{country?}/{bool?}`
- Historical percentage of renewables: `/energy/v1/renewables/history/{country?}/{year?}/{year?}/{bool?}`

***Note:*** All renewable endpoints are case-insensitive, so all parameter and values can be entered in any case.

---

##  - Current percentage of renewables endpoint

---
### Description 

The current percentage of renewables endpoint focuses on returning the current percentage of renewables in the energy
mix, for individual or selections of countries. It is possible to get one country, one country with its
neighbours, or all countries in the dataset.

### - Request
```
Method: GET
Path: /energy/v1/renewables/current/{country?}{neighbours=bool?}

or the alternate form which should work in the same way as the above:

Path: /energy/v1/renewables/current/country/bool
```
There are two optional parameters that can be used to filter the results:
- `country`: The country for which the current percentage of renewables should be returned. The country must either be
  It can be omitted completely, be written as a three-letter country code, or as the full name of the country. If no
  country is specified, the data for all countries is returned.
- `neighbours`: A boolean value indicating whether the results should include the neighbours of the specified country.
  If no value is specified, the neighbours are not included.

### - Response

The response is a JSON object containing the current percentage of renewables for the specified country, or all
countries if no country is specified. The data is returned as an array of objects, each containing the following
fields:

The responese differs depending on whether the country parameter is specified or not. If no country is specified, the
api will return all countries in the dataset, and the neighbours field will be omitted.

Example request: http://localhost:8080/energy/v1/renewables/current/nor

Will return the current percentage of renewables for Norway.

***Note:*** As Türkiye is recently changed it name from Turkey, the dataset is not updated to reflect this change. Therefore
it is needed to use the old name for Türkiye, which is Turkey, in order to get the correct data for this country.  
*Altenatively:* the three-letter country code can be used, which is TUR.

```
[
    {
        "name": "Norway",
        "isoCode": "NOR",
        "year": 2021,
        "percentage": 71.558365
    }
]
```   

- Example request: http://localhost:8080/energy/v1/renewables/current/nor?neighbours=true
- alternative form: http://localhost:8080/energy/v1/renewables/current/nor/true

Both examples will return the current percentage of renewables for Norway and its bordering countries.
```
[
    {
        "name": "Norway",
        "isoCode": "NOR",
        "year": 2021,
        "percentage": 71.558365
    },
    {
        "name": "Finland",
        "isoCode": "FIN",
        "year": 2021,
        "percentage": 41.666668
    },
    {
        "name": "Russia",
        "isoCode": "RUS",
        "year": 2021,
        "percentage": 17.857143
    },
    {
        "name": "Sweden",
        "isoCode": "SWE",
        "year": 2021,
        "percentage": 57.142857
    }
]
```
---
Example request: http://localhost:8080/energy/v1/renewables/current/

Will return the current percentage for all countries in dataset.
```
[
{
"name": "Algeria",
"isoCode": "DZA",
"year": 2021,
"percentage": 0.26136735
},
{
"name": "Argentina",
"isoCode": "ARG",
"year": 2021,
"percentage": 11.329249
},
...
{
"name": "Vietnam",
"isoCode": "VNM",
"year": 2021,
"percentage": 22.734407
}
]
```






##  - Historical percentage of renewables endpoint

---
### Description

The historical percentage of renewables endpoint focuses on returning historical percentages of renewables in the 
energy mix, including individual levels, as well as mean values for individual or selections of countries.

### - Request
```
Method: GET
Path: /energy/v1/renewables/history/{country?}{?begin=year&end=year?}{?sortByValue=bool?}

or the alternate form which should work in the same way as the above:

Path: /energy/v1/renewables/history/country/begin/end/bool
```
There are three optional parameters that can be used to filter the results:
- `country`: The country for which the historical data should be returned. The country must either be omitted completely
  or be entered as a three-letter country code, no more, no less. If no country is specified, the data for all
  countries is returned.
- `begin`: The year from which the historical data should be returned. If no begin year is specified, the data from the
  earliest year is returned.
- `end`: The year until which the historical data should be returned. If no end year is specified, the data until the
  latest year is returned.
- `sortByValue`: A boolean value indicating whether the results should be sorted by the percentage of renewables in the
  energy mix. If no value is specified, the sorting order is as it was in the original data set.

### - Response

The response is a JSON object containing the historical data for the specified country, or all countries if no country
is specified. The data is returned as an array of objects, each containing the following fields:

The response differs depending on whether the country parameter is specified or not. If no country is specified, the
year field is omitted, and the percentage field will display the mean percentage of renewables in the energy mix for
all countries in the specified time period.


Example request: http://localhost:8080/energy/v1/renewables/history/?country=nor

Will return every year that Norway has data for in the dataset.

Response:
```
[
    {
        "name": "Norway",
        "isoCode": "NOR",
        "year": 1965,
        "percentage": 67.87996
    },
    {
        "name": "Norway",
        "isoCode": "NOR",
        "year": 1966,
        "percentage": 65.3991
    },
    
    ...
    
    {
        "name": "Norway",
        "isoCode": "NOR",
        "year": 2021,
        "percentage": 71.558365
    }
]
```

Example request: http://localhost:8080/energy/v1/renewables/history/?country=nor&begin=2020

Will return every year for Norway from 2020 to the latest year there is data for in the dataset.

Response:
```
[
    {
        "name": "Norway",
        "isoCode": "NOR",
        "year": 2020,
        "percentage": 70.96306
    },
    {
        "name": "Norway",
        "isoCode": "NOR",
        "year": 2021,
        "percentage": 71.558365
    }
]
```

Example request: http://localhost:8080/energy/v1/renewables/history/?country=nor&end=1966

Will return every year for Norway from the earliest year there is data for in the dataset to 1966.

Response:
```
[
    {
        "name": "Norway",
        "isoCode": "NOR",
        "year": 1965,
        "percentage": 67.87996
    },
    {
        "name": "Norway",
        "isoCode": "NOR",
        "year": 1966,
        "percentage": 65.3991
    }
]
```

Example request: http://localhost:8080/energy/v1/renewables/history/?country=nor&begin=2000&end=2002

Will return every year for Norway from the dataset between 2000 and 2002.

Response:
```
[
    {
        "name": "Norway",
        "isoCode": "NOR",
        "year": 2000,
        "percentage": 72.39789
    },
    {
        "name": "Norway",
        "isoCode": "NOR",
        "year": 2001,
        "percentage": 67.58246
    },
    {
        "name": "Norway",
        "isoCode": "NOR",
        "year": 2002,
        "percentage": 69.30982
    }
]
```

Example request: http://localhost:8080/energy/v1/renewables/history/?country=nor&begin=2011&end=2013&sortByValue=true

Note that the sorting order is ascending, so the lowest percentage is first.
Will return every year for Norway from the dataset between 2011 and 2013, sorted by the percentage of renewables in the
energy mix.

Response:
```
[
    {
        "name": "Norway",
        "isoCode": "NOR",
        "year": 2011,
        "percentage": 66.30012
    },
    {
        "name": "Norway",
        "isoCode": "NOR",
        "year": 2013,
        "percentage": 67.50864
    },
    {
        "name": "Norway",
        "isoCode": "NOR",
        "year": 2012,
        "percentage": 70.095116
    }
]
```
***

For selecting the data for a specific year, just set the begin and end parameters to the same year.

Example request: http://localhost:8080/energy/v1/renewables/history/nor/1983/1983/

Response:
```
[
    {
        "name": "Norway",
        "isoCode": "NOR",
        "year": 1983,
        "percentage": 71.88228
    }
]
```

***

If no country is specified, the response will contain the mean percentage of renewables in the energy mix for all
countries in the specified time period.

Example request: http://localhost:8080/energy/v1/renewables/history/

Will return the mean percentage of renewables in the energy mix for all countries in the dataset.

Response:
```
[
    {
        "name": "Croatia",
        "isoCode": "HRV",
        "percentage": 20.009953
    },
    {
        "name": "Romania",
        "isoCode": "ROU",
        "percentage": 8.375739
    },
    
    ...
]
```

## - Notification endpoint

---
###  - Description
Notification is the endpoint where the user is able to register webhooks that are triggered by the service
based on specified events, often tied to a URL where the user wishes to be notified.

The user can register multiple webhooks, view all or a specific webhook and delete a webhook. The webhooks are
persistent, and therfore survive after service restart.

##### Root path
The root path of this endpoint is as follows: `/energy/v1/notifications`

##### Parameters
The enpoint has the following parameter: `{id}`

Where id is the webhook id. The id is, in some cases optional. For instance, when wanting to return all the 
registered webhooks, one can omit the id. When deleting an id, however, the id field is needed. 


##### Full path
Thus, the full path of the endpoint is: `/energy/v1/notifications/{id}`



### Registration of webhook
#### - Request
```
Method: POST
Path: /energy/v1/notifications
```
Where content type must be `application/json`

The body should contain
- The URL that is triggered upon an event (i.e. the URL the user wants to invoke)
- The country the trigger applies to (if empty -> triggered from any invocation)
- The number of invocations needed for triggering a notification.

Example of body request:
```
{
  "url": "https://localhost:8080/client/",
  "country": "NOR",
  "calls": 10
}
```

#### - Response

The response given back to a user POST request, is the registration id of the webhook.
This id may later be used to view this specific webhooks or delete it from the database.
The id of the webhooks is computed via a hashing utility, which takes into account the 
webhooks URL, country, and calls. When registering a webhook, the service will make sure
no equal webhooks can be registered.

- Content type: `application/json`
- Status codes
  - `201 Created`: Webhook was created successfully.
  - `204 No content`: Given if none of the other status codes are given.
  - `400 Bad request`: Something was wrong with the request body.
  - `500 Internal server error`: Something unexpected happened on the server side
- The webhook id is also returned.

Example of body response:

```
{
  "webhook_id": "103c1190ed8ea0fa3"
}
```

### Deletion of webhook

#### - Request
```
Method: DELETE
Path: /energy/v1/notifications/{id}
```
- where `{id}` is the id of the webhooks returned during webhook registration.

#### - Response

- Status codes:
  - `200 OK`: Webhook was successfully deleted
  - `204 No content`: Database is empty
  - `304 Not Modified`: If no webhook with the given id was found
  - `400 Bad request`: The id does not exist in the database
  - `500 Internal server error`: Something unexpected happened on the server side

### View single registered webhook
#### - Request
```
Method: GET
Path: /energy/v1/notifications/{id}
```
- where `{id}` is the id of the webhooks returned during webhook registration.

#### - Response

- Content type : `application/json`
- Status codes:
  - `200 OK`: Webhook was successfully fetched
  - `204 No content`: Database is empty
  - `400 Bad request`: The id is not found in the database
  - `500 Internal server error`: Something unexpected happened on the server side

Example of body response:
```
{
  "webhook_id": "103c1190ed8ea0fa3",
  "url": "https://localhost:8080/client/",
  "country": "NOR", 
  "calls": 5
}
```


### View all registered webhooks
#### - Request
```
Method: GET
Path: /energy/v1/notifications/
```
#### - Response

- Content type : `application/json`
- Status codes:
  - `200 OK`: All webhooks were fetched
  - `204 No content`: Database is empty
  - `500 Internal server error`: Something unexpected happened on the server side


Example of body response:

```
[
   {
      "webhook_id": "103c1190ed8ea0fa3",
      "url": "https://localhost:8080/client/",
      "country": "NOR",
      "calls": 5
   },
   {
      "webhook_id": "0cef6556124ca6af3f",
      "url": "https://localhost:8081/anotherClient/",
      "country": "FIN",
      "calls": 10
   },
   ...
]
```


## Webhook invocation

---
Webhooks will notify the user when it has been called the same amount as specified in `calls`.
Upon webhook invocation (the correct amount of times), the server will send notification information as follows:

```
Method: POST
Path: <URL specified by user in webhooks registration>
```
- Content type : `application/json`

Example body:
```
{
  "webhook_id": "103c1190ed8ea0fa3",
  "country": "NOR", 
  "calls": 5
}
```

## - Status endpoint

---
The status endpoint provides information on whether the service is up and running or not.
It does this by checking the status codes received from the different service endpoints 
we have defined above. It will, in addition, provide the number of webhooks currently registered
in the database.

The root path of this endpoint is: `energy/v1/status/`.

### - Request
```
Method: GET
Path: energy/v1/status
```

### - Response



- Example of body response:
```
{
    "countries_api": "200 OK",
    "markdown_html_api": "200 OK",
    "notification_db": "200 OK",
    "webhooks": 6,
    "version": "v1",
    "uptime": 86
}
```
Where:
- `countries_api`: status of the third party country API
- `markdown_html_api`: status of the markdown API
- `notification_db`: status of the webhooks database 
- `webhooks`: the amount of webhooks in the database
- `uptime`: the service uptime in seconds

## Deployment of application
___
This application is deployed on the following URL: 

- http://10.212.170.218:8080/ 

This is runs in docker on a VM in SkyHigh. This is only available on the 
NTNU internal network, and can be required with VPN or being on campus.

If you wish to run this application locally, you can create a Docker container or just run the source code. 


### Service Compilation requirements

There are some files needed for the application to compile and work.

#### Firestore service account key

Name of key: `assignment-2-key.json`

Location: `/assignment-2/assignment-2-key.json`

This is a key used to connect with Google Firestore, and has to be manually put into the project for
the application to compile and run. This can be acquired from the Google Firestore Console.

#### Hashing Secret

Filename: `hashingSecret.go`

Location: `/assignment-2/utils/hashing-utility/hashingSecret.go`

Where the content of the file is to be: 

```
package hashing_utility


var secret = []byte("WRITE-YOUR-SECRET-HERE")

func getSecret() []byte {
	return secret
}
```

### Docker deployment
___

To deploy the application:
1. Clone the repository
2. Meet the requirements of complation
3. You have to install these following packages on your computer:
   4. `docker.io`
   5. `docker-compose`
6. Then open the terminal and navigate to the `/assignment-2/` directory.
7. Run this command to create and run Docker-container
   8. `docker compose up -d`


The container is now up and running on port `8080` on your local computer. 

## Known issues
___

- The application could be tested better various places in the code, such as in the Firestore-database.
- The use of stubbing in testers should be more of a priority, as we take data straight from the API. (Noted that this application do not use large parts of data, which comprehensively improves the testing more than using the whole API).
- Error-handling could be more comprehensive, and the use of constant-error-messages could be more comprehensively used.
- In out database there is an empty struct, that we were unable to delete. 