package leetcode

import (
	"reflect"
	"testing"
)

func Test_findRestaurant(t *testing.T) {
	type args struct {
		list1 []string
		list2 []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"[Test Case 1]",
			args{
				[]string{
					"Shogun", "Tapioca Express", "Burger King", "KFC",
				},
				[]string{
					"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun",
				},
			},
			[]string{
				"Shogun",
			},
		},
		{
			"[Test Case 2]",
			args{
				[]string{
					"Shogun", "Tapioca Express", "Burger King", "KFC",
				},
				[]string{
					"KFC", "Shogun", "Burger King",
				},
			},
			[]string{
				"Shogun",
			},
		},
		{
			"[Test Case 3]",
			args{
				[]string{
					"Shogun", "Tapioca Express", "Burger King", "KFC",
				},
				[]string{
					"Tapioca Express", "Shogun", "Burger King",
				},
			},
			[]string{
				"Tapioca Express", "Shogun",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRestaurant(tt.args.list1, tt.args.list2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findRestaurant() = %v, want %v", got, tt.want)
			}
		})
	}
}
