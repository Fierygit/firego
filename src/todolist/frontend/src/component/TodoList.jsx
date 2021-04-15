import { useState, useCallback, useEffect } from 'react';
import { Todo } from './Todo';
import axios from 'axios';
import './index.css';
import moment from 'moment';

export function TodoList() {

    const [todoList, setTodoList] = useState([]);

    const [todoName, setTodoName] = useState("");

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

    const addTodo = useCallback(async () => {
        if (todoName === '') {
            alert(`todo can not be empty!!!`);
            return;
        }

        await axios.post('/todo', {
            todo: todoName
        });

        setTodoName('');

        getTodolist();
    }, [todoName, getTodolist]);

    const removeTodo = useCallback(async (id) => {
        await axios.post('/todo/delete', {
            id
        });
        getTodolist();
    }, [getTodolist]);

    const onKeyPress = useCallback((e) => {
        if (e.which !== 13) return;
        addTodo();
    }, [addTodo]);

    return (
        <div className="h-full w-full flex flex-col justify-start items-center">
            <h1 className="text-4xl sm:text-5xl md:text-7xl text-black dark:text-white font-mono font-black select-none">TodoList</h1>
            <h5 className="text-black dark:text-white font-mono select-none">{now}</h5>
            {
                todoList.map(todo => <Todo todo={todo} key={todo.Id} removeTodo={removeTodo} />)
            }
            <div className="flex items-center justify-center m-6 w-full h-12 md:h-10">
                <input className="text-2xl text-white dark:text-black focus:ring-2 focus:ring-blue-600 select-none h-full w-2/3 md:w-1/2 rounded-lg shadow-2xl bg-gray-700 dark:bg-white outline-none" placeholder=" add todo here..." value={todoName}
                    onInput={e => setTodoName(e.target.value)} onKeyPress={onKeyPress}></input>
                <button className="transform shadow-lg ml-3 text-xl h-full rounded-2xl bg-green-400 active:bg-green-600 hover:scale-110 font-bold w-20 md:w-40 text-gray-200" onClick={addTodo}>Add</button>
            </div>
        </div>)
}