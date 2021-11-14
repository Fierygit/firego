import { useCallback, useRef, useState, useEffect, Fragment } from "react"
import ReactDOM from 'react-dom';
import { useHistory } from "react-router-dom";

function Modal({ todo, visible, confirm, cancel }) {
    const history = useHistory();
    const btnRef = useRef(null);
    const [daily, setDaily] = useState(todo.Daily);

    const onClick = useCallback(async (e) => {
        btnRef.current.disabled = true;

        await confirm();
        setDaily(!daily);
    }, [confirm, btnRef, daily]);

    const checkDailyRecord = useCallback(async (e) => {
        history.push(`/daily/${todo.Id}`);
    }, [todo.Id, history]);

    return visible && ReactDOM.createPortal(
        <div className={visible ? 'block' : 'hidden'}>
            <div className='w-full min-h-screen z-10 fixed top-0 left-0 flex items-center justify-center bg-gray-900 bg-opacity-80' onClick={cancel}>
                <div className='flex items-center z-20 justify-center py-6 w-11/12 max-w-md flex-col text-black dark:text-gray-100 bg-gray-100 dark:bg-gray-800 rounded-xl' onClick={e => e.stopPropagation()}>
                    <div className='flex items-center justify-evenly w-full text-gray-100 dark:text-black'>
                        <button className='rounded-lg px-3 max-w-1/5 h-7 bg-indigo-600 font-bold select-none focus:outline-none' onClick={checkDailyRecord}>check record</button>
                        <button ref={btnRef} className='disabled:opacity-50 disabled:cursor-wait rounded-lg px-3 max-w-1/5 h-7 bg-blue-600 dark:bg-green-600 font-bold select-none focus:outline-none' onClick={onClick}>
                            {
                                !daily ?
                                    <span className='font-bold md:text-lg select-none'>
                                        mark as daily
                                    </span>
                                    :
                                    <span className='font-bold md:text-lg select-none'>
                                        unmark
                                    </span>
                            }
                        </button>
                        <button className='bg-gray-500 h-7 dark:bg-gray-100 px-3 max-w-1/5 rounded-lg font-bold select-none focus:outline-none' onClick={(_) => cancel()}>cancel</button>
                    </div>
                </div>
            </div>
        </div>
        , document.body);
}

export function DailyButton({ todo, setDailyTodo }) {
    const [visible, setVisible] = useState(false);

    useEffect(() => {
        const escape = event => { if (event.key === 'Escape') setVisible(false) };
        document.addEventListener('keydown', escape);

        return () => { document.removeEventListener('keydown', escape) };
    }, [setVisible]);

    const confirm = async (newName) => {
        await setDailyTodo();
        setVisible(false);
    };

    const cancel = (_) => {
        setVisible(false);
    };

    return (
        <Fragment>
            <Modal todo={todo} visible={visible} confirm={confirm} cancel={cancel} />
            <button title='daily todo' className='hidden group-hover:block disabled:opacity-50 disabled:cursor-wait transform w-8 md:w-10 h-4/5 hover:scale-125 select-none focus:outline-none' onClick={_ => setVisible(!visible)}>
                <svg className='w-full h-full fill-current' viewBox="0 -20 400 380">
                    <g>
                        <g>
                            <path d="M230.902,167.822c0-5.44,0.276-11.013,0.709-16.876h-0.564c-2.942,5.864-5.297,11.151-8.371,16.876l-12.13,19.804v0.275
			h20.356V167.822z"/>
                            <path d="M180.102,0C80.81,0,0.006,80.797,0.006,180.09c0,99.289,80.798,180.096,180.096,180.096
			c99.296,0,180.078-80.795,180.078-180.09C360.18,80.803,279.397,0,180.102,0z M323.064,187.938
			c-2.294,2.654-5.537,4.288-9.05,4.54c-0.318,0.023-0.643,0.048-0.961,0.048c-0.078,0-0.149-0.048-0.216-0.048
			c-0.246,0.012-0.475,0.072-0.727,0.072c-4.744,0-8.869-2.517-11.205-6.281l-24.02-20.741c-4.047-3.498-5.59-9.082-3.915-14.166
			c1.688-5.074,6.27-8.641,11.596-9.034l8.353-0.606c-15.913-49.101-62.029-84.728-116.364-84.728
			c-67.455,0-122.333,54.878-122.333,122.339c0,67.452,54.877,122.336,122.333,122.336c39.046,0,76.083-18.903,99.07-50.549
			c4.288-5.896,12.556-7.229,18.471-2.93c5.909,4.287,7.218,12.562,2.931,18.471c-27.946,38.484-72.982,61.453-120.471,61.453
			c-82.041,0-148.787-66.738-148.787-148.781c0-82.043,66.747-148.784,148.787-148.784c68.307,0,125.887,46.342,143.278,109.204
			l16.026-1.168c5.267-0.454,10.382,2.459,12.778,7.25c2.42,4.777,1.688,10.545-1.808,14.586L323.064,187.938z M190.742,203.79
			v-13.787l34.3-55.231h25.935v53.143h10.887v15.889h-10.887v21.623h-20.086v-21.623h-40.148V203.79z M142.44,150.249
			c-7.81,0-14.634,3.903-19.375,7.539l-6.002-15.208c6.839-5.146,17.438-9.353,29.706-9.353c20.5,0,31.795,11.998,31.795,28.451
			c0,15.207-11.013,27.34-24.121,39.049l-8.368,6.984v0.27h34.167v17.438h-64.291v-12.839l11.707-10.592
			c19.812-17.711,29.436-27.896,29.706-38.488C157.371,156.109,152.912,150.249,142.44,150.249z"/>
                        </g>
                    </g>
                </svg>
            </button>
        </Fragment>
    )
}