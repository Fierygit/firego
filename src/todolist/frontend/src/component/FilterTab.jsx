import { useState } from 'react';
import { RemindTodo } from "./RemindTodo";
import { isBefore1day } from '../util';

export function FilterTab({ todoList }) {
    const [showFinished, setShowFinished] = useState(false);

    const dailyTodoList = [];
    let filteredTodoList = todoList.filter((todo) => {
        if (todo.Daily) {
            dailyTodoList.push(todo);
            return false;
        }
        if (showFinished) return todo.Finished && isBefore1day(todo.Id);

        return !todo.Finished || (todo.Finished && !isBefore1day(todo.Id));
    });
    if (showFinished) filteredTodoList.reverse();
    filteredTodoList = [...dailyTodoList, ...filteredTodoList];

    return (
        <div className='pan relative top-3 w-full sm:w-11/12 md:w-3/4 lg:w-2/3'>
            <button className='relative bottom-2 left-3 ring-1 ring-green-500 dark:ring-green-700 text-white font-bold bg-green-400 dark:bg-green-600 px-3 rounded-sm focus:outline-none select-none' onClick={(_) => { setShowFinished(!showFinished); }}>{showFinished ? 'back' : 'check finished'}</button>
            <RemindTodo />
        </div>
    )
}
