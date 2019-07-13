package sumoftwo

import (
	"reflect"
	"testing"
)

func TestSumOfn(t *testing.T) {
	type args struct {
		arr []int
		n   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{name:"example01", args: args{arr:[]int{3, 24, 30}, n:54}, want: []int{1, 2}},
		{name:"example02", args: args{arr:[]int{3, 7, 8}, n:54}, want: nil}, 
	}
	for _, tt := range tests {
		tt := tt // 平行測試
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel() // 平行測試
			if got := SumOfn(tt.args.arr, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SumOfn() = %v, want %v", got, tt.want)
			}
		})
	}
}
