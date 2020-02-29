package sum

import "testing"
import "reflect"

func TestSum(t *testing.T) {

    t.Run("collection of 5 nums", func(t *testing.T) {
        t.Skip("Skip this too")
        nums := []int{1, 2, 3, 4, 5}

        got := Sum(nums)
        want := 15

        if got != want {
            t.Errorf("got %d want %d given, %v", got, want, nums)
        }
    })

    t.Run("collection of any size", func(t *testing.T) {
        t.Skip("skip this")
        nums := []int{1, 2, 3}

        got := Sum(nums)
        want := 6

        if got != 5 {
            t.Errorf("got %d want %d given, %v", got, want, nums)
        }
    })

    t.Run("add 1 or more lists of nums", func(t *testing.T) {
        got := SumAll([]int{2, 2}, []int{4, 4})
        want := []int{4, 8}

        if !reflect.DeepEqual(got, want) {
            t.Errorf("got %v want %v", got, want)
        }
    })

}
