package sumoftwo

type di struct {
	detect     bool
	labelindex int
}

// SumOfn 計算合條件的index
func SumOfn(arr []int, n int) []int {
	countn := make([]di, n+1)
	for _, v := range countn {
		v.detect = false
	}

	var ansOne, ansTwo int

	for i := 0; i < len(arr); i++ {
		if arr[i] <= n && countn[n-arr[i]].detect {
			ansOne = countn[n-arr[i]].labelindex
			ansTwo = i
			return []int{ansOne, ansTwo}
		}

		if arr[i] <= n {
			countn[arr[i]].detect = true
			countn[arr[i]].labelindex = i
		}

	}

	return nil

}
