import {createContext, ReactNode, useState} from "react";

export type ResultDataContextState = {
    extra: ReactNode | null
    setExtra: (extra: ReactNode | null) => void
}

export const Context =
    createContext<ResultDataContextState>({} as ResultDataContextState)

type ResultDataContextProps = {
    children: ReactNode
}

export const ResultDataContext = ({children}: ResultDataContextProps) => {
    const [extra, setExtra] = useState<ReactNode | null>()

    // useEffect(() => {
    //     if (location.state === null && location.pathname !== `/${RoutePaths.Result}`) {
    //         setExtra(null)
    //     }
    // }, [extra, location.pathname, location.state])

    return (
        <Context.Provider value={{ extra, setExtra }}>
            {children}
        </Context.Provider>
    )
}