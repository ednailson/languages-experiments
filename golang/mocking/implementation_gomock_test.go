package mocking

import (
	"errors"
	"github.com/ednailson/languages-experiments/golang/mocking/gomock_mocks"
	gomocklib "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestMethodAGoMock(t *testing.T) {
	// Initiating a new mock
	ctrl := gomocklib.NewController(t)
	mock := gomock.NewMockStore(ctrl)

	// Setting up the functions that will be called on MethodA

	now := time.Now().UTC()
	// Here we are setting up the first called method that it is Open
	// now will be matched in this case, if MethodA does not pass now as a parameter the test will fail
	// Open will return nil in this scenario
	mock.EXPECT().Open(gomocklib.Eq(now)).Return(nil)

	// Here we are setting up the second called method that is Sell
	// in this case, we will check the passed parameters if it is "a" and 2.
	// Sell will return nil in this scenario
	mock.EXPECT().Sell("a", 2).Return(nil)

	// Here we are setting up the second called method that is Close
	// in this case, we do not want to check the passed parameters, so we use mock.Anything
	// Close will return an error in this scenario
	mock.EXPECT().Close(gomocklib.Any()).Return(errors.New("error test"))

	sut := NewImplementation(mock)

	// Close method returns error (we set up this on line 33) so MethodA will return error
	err := sut.MethodA(now)

	assert.Error(t, err)
}
