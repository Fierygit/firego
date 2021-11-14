import { Todo } from './Todo';
import { AddDialog } from "./dialog/AddDialog";
import { TimeTitle } from "./TimeTitle";
import { FilterTab } from "./FilterTab";
import { useTodoList } from "../hook/useTodoList";

export function TodoList() {
    const [state, dispatch] = useTodoList();

    return (
        <div className="min-h-screen w-full flex flex-col justify-start items-center overflow-hidden">
            <h1 className="pan_y text-5xl sm:text-6xl lg:text-8xl text-black dark:text-white font-mono font-black select-none">Todo</h1>
            <TimeTitle />
            <FilterTab dispatch={dispatch} />
            {
                state.todoList.map((todo, i) => <Todo index={i} todo={todo} key={todo.Id} dispatch={dispatch} />)
            }
            <AddDialog dispatch={dispatch} />
        </div>
    )
}
