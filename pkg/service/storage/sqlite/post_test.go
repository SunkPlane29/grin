package sqlite

// func TestStorageCreatePost(t *testing.T) {
// 	type fields struct {
// 		db *sql.DB
// 	}
// 	type args struct {
// 		ctx context.Context
// 		p   post.Post
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		want    *post.Post
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			s := &Storage{
// 				db: tt.fields.db,
// 			}
// 			got, err := s.CreatePost(tt.args.ctx, tt.args.p)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Storage.CreatePost() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Storage.CreatePost() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
