package jwt

import "testing"

func Test_tokenProcess(t *testing.T) {
	type args struct {
		userId int64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test-1", args: args{userId: 114514}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tokenProcess(tt.args.userId); got != tt.want {
				t.Errorf("tokenProcess() = %v, want %v", got, tt.want)
			}
		})
	}
}
