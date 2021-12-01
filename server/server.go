package server

import (
	"context"
	"github.com/google/uuid"
	pb "github.com/zchary-ma/grpc-server/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"
	"net"
	"time"
)

type Server struct {
	pb.UnimplementedNoteServiceServer
}

var NoteMap map[string]*pb.Note

func NewServer() *Server {
	var s = &Server{}
	t := time.Now()
	NoteMap = map[string]*pb.Note{
		"1": {
			Id:    "1",
			Title: "note1",
			Contents: []*pb.Note_Content{
				{
					Type: pb.Note_Content_TEXT,
					Text: "text1",
				},
			},
			CreatedAt: timestamppb.New(t), // NOTE timezone? not match with time.Now()
			UpdatedAt: timestamppb.New(t.Add(time.Hour * 10)),
		},
		"2": {
			Id:    "2",
			Title: "note2",
			Contents: []*pb.Note_Content{
				{
					Type: pb.Note_Content_TEXT,
					Text: "text2",
				},
			},
			CreatedAt: timestamppb.New(t),
			UpdatedAt: timestamppb.New(t.Add(time.Hour * 10)),
		},
	}
	return s
}

func (s Server) GetNote(ctx context.Context, set *pb.IdSet) (*pb.NoteList, error) {
	var list = &pb.NoteList{}
	for _, id := range set.Ids {
		list.Notes = append(list.Notes, NoteMap[id])
	}
	return list, nil
}

func (s Server) CreateNote(ctx context.Context, note *pb.Note) (*pb.Id, error) {
	id := uuid.NewString()
	note.Id = id
	NoteMap[id] = note
	return &pb.Id{Id: id}, nil
}

func (s Server) UpdateNote(ctx context.Context, note *pb.Note) (*pb.Note, error) {
	NoteMap[note.Id] = note
	return note, nil
}

func (s Server) DeleteNote(ctx context.Context, set *pb.IdSet) (*pb.IdSet, error) {
	for _, id := range set.Ids {
		delete(NoteMap, id)
	}
	return set, nil
}

func (s *Server) ListenAndServe(port string) error {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}
	srv := grpc.NewServer()
	pb.RegisterNoteServiceServer(srv, Server{})

	reflection.Register(srv)
	return srv.Serve(listener)
}
