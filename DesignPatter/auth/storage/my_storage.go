package storage


type MyStorage struct {
	appid string
	pwd   string
}

func (s *MyStorage) GetStorage(appid string) string {
	return "123"
}
