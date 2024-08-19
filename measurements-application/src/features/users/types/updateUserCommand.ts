export type UpdateUserCommand = {
    id: number
    login: string
    roleId: number
    permissionIds: number[]
    columnIds: number[]
    postIds: number[]
}
