import {ReactNode} from "react";
import {useTypedSelector} from "../../../hooks/redux.ts";
import {Roles} from "../../../shared/enums/roles.ts";
import {useNavigate} from "react-router-dom";
import {RoutePaths} from "../enums/routePaths.ts";
import {getUrlFromRoute} from "../../../utils/getUrlFromRoute.ts";
import {useGetMeQuery} from "../../../features/users/api/userApi.ts";

type AdminRouteProps = {
    children: ReactNode
}

export const AdminRoute = (props: AdminRouteProps) => {
    const { children } = props

    const { isAuthorized } = useTypedSelector(state => state.auth)
    const {data: user} = useGetMeQuery()
    const navigate = useNavigate()

    if (!isAuthorized || !user || user.role.id !== Roles.Admin) {
        navigate(getUrlFromRoute(RoutePaths.Empty))
        return null;
    }

    return children
}