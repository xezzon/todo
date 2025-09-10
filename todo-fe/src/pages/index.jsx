import { todoClient } from "@/api/todo";
import { useEffect, useState } from "react";

export default function Todo() {

  const [todos, setTodos] = useState([]);
  const [input, setInput] = useState("");

  const fetchTodos = () => {
    todoClient.getTasks()
      .then(({ data }) => setTodos(data))
  }

  const addTodo = () => {
    if (input.trim() === "") return;
    todoClient.addTask({
      content: input
    })
      .then(() => {
        fetchTodos()
        setInput("");
      })
  };

  const deleteTodo = (id) => {
    todoClient.deleteTask({
      id,
    })
      .then(() => {
        fetchTodos()
      })
  };

  useEffect(() => {
    fetchTodos()
  }, [])

  return (
    <div style={{ maxWidth: 400, margin: "40px auto", fontFamily: "Segoe UI, Arial" }}>
      <h2 style={{ textAlign: "center" }}>TODO 列表</h2>
      <div style={{ display: "flex", gap: 8 }}>
        <input
          value={input}
          onChange={e => setInput(e.target.value)}
          placeholder="添加新任务..."
          style={{ flex: 1, padding: 8, borderRadius: 4, border: "1px solid #ccc" }}
        />
        <button onClick={addTodo} style={{ padding: "8px 16px", borderRadius: 4, border: "none", background: "#0078d4", color: "#fff" }}>
          添加
        </button>
      </div>
      <ul style={{ listStyle: "none", padding: 0, marginTop: 24 }}>
        {todos.map(todo => (
          <li key={todo.id} style={{ display: "flex", alignItems: "center", justifyContent: "space-between", padding: "8px 0", borderBottom: "1px solid #eee" }}>
            <span>{todo.content}</span>
            <button onClick={() => deleteTodo(todo.id)} style={{ background: "none", border: "none", color: "#d83b01", cursor: "pointer", fontSize: 16 }}>
              删除
            </button>
          </li>
        ))}
      </ul>
    </div>
  );
}