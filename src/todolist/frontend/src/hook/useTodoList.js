import { useReducer, useEffect } from "react";
import axios from 'axios';

export const Action = {
    SET_TODO_LIST: 'getTodoList',
    CLEAR_TODO_LIST: 'cleatTodoList',
    REMOVE_TODO: 'removeTodo',
    ADD_TODO: 'addTodo',
    UPDATE_TODO: 'updateTodo',
}

function reducer(state, action) {
    switch (action.type) {
        case Action.SET_TODO_LIST:
            return { todoList: action.payload.todoList };
        case Action.CLEAR_TODO_LIST:
            return { todoList: [] };
        case Action.REMOVE_TODO:
            return { todoList: state.todoList.filter((todo) => todo.Id !== action.payload.id) };
        case Action.ADD_TODO:
            return { todoList: [...state.todoList, action.payload.todo] }
        case Action.UPDATE_TODO:
            const filtered = state.todoList.map((todo) => {
                if (todo.Id !== action.payload.todo.Id)
                    return todo;

                return action.payload.todo;
            });
            return { todoList: [...filtered] }
        default:
            throw new Error();
    }
}

export function useTodoList() {
    const [state, dispatch] = useReducer(reducer, { todoList: [] });

    useEffect(() => {
        const getTodolist = async () => {
            const res = await axios.get('/todo?type=unfinished');
            dispatch({ type: Action.SET_TODO_LIST, payload: { todoList: res.data } });
        };
        getTodolist();
    }, []);

    return [state, dispatch];
}