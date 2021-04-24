import { useState, useCallback, useEffect } from 'react';
import { Todo } from './Todo';
import { AddDialog } from "./AddDialog";
import axios from 'axios';
import moment from 'moment';
import { RefreshButton } from './RefreshButton';
import { isBefore1day, checkRemindTodo } from '../util';

export function TodoList() {

    const [todoList, setTodoList] = useState([]);

    const [showFinished, setShowFinished] = useState(false);

    const [remindTodo, setRemindTodo] = useState(0);

    const [now, setNow] = useState(moment().format("MMMM Do YYYY, H:mm:ss"));

    const clearTodoList = useCallback(() => {
        setTodoList([]);
    }, [setTodoList]);

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
        setInterval(() => {
            setRemindTodo(checkRemindTodo());
        }, 5 * 1000);
    }, [setNow]);

    const removeTodo = useCallback(async (id) => {
        await axios.post('/todo/delete', {
            id
        });
        getTodolist();
    }, [getTodolist]);

    const filteredTodoList = todoList.filter((todo) => {
        if (showFinished) {
            return todo.Finished && isBefore1day(todo.Id);
        }
        return !todo.Finished || (todo.Finished && !isBefore1day(todo.Id));
    });
    if (showFinished) filteredTodoList.reverse();

    return (
        <div className="min-h-screen w-full flex flex-col justify-start items-center overflow-hidden">
            <h1 className="pan_y text-4xl sm:text-5xl md:text-7xl text-black dark:text-white font-mono font-black select-none">TodoList</h1>
            <h5 className="pan_y text-black dark:text-white font-mono text-xs md:text-base select-none">{now}</h5>
            <RefreshButton clearTodoList={clearTodoList} getTodolist={getTodolist} />
            <div className='pan relative top-3 w-full sm:w-11/12 md:w-3/4 lg:w-2/3'>
                <button className='relative bottom-2 left-3 ring-1 ring-green-500 dark:ring-green-700 text-white font-bold bg-green-400 dark:bg-green-600 px-3 rounded-sm focus:outline-none select-none' onClick={(_) => { setShowFinished(!showFinished); getTodolist() }}>{showFinished ? 'back' : 'check finished'}</button>
                {
                    remindTodo.length > 0 ?
                        (<span className='relative bottom-2 left-4 ring-1 ring-green-500 dark:ring-green-700 text-white font-bold bg-blue-400 dark:bg-blue-600 px-3 rounded-sm focus:outline-none select-none'>{`has ${remindTodo.length} remind todos`}</span>) : null
                }
            </div>
            {
                filteredTodoList.map((todo, i) => <Todo index={i} todo={todo} key={todo.Id} removeTodo={removeTodo} getTodolist={getTodolist} />)
            }
            <AddDialog getTodolist={getTodolist} />
        </div>
    )
}