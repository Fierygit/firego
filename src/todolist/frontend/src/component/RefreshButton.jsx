
import { useCallback, useRef } from 'react';

export function RefreshButton({ clearTodoList, getTodolist }) {
    const btnRef = useRef(null);

    const click = useCallback(async (e) => {
        btnRef.current.disabled = true;
        clearTodoList();
        await getTodolist();
        btnRef.current.disabled = false;
    }, [clearTodoList, getTodolist]);

    return (
        <div className='relative w-full sm:w-11/12 md:w-3/4 lg:w-2/3 h-6'>
            <button ref={btnRef} className='absolute right-5 bottom-1 w-5 h-5 md:w-6 md:h-6 text-black dark:text-white disabled:opacity-50 focus:outline-none' onClick={click}>
                <svg className='w-full h-full fill-current' viewBox="0 0 512 512" >
                    <g>
                        <g>
                            <path d="M493.815,70.629c-11.001-1.003-20.73,7.102-21.733,18.102l-2.65,29.069C424.473,47.194,346.429,0,256,0
			C158.719,0,72.988,55.522,30.43,138.854c-5.024,9.837-1.122,21.884,8.715,26.908c9.839,5.024,21.884,1.123,26.908-8.715
			C102.07,86.523,174.397,40,256,40c74.377,0,141.499,38.731,179.953,99.408l-28.517-20.367c-8.989-6.419-21.48-4.337-27.899,4.651
			c-6.419,8.989-4.337,21.479,4.651,27.899l86.475,61.761c12.674,9.035,30.155,0.764,31.541-14.459l9.711-106.53
			C512.919,81.362,504.815,71.632,493.815,70.629z"/>
                        </g>
                    </g>
                    <g>
                        <g>
                            <path d="M472.855,346.238c-9.838-5.023-21.884-1.122-26.908,8.715C409.93,425.477,337.603,472,256,472
			c-74.377,0-141.499-38.731-179.953-99.408l28.517,20.367c8.989,6.419,21.479,4.337,27.899-4.651
			c6.419-8.989,4.337-21.479-4.651-27.899l-86.475-61.761c-12.519-8.944-30.141-0.921-31.541,14.459l-9.711,106.53
			c-1.003,11,7.102,20.73,18.101,21.733c11.014,1.001,20.731-7.112,21.733-18.102l2.65-29.069C87.527,464.806,165.571,512,256,512
			c97.281,0,183.012-55.522,225.57-138.854C486.594,363.309,482.692,351.262,472.855,346.238z"/>
                        </g>
                    </g >
                </svg>
            </button>
        </div>
    );
}
