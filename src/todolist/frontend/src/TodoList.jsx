import { useState, useCallback, useEffect } from 'react';
import { Todo } from './Todo';
import './index.css';
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

    return (<div className="h-full w-full flex flex-col justify-start items-center">
        <h1 className="text-4xl sm:text-5xl md:text-7xl text-white font-mono font-black mb-5">TodoList App</h1>
        {
            todoList.map(todo => <Todo todo={todo} key={todo.Id} removeTodo={removeTodo} />)
        }
        <div className="flex items-center justify-center m-6 w-full h-12 md:h-10">
            <input className="text-2xl h-full w-2/3 md:w-1/2 rounded-lg shadow-2xl outline-none" placeholder="add todo here..." value={todoName}
                onInput={e => setTodoName(e.target.value)} onKeyPress={onKeyPress}></input>
            <button className="ml-3 text-xl h-full rounded-2xl bg-green-400 font-bold w-20 md:w-40" onClick={addTodo}>Add</button>
        </div>
    </div>)
}