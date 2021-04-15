import { Fragment, useState } from 'react';
import ReactDOM from 'react-dom';

function Modal({ oldName, visible, confirm, cancel }) {
    const [todoName, setTodoName] = useState(oldName);

    return visible && ReactDOM.createPortal(
        <div className={visible ? 'block' : 'hidden'}>
            <div className='w-full min-h-screen z-10 fixed top-0 left-0 flex items-center justify-center bg-gray-900 bg-opacity-80' onClick={cancel}>
                <div className='flex items-center z-20 justify-center py-6 w-11/12 md:w-1/2 lg:w-1/3 flex-col bg-gray-100 dark:bg-gray-800 rounded-xl' onClick={e => e.stopPropagation()}>
                    <div className='w-full flex items-center justify-center mb-5'>
                        <span className='font-bold md:text-lg dark:text-gray-100 select-none'>change to:</span>
                        <input className='text-white font-bold dark:text-gray-700 bg-gray-700 dark:bg-white ml-1 md:ml-5 outline-none pl-2 h-9 w-3/5 md:text-xl rounded-xl' type="text" value={todoName} onInput={e => setTodoName(e.target.value)} />
                    </div>
                    <div className='flex items-center justify-evenly w-full'>
                        <button className='rounded-lg w-1/4 h-7 bg-green-400 text-gray-100 dark:text-black font-bold select-none' onClick={(_) => confirm(todoName)}>confirm</button>
                        <button className='bg-gray-500 h-7 dark:bg-gray-100 w-1/4 text-gray-100 dark:text-black rounded-lg font-bold select-none' onClick={(_) => cancel()}>cancel</button>
                    </div>
                </div>
            </div>
        </div>
        , document.body);
}

export function EditDialog({ oldName, editTodo }) {

    const [visible, setVisible] = useState(false);

    const confirm = async (newName) => {
        await editTodo(newName);
        setVisible(false);
    };

    const cancel = (_) => {
        setVisible(false);
    };

    return (
        <Fragment>
            <Modal oldName={oldName} visible={visible} confirm={confirm} cancel={cancel} />
            <button className='absolute hidden group-hover:block transform right-6 md:right-9 w-8 md:w-10 h-4/5 hover:scale-125 outline-none' onClick={() => setVisible(!visible)}>
                <svg className='w-full h-full fill-current' viewBox="0 -1 401.52289 401"><path d="m370.589844 250.972656c-5.523438 0-10 4.476563-10 10v88.789063c-.019532 16.5625-13.4375 29.984375-30 30h-280.589844c-16.5625-.015625-29.980469-13.4375-30-30v-260.589844c.019531-16.558594 13.4375-29.980469 30-30h88.789062c5.523438 0 10-4.476563 10-10 0-5.519531-4.476562-10-10-10h-88.789062c-27.601562.03125-49.96875 22.398437-50 50v260.59375c.03125 27.601563 22.398438 49.96875 50 50h280.589844c27.601562-.03125 49.96875-22.398437 50-50v-88.792969c0-5.523437-4.476563-10-10-10zm0 0" /><path d="m376.628906 13.441406c-17.574218-17.574218-46.066406-17.574218-63.640625 0l-178.40625 178.40625c-1.222656 1.222656-2.105469 2.738282-2.566406 4.402344l-23.460937 84.699219c-.964844 3.472656.015624 7.191406 2.5625 9.742187 2.550781 2.546875 6.269531 3.527344 9.742187 2.566406l84.699219-23.464843c1.664062-.460938 3.179687-1.34375 4.402344-2.566407l178.402343-178.410156c17.546875-17.585937 17.546875-46.054687 0-63.640625zm-220.257812 184.90625 146.011718-146.015625 47.089844 47.089844-146.015625 146.015625zm-9.40625 18.875 37.621094 37.625-52.039063 14.417969zm227.257812-142.546875-10.605468 10.605469-47.09375-47.09375 10.609374-10.605469c9.761719-9.761719 25.589844-9.761719 35.351563 0l11.738281 11.734375c9.746094 9.773438 9.746094 25.589844 0 35.359375zm0 0" /></svg>
            </button>
        </Fragment>
    )
}
