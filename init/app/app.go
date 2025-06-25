package app

type App struct {
	config *config.Config
}

// 모든 Layer 간 연결이 목적
func NewApp(config *config.Config) *App {
	// 외부 값을 바로 할당하는 경우 코드가 이뻐
	a := &App{
		config: config,
	}

	return a 
}