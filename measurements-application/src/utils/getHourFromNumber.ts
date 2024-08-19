export const getHourFromNumber = (hour: number) => {
    const lastDigit = hour % 10;
    const lastTwoDigits = hour % 100;

    if (lastTwoDigits >= 11 && lastTwoDigits <= 14) {
        return 'часов';
    } else if (lastDigit === 1) {
        return 'час';
    } else if (lastDigit >= 2 && lastDigit <= 4) {
        return 'часа';
    }

    return 'часов';
}