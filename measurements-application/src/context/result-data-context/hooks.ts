import {useContext} from "react";
import {Context, ResultDataContextState} from "./ResultDataContext.tsx";

export const useResultDataContext = () => useContext<ResultDataContextState>(Context)
