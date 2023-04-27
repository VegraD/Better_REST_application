package database

import (
	"assignment-2/structs"
	hashing_utility "assignment-2/utils/hashing-utility"
	"reflect"
	"testing"
)

func TestAddWebhook(t *testing.T) {
	type args struct {
		url     string
		country string
		noCalls int
	}
	tests := []struct {
		name            string
		args            args
		wantedWebhookID string
		wantedErr       error
	}{
		{
			name: "Test to add webhook",
			args: args{
				url:     "http://testURL1.com",
				country: "Norway",
				noCalls: 5,
			},
			wantedWebhookID: hashing_utility.HashingTheWebhook("http://testURL1.com", "Norway", 5),
			wantedErr:       nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AddWebhook(tt.args.url, tt.args.country, tt.args.noCalls)
			if err != nil && err.Error() != tt.wantedErr.Error() {
				t.Errorf("AddWebhook() error = %v, wantErr %v", err, tt.wantedErr)
				return
			}
			if got != tt.wantedWebhookID {
				t.Errorf("AddWebhook() got = %v, want %v", got, tt.wantedWebhookID)
			}
		})
	}
}

func TestClearDB(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ClearDB()
		})
	}
}

/*
func TestDatabaseSetup(t *testing.T) {
	tests := []struct {
		name string
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TestDatabaseSetup(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TestDatabaseSetup() = %v, want %v", got, tt.want)
			}
		})
	}
}

*/

func TestDeletionOfWebhook(t *testing.T) {
	type args struct {
		webhookID string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeletionOfWebhook(tt.args.webhookID); (err != nil) != tt.wantErr {
				t.Errorf("DeletionOfWebhook() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetAllWebhooks(t *testing.T) {
	tests := []struct {
		name    string
		want    []structs.RegisteredWebhook
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAllWebhooks()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllWebhooks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllWebhooks() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAndDisplayWebhook(t *testing.T) {
	type args struct {
		webhookID string
	}
	tests := []struct {
		name    string
		args    args
		want    structs.RegisteredWebhook
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAndDisplayWebhook(tt.args.webhookID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAndDisplayWebhook() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAndDisplayWebhook() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetWebhookAmount(t *testing.T) {
	tests := []struct {
		name    string
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetWebhookAmount()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetWebhookAmount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetWebhookAmount() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateWebhooks(t *testing.T) {
	type args struct {
		url     string
		country string
		noCalls int
		count   int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UpdateWebhooks(tt.args.url, tt.args.country, tt.args.noCalls, tt.args.count)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateWebhooks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UpdateWebhooks() got = %v, want %v", got, tt.want)
			}
		})
	}
}
