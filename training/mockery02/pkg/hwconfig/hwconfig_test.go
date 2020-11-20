package hwconfig

import (
	"testing"

	mockInterface "./mocks"
	"github.com/stretchr/testify/assert"
)

func TestRunHwconfigMockNo(t *testing.T) {

	dt2 := DatabaseTable2{"def", 2}
	r := RunHwconfig(dt2)
	assert.Equal(t, 2006, r)

}

func TestRunHwconfigMockYes(t *testing.T) {

	mock2 := mockInterface.Service{}
	mock2.On("GetCount", 4).Return(777)

	r1 := RunHwconfig(&mock2)
	assert.Equal(t, 777, r1)

}
