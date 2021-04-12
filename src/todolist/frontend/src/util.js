import Long from "long";
import moment from "moment";

const twitter_epoch = Long.fromString("1288834974657");

export function snowflake2unixtime(snowflake) {
    let long = Long.fromString(snowflake).shiftRight(22).add(twitter_epoch);

    return long;
}

export function snowflake2moment(snowflake) {
    let create_time = moment.unix(snowflake2unixtime(snowflake).toNumber() / 1000);

    // return create_time.fromNow()
    return create_time.calendar()
}