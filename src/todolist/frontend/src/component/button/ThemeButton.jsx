import { useCallback, useRef, useEffect } from 'react';

export function ThemeButton() {
    const remindBtnRef = useRef(null);

    const toggleTheme = useCallback((e) => {
        let nowTheme = localStorage.theme === 'dark' ? 'light' : 'dark';
        localStorage.theme = nowTheme;

        if (nowTheme === 'dark') {
            document.querySelector('html').classList.add('dark')
        } else {
            document.querySelector('html').classList.remove('dark')
        }
    }, []);

    const openNotification = useCallback(async () => {
        if (Notification.permission !== 'granted') {
            remindBtnRef.current.classList.remove('hidden');
            let permission = await Notification.requestPermission();

            if (permission === 'granted')
                remindBtnRef.current.classList.add('hidden');
        }
    }, []);

    useEffect(() => {
        openNotification();
    }, [openNotification]);

    return (
        <div className='absolute text-black dark:text-white top-0 w-full h-0 cursor-pointer'>
            <svg className='absolute top-3 left-3 w-6 h-6 fill-current' onClick={toggleTheme}>
                <path d="M12,3c-4.97,0-9,4.03-9,9s4.03,9,9,9s9-4.03,9-9c0-0.46-0.04-0.92-0.1-1.36c-0.98,1.37-2.58,2.26-4.4,2.26 c-2.98,0-5.4-2.42-5.4-5.4c0-1.81,0.89-3.42,2.26-4.4C12.92,3.04,12.46,3,12,3L12,3z" />
            </svg>
            <svg ref={remindBtnRef} className='absolute top-3 left-10 w-6 h-6 fill-current hidden' onClick={openNotification}>
                <path d="M12 22c1.1 0 2-.9 2-2h-4c0 1.1.89 2 2 2zm6-6v-5c0-3.07-1.64-5.64-4.5-6.32V4c0-.83-.67-1.5-1.5-1.5s-1.5.67-1.5 1.5v.68C7.63 5.36 6 7.92 6 11v5l-2 2v1h16v-1l-2-2z" />
            </svg>
        </div>
    );
}
