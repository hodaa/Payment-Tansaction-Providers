package main

import "net/http"
import "net/http/httptest"
import "testing"
import "encoding/json"

func TestGetEntryByID(t *testing.T) {

	req, err := http.NewRequest("GET", "api/v1/payment/transaction", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("id", "1")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(get)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	if !JSON(rr.Body.String()) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), "hoda")
	}
	// Check the response body is what we expect.
	//	expected := `{"id":1,"first_name":"Krish","last_name":"Bhanushali","email_address":"krishsb2405@gmail.com","phone_number":"0987654321"}`
	//	if rr.Body.String() != expected {
	//		t.Errorf("handler returned unexpected body: got %v want %v",
	//			rr.Body.String(), expected)
	//	}
}

func JSON(str string) bool {
	var jsonStr map[string]interface{}
	err := json.Unmarshal([]byte(str), &jsonStr)
	return err == nil

}
