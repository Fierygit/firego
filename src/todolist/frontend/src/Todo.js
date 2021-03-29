import { useCallback, useState } from 'react';
import axios from 'axios';
import { URL } from "./const";

export function Todo({ todo, removeTodo }) {

    const [finished, setFinished] = useState(todo.Finished);

    const [className, setClassName] = useState(todo.Finished ? "todo line-through" : "todo");

    const finishTodo = useCallback(async (e) => {
        if (finished === true) {
            setClassName("todo");
        } else {
            setClassName("todo line-through");
        }

        await axios.post(`${URL}/todo/finish`, {
            id: todo.Id,
            finished: !finished
        });

        setFinished(!finished);
    }, [finished, todo.Id]);

    return (
        <div className={className}>
            <input checked={finished} className="todo-finish" type="checkbox" onChange={finishTodo}></input>
            <span className="todo-name">
                {todo.Name}
            </span>
            <button className="remove-todo-btn" onClick={() => removeTodo(todo.Id)}> delete </button>
        </div>
    );
}