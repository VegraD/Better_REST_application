# assignment-2

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

## - File structure
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
+---.idea
|   |   .gitignore
|   |   assignment-2.iml
|   |   misc.xml
|   |   modules.xml
|   |   vcs.xml
|   |   workspace.xml
|   |
|   \---inspectionProfiles
|           Project_Default.xml
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
## Service Compilation requirements

---

There are some files needed for the application to compile and work.

#### Firestore service account key

Name of key: `assignment-2-key.json`

Location: `/assignment-2/assignment-2-key.json`

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




