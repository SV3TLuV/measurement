export const getMonthFromNumber = (day: number): string => {
    const lastDigit = day % 10;
    const lastTwoDigits = day % 100;

    if (lastTwoDigits >= 11 && lastTwoDigits <= 14) {
        return 'месяцев';
    } else if (lastDigit === 1) {
        return 'месяц'
    } else if (lastDigit >= 2 && lastDigit <= 4) {
        return 'месяца';
    }

    return 'месяцев';
}