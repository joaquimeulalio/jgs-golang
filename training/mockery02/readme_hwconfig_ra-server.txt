========= REPO rmdb ========= 
rmdb/apio/service/mysql/ mysql.go :
	package mysql
	type rmdbService struct {
		connection *db.SQLHandlerConn
		Log        vflog.Logger
	}

rmdb/apio/service/mysql/ hw_config.go:
	package mysql
	func (r *rmdbService) GetHWConfigRMByDeviceID(deviceID int64) ([]byte, error) {
	}


rmdb/api/rmdb.go:
	package api (???)
	...
	type Service interface {
		...
		GetHWConfigRMByDeviceID(int64) ([]byte, error)
		...
	}


========= REPO ra-server ========= 
server.go:
	package attestation

	func (vf *VerifierServer) verifyHWConfig(
		originalSCL, err := vf.rmdb.GetHWConfigRMByDeviceID(deviceID)
	)

	// VerifierServer has the reference to the services used during an attestation
	type VerifierServer struct {
		svc               *Service
		policy            policy.Service
		device            device.Service
		profile           profile.Service
		attestationResult attResult.Service
		rmdb              rmdb.Service
		ca                ca.Service
		sessions          *sessions.CookieStore
		metrics           *attMetrics.AttestationMetrics
	}

"vf *VerifierServer" é o receiver de verifyHWConfig(). E é um concrete type em vez de *interface*. Logo verifyHWConfig *não* pode ser mockado. Este entendimento esta correto ? 
Este aqui e o codigo auto gerado pelo mockery para a interface rmdb.Service ? 
	rmdb/api/service/mocks/Service.go:283:func (_m *Service) GetHWConfigRMByDeviceID(_a0 int64) ([]byte, error) {

ra-server/cmd/rask/main.go:
	vf := attestation.NewVerifier(conf, log)

