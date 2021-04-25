import Long from "long";
import moment from "moment";

const twitter_epoch = Long.fromString("1288834974657");

export function snowflake2unixtime(snowflake) {
    let long = Long.fromString(snowflake).shiftRight(22).add(twitter_epoch);

    return long;
}

export function snowflake2moment(snowflake) {
    let create_time = moment.unix(snowflake2unixtime(snowflake).toNumber() / 1000);

    return create_time
}

export function isBefore1day(snowflake) {
    let create_time = snowflake2moment(snowflake);
    let yesterday = moment(new Date()).add(-1, 'days');

    return create_time.isBefore(yesterday);
}

export function addRemindTodo(todo) {
    let remindTodos = localStorage.remindTodos === undefined ? { todos: [] } : JSON.parse(localStorage.remindTodos);

    remindTodos.todos.push(todo);
    localStorage.remindTodos = JSON.stringify(remindTodos);
}

export function removeRemindTodo(todoId) {
    let remindTodos = localStorage.remindTodos === undefined ? { todos: [] } : JSON.parse(localStorage.remindTodos);

    remindTodos.todos = remindTodos.todos.filter((t) => t.Id !== todoId);
    localStorage.remindTodos = JSON.stringify(remindTodos);
}

function showRemindTodo(todo) {
    if (Notification.permission !== 'granted') {
        Notification.requestPermission(() => { });
        return;
    }

    if (moment().isBefore(todo.expire))
        return;

    const config = {
        body: todo.Name,
        icon: 'android-chrome-192x192.png',
        tag: todo.Id,
    }

    const notification = new Notification('todo提醒', config);
    notification.onshow = (e) => {
        e.preventDefault();
        removeRemindTodo(todo.Id);
        if (todo.retry < 3) {
            todo.retry = todo.retry + 1;
            todo.expire = moment(new Date()).add(5, 'minutes');
            addRemindTodo(todo);
        }
    };
    notification.onclick = (e) => {
        e.preventDefault();
        removeRemindTodo(todo.Id);
        // window.open('https://todo.firego.cn', '');
    };
    notification.onclose = (e) => {
        e.preventDefault();
        removeRemindTodo(todo.Id);
    };
}

export function checkRemindTodo() {
    if (localStorage.remindTodos === undefined) return [];
    let remindTodos = JSON.parse(localStorage.remindTodos);

    if (remindTodos.todos.length === 0) return [];

    for (let todo of remindTodos.todos) {
        todo.expire = moment(todo.expire);
        showRemindTodo(todo);
    }

    return remindTodos.todos;
}

export function getTodayFormat() {
    let date = new Date();
    let year = date.getUTCFullYear();
    let month = date.getUTCMonth() + 1;
    let day = date.getUTCDate();

    if (year < 10) year = '0' + year;
    if (month < 10) month = '0' + month;
    if (day < 10) day = '0' + day;

    return `${year}-${month}-${day}`
}