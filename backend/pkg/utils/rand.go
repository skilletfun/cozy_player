package utils

import "math/rand/v2"

// Shuffle get slice of any type and shuffle it in place.
func Shuffle[T any](items *[]T) {
	rand.Shuffle(len(*items), func(i, j int) {
		(*items)[i], (*items)[j] = (*items)[j], (*items)[i]
	})
}
