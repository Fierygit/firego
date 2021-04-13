import { useCallback, useState } from 'react';
import axios from 'axios';
import { snowflake2moment } from '../util';

export function Todo({ todo, removeTodo }) {

    const [finished, setFinished] = useState(todo.Finished);

    const finishTodo = useCallback(async (e) => {
        await axios.post('/todo/finish', {
            id: todo.Id,
            finished: !finished
        });

        setFinished(!finished);
    }, [finished, todo.Id]);

    return (
        <div className={`flex relative text-black w-11/12 md:w-4/5 m-1 bg-blue-300 hover:bg-blue-500 dark:bg-yellow-500 dark:hover:bg-yellow-700 h-8 md:h-10 items-center rounded-xl border-black dark:border-white border-2 border-opacity-20 ${finished ? 'line-through' : ''}`}>
            <input className="relative mx-2 checked:bg-blue-600 checked:border-transparent h-4 w-4 md:h-6 md:w-6" checked={finished} type="checkbox" onChange={finishTodo}></input>
            <span className="font-bold text-lg md:text-xl">
                {todo.Name}
            </span>
            <span className={'absolute text-xs md:text-base right-8 md:right-12'}>
                {snowflake2moment(todo.Id)}
            </span>
            <button className="absolute right-0 w-8 md:w-10 font-bold text-xl text-white bg-red-600 hover:bg-red-800 h-full rounded-xl" onClick={() => removeTodo(todo.Id)}> X </button>
        </div>
    );
}