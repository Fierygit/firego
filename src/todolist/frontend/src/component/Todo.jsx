import { useCallback, useRef, useState } from 'react';
import { snowflake2moment } from '../util';
import { EditDialog } from "./EditDialog";
import { ConfirmDialog } from "./ConfirmDIalog";
import { ClockDialog } from "./ClockDialog";
import axios from 'axios';

export function Todo({ index, todo, removeTodo, getTodolist }) {

    const [todoName, setTodoName] = useState(todo.Name);

    const [finished, setFinished] = useState(todo.Finished);

    const editBtnRef = useRef(null);

    const finishTodo = useCallback(async (_) => {
        await axios.post('/todo/finish', {
            id: todo.Id,
            finished: !finished
        });

        setFinished(!finished);
    }, [finished, todo.Id]);

    const editTodo = useCallback(async (newName) => {
        await axios.post('/todo/edit', {
            id: todo.Id,
            todo: newName
        });

        setTodoName(newName);
    }, [todo.Id, setTodoName]);

    return (
        <div style={{ "animation": `pan ${index * 0.05 + 0.1}s linear 1` }} className={`flex relative group transform text-black dark:text-gray-200 w-full sm:w-11/12 md:w-3/4 lg:w-2/3 top-3 mb-1 sm:mb-2 shadow-lg bg-white hover:bg-blue-200 dark:bg-black dark:hover:bg-blue-700 min-h-full md:h-10 items-center justify-start sm:rounded-lg border-black hover:translate-x-2 border-2 border-opacity-20 ${finished ? 'line-through' : ''}`}>
            <input className="form-tick appearance-none checked:bg-green-400 dark:checked:bg-green-600 border border-gray-600 dark:border-gray-400 select-none cursor-pointer mx-1 rounded-md h-4 w-4 md:h-6 md:w-6 focus:outline-none" checked={finished} type="checkbox" onChange={finishTodo}></input>
            <span className="font-bold font-mono text-sm sm:text-lg md:text-xl block cursor-pointer" onClick={(_) => editBtnRef.current.click()}>
                {todoName}
            </span>
            <span className='absolute text-xs hidden sm:inline md:text-base right-2 md:right-4 select-none font-medium group-hover:hidden'>
                {snowflake2moment(todo.Id).calendar()}
            </span>
            <ClockDialog todo={todo}></ClockDialog>
            <EditDialog btnRef={editBtnRef} oldName={todoName} editTodo={editTodo} />
            <ConfirmDialog id={todo.Id} name={todoName} callback={removeTodo}></ConfirmDialog>
        </div>
    );
}
