package domain

import(
	"testing"
	"github.com/stretchr/testify/assert"
	"net/http"
)

func TestGetUserNoUserFound(t *testing.T){

	user, err := GetUser(0)

	assert.Nil(t, user,"We were not expecting this when user id is 0")
	assert.NotNil(t,err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)

	if user != nil{
		t.Error("We were not expecting this when user id is 0")
	}

	if err == nil{
		t.Error("We were expecting an error when user id is 0")
	}
}

func TestGetUserNoError(t *testing.T){

	user, err := GetUser(123)

	assert.Nil(t, err, "We were not expecting the error ")
	assert.NotNil(t, user)
}