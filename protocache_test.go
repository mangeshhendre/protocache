package protocache

import (
	"testing"
	"time"
)

var (
	servers = []string{"localhost:11211"}
)

func TestPC_GetVersionedKey(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "First",
			args: args{
				key: "Bar",
			},
			want: "f7a61c3b31b472f5563aeb62b638041ddd8128031354f6060a6c3dedb65a9061",
			// want:    "Bar|1",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pc := New(2*time.Second, "UnitTests", servers...)
			got, err := pc.getVersionedKey(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("PC.GetVersionedKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PC.GetVersionedKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPC_GetSPSK(t *testing.T) {
	type args struct {
		primaryContext   string
		secondaryContext string
		key              string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "First",
			args: args{
				primaryContext:   "Pri",
				secondaryContext: "Sec",
				key:              "Key",
			},
			want:    "be43d3305e7b93afc886d779fac762553d0406f0b49cb1a2796779636cb222df",
			wantErr: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pc := New(2*time.Second, "UnitTests", servers...)
			got, err := pc.getSPSK(tt.args.primaryContext, tt.args.secondaryContext, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("PC.GetSPSK() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PC.GetSPSK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPC_HashKey(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "First",
			args: args{
				key: "Foo",
			},
			want: "1cbec737f863e4922cee63cc2ebbfaafcd1cff8b790d8cfd2e6a5d550b648afa",
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pc := New(2*time.Second, "UnitTests", servers...)
			if got := pc.HashKey(tt.args.key); got != tt.want {
				t.Errorf("PC.HashKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func TestPC_Set(t *testing.T) {
// 	type fields struct {
// 		Scope    string
// 		Memcache *memcache.Client
// 		Logger   logxi.Logger
// 	}
// 	type args struct {
// 		primaryContext   string
// 		secondaryContext string
// 		key              string
// 		value            proto.Message
// 		expiration       time.Duration
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		wantErr bool
// 	}{
// 		{
// 			name: "First",
// 			args: args{
// 				primaryContext:   "Foo",
// 				secondaryContext: "Bar",
// 				key:              "Baz",
// 			},
// 			wantErr: false,
// 		},
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			pc := New(2*time.Second, "UnitTests", servers...)
// 			if err := pc.Set(tt.args.primaryContext, tt.args.secondaryContext, tt.args.key, tt.args.value, tt.args.expiration); (err != nil) != tt.wantErr {
// 				t.Errorf("PC.Set() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }
