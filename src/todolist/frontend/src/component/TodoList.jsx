import { useState } from 'react';
import { Todo } from './Todo';
import { AddDialog } from "./dialog/AddDialog";
import { RefreshButton } from './button/RefreshButton';
import { RemindTodo } from "./RemindTodo";
import { TimeTitle } from "./TimeTitle";
import { isBefore1day } from '../util';
import { useTodoList } from "../hook/useTodoList";

export function TodoList() {
    const [state, dispatch] = useTodoList();

    const [showFinished, setShowFinished] = useState(false);

    const dailyTodoList = [];
    let filteredTodoList = state.todoList.filter((todo) => {
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
        <div className="min-h-screen w-full flex flex-col justify-start items-center overflow-hidden">
            <h1 className="pan_y text-5xl sm:text-6xl lg:text-8xl text-black dark:text-white font-mono font-black select-none">Todo</h1>
            <TimeTitle />
            <RefreshButton dispatch={dispatch} />
            <div className='pan relative top-3 w-full sm:w-11/12 md:w-3/4 lg:w-2/3'>
                <button className='relative bottom-2 left-3 ring-1 ring-green-500 dark:ring-green-700 text-white font-bold bg-green-400 dark:bg-green-600 px-3 rounded-sm focus:outline-none select-none' onClick={(_) => { setShowFinished(!showFinished); }}>{showFinished ? 'back' : 'check finished'}</button>
                <RemindTodo />
            </div>
            {
                filteredTodoList.map((todo, i) => <Todo index={i} todo={todo} key={todo.Id} dispatch={dispatch} />)
            }
            <AddDialog dispatch={dispatch} />
        </div>
    )
}
