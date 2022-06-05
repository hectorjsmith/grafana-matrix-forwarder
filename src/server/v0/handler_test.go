package v0

import (
	"net/url"
	"reflect"
	"testing"
)

func Test_getRoomIDsFromURL(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name:    "GIVEN url with no room id WHEN get room ids THEN error returned",
			args:    args{url: "http://localhost/"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "GIVEN url with a single room id WHEN get room ids THEN array with one room id returned",
			args:    args{url: "http://localhost/?roomId=test"},
			want:    []string{"test"},
			wantErr: false,
		},
		{
			name:    "GIVEN url with a multiple room ids WHEN get room ids THEN array of all room ids returned",
			args:    args{url: "http://localhost/?roomId=test1&roomId=test2&somethingElse=test3"},
			want:    []string{"test1", "test2"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputUrl, err := url.Parse(tt.args.url)
			if err != nil {
				t.Fatalf("Invalid test data - not a valid url")
			}
			got, err := getRoomIDsFromURL(inputUrl)
			if (err != nil) != tt.wantErr {
				t.Errorf("getRoomIDsFromURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getRoomIDsFromURL() got = %v, want %v", got, tt.want)
			}
		})
	}
}
