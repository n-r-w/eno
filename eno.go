// Package eno - общие коды ошибок
package eno

type ErrNo int

const (
	ErrNotImplemented   = iota + ErrNo(1) // не реализовано
	ErrInternal                           // загадочная внутренняя ошибка
	ErrNoAccess                           // нет доступа
	ErrTooManyRequests                    // слишком много запросов
	ErrNotAuthenticated                   // не залогинен
	ErrObjectNotFound                     // объект не найден
	ErrObjectExist                        // объект уже существует
	ErrNotReady                           // операция не может быть выполнена, т.к. не готов сервер т.п.
	ErrBadRequest                         // некорректный запрос
	ErrDeadlineExceeded                   // закончилось отведенное время

	ErrUser ErrNo = 10000 // с какого номера начинать свои собственные сообщения
)

var errDescr map[ErrNo]string = map[ErrNo]string{
	ErrNotImplemented:   "ErrNotImplemented",
	ErrInternal:         "ErrInternal",
	ErrNoAccess:         "ErrNoAccess",
	ErrTooManyRequests:  "ErrTooManyRequests",
	ErrNotAuthenticated: "ErrNotAuthenticated",
	ErrObjectNotFound:   "ErrObjectNotFound",
	ErrObjectExist:      "ErrObjectExist",
	ErrNotReady:         "ErrNotReady",
	ErrBadRequest:       "ErrBadRequest",
	ErrDeadlineExceeded: "ErrDeadlineExceeded",
}

func Name(e ErrNo) string {
	if n, ok := errDescr[e]; ok {
		return n
	}
	return ""
}
