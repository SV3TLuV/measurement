import React, {useState} from "react";

type useOptionsProps<T extends React.Key> = {
    defaultValues: T[]
    getTitle: (value: T) => string
}

type OptionProps<T extends React.Key> = {
    key: number,
    value: T,
    title: string
}

type useOptionsReturn<T extends React.Key> = {
    options: OptionProps<T>[]
    reset: () => void
    onSearch: (value: T) => void
}

const compareOptions = <T extends React.Key>(a: T, b: T): number => {
    if (typeof a === 'number' && typeof b === 'number') {
        return a - b;
    } else if (typeof a === 'string' && typeof b === 'string') {
        return a.localeCompare(b);
    }

    return 0;
}

const getOptions = <T extends React.Key>(values: T[], getTitle: (value: T) => string): OptionProps<T>[] => {
    return values.map(value => ({
        key: value,
        value: value,
        title: getTitle(value)
    } as OptionProps<T>))
}

export const useOptions = <T extends React.Key>(
    {
        defaultValues,
        getTitle
    }: useOptionsProps<T>): useOptionsReturn<T> => {
    const [options, setOptions] =
        useState<OptionProps<T>[]>(getOptions(defaultValues, getTitle))

    const reset = () => setOptions(getOptions(defaultValues, getTitle))

    const onSearch = (value: T) => setOptions(getOptions(
        [...new Set(defaultValues.concat(value))].sort(compareOptions),
        getTitle))

    return { options, reset, onSearch }
}