# assignment-2



## - Historical percentage of renewables endpoint

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

***Note:*** The endpoint is case-insensitive, so the country parameter can be entered in any case.

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

