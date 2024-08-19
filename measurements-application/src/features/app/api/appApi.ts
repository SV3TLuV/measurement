import {baseApi} from "../../../lib/api/baseApi.ts";
import {ApiTags} from "../../../lib/api/enums/apiTags.ts";
import HTTPMethod from "http-method-enum";
import {Configuration} from "../types/configuration.ts";
import {available, unavailable} from "../stores/appSlice.ts";

export const appApi = baseApi.injectEndpoints({
    endpoints: builder => ({
        getAppIsOnline: builder.query<boolean, void | undefined>({
            query: () => ({
                url: `online`,
                method: HTTPMethod.GET
            }),
            async onQueryStarted(_, {dispatch, queryFulfilled}) {
                try {
                    const { data } = await queryFulfilled

                    if (data) {
                        dispatch(available())
                    } else {
                        dispatch(unavailable())
                    }
                } catch { /* empty */ }
            },
        }),
        getConfiguration: builder.query<Configuration, void | undefined>({
            query: () => ({
                url: `configuration`,
                method: HTTPMethod.GET
            }),
            providesTags: () => [
                {type: ApiTags.Configuration }
            ]
        }),
        updateConfiguration: builder.mutation<void, Configuration>({
            query: command => ({
                url: `configuration`,
                method: HTTPMethod.PUT,
                body: command,
            }),
            invalidatesTags: [{type: ApiTags.Configuration}]
        }),
    })
})

export const {
    useGetAppIsOnlineQuery,
    useGetConfigurationQuery,
    useUpdateConfigurationMutation
} = appApi