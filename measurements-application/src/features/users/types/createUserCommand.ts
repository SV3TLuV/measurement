export type CreateUserCommand = {
    login: string
    password: string
    roleId: number
    permissionIds: number[]
    columnIds: number[]
    postIds: number[]
}