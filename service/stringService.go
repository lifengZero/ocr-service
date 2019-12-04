package service

//接口
type StringService interface {
	Uppercase(string) (string, error)
	Count(string) int
}
