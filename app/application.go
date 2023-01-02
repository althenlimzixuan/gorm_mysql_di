package app

type AppService struct {
}

type AppServiceInt interface {
}

func ProvideAppService() *AppService {
	return &AppService{}
}

func (svc *AppService) Initialize() {

}

func (svc *AppService) StartApp() {

}

func (svc *AppService) StopApp() {

}
