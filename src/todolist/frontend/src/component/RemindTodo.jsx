import { useEffect, useState } from 'react'
import { checkRemindTodo } from "../util";

export function RemindTodo() {
    const [remindTodo, setRemindTodo] = useState([]);

    useEffect(() => {
        let timer = setInterval(() => {
            setRemindTodo(checkRemindTodo());
        }, 1000);

        return () => { clearInterval(timer) };
    }, []);

    return remindTodo.length > 0 ?
        <span className='relative bottom-2 left-4 ring-1 ring-green-500 dark:ring-green-700 text-white font-bold bg-blue-400 dark:bg-blue-600 px-3 rounded-sm focus:outline-none select-none'>
            {`has ${remindTodo.length} remind todos`}
        </span>
        :
        null
}
