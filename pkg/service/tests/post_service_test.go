package tests

import (
	"context"
	"testing"

	"github.com/SunkPlane29/grin/pkg/service/post"
	"github.com/SunkPlane29/grin/pkg/service/storage/memory"
	"github.com/segmentio/ksuid"
)

func TestServiceCreatePost(t *testing.T) {
	creatorID := ksuid.New().String()

	type args struct {
		ctx       context.Context
		creatorID string
		p         post.Post
	}
	tests := []struct {
		name    string
		args    args
		want    *post.Post
		wantErr bool
	}{
		{
			name: "normal case",
			args: args{
				ctx:       context.Background(),
				creatorID: creatorID,
				p:         post.Post{Content: "lorem ipsum"},
			},
			want: &post.Post{
				CreatorID: creatorID,
				Content:   "lorem ipsum",
			},
			wantErr: false,
		}, //TODO: add more tests cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := memory.NewPostStorage()

			s := post.New(store)

			got, err := s.CreatePost(tt.args.ctx, tt.args.creatorID, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.CreatePost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got.CreatorID != tt.want.CreatorID || got.Content != tt.want.Content {
				t.Errorf("service.CreatePost() = CreatorID: %s Content: %s, want CreatorID: %s, Content: %s", got.CreatorID, got.Content, tt.want.CreatorID, tt.want.Content)
			}

			if _, err := ksuid.Parse(got.ID); err != nil {
				t.Errorf("service.CreatePost() error = %v, invalid ID %s", err, got.ID)
			}
		})
	}
}

func TestServiceGetPosts(t *testing.T) {
	creatorID := ksuid.New().String()

	type args struct {
		ctx       context.Context
		creatorID string
		posts     []post.Post
	}
	tests := []struct {
		name    string
		args    args
		want    *[]post.Post
		wantErr bool
	}{
		{
			name: "normal case",
			args: args{
				ctx:       context.Background(),
				creatorID: creatorID,
				posts:     []post.Post{{Content: "lorem ipsum"}},
			},
			want: &[]post.Post{
				{
					CreatorID: creatorID,
					Content:   "lorem ipsum",
				},
			},
			wantErr: false,
		}, //TODO: add more tests cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := memory.NewPostStorage()

			s := post.New(store)

			for _, post := range tt.args.posts {
				_, err := s.CreatePost(tt.args.ctx, tt.args.creatorID, post)
				if err != nil {
					t.Errorf("service.GetPosts() unexpected error: %v", err)
					return
				}
			}

			got, err := s.GetPosts(tt.args.ctx, tt.args.creatorID)
			if err != nil && !tt.wantErr {
				t.Errorf("service.GetPosts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if len(*got) != len(*tt.want) {
				t.Errorf("service.GetPosts() unmatched returned slices from usecase: got len: %d, want len: %d", len(*got), len(*tt.want))
			} //TODO: find a better method for comparing the result and the requirement
		})
	}
}

func TestServiceGetPost(t *testing.T) {
	creatorID := ksuid.New().String()

	type args struct {
		ctx       context.Context
		creatorID string
		p         post.Post
	}
	tests := []struct {
		name    string
		args    args
		want    *post.Post
		wantErr bool
	}{
		{
			name: "normal case",
			args: args{
				ctx:       context.Background(),
				creatorID: creatorID,
				p:         post.Post{Content: "lorem ipsum"},
			},
			want: &post.Post{
				CreatorID: creatorID,
				Content:   "lorem ipsum",
			},
			wantErr: false,
		}, //TODO: add more tests cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := memory.NewPostStorage()

			s := post.New(store)

			created, err := s.CreatePost(tt.args.ctx, tt.args.creatorID, tt.args.p)
			if err != nil {
				t.Errorf("service.GetPost() unexpected error = %v", err)
				return
			}

			got, err := s.GetPost(tt.args.ctx, tt.args.creatorID, created.ID)
			if err != nil && !tt.wantErr {
				t.Errorf("service.GetPost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got.ID != created.ID || got.CreatorID != tt.args.creatorID || got.Content != tt.args.p.Content {
				t.Errorf("service.GetPost() = got ID: %s, CreatorID: %s, Content: %s, want ID: %s, CreatorID: %s, Content: %s", got.ID, got.CreatorID, got.Content, created.ID, tt.args.creatorID, tt.args.p.Content)
			}
		})
	}
}
