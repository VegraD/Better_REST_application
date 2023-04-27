package currentHandler

import (
	"assignment-2/constants"
	"assignment-2/database"
	"assignment-2/structs"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
)

var currentEndpoint *httptest.Server

/*
Test function for current endpoint.
*/
func TestMain(m *testing.M) {
	workDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	err2 := os.Chdir(workDir + "/../../..")
	if err2 != nil {
		return
	}
	database.InitFirestore()
	defer func() {
		err := database.CloseDB()
		if err != nil {
			log.Printf("Error in closing database: %s", err)
		}
	}()
	//database.InitFirestore()
	server := httptest.NewServer(http.HandlerFunc(CurrentHandler))
	defer server.Close()

	currentEndpoint = httptest.NewServer(http.HandlerFunc(CurrentHandler))
	defer currentEndpoint.Close()

	defer database.ClearDB()

	m.Run()

}

func TestSingleCountry(t *testing.T) {

	type args struct {
		url string
	}
	tests := []struct {
		name               string
		args               args
		method             string
		expectedStatusCode int
		expectedResponse   []structs.CountryInfo
	}{
		{
			name: "Valid Get Request to Current Endpoint with Country as ISO Code",
			args: args{
				url: currentEndpoint.URL + constants.CurrentEP + "NOR",
			},
			method:             http.MethodGet,
			expectedStatusCode: http.StatusOK,
			expectedResponse: []structs.CountryInfo{{Country: "Norway", IsoCode: "NOR", Year: 2021,
				Percentage: 71.558365}},
		},
		{
			name: "Valid Get Request to Current Endpoint with Country as case-mixed ISO Code",
			args: args{
				url: currentEndpoint.URL + constants.CurrentEP + "NoR",
			},
			method:             http.MethodGet,
			expectedStatusCode: http.StatusOK,
			expectedResponse: []structs.CountryInfo{{Country: "Norway", IsoCode: "NOR", Year: 2021,
				Percentage: 71.558365}},
		}, {
			name: "Valid Get Request to Current Endpoint with full country name",
			args: args{
				url: currentEndpoint.URL + constants.CurrentEP + "Norway",
			},
			method:             http.MethodGet,
			expectedStatusCode: http.StatusOK,
			expectedResponse: []structs.CountryInfo{{Country: "Norway", IsoCode: "NOR", Year: 2021,
				Percentage: 71.558365}},
		}, {
			name: "Valid Get Request to Current Endpoint with case-mixed full country name",
			args: args{
				url: currentEndpoint.URL + constants.CurrentEP + "NoRWaY",
			},
			method:             http.MethodGet,
			expectedStatusCode: http.StatusOK,
			expectedResponse: []structs.CountryInfo{{Country: "Norway", IsoCode: "NOR", Year: 2021,
				Percentage: 71.558365}},
		}, {
			name: "Invalid Post Request",
			args: args{
				url: currentEndpoint.URL + "/NOR",
			},
			method:             http.MethodPost,
			expectedStatusCode: 501,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(tt.method, tt.args.url, nil)
			if err != nil {
				t.Fatal(err)
			}
			res, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Fatal(err)
			}
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					t.Fatal(err)
				}
			}(res.Body)
			if res.StatusCode != tt.expectedStatusCode {
				t.Errorf("expected status code %v, got %v", tt.expectedStatusCode, res.StatusCode)
			} else if res.StatusCode == http.StatusOK {
				var got []structs.CountryInfo
				err = json.NewDecoder(res.Body).Decode(&got)
				if err != nil {
					t.Fatal(err)
				}
				if !reflect.DeepEqual(got, tt.expectedResponse) {
					t.Errorf("expected response %v, got %v", tt.expectedResponse, got)
				}
			}
		})
	}
}

func TestNeighbours(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name               string
		args               args
		method             string
		expectedStatusCode int
		expectedResponse   []structs.CountryInfo
	}{
		{
			name: "Valid Get Request to Current Endpoint with Country as ISO Code and neighbours set to true",
			args: args{
				url: currentEndpoint.URL + constants.CurrentEP + "NOR?neighbours=true",
			},
			method:             http.MethodGet,
			expectedStatusCode: http.StatusOK,
			expectedResponse: []structs.CountryInfo{
				{
					Country:    "Norway",
					IsoCode:    "NOR",
					Year:       2021,
					Percentage: 71.558365,
				},
				{
					Country:    "Finland",
					IsoCode:    "FIN",
					Year:       2021,
					Percentage: 34.61129,
				},
				{
					Country:    "Russia",
					IsoCode:    "RUS",
					Year:       2021,
					Percentage: 6.6202893,
				},
				{
					Country:    "Sweden",
					IsoCode:    "SWE",
					Year:       2021,
					Percentage: 50.924007,
				},
			},
		}, {
			name: "Valid Get Request to Current Endpoint with Country as mixed-case ISO Code and neighbours set " +
				"to true",
			args: args{
				url: currentEndpoint.URL + constants.CurrentEP + "NoR?neighbours=true",
			},
			method:             http.MethodGet,
			expectedStatusCode: http.StatusOK,
			expectedResponse: []structs.CountryInfo{
				{
					Country:    "Norway",
					IsoCode:    "NOR",
					Year:       2021,
					Percentage: 71.558365,
				},
				{
					Country:    "Finland",
					IsoCode:    "FIN",
					Year:       2021,
					Percentage: 34.61129,
				},
				{
					Country:    "Russia",
					IsoCode:    "RUS",
					Year:       2021,
					Percentage: 6.6202893,
				},
				{
					Country:    "Sweden",
					IsoCode:    "SWE",
					Year:       2021,
					Percentage: 50.924007,
				},
			},
		}, {
			name: "Valid Get Request to Current Endpoint with Country as ISO Code and neighbours set to true in path " +
				"format",
			args: args{
				url: currentEndpoint.URL + constants.CurrentEP + "NOR/true",
			},
			method:             http.MethodGet,
			expectedStatusCode: http.StatusOK,
			expectedResponse: []structs.CountryInfo{
				{
					Country:    "Norway",
					IsoCode:    "NOR",
					Year:       2021,
					Percentage: 71.558365,
				},
				{
					Country:    "Finland",
					IsoCode:    "FIN",
					Year:       2021,
					Percentage: 34.61129,
				},
				{
					Country:    "Russia",
					IsoCode:    "RUS",
					Year:       2021,
					Percentage: 6.6202893,
				},
				{
					Country:    "Sweden",
					IsoCode:    "SWE",
					Year:       2021,
					Percentage: 50.924007,
				},
			},
		}, {
			name: "Valid Get Request to Current Endpoint with Country as mixed-case ISO Code and neighbours set to " +
				"mixed-case true in path format",
			args: args{
				url: currentEndpoint.URL + constants.CurrentEP + "NOR/true",
			},
			method:             http.MethodGet,
			expectedStatusCode: http.StatusOK,
			expectedResponse: []structs.CountryInfo{
				{
					Country:    "Norway",
					IsoCode:    "NOR",
					Year:       2021,
					Percentage: 71.558365,
				},
				{
					Country:    "Finland",
					IsoCode:    "FIN",
					Year:       2021,
					Percentage: 34.61129,
				},
				{
					Country:    "Russia",
					IsoCode:    "RUS",
					Year:       2021,
					Percentage: 6.6202893,
				},
				{
					Country:    "Sweden",
					IsoCode:    "SWE",
					Year:       2021,
					Percentage: 50.924007,
				},
			},
		},
		{
			name: "Valid Get Request to Current Endpoint with Country as mixed-case Full Name and neighbours set to " +
				"mixed-case true in path format",
			args: args{
				url: currentEndpoint.URL + constants.CurrentEP + "NORwaY/true",
			},
			method:             http.MethodGet,
			expectedStatusCode: http.StatusOK,
			expectedResponse: []structs.CountryInfo{
				{
					Country:    "Norway",
					IsoCode:    "NOR",
					Year:       2021,
					Percentage: 71.558365,
				},
				{
					Country:    "Finland",
					IsoCode:    "FIN",
					Year:       2021,
					Percentage: 34.61129,
				},
				{
					Country:    "Russia",
					IsoCode:    "RUS",
					Year:       2021,
					Percentage: 6.6202893,
				},
				{
					Country:    "Sweden",
					IsoCode:    "SWE",
					Year:       2021,
					Percentage: 50.924007,
				},
			},
		}, {
			name: "Valid Get Request to Current Endpoint with Country as mixed-case full country name and neighbours " +
				"set to true",
			args: args{
				url: currentEndpoint.URL + constants.CurrentEP + "NoRWaY?neighbours=tRUe",
			},
			method:             http.MethodGet,
			expectedStatusCode: http.StatusOK,
			expectedResponse: []structs.CountryInfo{
				{
					Country:    "Norway",
					IsoCode:    "NOR",
					Year:       2021,
					Percentage: 71.558365,
				},
				{
					Country:    "Finland",
					IsoCode:    "FIN",
					Year:       2021,
					Percentage: 34.61129,
				},
				{
					Country:    "Russia",
					IsoCode:    "RUS",
					Year:       2021,
					Percentage: 6.6202893,
				},
				{
					Country:    "Sweden",
					IsoCode:    "SWE",
					Year:       2021,
					Percentage: 50.924007,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(tt.method, tt.args.url, nil)
			if err != nil {
				t.Fatal(err)
			}
			res, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Fatal(err)
			}
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					t.Fatal(err)
				}
			}(res.Body)
			if res.StatusCode != tt.expectedStatusCode {
				t.Errorf("expected status code %v, got %v", tt.expectedStatusCode, res.StatusCode)
			}
			var got []structs.CountryInfo
			err = json.NewDecoder(res.Body).Decode(&got)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(got, tt.expectedResponse) {
				t.Errorf("expected response %v, got %v", tt.expectedResponse, got)
			}
		})
	}
}
