import { useParams } from "react-router-dom";
import { useCallback, useEffect, useState } from "react";
import Calendar from 'react-github-contribution-calendar';
import axios from 'axios';
import { getTodayFormat, snowflake2moment } from "../util";

export function DailyRecord() {
    let { todo_id } = useParams();

    const [data, setData] = useState({ todo: "", records: [] });

    const getDailyRecord = useCallback(async () => {
        const res = await axios.get(`/todo/daily/${todo_id}`);
        setData(res.data)
    }, [todo_id]);

    useEffect(() => {
        getDailyRecord();
    }, [getDailyRecord]);

    const values = {}

    data.records.forEach(element => {
        values[element] = 1
    });

    const until = getTodayFormat();
    var panelColors = [
        '#FFFFFF',
        '#047857',
    ];

    return (
        <div className='min-h-screen w-full flex flex-col items-center justify-items-center'>
            <div className='flex flex-col items-center justify-items-center w-full md:w-3/4 xl:w-3/5 dark:text-white '>
                <h1 className='block text-3xl md:text-5xl lg:text-7xl w-full my-6 font-mono font-bold text-center select-none'>
                    daily todo records
                </h1>
                <h2 className='block mb-12 text-2xl md:text-4xl lg:text-6xl w-full font-mono font-bold text-center select-none'>
                    "{data.todo}"
                </h2>
                <span className='self-end relative select-none'>{snowflake2moment(todo_id).calendar()}</span>
                <div className='w-full select-none bg-gray-700 dark:bg-black p-2 md:p-5 rounded-md'>
                    <Calendar values={values} until={until} panelColors={panelColors} />
                </div>
            </div>
        </div>
    );
}
