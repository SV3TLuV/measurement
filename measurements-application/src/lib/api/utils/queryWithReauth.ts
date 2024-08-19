import {Mutex} from "async-mutex";
import {BaseQueryFn, FetchArgs, FetchBaseQueryError} from "@reduxjs/toolkit/dist/query/react";
import {AppState} from "../../store";
import {logout} from "../../../features/auth/stores/authSlice.ts";
import {RefreshCommand} from "../../../features/auth/types/refreshCommand.ts";
import {authApi} from "../../../features/auth/api/authApi.ts";
import {fetchQueryBase} from "./fetchQueryBase.ts";

const mutex = new Mutex()

export const fetchQueryWithReauth: BaseQueryFn<
    string | FetchArgs,
    unknown,
    FetchBaseQueryError
> = async (args, api, extraOptions) => {
    await mutex.waitForUnlock()
    let result = await fetchQueryBase(args, api, extraOptions)

    if (result.error) {
        if (result.error.status === 401) {
            if (!mutex.isLocked()) {
                const release = await mutex.acquire()

                try {
                    const authState = (api.getState() as AppState).auth;

                    const refreshToken = authState.refreshToken

                    if (refreshToken) {
                        await api.dispatch(authApi.endpoints.refresh.initiate({
                            refreshToken: refreshToken
                        } as RefreshCommand))
                    }

                    result = await fetchQueryBase(args, api, extraOptions)

                    if (result.error && result.error.status === 401) {
                        api.dispatch(logout())
                    }
                }
                finally {
                    await release()
                }
            } else {
                await mutex.waitForUnlock()
                result = await fetchQueryBase(args, api, extraOptions)
            }
        }
    }

    return result
}