package sqlite

import (
	"context"
	"reflect"
	"testing"

	"github.com/SunkPlane29/grin/pkg/auth/core"
)

func TestStorageStoreUser(t *testing.T) {
	type args struct {
		ctx  context.Context
		user core.User
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "normal usecase",
			args: args{
				ctx:  context.Background(),
				user: core.User{ID: "lorem", Username: "ipsum", PasswordHash: []byte("sit dor amet")},
			},
			wantErr: false,
		},
		// TODO: Add test cases.
	}

	s, err := NewScratch(context.Background(), "./test.db")
	if err != nil {
		t.Fatal(err)
	} //FIXME: some issues testing with make, should create temp dir and temp files for testing

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := s.StoreUser(tt.args.ctx, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("Storage.StoreUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStorageGetUser(t *testing.T) {
	type args struct {
		ctx  context.Context
		user core.User
	}
	tests := []struct {
		name    string
		args    args
		want    *core.User
		wantErr bool
	}{
		{
			name: "normal usecase",
			args: args{
				ctx:  context.Background(),
				user: core.User{ID: "lorem", Username: "ipsum", PasswordHash: []byte("sit dor amet")},
			},
			want:    &core.User{ID: "lorem", Username: "ipsum", PasswordHash: []byte("sit dor amet")},
			wantErr: false,
		},
	}

	s, err := NewScratch(context.Background(), "./test.db")
	if err != nil {
		t.Fatal(err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := s.StoreUser(tt.args.ctx, tt.args.user)
			if err != nil {
				t.Errorf("Storage.GetUser() unexpected error = %v", err)
				return
			}

			got, err := s.GetUser(tt.args.ctx, tt.args.user.ID)
			if err != nil && !tt.wantErr {
				t.Errorf("Storage.GetUser() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Storage.GetUser() = User %v, want User %v", got, tt.want)
			}
		})
	}
}
