package user

import (
	"example/mock-library/mocks/gomock"
	mocks "example/mock-library/mocks/mockery"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestUserWithGoMock(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockActivity := mocks_gomocks.NewMockActivity(mockCtrl)
	testUser := &User{
		Activity: mockActivity,
	}

	// Expect Do to be called once with 1 and "Hello World" as parameters, and return nil from the mocked call.
	mockActivity.EXPECT().DoSomething(1, "Hello").Return(nil).Times(1)

	err := testUser.Use()
	if err != nil {
		return
	}
}

func TestUserWithTestifyMock(t *testing.T) {
	mockActivity := &mocks.Activity{}

	testUser := &User{Activity: mockActivity}

	// Expect Do to be called once with 1 and "Hello World" as parameters, and return nil from the mocked call.
	mockActivity.On("DoSomething", 1, "Hello World").Return(nil).Once()

	testUser.Use()
	mockActivity.AssertExpectations(t)
}
