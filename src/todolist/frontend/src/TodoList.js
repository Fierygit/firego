import { useState, useCallback, useEffect } from 'react';
import { Todo } from './Todo';
import './App.css';
import axios from 'axios';
import { URL } from "./const";

export function TodoList() {

    const [todoList, setTodoList] = useState([]);

    const [todoName, setTodoName] = useState("");

    const getTodolist = useCallback(async () => {
        const todos = await axios.get(`${URL}/todo`);
        setTodoList(todos.data);
    }, []);

    useEffect(() => {
        getTodolist();
    }, [getTodolist]);

    const addTodo = useCallback(async () => {
        if (todoName === '') {
            alert(`${todoName} illegal`);
            return;
        }

        await axios.post(`${URL}/todo`, {
            todo: todoName
        });

        getTodolist();
    }, [todoName, getTodolist]);

    const removeTodo = useCallback(async (id) => {
        await axios.post(`${URL}/todo/delete`, {
            id
        });
        getTodolist();
    }, [getTodolist]);

    const onKeyPress = useCallback((e) => {
        if (e.which !== 13) return;
        addTodo();
    }, [addTodo]);

    return (<div className="todolist">
        <h1>TodoList App</h1>
        {
            todoList.map(todo => <Todo todo={todo} key={todo.Id} removeTodo={removeTodo} />)
        }
        <div className="add-todo">
            <input className="add-todo-input" placeholder="add todo here..." value={todoName}
                onInput={e => setTodoName(e.target.value)} onKeyPress={onKeyPress}></input>
            <button className="add-todo-btn" onClick={addTodo}>Add Todo</button>
        </div>
    </div>)
}