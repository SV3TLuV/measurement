import {createApi} from "@reduxjs/toolkit/dist/query/react";
import {fetchQueryWithReauth} from "./utils/queryWithReauth.ts";
import {ApiTags} from "./enums/apiTags.ts";

export const baseApi = createApi({
    reducerPath: 'BaseApi',
    baseQuery: fetchQueryWithReauth,
    tagTypes: Object.values(ApiTags),
    keepUnusedDataFor: 0,
    refetchOnReconnect: true,
    refetchOnFocus: true,
    endpoints: () => ({}),
})