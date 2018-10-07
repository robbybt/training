package main

type logging struct {
	Detail map[string]interface{}
}

func GetLogging() *logging {
	return &logging{
		Detail: make(map[string]interface{}, 0),
	}
}

func (l *logging) Save(funcName string, detail interface{}) {
	l.Detail[funcName] = detail
}

func (l *logging) GetDetail() map[string]interface{} {
	return l.Detail
}
