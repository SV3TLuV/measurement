import {baseApi} from "../../../lib/api/baseApi.ts";
import {ApiTags} from "../../../lib/api/enums/apiTags.ts";
import HTTPMethod from "http-method-enum";
import {Permission} from "../types/permission.ts";

export const permissionApi = baseApi.injectEndpoints({
    endpoints: builder => ({
        getPermissions: builder.query<Permission[], undefined | void>({
            query: () => ({
                url: `/permissions`,
                method: HTTPMethod.GET,
            }),
            providesTags: permissions => [
                ...(permissions ?? []).map(({id}) => ({type: ApiTags.Permission, id} as const)),
            ]
        }),
    }),
})

export const {
    useGetPermissionsQuery,
} = permissionApi