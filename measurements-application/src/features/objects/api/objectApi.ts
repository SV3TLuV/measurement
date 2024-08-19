import {baseApi} from "../../../lib/api/baseApi.ts";
import {ApiTags} from "../../../lib/api/enums/apiTags.ts";
import {buildUrlArguments} from "../../../utils/buildUrlArguments.ts";
import HTTPMethod from "http-method-enum";
import {Facility} from "../types/facility.ts";
import {GetObjectListQuery} from "../types/getObjectListQuery.ts";
import {GetAndAddNewPostsQueryResult} from "../types/getAndAddNewPostsQueryResult.ts";

export const objectApi = baseApi.injectEndpoints({
    endpoints: builder => ({
        getObjects: builder.query<Facility[], GetObjectListQuery | void>({
            query: query => ({
                url: `/objects?${buildUrlArguments(query ?? {})}`,
                method: HTTPMethod.GET
            }),
            providesTags: result => [
                ...(result ?? []).map(({id}) => ({type: ApiTags.Object, id} as const)),
                { type: ApiTags.Object }
            ]
        }),
        getPost: builder.query<Facility, number>({
            query: id => ({
                url: `/objects/posts/${id}`,
                method: HTTPMethod.GET
            }),
            providesTags: () => [
                { type: ApiTags.Object }
            ]
        }),
        searchNewObjects: builder.mutation<GetAndAddNewPostsQueryResult, void | undefined>({
            query: () => ({
                url: `/objects/search-new`,
                method: HTTPMethod.GET
            }),
            invalidatesTags: [{ type: ApiTags.Object }]
        }),
        enablePostListened: builder.mutation<void, number>({
            query: id => ({
                url: `/objects/${id}/enable`,
                method: HTTPMethod.PUT
            }),
            invalidatesTags: [{ type: ApiTags.Object }]
        }),
        disablePostListened: builder.mutation<void, number>({
            query: id => ({
                url: `/objects/${id}/disable`,
                method: HTTPMethod.PUT
            }),
            invalidatesTags: [{ type: ApiTags.Object }]
        }),
    })
})

export const {
    useGetObjectsQuery,
    useGetPostQuery,
    useSearchNewObjectsMutation,
    useEnablePostListenedMutation,
    useDisablePostListenedMutation,
} = objectApi