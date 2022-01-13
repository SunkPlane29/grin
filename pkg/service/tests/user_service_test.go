package tests

import (
	"context"
	"reflect"
	"testing"

	"github.com/SunkPlane29/grin/pkg/service/storage/memory"
	"github.com/SunkPlane29/grin/pkg/service/user"
	"github.com/segmentio/ksuid"
)

func TestServiceCreateUser(t *testing.T) {
	userID := ksuid.New().String()

	type args struct {
		ctx    context.Context
		userID string
		user   user.User
	}
	tests := []struct {
		name    string
		args    args
		want    *user.User
		wantErr bool
	}{
		{
			name: "normal test case",
			args: args{
				ctx:    context.Background(),
				userID: userID,
				user: user.User{
					Username: "lorem",
					Alias:    "ipsum",
				},
			},
			want: &user.User{
				ID:       userID,
				Username: "lorem",
				Alias:    "ipsum",
			},
			wantErr: false,
		},
		// TODO: Add test cases.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := memory.NewUserStorage()

			s := user.New(store)

			got, err := s.CreateUser(tt.args.ctx, tt.args.userID, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.CreateUser() = User %v, want User %v", got, tt.want)
			}
		})
	}
}

func TestServiceCheckUserExists(t *testing.T) {
	userID := ksuid.New().String()

	type args struct {
		ctx    context.Context
		userID string
		user   user.User
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "normal test case",
			args: args{
				ctx:    context.Background(),
				userID: userID,
				user: user.User{
					Username: "lorem",
					Alias:    "ipsum",
				},
			},
			want: true,
		},
		// TODO: Add test cases.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := memory.NewUserStorage()

			s := user.New(store)

			if _, err := s.CreateUser(tt.args.ctx, tt.args.userID, tt.args.user); err != nil {
				t.Errorf("service.CreateUser() unexpected error = %v", err)
				return
			}

			got := s.CheckUserExists(tt.args.ctx, tt.args.userID)

			if got != tt.want {
				t.Errorf("service.CreateUser() = User %v, want User %v", got, tt.want)
			}
		})
	}
}
