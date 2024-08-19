import {Role} from "../../roles/types/role.ts";

export type User = {
    id: number
    login: string
    role: Role
    isBlocked: boolean
    permissions: number[]
    columns: number[]
    posts: number[]
}

