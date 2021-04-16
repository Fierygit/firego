import { useCallback, useEffect, useRef, useState } from 'react';
import { snowflake2moment } from '../util';
import { EditDialog } from "./EditDialog";
import { ConfirmDialog } from "./ConfirmDIalog";
import axios from 'axios';

export function Todo({ todo, removeTodo }) {

    const [todoName, setTodoName] = useState(todo.Name);

    const [finished, setFinished] = useState(todo.Finished);

    const editBtnRef = useRef(null);

    useEffect(() => {
        console.log(editBtnRef.current);
    }, []);

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
        <div className={`flex relative group transform text-black dark:text-gray-200 w-11/12 md:w-4/5 m-1 hover:m-3 shadow-lg bg-white hover:bg-blue-200 dark:bg-black dark:hover:bg-blue-700 min-h-full md:h-10 items-center justify-start rounded-xl border-black hover:translate-x-2 border-2 border-opacity-20 ${finished ? 'line-through' : ''}`}>
            <input className="select-none cursor-pointer mx-1 border border-gray-300 rounded-md h-4 w-4 md:h-6 md:w-6" checked={finished} type="checkbox" onChange={finishTodo}></input>
            <span className="font-bold text-lg md:text-xl block cursor-pointer" onClick={(_) => editBtnRef.current.click()}>
                {todoName}
            </span>
            <span className='absolute text-xs md:text-base right-16 md:right-20 select-none font-medium'>
                {snowflake2moment(todo.Id)}
            </span>
            <EditDialog btnRef={editBtnRef} oldName={todoName} editTodo={editTodo} />
            <ConfirmDialog id={todo.Id} name={todoName} callback={removeTodo}></ConfirmDialog>
        </div>
    );
}