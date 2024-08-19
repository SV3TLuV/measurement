export const getMinuteFromNumber = (minute: number): string => {
    const lastDigit = minute % 10;
    const lastTwoDigits = minute % 100;

    if (lastTwoDigits >= 11 && lastTwoDigits <= 14) {
        return 'минут';
    } else if (lastDigit === 1) {
        return 'минута';
    } else if (lastDigit >= 2 && lastDigit <= 4) {
        return 'минуты';
    }

    return 'минут';
}