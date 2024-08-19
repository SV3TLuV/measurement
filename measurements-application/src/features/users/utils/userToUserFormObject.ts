import {UserFormObject} from "../types/userFormObject.ts";
import {User} from "../types/user.ts";

export const userToUserFormObject = (user: User | null | undefined): UserFormObject => {
    if (user === null || user === undefined) {
        return {} as UserFormObject
    }

    console.log(user)

    return {
        id: user?.id,
        password: '',
        confirm: '',
        login: user?.login,
        roleId: user?.role?.id,
        columnIds: user?.columns ?? [],
        postIds: user?.posts ?? [],
        permissionIds: user?.permissions ?? [],
    } as UserFormObject
}