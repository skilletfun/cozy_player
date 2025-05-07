export function secondsToHumanString(value, short = false) {
    function numberEnding (number) {
        return (number > 1) ? 's' : '';
    }

    const hourLabel = short ? ' h' : ' hour';
    const minLabel = short ? ' min' : ' minute';

    let result = "";

    const hours = Math.floor(value / 3600);
    if (hours) {
        result = hours + hourLabel + numberEnding(hours) + ' ';
    }
    const minutes = Math.floor((value % 3600) / 60);
    if (minutes) {
        return result + minutes + minLabel + numberEnding(minutes);
    }
}