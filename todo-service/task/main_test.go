package task

import (
	"testing"
	taskPb "todo-service/gen/todo/task"

	"github.com/stretchr/testify/assert"
)

func TestTaskStore_Add_GetAll_Delete(t *testing.T) {
	ts := NewTaskStore()

	// Add
	task1 := ts.Add(&taskPb.AddTaskReq{Content: "Test1"})
	task2 := ts.Add(&taskPb.AddTaskReq{Content: "Test2"})
	assert.NotNil(t, task1)
	assert.NotNil(t, task2)
	assert.NotEqual(t, task1.Id, "")
	assert.NotEqual(t, task2.Id, "")

	// GetAll
	tasks := ts.GetAll()
	assert.Len(t, tasks, 2)
	var found1, found2 bool
	for _, t := range tasks {
		if t.Content == "Test1" {
			found1 = true
		}
		if t.Content == "Test2" {
			found2 = true
		}
	}
	assert.True(t, found1)
	assert.True(t, found2)

	// Delete
	ts.Delete(task1.Id)
	tasks = ts.GetAll()
	assert.Len(t, tasks, 1)
	assert.Equal(t, "Test2", tasks[0].Content)
}

func TestTaskStore_EdgeCases(t *testing.T) {
	ts := NewTaskStore()

	// 删除不存在的任务
	ts.Delete("non-existent-id")
	tasks := ts.GetAll()
	assert.Len(t, tasks, 0)

	// 添加空内容任务
	task := ts.Add(&taskPb.AddTaskReq{Content: ""})
	assert.NotNil(t, task)
	assert.Equal(t, "", task.Content)

	// 再次删除已删除的任务
	ts.Delete(task.Id)
	ts.Delete(task.Id) // 重复删除
	tasks = ts.GetAll()
	assert.Len(t, tasks, 0)

	// 添加多个空内容任务
	ts.Add(&taskPb.AddTaskReq{Content: ""})
	ts.Add(&taskPb.AddTaskReq{Content: ""})
	tasks = ts.GetAll()
	assert.Len(t, tasks, 2)
	for _, task := range tasks {
		assert.Equal(t, "", task.Content)
	}
}
