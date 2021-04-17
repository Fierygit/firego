import { Fragment, useState, useCallback, useRef, useEffect } from 'react';
import ReactDOM from 'react-dom';
import axios from "axios";

function Modal({ visible, confirm, cancel }) {
    const [todoName, setTodoName] = useState("");
    const btnRef = useRef(null);

    const onClick = useCallback(async (e) => {
        btnRef.current.classList.toggle('cursor-wait');
        btnRef.current.disabled = true;

        await confirm(todoName);
        setTodoName("");

    }, [confirm, setTodoName, todoName, btnRef]);

    const onKeyPress = useCallback(async (e) => {
        if (e.which !== 13) return;

        btnRef.current.click();
    }, []);

    return visible && ReactDOM.createPortal(
        <div className={visible ? 'block' : 'hidden'}>
            <div className='w-full min-h-screen z-10 fixed top-0 left-0 flex items-center justify-center bg-gray-900 bg-opacity-80' onClick={cancel}>
                <div className='flex items-center z-20 justify-center py-6 w-11/12 md:w-2/3 lg:w-1/2 flex-col bg-gray-100 dark:bg-gray-800 rounded-xl' onClick={e => e.stopPropagation()}>
                    <div className='w-full flex items-center justify-center mb-5'>
                        <span className='font-bold md:text-lg dark:text-gray-100 select-none'>add todo:</span>
                        <input className='text-white font-bold dark:text-gray-700 bg-gray-700 dark:bg-white ml-1 md:ml-5 outline-none pl-2 h-9 w-3/4 md:text-xl rounded-xl' autoFocus type="text" value={todoName} onInput={e => setTodoName(e.target.value)} onKeyPress={onKeyPress} />
                    </div>
                    <div className='flex items-center justify-evenly w-full'>
                        <button ref={btnRef} className='disabled:opacity-50 rounded-lg w-1/5 h-7 bg-green-400 dark:bg-green-600 text-gray-100 dark:text-black font-bold select-none' onClick={onClick}>confirm</button>
                        <button className='bg-gray-500 w-1/5 h-7 dark:bg-gray-100 text-gray-100 dark:text-black rounded-lg font-bold select-none' onClick={(_) => cancel()}>cancel</button>
                    </div>
                </div>
            </div>
        </div>
        , document.body);
}

export function AddDialog({ getTodolist }) {

    const [visible, setVisible] = useState(false);

    const addBtnRef = useRef(null);

    useEffect(() => {
        document.addEventListener('keydown', (event) => {
            const keyName = event.key;

            if (keyName === 'Control') {
                return;
            }

            event.key === 'Escpase' && setVisible(false)

            event.ctrlKey && keyName === 'c' && addBtnRef.current.click();

        }, false);

    }, [addBtnRef, setVisible]);

    const addTodo = useCallback(async (todoName) => {
        if (todoName === '') {
            alert(`todo can not be empty!!!`);
            return;
        }

        await axios.post('/todo', {
            todo: todoName
        });

        await getTodolist();

    }, [getTodolist]);

    const confirm = async (todoName) => {
        await addTodo(todoName);
        setVisible(false);
    };

    const cancel = (_) => {
        setVisible(false);
    };

    return (
        <Fragment>
            <Modal visible={visible} confirm={confirm} cancel={cancel} />
            <div className='sticky items-end bottom-0 mt-auto w-full opacity-100' >
                <button ref={addBtnRef} className="disabled:opacity-50 absolute right-7 md:right-10 bottom-7 md:bottom-10 transform shadow-lg text-xl rounded-full bg-green-400 dark:bg-green-600 hover:scale-110 font-bold px-3 py-3 h-12 w-12 md:h-16 md:w-16 text-gray-200 outline-none" onClick={(_) => setVisible(!visible)}>
                    <svg className='w-full h-full fill-current' viewBox="0 0 512 512" >
                        <path d="M492,236H276V20c0-11.046-8.954-20-20-20c-11.046,0-20,8.954-20,20v216H20c-11.046,0-20,8.954-20,20s8.954,20,20,20h216 v216c0,11.046,8.954,20,20,20s20-8.954,20-20V276h216c11.046,0,20-8.954,20-20C512,244.954,503.046,236,492,236z" />
                    </svg>
                    <span className='text-sm text-black dark:text-white pt-4 inline-block'>
                        ctrl+c
                    </span>
                </button>
            </div>
        </Fragment>
    )
}
