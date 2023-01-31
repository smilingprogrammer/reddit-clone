// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"context"
)

type Querier interface {
	ChangeBlockedStatus(ctx context.Context, arg ChangeBlockedStatusParams) error
	CreatePost(ctx context.Context, arg CreatePostParams) (Post, error)
	CreateSub(ctx context.Context, arg CreateSubParams) (Sub, error)
	CreateSubscriber(ctx context.Context, arg CreateSubscriberParams) (Subscriber, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeletePost(ctx context.Context, id int64) error
	DeleteSub(ctx context.Context, name string) error
	DeleteSubscriber(ctx context.Context, id int64) error
	DeleteUser(ctx context.Context, id int64) error
	GetPost(ctx context.Context, id int64) (Post, error)
	GetSub(ctx context.Context, name string) (Sub, error)
	GetSubscriber(ctx context.Context, arg GetSubscriberParams) (Subscriber, error)
	GetUser(ctx context.Context, id int64) (User, error)
	ListPosts(ctx context.Context, arg ListPostsParams) ([]Post, error)
	ListSub(ctx context.Context, arg ListSubParams) ([]Sub, error)
	ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error)
	UpdatePost(ctx context.Context, arg UpdatePostParams) (Post, error)
	UpdateSub(ctx context.Context, arg UpdateSubParams) (Sub, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
}

var _ Querier = (*Queries)(nil)
