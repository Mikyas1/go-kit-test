package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kit/kit/sd/etcd"
	"google.golang.org/grpc"
	"log"
	pb "users/pkg/grpc"
)

// UsersService describes the service.
type UsersService interface {
	Create(ctx context.Context, email string) error
}

type basicUsersService struct {
	notificatorServiceClient pb.NotificatorClient
}

func (b *basicUsersService) Create(ctx context.Context, email string) error {
	reply, err := b.notificatorServiceClient.SendEmail(
		context.Background(),
		&pb.SendEmailRequest{
			Email:   email,
			Content: "Hi! Thank you for registration",
		},
	)
	if err != nil {
		return errors.New("no reply form notification server")
	}
	if reply != nil {
		log.Printf("Email ID: %s", reply.Id)
		return nil
	}
	return nil

}

// NewBasicUsersService returns a naive, stateless implementation of UsersService.
func NewBasicUsersService() UsersService {
	var (
		etcdServer = "http://etcd:2379"
		prefix     = "/services/notificator"
	)

	client, err := etcd.NewClient(context.Background(), []string{etcdServer}, etcd.ClientOptions{})
	if err != nil {
		log.Printf("unable to connect to etcd: %s", err.Error())
		return new(basicUsersService)
	}

	entries, err := client.GetEntries(prefix)
	if err != nil || len(entries) == 0 {
		log.Printf("unable to get entries: %s", err.Error())
		return new(basicUsersService)
	}

	fmt.Println("=======================")
	fmt.Println(entries)

	conn, err := grpc.Dial(entries[0], grpc.WithInsecure())
	if err != nil {
		log.Printf("unable to connect to notificator: %s", err.Error())
		return new(basicUsersService)
	}
	return &basicUsersService{
		notificatorServiceClient: pb.NewNotificatorClient(conn),
	}
}

// New returns a UsersService with all of the expected middleware wired in.
func New(middleware []Middleware) UsersService {
	var svc UsersService = NewBasicUsersService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
