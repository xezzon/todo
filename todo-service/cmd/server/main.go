package main

import (
	"context"
	"net/http"

	"connectrpc.com/connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	todo "todo-service/gen/todo"
	taskPb "todo-service/gen/todo/task"
	"todo-service/gen/todo/task/taskconnect"

	emptypb "google.golang.org/protobuf/types/known/emptypb"

	taskStore "todo-service/task"
)

type TodoServiceServer struct {
	TaskStore taskStore.TaskStore
}

func newTodoServiceServer() *TodoServiceServer {
	return &TodoServiceServer{
		TaskStore: taskStore.NewTaskStore(),
	}
}

func (s *TodoServiceServer) AddTask(
	ctx context.Context,
	req *connect.Request[taskPb.AddTaskReq],
) (*connect.Response[todo.Id], error) {
	task := s.TaskStore.Add(req.Msg)
	res := connect.NewResponse(&todo.Id{
		Id: task.Id,
	})
	return res, nil
}

func (s *TodoServiceServer) GetTasks(
	ctx context.Context,
	req *connect.Request[emptypb.Empty],
) (*connect.Response[taskPb.TaskList], error) {
	res := connect.NewResponse(&taskPb.TaskList{
		Data: s.TaskStore.GetAll(),
	})
	return res, nil
}

func (s *TodoServiceServer) DeleteTask(
	ctx context.Context,
	req *connect.Request[todo.Id],
) (*connect.Response[emptypb.Empty], error) {
	s.TaskStore.Delete(req.Msg.Id)
	res := connect.NewResponse(&emptypb.Empty{})
	return res, nil
}

func main() {
	server := newTodoServiceServer()
	mux := http.NewServeMux()
	path, handler := taskconnect.NewTodoServiceHandler(server)
	mux.Handle(path, handler)
	http.ListenAndServe(
		"localhost:8080",
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
