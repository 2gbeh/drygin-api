package slice_util

func DeleteIndex[S ~[]V, V any](slice S, index int) S {
	return append(slice[:index], slice[index + 1:]...)
}
