import { useCallback, useRef } from 'react';
import { snowflake2moment } from '../util';
import { EditDialog } from "./dialog/EditDialog";
import { ConfirmDialog } from "./dialog/ConfirmDialog";
import { ClockDialog } from "./dialog/ClockDialog";
import { DailyButton } from "./button/DailyButton";
import { Action } from "../hook/useTodoList";
import axios from 'axios';

export function Todo({ index, todo, dispatch }) {
    const editBtnRef = useRef(null);

    const finishTodo = useCallback(async (_) => {
        const res = await axios.post('/todo/finish', {
            id: todo.Id,
            finished: !todo.Finished
        });

        dispatch({ type: Action.UPDATE_TODO, payload: { todo: res.data } });
    }, [todo.Id, todo.Finished, dispatch]);

    const removeTodo = useCallback(async () => {
        await axios.post('/todo/delete', {
            id: todo.Id
        });
        dispatch({ type: Action.REMOVE_TODO, payload: { id: todo.Id } });
    }, [dispatch, todo.Id]);

    const editTodo = useCallback(async (newName) => {
        const res = await axios.post('/todo/edit', {
            id: todo.Id,
            todo: newName
        });

        dispatch({ type: Action.UPDATE_TODO, payload: { todo: res.data } });
    }, [todo.Id, dispatch]);

    const setDailyTodo = useCallback(async () => {
        const res = await axios.post('/todo/daily', {
            id: todo.Id,
            daily: !todo.Daily
        });

        dispatch({ type: Action.UPDATE_TODO, payload: { todo: res.data } });
    }, [todo.Id, todo.Daily, dispatch]);

    return (
        <div style={{ "animation": `pan ${index * 0.05 + 0.1}s ease-out 1` }} className={`flex relative group transform text-black dark:text-gray-200 w-full sm:w-11/12 md:w-3/4 lg:w-2/3 top-3 mb-1 sm:mb-2 shadow-sm hover:shadow-xl bg-white hover:bg-blue-200 dark:bg-black dark:hover:bg-blue-700 min-h-full md:h-10 items-center justify-start sm:rounded-lg border-black hover:translate-x-2 border-2 border-opacity-20 ${todo.Finished ? 'line-through' : ''}`}>
            <input className="form-tick appearance-none checked:bg-green-400 dark:checked:bg-green-600 border border-gray-600 dark:border-gray-400 select-none cursor-pointer mx-1 rounded-md h-6 w-6 focus:outline-none" checked={todo.Finished} type="checkbox" onChange={finishTodo}></input>
            <span className="font-bold font-sans text-lg md:text-xl block cursor-pointer" onClick={(_) => editBtnRef.current.click()}>
                {todo.Name}
            </span>
            <span className='absolute text-xs hidden sm:inline md:text-base right-2 md:right-4 select-none font-medium group-hover:hidden'>
                {snowflake2moment(todo.Id).calendar()}
            </span>
            <span className='absolute right-0 flex items-center justify-items-center h-full'>
                <DailyButton todo={todo} setDailyTodo={setDailyTodo} />
                <ClockDialog todo={todo} />
                <EditDialog btnRef={editBtnRef} oldName={todo.Name} editTodo={editTodo} />
                <ConfirmDialog id={todo.Id} name={todo.Name} callback={removeTodo} />
            </span>
            {
                todo.Daily && !todo.Finished ?
                    <span className="flex absolute h-3 w-3 top-0 right-0 -mt-1 -mr-1">
                        <span className="animate-ping absolute inline-flex h-full w-full rounded-full bg-red-400 dark:bg-red-600 opacity-75"></span>
                        <span className="relative inline-flex rounded-full h-3 w-3 bg-red-600 dark:bg-red-800"></span>
                    </span>
                    :
                    null
            }
            {
                todo.Daily && todo.Finished ?
                    <span className="flex absolute h-3 w-3 top-0 right-0 -mt-1 -mr-1">
                        <span className="relative inline-flex rounded-full h-3 w-3 bg-green-600 dark:bg-green-800"></span>
                    </span>
                    :
                    null
            }
        </div>
    );
}
