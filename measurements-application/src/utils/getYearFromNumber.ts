export const getYearFromNumber = (day: number): string => {
    const lastDigit = day % 10;
    const lastTwoDigits = day % 100;

    if (lastTwoDigits >= 11 && lastTwoDigits <= 14) {
        return 'лет';
    } else if (lastDigit === 1) {
        return 'год'
    } else if (lastDigit >= 2 && lastDigit <= 4) {
        return 'года';
    }

    return 'лет';
}