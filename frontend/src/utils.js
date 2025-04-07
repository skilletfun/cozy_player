export function secondsToHumanString(value) {
    function numberEnding (number) {
        return (number > 1) ? 's' : '';
    }

    var result = "";

    const hours = Math.floor(value / 3600);
    if (hours) {
        result = hours + ' hour' + numberEnding(hours) + ' ';
    }
    const minutes = Math.floor((value % 3600) / 60);
    if (minutes) {
        return result + minutes + ' minute' + numberEnding(minutes);
    }
}