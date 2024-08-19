import {baseApi} from "../../../lib/api/baseApi.ts";
import {ApiTags} from "../../../lib/api/enums/apiTags.ts";
import {buildUrlArguments} from "../../../utils/buildUrlArguments.ts";
import HTTPMethod from "http-method-enum";
import {GetUserListQuery} from "../types/getUserListQuery.ts";
import {Column} from "../../columns/types/column.ts";
import {ChangePasswordCommand} from "../types/changePasswordCommand.ts";
import {CreateUserCommand} from "../types/createUserCommand.ts";
import {UpdateUserCommand} from "../types/updateUserCommand.ts";
import {User} from "../types/user.ts";
import {PagedList} from "../../../lib/api/models/pagedList.ts";
import {Permission} from "../../permissions/types/permission.ts";
import {Facility} from "../../objects/types/facility.ts";

export const userApi = baseApi.injectEndpoints({
    endpoints: builder => ({
        getUsers: builder.query<PagedList<User>, GetUserListQuery | void>({
            query: query => ({
                url: `/users?${buildUrlArguments(query ?? {
                    page: 1,
                    pageSize: 50
                })}`,
                method: HTTPMethod.GET,
            }),
            providesTags: result => [
                ...(result?.items ?? []).map(({id}) => ({type: ApiTags.User, id} as const)),
            ]
        }),
        getMe: builder.query<User, undefined | void>({
            query: () => ({
                url: `/users/me`,
                method: HTTPMethod.GET,
            }),
            providesTags: () => [
                ({type: ApiTags.User, id: "me"})
            ]
        }),
        getUserColumns: builder.query<Column[], number>({
            query: id => `/users/${id}/columns`,
        }),
        getUserObjects: builder.query<Facility[], number>({
            query: id => `/users/${id}/objects`,
            providesTags: [{ type: ApiTags.Object }]
        }),
        getUserPermissions: builder.query<Permission[], number>({
            query: id => `/users/${id}/permissions`,
            providesTags: [{ type: ApiTags.Permission }]
        }),
        banUser: builder.mutation<void, number>({
            query: id => ({
                url: `/users/${id}/ban`,
                method: HTTPMethod.PUT
            }),
            invalidatesTags: [{type: ApiTags.User}]
        }),
        unbanUser: builder.mutation<void, number>({
            query: id => ({
                url: `/users/${id}/unban`,
                method: HTTPMethod.PUT
            }),
            invalidatesTags: [{type: ApiTags.User}]
        }),
        changePassword: builder.mutation<void, ChangePasswordCommand>(  {
            query: command => ({
                url: `/users/change-password`,
                method: HTTPMethod.PUT,
                body: command
            })
        }),
        createUser: builder.mutation<User, CreateUserCommand>({
            query: command => ({
                url: `/users`,
                method: HTTPMethod.POST,
                body: command,
            }),
            invalidatesTags: [{type: ApiTags.User}]
        }),
        editUser: builder.mutation<void, UpdateUserCommand>({
            query: command => ({
                url: `/users`,
                method: HTTPMethod.PUT,
                body: {
                    userId: command.id,
                    login: command.login,
                    roleId: command.roleId,
                    postIds: command.postIds,
                    columnIds: command.columnIds,
                    permissionIds: command.permissionIds,
                },
            }),
            invalidatesTags: [{type: ApiTags.User}]
        }),
        deleteUser: builder.mutation<void, number>({
            query: id => ({
                url: `/users/${id}`,
                method: HTTPMethod.DELETE,
            }),
            invalidatesTags: [{type: ApiTags.User}]
        }),
    }),
})

export const {
    useGetUsersQuery,
    useGetMeQuery,
    useGetUserColumnsQuery,
    useGetUserObjectsQuery,
    useGetUserPermissionsQuery,
    useBanUserMutation,
    useUnbanUserMutation,
    useChangePasswordMutation,
    useCreateUserMutation,
    useEditUserMutation,
    useDeleteUserMutation,
} = userApi;