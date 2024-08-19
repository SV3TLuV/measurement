import {baseApi} from "../../../lib/api/baseApi.ts";
import {ApiTags} from "../../../lib/api/enums/apiTags.ts";
import HTTPMethod from "http-method-enum";
import {Column} from "../types/column.ts";

export const columnApi = baseApi.injectEndpoints({
    endpoints: builder => ({
        getColumns: builder.query<Column[], void | undefined>({
            query: () => ({
                url: `/columns`,
                method: HTTPMethod.GET
            }),
            providesTags: columns => [
                ...(columns ?? []).map(({id}) => ({type: ApiTags.Column, id} as const)),
            ]
        })
    })
})

export const {
    useGetColumnsQuery,
} = columnApi