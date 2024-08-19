import {baseApi} from "../../../lib/api/baseApi.ts";
import {ApiTags} from "../../../lib/api/enums/apiTags.ts";
import HTTPMethod from "http-method-enum";
import {Role} from "../types/role.ts";

export const roleApi = baseApi.injectEndpoints({
    endpoints: builder => ({
        getRoles: builder.query<Role[], void | undefined>({
            query: () => ({
                url: `/roles`,
                method: HTTPMethod.GET
            }),
            providesTags: roles => [
                ...(roles ?? []).map(({id}) => ({type: ApiTags.Role, id} as const)),
            ]
        })
    })
})

export const {
    useGetRolesQuery
} = roleApi