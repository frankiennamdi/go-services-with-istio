package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/frankiennamdi/go-services-with-istio/patientservice/config"

	"github.com/frankiennamdi/go-services-with-istio/patientservice/controllers"
	"github.com/frankiennamdi/go-services-with-istio/patientservice/models"
	"github.com/gorilla/mux"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDetailsHandler(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	testPatientDetail := models.PatientDetail{PatientName: "Franklin Ejoh", PatientAge: 5,
		PatientInsurance: "Blue Cross"}

	patientDetailServiceUrl := "http://patient-details-service:3000"

	expectedPatientDetail, err := json.Marshal(testPatientDetail)
	expectedPatient, err := json.Marshal(models.PatientInfo{PatientDetail: testPatientDetail})

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/patient/%d", patientDetailServiceUrl, 2),
		httpmock.NewStringResponder(200, string(expectedPatientDetail)))

	req, err := http.NewRequest("GET", "/data-service/patients/2", nil)
	req = mux.SetURLVars(req, map[string]string{
		"id": "2",
	})
	require.NoError(t, err)
	rr := httptest.NewRecorder()

	config := &config.Config{PatientDetailServiceUrl: patientDetailServiceUrl}
	httpHandler := controllers.Handler{Config: config, Handle: controllers.Details}
	httpHandler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusOK, "handler returned wrong status code: got %v want %v",
		rr.Code, http.StatusOK)

	expected := string(expectedPatient)
	same, _ := AreEqualJSON(rr.Body.String(), string(expectedPatient))
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
