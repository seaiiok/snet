package pkg

import (
	"os"
)

type snet struct {
	Conf map[string]string
}

func (s *snet) GetConfig() map[string]string {
	return s.GetConfig()
}

// func (s *service) NewServer() {

// }

// func (s *service) NewClient() {

// }

type server struct {
}

func isNotExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return false
		}
		return true
	}
	return false
}

// 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// 判断所给路径是否为文件
func IsFile(path string) bool {
	return !IsDir(path)
}
