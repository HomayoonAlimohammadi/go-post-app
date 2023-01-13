package core

import (
	"context"
	"sync"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/gofrs/uuid"
	"github.com/homayoonalimohammadi/go-post-app/postapp/gen/go/postapp"
)

type PostApp struct {
	logger grpclog.LoggerV2
	posts  map[string]*postapp.Post
	mutex  sync.Mutex
	postapp.UnimplementedPostAppServer
}

func New(logger grpclog.LoggerV2) *PostApp {
	return &PostApp{
		logger: logger,
		posts:  make(map[string]*postapp.Post),
		mutex:  sync.Mutex{},
	}
}

func (p *PostApp) GetPost(ctx context.Context,
	req *postapp.GetPostRequest) (*postapp.GetPostResponse, error) {

	post, ok := p.posts[req.GetToken()]
	if !ok {
		return nil, status.Error(codes.NotFound, "post not found")
	}

	resp := &postapp.GetPostResponse{
		Post: post,
	}
	return resp, nil
}

func (p *PostApp) GetPosts(ctx context.Context, _ *emptypb.Empty) (*postapp.GetPostsResponse, error) {

	resp := &postapp.GetPostsResponse{}
	for _, post := range p.posts {
		resp.Posts = append(resp.Posts, post)
	}
	return resp, nil
}

func (p *PostApp) CreatePost(ctx context.Context,
	req *postapp.CreatePostRequest) (*emptypb.Empty, error) {

	p.mutex.Lock()
	defer p.mutex.Unlock()
	var newToken string
	for newToken == "" || p.posts[newToken] != nil {
		newToken = uuid.Must(uuid.NewV4()).String()[:5]
	}
	req.Post.Id = newToken
	p.posts[newToken] = req.Post
	p.logger.Infof("added post %s", newToken)

	return &emptypb.Empty{}, nil
}
