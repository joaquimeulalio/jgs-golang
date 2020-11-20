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

