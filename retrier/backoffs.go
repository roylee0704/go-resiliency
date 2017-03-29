package retrier

import "time"

// ConstantBackoff generates simple back-off strategy of retrying 'n' times,
// and waiting 'amount' time after each one
func ConstantBackoff(n int, amount time.Duration) []time.Duration {
	ret := make([]time.Duration, n)
	for i := range ret {
		ret[i] = amount
	}
	return ret
}

// ExponentialBackoff generates simple back-off strategy of retrying 'n' times,
// and waiting double 'amount' time after each one.
func ExponentialBackoff(n int, amount time.Duration) []time.Duration {
	ret := make([]time.Duration, n)
	next := amount
	for i := range ret {
		ret[i] = next
		next *= 2
	}
	return ret
}
