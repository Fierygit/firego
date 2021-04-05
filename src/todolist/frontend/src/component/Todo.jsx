import { useCallback, useState } from 'react';
import axios from 'axios';

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
        <div className={`flex relative text-black w-11/12 md:w-4/5 m-1 bg-yellow-500 h-8 md:h-10 items-center rounded-xl border-2 border-opacity-20 ${finished ? 'line-through' : ''}`}>
            <input className="relative ml-2 checked:bg-blue-600 checked:border-transparent h-4 w-4 md:h-6 md:w-6" checked={finished} type="checkbox" onChange={finishTodo}></input>
            <span className="font-bold text-lg md:text-xl">
                {todo.Name}
            </span>
            <button className="absolute right-0 w-10 font-bold text-xl text-white bg-red-600 h-full rounded-xl" onClick={() => removeTodo(todo.Id)}> X </button>
        </div>
    );
}