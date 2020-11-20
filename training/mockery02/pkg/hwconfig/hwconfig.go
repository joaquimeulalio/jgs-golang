package hwconfig

import (
	"fmt"
)

// DatabaseTable1 DatabaseTable1
type DatabaseTable1 struct {
	Name  string
	Count int
}

// GetCount GetCount
func (dt *DatabaseTable1) GetCount(delta int) int {
	return 1000 + dt.Count + delta
}

// DatabaseTable2 DatabaseTable1
type DatabaseTable2 struct {
	Name  string
	Count int
}

// GetCount GetCount
func (dt DatabaseTable2) GetCount(delta int) int {
	return 2000 + dt.Count + delta
}

// Service interface
type Service interface {
	GetCount(delta int) int
}

// RunHwconfig RunHwconfig
func RunHwconfig(i Service) int {
	fmt.Println("RunHwconfig()")

	n := i.GetCount(4)
	fmt.Printf("n=%d\n", n)
	return n
}
