import { useState, useEffect, useCallback } from 'react';
import { RemindTodo } from "./RemindTodo";
import { RefreshButton } from './button/RefreshButton';
import { Action } from "../hook/useTodoList";
import axios from 'axios';

export function FilterTab({ dispatch }) {
    const [filterType, setFilterType] = useState('unfinished');

    const getTodolist = useCallback(async () => {
        dispatch({ type: Action.CLEAR_TODO_LIST });
        switch (filterType) {
            case 'all': {
                const res = await axios.get('/todo?type=all');
                dispatch({ type: Action.SET_TODO_LIST, payload: { todoList: res.data } });
                break;
            }
            case 'finished': {
                const res = await axios.get('/todo?type=finished');
                dispatch({ type: Action.SET_TODO_LIST, payload: { todoList: res.data } });
                break;
            }
            case 'unfinished':
            default: {
                const res = await axios.get('/todo?type=unfinished');
                dispatch({ type: Action.SET_TODO_LIST, payload: { todoList: res.data } });
                break;
            }
        }
    }, [dispatch, filterType]);

    useEffect(() => {
        getTodolist();
    }, [dispatch, filterType, getTodolist]);

    return (
        <div className='pan relative top-1 text-sm md:text-base w-full sm:w-11/12 md:w-3/4 lg:w-2/3 left-3'>
            <button disabled={filterType === 'unfinished'} className='disabled:opacity-50 disabled:cursor-default relative ring-1 ring-blue-700 dark:ring-green-700 text-white font-bold bg-blue-600 dark:bg-green-600 px-3 rounded-sm focus:outline-none select-none' onClick={(_) => { setFilterType('unfinished'); }}>
                unfinished
            </button>
            <button disabled={filterType === 'finished'} className='disabled:opacity-50 disabled:cursor-default relative ml-3 ring-1 ring-blue-500 dark:ring-green-700 text-white font-bold bg-blue-600 dark:bg-green-600 px-3 rounded-sm focus:outline-none select-none' onClick={(_) => { setFilterType('finished'); }}>
                finished
            </button>
            <button disabled={filterType === 'all'} className='disabled:opacity-50 disabled:cursor-default relative ml-3 ring-1 ring-blue-500 dark:ring-green-700 text-white font-bold bg-blue-600 dark:bg-green-600 px-3 rounded-sm focus:outline-none select-none' onClick={(_) => { setFilterType('all'); }}>
                all
            </button>
            <RemindTodo />
            <RefreshButton getTodolist={getTodolist} />
        </div>
    )
}
