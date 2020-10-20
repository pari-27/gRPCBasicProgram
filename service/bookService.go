package service

import (
	"context"
	"log"
	"strconv"
)

type BookServiceGrpcImpl struct {
}

//NewRepositoryServiceGrpcImpl returns the pointer to the implementation.
func NewBookServiceGrpcImpl() *BookServiceGrpcImpl {
	return &BookServiceGrpcImpl{}
}

//Add function implementation of gRPC Service.
func (serviceImpl *BookServiceGrpcImpl) Add(ctx context.Context, in *demoPB.Book) (*service.GetBookResponse, error) {
	log.Println("Received request for getting repository with id " + strconv.FormatInt(in.Id, 10))

	//Logic to persist to database or storage.
	log.Println("Repository persisted to the storage")

	return &service.GetBookResponse{
		listBook: in,
		Error:    nil,
	}, nil
}
