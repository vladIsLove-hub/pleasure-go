package pleasure_go

func Reduce[T any, U any](entity []T, cb func(acc U, item T) U, structure U) U {
	var acc U = structure

	for i := 0; i < len(entity); i++ {
		acc = cb(acc, entity[i])
	}

	return acc
}