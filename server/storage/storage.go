package storage

type Storage struct {
	data map[string]string
}

func CreateStorage() *Storage {
	return &Storage{make(map[string]string)}
}

func (s *Storage) Save(key, value string) (string, error) {
	replaced := s.data[key]
	s.data[key] = value
	return replaced, nil
}

func (s *Storage) Delete(key string) (string, error) {
	removed := s.data[key]
	delete(s.data, key)
	return removed, nil
}
