package encodex

import "testing"

func Test_CheckPassPassword(t *testing.T) {
	type args struct {
		password      string
		checkPassword string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{
				password:      "123456",
				checkPassword: "123456",
			},
			want: true,
		},
		{
			name: "test2",
			args: args{
				password:      "123456",
				checkPassword: "1234567",
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hash, err := hashPassword(tt.args.password)
			if err != nil {
				t.Errorf("hashPassword() error = %v", err)
				return
			}
			if checkPasswordHash(tt.args.checkPassword, hash) != tt.want {
				t.Errorf("checkPasswordHash() error")
				return
			}
		})
	}

}
