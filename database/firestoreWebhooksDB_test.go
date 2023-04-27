package database

import (
	hashing_utility "assignment-2/utils/hashing-utility"
	"testing"
)

var webhookIDs []string

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

func TestDeletionOfWebhook(t *testing.T) {
	type args struct {
		webhookID string
	}
	tests := []struct {
		name      string
		args      args
		wantedErr error
	}{
		{
			name: "Test to delete a webhook",
			args: args{
				webhookID: webhookIDs[1],
			},
			wantedErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := DeletionOfWebhook(tt.args.webhookID)
			if err != nil && err.Error() != tt.wantedErr.Error() {
				t.Errorf("DeletionOfWebhook() error = %v, wantErr %v", err, tt.wantedErr)
			}
		})
	}
}

/*
func TestGetAndDisplayWebhook(t *testing.T) {
	type args struct {
		webhookID string
	}
	tests := []struct {
		name            string
		args    	    args
		wantedWebhook   structs.RegisteredWebhook
		wantedErr       error
	}{
		{
			name: "",
			args: args{
				webhookID: webhookIDs[3],
			},
			wantedWebhook: ,

		}
	},
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAndDisplayWebhook(tt.args.webhookID)
			if err != nil && err.Error() != tt.wantedErr.Error() {
				t.Errorf("GetAndDisplayWebhook() error = %v, wantErr %v", err, tt.wantedErr)
				return
			}
			if !reflect.DeepEqual(got, tt.wantedWebhook) {
				t.Errorf("GetAndDisplayWebhook() got = %v, want %v", got, tt.wantedWebhook)
			}
		})
	}
}
*/

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
