package currentHandler

import (
	"assignment-2/constants"
	"assignment-2/structs"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

var currentEndpoint *httptest.Server

/*
Test function for current endpoint.
*/
func TestMain(m *testing.M) {

	server := httptest.NewServer(http.HandlerFunc(CurrentHandler))
	defer server.Close()

	//borders, err := utils.GetBordersFromJson()
	//records, err := utils.GetCountriesFromCsv()
	//if err != nil {
	//	return
	//}

	//client := http.Client{}

	// URL under which server is instantiated
	fmt.Println("URL: ", server.URL+constants.CurrentEP)

	//res, err := client.Get(server.URL + constants.CurrentEP)
	//if err != nil {
	//	t.Fatal("Get request to URL failed:", err.Error())
	//}

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
		excpectedResponse  structs.CountryInfo
	}{
		{
			name: "Valid Get Request to Current Endpoint with Country as ISO Code",
			args: args{
				url: currentEndpoint.URL + constants.CurrentEP + "/NOR",
			},
			method:             http.MethodGet,
			expectedStatusCode: http.StatusOK,
			excpectedResponse: structs.CountryInfo{
				Country:    "Norway",
				IsoCode:    "NOR",
				Year:       2021,
				Percentage: 71.558365,
			},
		},
		{
			name: "Valid Get Request to Current Endpoint with Country as case-mixed ISO Code",
			args: args{
				url: currentEndpoint.URL + constants.CurrentEP + "/NoR",
			},
			method:             http.MethodGet,
			expectedStatusCode: http.StatusOK,
			excpectedResponse: structs.CountryInfo{
				Country:    "Norway",
				IsoCode:    "NOR",
				Year:       2021,
				Percentage: 71.558365,
			},
		}, {
			name: "Valid Get Request to Current Endpoint with full country name",
			args: args{
				url: currentEndpoint.URL + constants.CurrentEP + "/Norway",
			},
			method:             http.MethodGet,
			expectedStatusCode: http.StatusOK,
			excpectedResponse: structs.CountryInfo{
				Country:    "Norway",
				IsoCode:    "NOR",
				Year:       2021,
				Percentage: 71.558365,
			},
		}, {
			name: "Valid Get Request to Current Endpoint with case-mixed full country name",
			args: args{
				url: currentEndpoint.URL + constants.CurrentEP + "/NoRWaY",
			},
			method:             http.MethodGet,
			expectedStatusCode: http.StatusOK,
			excpectedResponse: structs.CountryInfo{
				Country:    "Norway",
				IsoCode:    "NOR",
				Year:       2021,
				Percentage: 71.558365,
			},
		}, {
			name: "Invalid Post Request",
			args: args{
				url: currentEndpoint.URL + "/corona/v1/cases/Norway",
			},
			method:             http.MethodPost,
			expectedStatusCode: http.StatusMethodNotAllowed,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := http.Client{}
			req, err := http.NewRequest(tt.method, tt.args.url, nil)
			if err != nil {
				t.Fatal("Failed to create request:", err.Error())
			}
			res, err := client.Do(req)
			if err != nil {
				t.Fatal("Get request to URL failed:", err.Error())
			}
			if res.StatusCode != tt.expectedStatusCode {
				t.Errorf("Expected status code %v, got %v", tt.expectedStatusCode, res.StatusCode)
			}
			if res.StatusCode == http.StatusOK {
				var actualResponse structs.CountryInfo
				err = json.NewDecoder(res.Body).Decode(&actualResponse)
				if err != nil {
					t.Fatal("Failed to decode response:", err.Error())
				}
				if !reflect.DeepEqual(actualResponse, tt.excpectedResponse) {
					t.Errorf("Expected response %v, got %v", tt.excpectedResponse, actualResponse)
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
				url: currentEndpoint.URL + constants.CurrentEP + "/NOR?neighbours=true",
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
				url: currentEndpoint.URL + constants.CurrentEP + "/NoR?neighbours=true",
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
				url: currentEndpoint.URL + constants.CurrentEP + "/NOR/true",
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
				url: currentEndpoint.URL + constants.CurrentEP + "/NOR/true",
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
				url: currentEndpoint.URL + constants.CurrentEP + "/NORwaY/true",
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
				url: currentEndpoint.URL + constants.CurrentEP + "/NoRWaY?neighbours=trUe",
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
			defer res.Body.Close()
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
