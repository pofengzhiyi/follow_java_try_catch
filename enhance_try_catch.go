package main

import (
	"log"
)

type Exception struct {
	Id int       // exception id
	Msg string   // exception msg
}

type TryStruct struct {
	catches map[int]ExceptionHandler
	try   func()
}

//类似java的try
func Try(tryHandler func()) *TryStruct {
	tryStruct := TryStruct{
		catches: make(map[int]ExceptionHandler),
		try: tryHandler,
	}
	return &tryStruct
}


type ExceptionHandler func(Exception)

//java中的catch
func (this *TryStruct) Catch(exceptionId int, catch func(Exception)) *TryStruct {
	this.catches[exceptionId] = catch
	return this
}

//java的finally
func (this *TryStruct) Finally(finally func()) {
	defer func() {
		if e := recover(); nil != e {

			exception := e.(Exception)

			if catch, ok := this.catches[exception.Id]; ok {
				catch(exception)
			}

			finally()
		}
	}()

	this.try()
}


func Throw(id int, msg string) Exception {
	panic(Exception{id,msg})
}

func main() {

	Try(func() {
		log.Println("try...")
		//  指定了异常代码为2，错误信息为error2
		Throw(2,"error2")
	}).Catch(1, func(e Exception) {
		log.Println(e.Id,e.Msg)
	}).Catch(2, func(e Exception) {
		log.Println(e.Id,e.Msg)
	}).Finally(func() {
		log.Println("finally")
	})
}
