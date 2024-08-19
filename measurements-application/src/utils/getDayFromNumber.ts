export const getDayFromNumber = (day: number): string => {
    const lastDigit = day % 10;
    const lastTwoDigits = day % 100;

    if (lastTwoDigits >= 11 && lastTwoDigits <= 14) {
        return 'дней';
    } else if (lastDigit === 1) {
        return 'день'
    } else if (lastDigit >= 2 && lastDigit <= 4) {
        return 'дня';
    }

    return 'дней';
}