import {BaseQueryFn, FetchArgs, FetchBaseQueryError} from "@reduxjs/toolkit/dist/query/react";
import {baseQuery} from "./baseQuery.ts";
import {available, unavailable} from "../../../features/app/stores/appSlice.ts";
import {AppState} from "../../store";

export const fetchQueryBase: BaseQueryFn<
    string | FetchArgs,
    unknown,
    FetchBaseQueryError
> = async (args, api, extraOptions) => {
    const { available: online } = (api.getState() as AppState).app

    const result = await baseQuery(args, api, extraOptions)

    if (result.error) {
        if (result.error.status === 'FETCH_ERROR') {
            api.dispatch(unavailable())
        }
    } else if (!online) {
        api.dispatch(available())
    }

    return result
}