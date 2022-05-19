package sol

import (
	"reflect"
	"testing"
)

func BenchmarkTest(b *testing.B) {
	commands := &[]string{"Trie", "insert", "search", "search", "startsWith", "insert", "search"}
	val := &[][]string{{}, {"apple"}, {"apple"}, {"app"}, {"app"}, {"app"}, {"app"}}
	for idx := 0; idx < b.N; idx++ {
		Run(commands, val)
	}
}
func TestRun(t *testing.T) {
	type args struct {
		commands *[]string
		val      *[][]string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "input:[\"Trie\", \"insert\", \"search\", \"search\", \"startsWith\", \"insert\", \"search\"]," +
				"values:[[], [\"apple\"], [\"apple\"], [\"app\"], [\"app\"], [\"app\"], [\"app\"]]",
			args: args{commands: &[]string{"Trie", "insert", "search", "search", "startsWith", "insert", "search"},
				val: &[][]string{{}, {"apple"}, {"apple"}, {"app"}, {"app"}, {"app"}, {"app"}}},
			want: []string{"null", "null", "true", "false", "true", "null", "true"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Run(tt.args.commands, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Run() = %v, want %v", got, tt.want)
			}
		})
	}
}
