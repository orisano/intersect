package intersect

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

func TestIntersect(t *testing.T) {

	t.Run("Shotgun1", func(t *testing.T) {
		testIntersect(t, Shotgun1)
	})
	t.Run("Shotgun4", func(t *testing.T) {
		testIntersect(t, Shotgun4)
	})
	t.Run("Galloping", func(t *testing.T) {
		testIntersect(t, Galloping)
	})
	t.Run("Naive", func(t *testing.T) {
		testIntersect(t, Naive)
	})
}

func testIntersect(t *testing.T, fn func(dest, small, large []int) []int) {
	t.Helper()
	type args struct {
		dest  []int
		small []int
		large []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				small: []int{1, 8},
				large: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			},
			want: []int{1, 8},
		},
		{
			args: args{
				small: []int{1, 8},
				large: []int{1, 2, 3, 4, 5, 6, 7, 8},
			},
			want: []int{1, 8},
		},
		{
			args: args{
				small: []int{8},
				large: []int{1, 2, 3, 4, 5, 6, 7},
			},
			want: nil,
		},
		{
			args: args{
				small: []int{3, 4},
				large: []int{1, 2, 3, 4, 5, 6, 7},
			},
			want: []int{3, 4},
		},
		{
			args: args{
				small: []int{-1, 4},
				large: []int{1, 2, 3, 4, 5, 6, 7},
			},
			want: []int{4},
		},
		{
			args: args{
				small: []int{1, 3, 4, 5, 6, 7},
				large: []int{1, 2, 3, 4, 5, 6, 7},
			},
			want: []int{1, 3, 4, 5, 6, 7},
		},
		{
			args: args{
				small: []int{1, 1, 1, 4, 5, 6, 7},
				large: []int{1, 1, 3, 4, 5, 6, 7},
			},
			want: []int{1, 1, 4, 5, 6, 7},
		},
		{
			args: args{
				small: []int{1, 1, 1, 1, 1, 6, 7, 8},
				large: []int{1, 1, 1, 1, 5, 6, 7, 8},
			},
			want: []int{1, 1, 1, 1, 6, 7, 8},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fn(tt.args.dest, tt.args.small, tt.args.large); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkIntersect(b *testing.B) {
	params := []struct {
		n     int
		nTest int
		mMax  int
	}{
		{
			n:     100000000,
			nTest: 32,
			mMax:  1000000000,
		},
		{
			n:     10000,
			nTest: 7000,
			mMax:  1000000000,
		},
	}

	for _, param := range params {
		b.Run(fmt.Sprintf("n=%d,nTest=%d,mMax=%d", param.n, param.nTest, param.mMax), func(b *testing.B) {
			rng := rand.New(rand.NewSource(0))
			a := make([]int, param.n)
			for i := range a {
				a[i] = rng.Intn(param.mMax)
			}
			sort.Ints(a)
			q := make([]int, param.nTest)
			for i := range q {
				q[i] = a[rng.Intn(len(a))]
			}
			sort.Ints(q)
			intersection := make([]int, 0, param.nTest)

			b.ResetTimer()
			b.Run("Naive", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					intersection = Naive(intersection[:0], q, a)
				}
			})
			b.Run("Galloping", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					intersection = Galloping(intersection[:0], q, a)
				}
			})
			b.Run("Shotgun1", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					intersection = Shotgun1(intersection[:0], q, a)
				}
			})
			b.Run("Shotgun4", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					intersection = Shotgun4(intersection[:0], q, a)
				}
			})
		})
	}
}
