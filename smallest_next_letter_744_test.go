package roman_number

import "testing"

func Test_nextGreatestLetter(t *testing.T) {
	type args struct {
		letters []byte
		target  byte
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{"Example 0", args{[]byte{'a', 'b'}, 'z'}, 'a'},
		{"Example 1", args{[]byte{'c', 'f', 'j'}, 'a'}, 'c'},
		{"Example 2", args{[]byte{'c', 'f', 'j'}, 'c'}, 'f'},
		{"Example 3", args{[]byte{'c', 'f', 'j'}, 'd'}, 'f'},
		{"Submit Failed 1", args{[]byte{'c', 'f', 'j'}, 'j'}, 'c'},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreatestLetter(tt.args.letters, tt.args.target); got != tt.want {
				t.Errorf("nextGreatestLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}
