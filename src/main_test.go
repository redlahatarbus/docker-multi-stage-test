package main

import "io/ioutil"
import "testing"
import "net/http"
import "net/http/httptest"

func TestGetCars(t *testing.T) {
	rw := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/cars", nil)

	GetCars(rw, r)

	resp := rw.Result()

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected http status %v does not match actual http status %v", http.StatusOK, resp.StatusCode)
	} else {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		bodyString := string(bodyBytes)

		if bodyString == "" {
			t.Errorf("body should not be empty")
		}
	}
}
