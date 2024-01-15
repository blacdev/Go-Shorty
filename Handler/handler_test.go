package handler

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestMapHandler(t *testing.T) {
	// Define a pathsToUrls map for testing
	pathsToUrls := map[string]string{
		"/test-path": "http://example.com",
	}

	// Define a fallback handler for testing
	fallback := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Fallback", http.StatusNotFound)
	})

	// Create a MapHandler
	handler := MapHandler(pathsToUrls, fallback)

	// Create a request to pass to our handler
	req, err := http.NewRequest("GET", "/test-path", nil)

	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the MapHandler
	handler(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusPermanentRedirect {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusPermanentRedirect)
	}

	// Check the redirect location
	if location := rr.Header().Get("Location"); location != "http://example.com" {
		t.Errorf("handler returned wrong location: got %v want %v", location, "http://example.com")
	}
}

func TestYAMLHandler(t *testing.T) {
	type args struct {
		yml      []byte
		fallback http.Handler
	}
	tests := []struct {
		name    string
		args    args
		want    http.HandlerFunc
		wantErr bool
	}{
		// TODO: Add test cases.

}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := YAMLHandler(tt.args.yml, tt.args.fallback)
			if (err != nil) != tt.wantErr {
				t.Errorf("YAMLHandler() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("YAMLHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ReadFile(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		test_name     string
		args     args
		wantData []byte
	}{
		// TODO: Add test cases.
		{
			test_name: "Reading a file",
			args: args{
				name: "dummyfile.txt",
			},
			wantData: []byte("testing"),
	},
}
for _, tt := range tests {
	t.Run(tt.test_name, func(t *testing.T) {
		if gotData := ReadFile(tt.args.name); !reflect.DeepEqual(gotData, tt.wantData) {

				t.Errorf("reafFile() = %v, want %v", gotData, tt.wantData)
			}
		})
	}
}
