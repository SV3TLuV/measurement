import HTTPMethod from "http-method-enum";
import {login, logout} from "../stores/authSlice.ts";
import {createApi, FetchBaseQueryError} from "@reduxjs/toolkit/dist/query/react";
import {AppState} from "../../../lib/store";
import {RefreshCommand} from "../types/refreshCommand.ts";
import {AuthResponse} from "../types/authResponse.ts";
import {LoginCommand} from "../types/loginCommand.ts";
import {fetchQueryBase} from "../../../lib/api/utils/fetchQueryBase.ts";
import {baseApi} from "../../../lib/api/baseApi.ts";

export const authApi = createApi({
    reducerPath: 'AuthApi',
    baseQuery: fetchQueryBase,
    refetchOnReconnect: true,
    refetchOnFocus: true,
    keepUnusedDataFor: 0,
    endpoints: builder => ({
        checkSession: builder.query<void, void>({
            query: () => ({
                url: `/auth/session-alive`,
                method: HTTPMethod.GET
            }),
        }),
        login: builder.mutation<AuthResponse, LoginCommand>({
            query: command => ({
                url: `/auth/login`,
                method: HTTPMethod.POST,
                body: command,
            }),
            async onQueryStarted(_, {dispatch, queryFulfilled}) {
                try {
                    const { data } = await queryFulfilled
                    dispatch(login(data))
                } catch { /* empty */ }
            },
        }),
        logout: builder.mutation<void, void>({
            queryFn: async (_, api, extraOptions) => {
                try {
                    const authState = (api.getState() as AppState).auth;

                    const response = await fetchQueryBase({
                        url: `/auth/logout`,
                        method: HTTPMethod.POST,
                        body: {
                            refreshToken: authState.refreshToken
                        },
                    }, api, extraOptions)

                    return { error: response.error as FetchBaseQueryError }
                } finally {
                    api.dispatch(baseApi.util.resetApiState())
                    api.dispatch(logout())
                }
            },
        }),
        refresh: builder.mutation<AuthResponse, RefreshCommand>({
            queryFn: async (command, api, extraOptions) => {
                const response = await fetchQueryBase({
                    url: `/auth/refresh`,
                    method: HTTPMethod.PUT,
                    body: command,
                }, api, extraOptions)

                if (response.data) {
                    const data = response.data as AuthResponse

                    api.dispatch(login(data))

                    return { data: data }
                }

                return { error: response.error as FetchBaseQueryError }
            },
        }),
    }),
})

export const {
    useCheckSessionQuery,
    useLoginMutation,
    useLogoutMutation,
} = authApi;
