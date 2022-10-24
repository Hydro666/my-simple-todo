package internal

type App struct {
	*TaskDB
}

func NewApp() (*App, error) {
	t, err := NewTaskDB()
	if err != nil {
		return nil, err
	}
	return &App{t}, nil
}
