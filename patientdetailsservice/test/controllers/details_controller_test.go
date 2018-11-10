package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	controllers "github.com/frankiennamdi/go-services-with-istio/patientdetailsservice/controllers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDetailsHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/patient/details/5", nil)
	require.NoError(t, err)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.Details)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, rr.Code, http.StatusOK, "handler returned wrong status code: got %v want %v",
		rr.Code, http.StatusOK)
	expected := `{"patientName": "Franklin Ejoh","patientAge": 5,"patientInsurance": "Blue Cross"}`
	same, _ := AreEqualJSON(rr.Body.String(), expected)
	assert.Equal(t, same, true, "handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
}

func AreEqualJSON(s1, s2 string) (bool, error) {
	var o1 interface{}
	var o2 interface{}

	var err error
	err = json.Unmarshal([]byte(s1), &o1)
	if err != nil {
		return false, fmt.Errorf("Error mashalling string 1 :: %s", err.Error())
	}
	err = json.Unmarshal([]byte(s2), &o2)
	if err != nil {
		return false, fmt.Errorf("Error mashalling string 2 :: %s", err.Error())
	}
	return reflect.DeepEqual(o1, o2), nil
}
