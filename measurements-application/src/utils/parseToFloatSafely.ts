export const parseToFloatSafely = (value: string): number | null => {
    const result = parseFloat(value);
    if (isNaN(result)) {
        return null;
    }
    return result;
};
