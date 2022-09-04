// package delislice
// remove item from slice on index or value

package delislice

import (
	"errors"
	"fmt"
	"reflect"
)

type Item interface {
	interface{}
}

// remove item in slice on index
func DelItemInd[T Item](i uint, sl []T) ([]T, error) {
	if len(sl) > int(i) {
		return append(sl[:i], sl[i+1:]...), nil
	}
	return sl, errors.New("index out of range, len sl: " + fmt.Sprintf("%d", len(sl)))
}

// remove item in slice on value
// where val is interface{}
func DelItemVal[T Item](val T, sl []T) ([]T, error) {

	if len(sl) > 0 {
		for i, v := range sl {
			if reflect.DeepEqual(v, val) {
				sl, _ := DelItemInd(uint(i), sl)
				return DelItemVal(val, sl)
			}
		}
		return sl, nil
	}
	return sl, errors.New("slice is empty")
}
