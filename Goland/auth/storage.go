package auth


type MyStorage struct {
	appid string
	pwd   string
}

func (s *MyStorage) getStorage(appid string) string {
	return "123"
}
