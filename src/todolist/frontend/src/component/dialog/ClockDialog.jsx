import { Fragment, useState, useCallback, useRef, useEffect } from 'react';
import ReactDOM from 'react-dom';
import { addRemindTodo } from "../../util";
import moment from 'moment';

function Modal({ visible, confirm, cancel }) {
    const btnRef = useRef(null);
    const timeRef = useRef(null);
    const timeUnitRef = useRef(null);

    const onClick = useCallback(async (e) => {
        btnRef.current.classList.toggle('cursor-wait');
        btnRef.current.disabled = true;

        let time = parseInt(timeRef.current.value);
        let time_unit = timeUnitRef.current.value;
        let now = moment(new Date());

        if (isNaN(time) || time === 0) {
            alert('invalid input');
            btnRef.current.disabled = false;
            return;
        }

        switch (time_unit) {
            case 'second':
                now.add(time, 'seconds')
                break;
            case 'minute':
                now.add(time, 'minutes')
                break;
            case 'hour':
                now.add(time, 'hours')
                break;
            case 'day':
                now.add(time, 'days')
                break;
            default:
                break;
        }

        await confirm(now);

    }, [confirm, btnRef]);

    return visible && ReactDOM.createPortal(
        <div className={visible ? 'block' : 'hidden'}>
            <div className='w-full min-h-screen z-10 fixed top-0 left-0 flex items-center justify-center bg-gray-900 bg-opacity-80' onClick={cancel}>
                <div className='flex items-center z-20 justify-center py-6 w-11/12 md:w-2/3 lg:w-1/2 flex-col bg-gray-100 dark:bg-gray-800 rounded-xl' onClick={e => e.stopPropagation()}>
                    <div className='w-full flex items-center justify-center mb-5'>
                        <div>
                            <label htmlFor="time" className="block text-base font-medium text-black dark:text-white select-none">⏰ remind me later:</label>
                            <div className="mt-1 relative rounded-md shadow-sm">
                                <input
                                    ref={timeRef}
                                    type="text"
                                    name="time"
                                    id="time"
                                    className="ring-2 ring-indigo-500 dark:ring-gray-500  block w-full pl-7 pr-12 text-lg  dark:border-black rounded-md"
                                    placeholder="0"
                                />
                                <div className="absolute inset-y-0 right-0 flex items-center">
                                    <label htmlFor="time_unit" className="sr-only">time unit</label>
                                    <select
                                        ref={timeUnitRef}
                                        id="time_unit"
                                        name="time_unit"
                                        defaultValue={'minute'}
                                        className="h-full py-0 pl-2 pr-7 border-transparent bg-transparent text-gray-500 text-base font-bold rounded-md select-none"
                                    >
                                        <option>second</option>
                                        <option>minute</option>
                                        <option>hour</option>
                                        <option>day</option>
                                    </select>
                                </div>
                            </div>
                        </div>
                    </div>
                    <div className='flex items-center justify-evenly w-full'>
                        <button ref={btnRef} className='disabled:opacity-50 disabled:cursor-wait rounded-lg w-1/5 h-7 bg-blue-600 dark:bg-green-600 text-gray-100 dark:text-black font-bold select-none focus:outline-none' onClick={onClick}>confirm</button>
                        <button className='bg-gray-500 w-1/5 h-7 dark:bg-gray-100 text-gray-100 dark:text-black rounded-lg font-bold select-none focus:outline-none' onClick={(_) => cancel()}>cancel</button>
                    </div>
                </div>
            </div>
        </div>
        , document.body);
}

export function ClockDialog({ todo }) {

    const [visible, setVisible] = useState(false);

    useEffect(() => {
        const escape = event => { if (event.key === 'Escape') setVisible(false) };
        document.addEventListener('keydown', escape);

        return () => { document.removeEventListener('keydown', escape) };
    }, [setVisible]);

    const confirm = async (expire) => {
        todo.expire = expire;
        todo.retry = 0;
        addRemindTodo(todo);
        setVisible(false);
    };

    const cancel = (_) => {
        setVisible(false);
    };

    return (
        <Fragment>
            <Modal visible={visible} confirm={confirm} cancel={cancel} />
            <button title='clock' className='hidden lg:group-hover:block transform w-8 md:w-10 h-4/5 hover:scale-125 select-none focus:outline-none' onClick={() => setVisible(!visible)}>
                <svg className='w-full h-full fill-current' viewBox="0 0 500 500"><g>
                    <g>
                        <path d="M256.001,77.017c-107.656,0-195.244,87.589-195.244,195.244c0,107.662,87.589,195.25,195.244,195.25
			c107.662,0,195.244-87.589,195.244-195.25C451.245,164.606,363.657,77.017,256.001,77.017z M256.001,432.126
			c-88.143,0-159.853-71.715-159.853-159.858s71.709-159.853,159.853-159.853s159.853,71.71,159.853,159.853
			C415.854,360.416,344.144,432.126,256.001,432.126z"/>
                    </g>
                </g>
                    <g>
                        <g>
                            <path d="M310.268,266.363H263.08v-68.424c0-9.774-7.922-17.696-17.696-17.696c-9.774,0-17.696,7.922-17.696,17.696v86.12
			c0,9.774,7.922,17.696,17.696,17.696h64.885c9.774,0,17.696-7.922,17.696-17.696C327.964,274.285,320.042,266.363,310.268,266.363
			z"/>
                        </g>
                    </g>
                    <g>
                        <g>
                            <path d="M155.766,398.911c-7.267-6.542-18.457-5.946-24.992,1.315l-53.088,58.986c-6.542,7.261-5.946,18.451,1.315,24.987
			c3.38,3.05,7.615,4.548,11.833,4.548c4.843,0,9.668-1.976,13.16-5.863l53.088-58.986
			C163.623,416.636,163.027,405.446,155.766,398.911z"/>
                        </g>
                    </g>
                    <g>
                        <g>
                            <path d="M434.322,459.218l-53.088-58.986c-6.524-7.267-17.719-7.857-24.987-1.315c-7.267,6.536-7.851,17.725-1.315,24.987
			l53.088,58.986c3.486,3.881,8.311,5.857,13.154,5.857c4.212,0,8.447-1.498,11.833-4.542
			C440.274,477.669,440.858,466.479,434.322,459.218z"/>
                        </g>
                    </g>
                    <g>
                        <g>
                            <path d="M152.764,49.046c-35.162-34.43-91.841-34.377-126.342,0.13C9.256,66.335-0.123,89.039,0.001,113.105
			c0.13,23.777,9.556,46.039,26.55,62.685c3.445,3.368,7.91,5.049,12.381,5.049c4.53,0,9.06-1.734,12.511-5.179L152.899,74.204
			c3.339-3.344,5.203-7.881,5.179-12.605C158.055,56.873,156.143,52.355,152.764,49.046z M40.873,136.174
			c-3.545-7.143-5.439-15.047-5.48-23.258c-0.077-14.534,5.621-28.29,16.05-38.719c10.405-10.399,24.161-15.596,37.993-15.596
			c8.087,0,16.204,1.775,23.683,5.326L40.873,136.174z"/>
                        </g>
                    </g>
                    <g>
                        <g>
                            <path d="M485.581,49.17c-34.507-34.501-91.187-34.56-126.348-0.13c-3.374,3.309-5.291,7.828-5.315,12.552
			c-0.024,4.725,1.846,9.267,5.185,12.605l101.456,101.456c3.451,3.457,7.981,5.185,12.511,5.185c4.471,0,8.942-1.681,12.381-5.061
			c17-16.64,26.426-38.901,26.55-62.679C512.125,89.039,502.74,66.335,485.581,49.17z M471.13,136.174l-72.246-72.246
			c20.279-9.627,45.189-6.211,61.676,10.275c10.429,10.429,16.127,24.178,16.05,38.719
			C476.568,121.127,474.675,129.037,471.13,136.174z"/>
                        </g>
                    </g>
                </svg>
            </button>
        </Fragment>
    )
}
