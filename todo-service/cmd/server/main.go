package main

import (
	"context"
	"net/http"

	"connectrpc.com/connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	todo "todo-service/gen/todo"
	task "todo-service/gen/todo/task"
	"todo-service/gen/todo/task/taskconnect"

	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type TodoServiceServer struct {
}

func (s *TodoServiceServer) AddTask(
	ctx context.Context,
	req *connect.Request[task.AddTaskReq],
) (*connect.Response[todo.Id], error) {
	res := connect.NewResponse(&todo.Id{
		Id: "",
	})
	return res, nil
}

func (s *TodoServiceServer) GetTasks(
	ctx context.Context,
	req *connect.Request[emptypb.Empty],
) (*connect.Response[task.TaskList], error) {
	res := connect.NewResponse(&task.TaskList{
		Data: make([]*task.Task, 0),
	})
	return res, nil
}

func (s *TodoServiceServer) DeleteTask(
	ctx context.Context,
	req *connect.Request[todo.Id],
) (*connect.Response[emptypb.Empty], error) {
	res := connect.NewResponse(&emptypb.Empty{})
	return res, nil
}

func main() {
	server := &TodoServiceServer{}
	mux := http.NewServeMux()
	path, handler := taskconnect.NewTodoServiceHandler(server)
	mux.Handle(path, handler)
	http.ListenAndServe(
		"localhost:8080",
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
