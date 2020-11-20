package hwconfig

import (
	"testing"

	mockInterface "./mocks"
	"github.com/stretchr/testify/assert"
)

func TestRunHwconfigMockNo(t *testing.T) {

	dt2 := DatabaseTable2{"def", 2}
	r := RunHwconfig(dt2)
	assert.Equal(t, 2113, r)

}

func TestRunHwconfigMockMethod1(t *testing.T) {

	mock2 := mockInterface.Service{}
	mock2.On("GetDatabaseCount", 111).Return(777)

	r1 := RunHwconfig(&mock2)
	assert.Equal(t, 777, r1)

}

func TestRunHwconfigMockMethod2(t *testing.T) {

	mock2 := mockInterface.Service{}
	mock2.On("GetDatabaseCount", 222).Return(888)

	vf := VerifierServer{}
	vf.HwconfigService = &mock2
	r2 := vf.RunHwconfig()

	assert.Equal(t, 888, r2)

}
