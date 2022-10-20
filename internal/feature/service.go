package feature

func NewService(store Store) Service {
	return Service{store: store}
}

type Service struct {
	store Store
}
