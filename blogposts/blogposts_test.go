package blogposts_test

// import (
// 	"testing"
// 	"testing/fstest"

// 	"github.com/ashishj12/blogposts"
// )

// func TestNewBlogPosts(t *testing.T) {
// 	fs := fstest.MapFS{
// 		"hello world.md":  {Data: []byte("hi")},
// 		"hello-world2.md": {Data: []byte("hola")},
// 	}

// 	posts := blogposts.NewPostsFromFS(fs)

// 	if len(posts) != len(fs) {
// 		t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
// 	}
// }

// func assertPost(t *testing.T, got blogposts.Post, want blogposts.Post) {
// 	t.Helper()
// 	if !reflect.DeepEqual(got, want) {
// 		t.Errorf("got %+v, want %+v", got, want)
// 	}
// }
