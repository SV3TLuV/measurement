import {baseApi} from "../../../lib/api/baseApi.ts";
import {ApiTags} from "../../../lib/api/enums/apiTags.ts";
import HTTPMethod from "http-method-enum";
import {Quality} from "../types/quality.ts";

export const qualityApi = baseApi.injectEndpoints({
    endpoints: builder => ({
        getQualities: builder.query<Quality[], void | undefined>({
            query: () => ({
                url: `/qualities`,
                method: HTTPMethod.GET
            }),
            providesTags: qualities => [
                ...(qualities ?? []).map(({id}) => ({type: ApiTags.Quality, id} as const)),
            ]
        })
    })
})

export const {
    useGetQualitiesQuery
} = qualityApi