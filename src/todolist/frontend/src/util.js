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

export function removeRemindTodo(todo) {
    let remindTodos = localStorage.remindTodos === undefined ? { todos: [] } : JSON.parse(localStorage.remindTodos);

    remindTodos.todos = remindTodos.todos.filter((t) => t.Id !== todo.Id);
    localStorage.remindTodos = JSON.stringify(remindTodos);
}

function showRemindTodo(todo) {
    if (Notification.permission !== 'granted') {
        Notification.requestPermission(() => { });
        return false;
    }

    if (moment().isBefore(todo.expire))
        return false;

    const config = {
        body: todo.Name,
        icon: 'android-chrome-192x192.png',
        tag: todo.Id,
    }

    new Notification('todo提醒', config);

    return true;
}

export function checkRemindTodo() {
    if (localStorage.remindTodos === undefined) return;
    let remindTodos = JSON.parse(localStorage.remindTodos);

    if (remindTodos.todos.length === 0) return;

    for (let todo of remindTodos.todos) {
        todo.expire = moment(todo.expire);
        if (showRemindTodo(todo)){
            removeRemindTodo(todo);
        }
    }
}