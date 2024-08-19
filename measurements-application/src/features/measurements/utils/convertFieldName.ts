export const convertFieldName = (fieldName: string): string  => {
    return fieldName.replace(/_([a-z])/g, (_, letter) => letter.toUpperCase());
}