package service

import (
	"Notification/internal/core/model"
	"Notification/mocks"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNoficationService_SendEmail(t *testing.T) {
	mockNotfi := new(mocks.NotfiService)

	user := &model.SendMessageEmail{
		Email:   "amir.b@ba.com",
		Message: "Create Email",
	}

	jsonUser, _ := json.Marshal(user)

	mockNotfi.On("SendEmail", string(jsonUser)).Return(true, nil)

	result, err := mockNotfi.SendEmail(string(jsonUser))

	assert.NoError(t, err)

	assert.Equal(t, true, result)

	mockNotfi.AssertExpectations(t)

	if result {
		assert.True(t, true)
	} else {
		assert.Error(t, err)
	}

}
