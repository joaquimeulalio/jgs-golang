package hwconfig

import (
	"fmt"
)

// DatabaseTable1 DatabaseTable1
type DatabaseTable1 struct {
	Name  string
	Count int
}

// GetDatabaseCount GetDatabaseCount
func (dt *DatabaseTable1) GetDatabaseCount(delta int) int {
	return 1000 + dt.Count + delta
}

// DatabaseTable2 DatabaseTable1
type DatabaseTable2 struct {
	Name  string
	Count int
}

// GetDatabaseCount GetDatabaseCount
func (dt DatabaseTable2) GetDatabaseCount(delta int) int {
	return 2000 + dt.Count + delta
}

// Service interface
type Service interface {
	GetDatabaseCount(delta int) int
}

// RunHwconfig RunHwconfig Method 1: Service interface as parameter
func RunHwconfig(i Service) int {
	fmt.Println("RunHwconfig() Method 1: Service interface as parameter")

	n := i.GetDatabaseCount(111)
	fmt.Printf("n=%d\n", n)
	return n
}

// VerifierServer VerifierServer
type VerifierServer struct {
	HwconfigService Service
}

// RunHwconfig RunHwconfig Method 2: Struct receiver containing service interface inside it
func (vf *VerifierServer) RunHwconfig() int {
	fmt.Println("RunHwconfig() Method 2: Struct receiver containing service interface inside it")

	n := vf.HwconfigService.GetDatabaseCount(222)
	fmt.Printf("n=%d\n", n)
	return n
}
