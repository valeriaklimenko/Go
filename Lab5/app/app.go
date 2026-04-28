package app

type DB interface {
	CreateTable() error
	Insert(name string) error
	GetAll() ([]string, error)
	DeleteAll() error
}

type App struct {
	db DB
}

func New(db DB) *App {
	return &App{db: db}
}

func (a *App) Run(name string) ([]string, error) {
	if err := a.db.CreateTable(); err != nil {
		return nil, err
	}

	if name != "" {
		if err := a.db.Insert(name); err != nil {
			return nil, err
		}
	}

	users, err := a.db.GetAll()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (a *App) Clear() error {
	return a.db.DeleteAll()
}
