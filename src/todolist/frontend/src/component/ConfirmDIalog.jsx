import { Fragment, useState, useRef, useCallback, useEffect } from 'react';
import ReactDOM from 'react-dom';

function Modal({ name, visible, confirm, cancel }) {
    const btnRef = useRef(null);

    const onClick = useCallback(async (e) => {
        btnRef.current.disabled = true;

        await confirm();
    }, [confirm, btnRef]);

    return visible && ReactDOM.createPortal(
        <div className={visible ? 'block' : 'hidden'}>
            <div className='w-full min-h-screen z-10 fixed top-0 left-0 flex items-center justify-center bg-gray-900 bg-opacity-80' onClick={cancel}>
                <div className='flex items-center z-20 justify-center py-6 w-11/12 max-w-md flex-col bg-gray-100 dark:bg-gray-800 rounded-xl' onClick={e => e.stopPropagation()}>
                    <div className='w-full flex items-center justify-center mb-5'>
                        <span className='font-bold md:text-lg dark:text-gray-100 select-none'>delete todo: {name} ?</span>
                    </div>
                    <div className='flex items-center justify-evenly w-full'>
                        <button ref={btnRef} className='disabled:opacity-50 disabled:cursor-wait rounded-lg w-1/5 h-7 bg-red-600 dark:bg-red-800 text-gray-100 dark:text-black font-bold select-none focus:outline-none' onClick={onClick}>confirm</button>
                        <button className='bg-gray-500 h-7 dark:bg-gray-100 w-1/5 text-gray-100 dark:text-black rounded-lg font-bold select-none focus:outline-none' onClick={(_) => cancel()}>cancel</button>
                    </div>
                </div>
            </div>
        </div>
        , document.body);
}

export function ConfirmDialog({ id, name, callback }) {

    const [visible, setVisible] = useState(false);

    useEffect(() => {
        document.addEventListener('keydown', event => event.key === 'Escape' && setVisible(false), false);
    }, [setVisible]);

    const confirm = async () => {
        await callback(id);
        setVisible(false);
    };

    const cancel = (_) => {
        setVisible(false);
    };

    return (
        <Fragment>
            <Modal name={name} visible={visible} confirm={confirm} cancel={cancel} />
            <button title='delete' className="absolute hidden group-hover:block transform right-0 w-8 md:w-10 h-4/5 hover:scale-125 focus:outline-none" onClick={() => setVisible(!visible)}>
                <svg className='w-full h-full fill-current' viewBox="-40 0 427 427.00131"><path d="m232.398438 154.703125c-5.523438 0-10 4.476563-10 10v189c0 5.519531 4.476562 10 10 10 5.523437 0 10-4.480469 10-10v-189c0-5.523437-4.476563-10-10-10zm0 0" /><path d="m114.398438 154.703125c-5.523438 0-10 4.476563-10 10v189c0 5.519531 4.476562 10 10 10 5.523437 0 10-4.480469 10-10v-189c0-5.523437-4.476563-10-10-10zm0 0" /><path d="m28.398438 127.121094v246.378906c0 14.5625 5.339843 28.238281 14.667968 38.050781 9.285156 9.839844 22.207032 15.425781 35.730469 15.449219h189.203125c13.527344-.023438 26.449219-5.609375 35.730469-15.449219 9.328125-9.8125 14.667969-23.488281 14.667969-38.050781v-246.378906c18.542968-4.921875 30.558593-22.835938 28.078124-41.863282-2.484374-19.023437-18.691406-33.253906-37.878906-33.257812h-51.199218v-12.5c.058593-10.511719-4.097657-20.605469-11.539063-28.03125-7.441406-7.421875-17.550781-11.5546875-28.0625-11.46875h-88.796875c-10.511719-.0859375-20.621094 4.046875-28.0625 11.46875-7.441406 7.425781-11.597656 17.519531-11.539062 28.03125v12.5h-51.199219c-19.1875.003906-35.394531 14.234375-37.878907 33.257812-2.480468 19.027344 9.535157 36.941407 28.078126 41.863282zm239.601562 279.878906h-189.203125c-17.097656 0-30.398437-14.6875-30.398437-33.5v-245.5h250v245.5c0 18.8125-13.300782 33.5-30.398438 33.5zm-158.601562-367.5c-.066407-5.207031 1.980468-10.21875 5.675781-13.894531 3.691406-3.675781 8.714843-5.695313 13.925781-5.605469h88.796875c5.210937-.089844 10.234375 1.929688 13.925781 5.605469 3.695313 3.671875 5.742188 8.6875 5.675782 13.894531v12.5h-128zm-71.199219 32.5h270.398437c9.941406 0 18 8.058594 18 18s-8.058594 18-18 18h-270.398437c-9.941407 0-18-8.058594-18-18s8.058593-18 18-18zm0 0" /><path d="m173.398438 154.703125c-5.523438 0-10 4.476563-10 10v189c0 5.519531 4.476562 10 10 10 5.523437 0 10-4.480469 10-10v-189c0-5.523437-4.476563-10-10-10zm0 0" /></svg>
            </button>
        </Fragment >
    )
}
