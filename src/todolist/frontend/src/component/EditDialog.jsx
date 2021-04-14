import { Fragment, useState } from 'react';
// import ReactDOM from 'react-dom';

function Modal({ oldName, visible, confirm, cancel }) {
    const [todoName, setTodoName] = useState(oldName);

    // return visible && ReactDOM.createPortal(
    //     <div className={'fixed z-10 flex-col left-0 top-0 w-96 rounded-2xl bg-black' + visible ? 'block' : 'hidden'}>
    //         <div>
    //             <input className='text-black dark:text-gray-700 bg-gray-100 dark:bg-white border border-black outline-none w-96 h-9 text-xl rounded-xl mb-3' type="text" value={todoName} onInput={e => setTodoName(e.target.value)} />
    //         </div>
    //         <div className='flex items-center justify-center'>
    //             <button className='w-20 mr-16 rounded-lg bg-green-400' onClick={(_) => confirm(todoName)}>confirm</button>
    //             <button className='w-20 bg-gray-100 text-black rounded-lg' onClick={(_) => cancel()}>cancel</button>
    //         </div>
    //     </div>
    //     , document.body);
    return visible && (
        <div className={'fixed z-10 flex-col left-0 top-0 w-96 rounded-2xl bg-black' + visible ? 'block' : 'hidden'}>
            <div>
                <input className='text-black dark:text-gray-700 bg-gray-100 dark:bg-white border border-black outline-none w-96 h-9 text-xl rounded-xl mb-3' type="text" value={todoName} onInput={e => setTodoName(e.target.value)} />
            </div>
            <div className='flex items-center justify-center'>
                <button className='w-20 mr-16 rounded-lg bg-green-400' onClick={(_) => confirm(todoName)}>confirm</button>
                <button className='w-20 bg-gray-100 text-black rounded-lg' onClick={(_) => cancel()}>cancel</button>
            </div>
        </div>
    );
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
