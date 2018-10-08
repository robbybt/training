package main

//Logging will save all detail log
type Logging struct {
	Detail map[string]interface{}
}

//GetLogging get new instance
func GetLogging() *Logging {
	return &Logging{
		Detail: make(map[string]interface{}, 0),
	}
}

//Save will save the detail on map[string]interface{}
func (l *Logging) Save(funcName string, detail interface{}) {
	l.Detail[funcName] = detail
}

//GetDetail will pass the detail which was saved
func (l *Logging) GetDetail() map[string]interface{} {
	return l.Detail
}
