package handler

import (
	"log"
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
	fallback := fallbackTest(t)

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
	t.Run("Test yaml handler", func(t *testing.T) {

		ymlData := []byte(`
- path: /google
  url: https://www.google.com
- path: /github
  url: https://www.github.com
- path: /stackoverflow
  url: https://www.stackoverflow.com
`)
		// ymlData = ReadFile("sample.yml")
		fallback := fallbackTest(t)

		handler, err := YAMLHandler(ymlData, fallback)

		if err != nil {
			t.Errorf("%v", err)
		}
		req, err := http.NewRequest("GET", "/google", nil)

		if err != nil {
			t.Errorf("%v", err)
		}

		rr := httptest.NewRecorder()

		handler(rr, req)

			// Check the status code
		if status := rr.Code; status != http.StatusMovedPermanently {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusMovedPermanently)
		}

		// Check the redirect location
		if location := rr.Header().Get("Location"); location != "https://www.google.com" {
			t.Errorf("handler returned wrong location: got %v want %v", location, "https://www.google.com")
		}
	})

}

func Test_ReadFile(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		test_name string
		args      args
		wantData  []byte
	}{
		// TODO: Add test cases.
		{
			test_name: "Reading a file",
			args: args{
				name: "sample.yml",
			},
			wantData: []byte(
`- path: /google
  url: https://www.google.com
- path: /github
  url: https://www.github.com
- path: /stackoverflow
  url: https://www.stackoverflow.com`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.test_name, func(t *testing.T) {
			if gotData := ReadFile(tt.args.name); !reflect.DeepEqual(gotData, tt.wantData) {

				t.Errorf("reafFile() = %v, want %v", string(gotData), string(tt.wantData))
			}
		})
	}
}

func TestJsonHandler(t *testing.T) {
	t.Run("Test JSON Handler", func(t *testing.T) {

		jsonData := []byte(
			`[
					{"Path":"/test-path1", "URL":"http://example1.com"},
					{"Path":"/test-path2", "URL":"http://example2.com"}
			]`,
	)

		fallback := fallbackTest(t)

		handler, err := JsonHandler(jsonData, fallback)

		if err != nil {
			log.Fatal(err)
		}
		req, err := http.NewRequest("GET", "/test-path1", nil)

		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		handler(rr, req)

			// Check the status code
	if status := rr.Code; status != http.StatusMovedPermanently {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusMovedPermanently)
	}

	// Check the redirect location
	if location := rr.Header().Get("Location"); location != "http://example1.com" {
		t.Errorf("handler returned wrong location: got %v want %v", location, "http://example1.com")
	}
	})
}

func fallbackTest(t *testing.T) http.HandlerFunc{
	t.Helper()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Fallback", http.StatusNotFound)
	})
}
