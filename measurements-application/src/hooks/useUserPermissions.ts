import {useGetMeQuery, useGetUserPermissionsQuery} from "../features/users/api/userApi.ts";
import {useMemo} from "react";

type usePermissionsReturns = {
    canExport: boolean
    canFilter: boolean
}

export const useUserPermissions = () => {
    const {data: user} = useGetMeQuery()

    const {data: permissions = []} = useGetUserPermissionsQuery(user?.id ?? 0,  {
        skip: !user
    })

    const userPermissions: usePermissionsReturns = useMemo(() => {
        const userPermissions: usePermissionsReturns = {
            canExport: false,
            canFilter: false
        }

        permissions.forEach(permission => {
            switch (permission.name) {
                case "canExport":
                    userPermissions.canExport = true
                    break
                case "canFilter":
                    userPermissions.canFilter = true
                    break
            }
        })

        return userPermissions
    }, [permissions])

    return userPermissions
}