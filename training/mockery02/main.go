package main

import (
	"fmt"

	"./pkg/hwconfig"
)

func main() {

	/*
		cannot use dt1 (type hwconfig.DatabaseTable1) as type hwconfig.Service in argument to hwconfig.RunHwconfigv2:
		hwconfig.DatabaseTable1 does not implement hwconfig.Service (GetCount method has pointer receiver

		dt1 := hwconfig.DatabaseTable1{"abc", 1}
		n1 := hwconfig.RunHwconfig(dt1)
		fmt.Printf("n1=%d\n", n1)
	*/
	dt1 := hwconfig.DatabaseTable1{"abc", 1}
	n1 := hwconfig.RunHwconfig(&dt1) // <== pass address to fix
	fmt.Printf("n1=%d\n", n1)
	/*
		_/home/joaquim/suporte/github-publico/jgs-golang/training/mockery02/pkg/hwconfig.DatabaseTable2 composite literal uses unkeyed fields go-vet
	*/
	dt2 := hwconfig.DatabaseTable2{"def", 2}
	n2 := hwconfig.RunHwconfig(dt2)
	fmt.Printf("n2=%d\n", n2)
}

/*
How to install ?
	https://github.com/vektra/mockery#github-release
	go get "github.com/stretchr/objx"
	go get github.com/vektra/mockery/v2/.../
	go install mockery ??? (nao me lembro)

How to run ?
	cd <workdir>/pkg/hwconfig
	~/go/bin/mockery  --name=Service
	"Service" é o nome da interface a ser mockada, e não do package.
	o ~/go/bin deveria estar no PATH.

Comentarios:
	- precisa ser um package diferente do main
	- o truque parece ser: o serviço (ou interface) a ser mockado deve ser passado como parametro na função a ser testada
	- no mock.On() aparece o  nome da função do serviço (ou interface) a ser mockado, ex: "mock2.On("GetCount", 4).Return(777)"
*/
