package ini

type Ini struct{}

func (i *Ini) Marshal(v interface{}) ([]byte, error) {
	return nil, nil
}

func (i *Ini) Unmarshal(data []byte, v interface{}) error {
	return nil
}
