import { useState, useCallback, useEffect } from 'react';
import { Todo } from './Todo';
import { AddDialog } from "./AddDialog";
import axios from 'axios';
import moment from 'moment';

export function TodoList() {

    const [todoList, setTodoList] = useState([]);

    const [now, setNow] = useState(moment().format("MMMM Do YYYY, H:mm:ss"));

    const getTodolist = useCallback(async () => {
        const todos = await axios.get('/todo');
        setTodoList(todos.data);
    }, []);

    useEffect(() => {
        getTodolist();
    }, [getTodolist]);

    useEffect(() => {
        setInterval(() => {
            setNow(moment().format("MMMM Do YYYY, H:mm:ss"));
        }, 1000);
    }, [setNow]);

    const removeTodo = useCallback(async (id) => {
        await axios.post('/todo/delete', {
            id
        });
        getTodolist();
    }, [getTodolist]);

    return (
        <div className="min-h-screen w-full flex flex-col justify-start items-center">
            <h1 className="text-4xl sm:text-5xl md:text-7xl text-black dark:text-white font-mono font-black select-none">TodoList</h1>
            <h5 className="text-black dark:text-white font-mono select-none">{now}</h5>
            {
                todoList.map(todo => <Todo todo={todo} key={todo.Id} removeTodo={removeTodo} />)
            }
            <AddDialog getTodolist={getTodolist} ></AddDialog>
        </div>
    )
}