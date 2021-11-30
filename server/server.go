package server

import (
	"context"
	"github.com/zchary-ma/pre/pb"
	"google.golang.org/grpc"
	"net"
)

type server struct {
	pb.UnimplementedNoteServiceServer
	NoteMap map[string]*pb.Note
}

func (s server) GetNote(ctx context.Context, set *pb.IdSet) (*pb.NoteList, error) {
	panic("implement me")
}

func (s server) CreateNote(ctx context.Context, note *pb.Note) (*pb.Note, error) {
	panic("implement me")
}

func (s server) UpdateNote(ctx context.Context, note *pb.Note) (*pb.Note, error) {
	panic("implement me")
}

func (s server) DeleteNote(ctx context.Context, set *pb.IdSet) (*pb.IdSet, error) {
	panic("implement me")
}

func ListenAndServe(port string) error {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}
	s := grpc.NewServer()
	pb.RegisterNoteServiceServer(s, server{})

	return s.Serve(listener)
}
