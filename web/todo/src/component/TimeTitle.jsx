import { useState, useEffect } from 'react';
import React from 'react'
import moment from 'moment';

export function TimeTitle() {
    const [now, setNow] = useState(moment().format("MMMM Do YYYY, H:mm:ss"));

    useEffect(() => {
        let timer = setInterval(() => {
            setNow(moment().format("MMMM Do YYYY, H:mm:ss"));
        }, 1000);

        return () => { clearInterval(timer) };
    }, []);

    return (
        <h5 className="pan_y text-black dark:text-white font-mono text-xs md:text-base select-none">{now}</h5>
    )
}
