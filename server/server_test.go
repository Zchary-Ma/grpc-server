package server

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/zchary-ma/grpc-server/mock"
	pb "github.com/zchary-ma/grpc-server/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"testing"
	"time"
)

func TestServer_CreateNote(t *testing.T) {
	startTime := time.Now()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish() // as

	// init mock client
	mockNoteClient := mock.NewMockNoteServiceClient(ctrl)

	// req
	req := &pb.Note{
		Id:    "3",
		Title: "test note",
		Contents: []*pb.Note_Content{
			{
				Type: pb.Note_Content_TEXT,
				Text: "test note",
			},
		},
		CreatedAt: timestamppb.New(startTime),
		UpdatedAt: timestamppb.New(startTime.Add(time.Hour * 2)),
	}

	mockNoteClient.EXPECT().CreateNote(gomock.Any(), req).Return(&pb.Note{
		Id:    "3",
		Title: "test note",
		Contents: []*pb.Note_Content{
			{
				Type: pb.Note_Content_TEXT,
				Text: "test note",
			},
		},
		CreatedAt: timestamppb.New(startTime),
		UpdatedAt: timestamppb.New(startTime.Add(time.Hour * 2)),
	}, nil)

	// conn, err := grpc.Dial("50051", grpc.WithInsecure())
	// if err != nil {
	// 	t.Fatalf("did not connect: %v", err)
	// }
	// defer conn.Close()
	// client := pb.NewNoteServiceClient(conn)

	testserverCreateNote(t, mockNoteClient)

}

func testserverCreateNote(t *testing.T, client pb.NoteServiceClient) {
	ctx, canncel := context.WithTimeout(context.Background(), time.Second)
	defer canncel()

	note := pb.Note{
		Id:    "3",
		Title: "test note",
		Contents: []*pb.Note_Content{
			{
				Type: pb.Note_Content_TEXT,
				Text: "test note",
			},
		},
	}

	r, err := client.CreateNote(ctx, &note)
	if err != nil {
		t.Errorf("CreateNote() error = %v", err)
	}

	if r.Id != "3" {
		t.Errorf("CreateNote() error = %v", err)
	}

	log.Printf("Recv: %v\n", r)

}
