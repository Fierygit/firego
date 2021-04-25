import { useCallback, useRef, useState } from 'react';
import { snowflake2moment } from '../util';
import { EditDialog } from "./EditDialog";
import { ConfirmDialog } from "./ConfirmDIalog";
import { ClockDialog } from "./ClockDialog";
import { DailyButton } from "./DailyButton";
import axios from 'axios';

export function Todo({ index, todo, removeTodo, getTodolist }) {

    const [todoName, setTodoName] = useState(todo.Name);

    const [finished, setFinished] = useState(todo.Finished);

    const [daily, setDaily] = useState(todo.Daily);

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

    const setDailyTodo = useCallback(async () => {
        await axios.post('/todo/daily', {
            id: todo.Id,
            daily: !daily
        });

        setDaily(!daily);
    }, [todo.Id, daily]);

    return (
        <div style={{ "animation": `pan ${index * 0.05 + 0.1}s ease-out 1` }} className={`flex relative group transform text-black dark:text-gray-200 w-full sm:w-11/12 md:w-3/4 lg:w-2/3 top-3 mb-1 sm:mb-2 shadow-sm hover:shadow-xl bg-white hover:bg-blue-200 dark:bg-black dark:hover:bg-blue-700 min-h-full md:h-10 items-center justify-start sm:rounded-lg border-black hover:translate-x-2 border-2 border-opacity-20 ${finished ? 'line-through' : ''}`}>
            <input className="form-tick appearance-none checked:bg-green-400 dark:checked:bg-green-600 border border-gray-600 dark:border-gray-400 select-none cursor-pointer mx-1 rounded-md h-6 w-6 focus:outline-none" checked={finished} type="checkbox" onChange={finishTodo}></input>
            <span className="font-bold font-sans text-lg md:text-xl block cursor-pointer" onClick={(_) => editBtnRef.current.click()}>
                {todoName}
            </span>
            <span className='absolute text-xs hidden sm:inline md:text-base right-2 md:right-4 select-none font-medium group-hover:hidden'>
                {snowflake2moment(todo.Id).calendar()}
            </span>
            <span className='absolute right-0 flex items-center justify-items-center h-full'>
                <DailyButton todo={todo} setDailyTodo={setDailyTodo}/>
                <ClockDialog todo={todo} />
                <EditDialog btnRef={editBtnRef} oldName={todoName} editTodo={editTodo} />
                <ConfirmDialog id={todo.Id} name={todoName} callback={removeTodo} />
            </span>
            {
                daily && !finished ?
                    <span className="flex absolute h-3 w-3 top-0 right-0 -mt-1 -mr-1">
                        <span className="animate-ping absolute inline-flex h-full w-full rounded-full bg-red-400 dark:bg-red-600 opacity-75"></span>
                        <span className="relative inline-flex rounded-full h-3 w-3 bg-red-600 dark:bg-red-800"></span>
                    </span>
                    :
                    null
            }
            {
                daily && finished ?
                    <span className="flex absolute h-3 w-3 top-0 right-0 -mt-1 -mr-1">
                        <span className="relative inline-flex rounded-full h-3 w-3 bg-green-600 dark:bg-green-800"></span>
                    </span>
                    :
                    null
            }
        </div>
    );
}
