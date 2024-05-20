package handles

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	pb "github.com/Vova4o/todogrpc/todoproto/proto"
)

func NextDate(c pb.TodoProtoServiceClient) string {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := c.NextDate(ctx, &pb.NextDateRequest{
		Now:    "20210304",
		Date:   "20210302",
		Repeat: "d 2",
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Printf("Next date: %s\n", res.Date)
	return res.Date
}

func AddTask(c pb.TodoProtoServiceClient) int64 {
	log.Printf("Add task was invoked")

	record := &pb.AddTaskRequest{
		Date:    "20240517",
		Title:   "Test1",
		Comment: "No comment",
		Repeat:  "",
	}

	res, err := c.AddTask(context.Background(), record)
	if err != nil {
		log.Printf("Unexpected error: %v\n", err)
	}

	log.Printf("Task has been created: %v\n", res.Id)
	return res.Id
}

func AllTasks(c pb.TodoProtoServiceClient) {
	log.Printf("All tasks was invoked")

	stream, err := c.AllTasks(context.Background(), &pb.TaskRequest{
		Search: "",
	})
	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Unexpected error: %v\n", err)
		}

		log.Printf("Task: %v\n", res)
	}
}

func AllTasksByName(c pb.TodoProtoServiceClient) {
	log.Printf("All tasks by Name was invoked")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stream, err := c.AllTasks(ctx, &pb.TaskRequest{
		Search: "Test",
	})
	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Unexpected error: %v\n", err)
		}

		log.Printf("Task: %v\n", res)
	}
}

func AllTasksByDate(c pb.TodoProtoServiceClient) {
	log.Printf("All tasks by Date was invoked")

	stream, err := c.AllTasks(context.Background(), &pb.TaskRequest{
		Search: "17.05.2024",
	})
	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Unexpected error: %v\n", err)
		}

		log.Printf("Task: %v\n", res)
	}
}
