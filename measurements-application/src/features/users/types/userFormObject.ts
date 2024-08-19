export type UserFormObject = {
    id: number
    login: string
    password: string
    confirm: string
    roleId: number
    columnIds: number[]
    postIds: number[]
    permissionIds: number[]
}