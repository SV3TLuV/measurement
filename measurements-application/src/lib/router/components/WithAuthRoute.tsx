import {ReactNode} from "react";
import {useTypedSelector} from "../../../hooks/redux.ts";
import {useNavigate} from "react-router-dom";
import {RoutePaths} from "../enums/routePaths.ts";
import {getUrlFromRoute} from "../../../utils/getUrlFromRoute.ts";

type WithAuthRouteProps = {
    children: ReactNode
}

export const WithAuthRoute = (props: WithAuthRouteProps) => {
    const { children } = props

    const { isAuthorized } = useTypedSelector(state => state.auth)
    const navigate = useNavigate()

    if (!isAuthorized) {
        setTimeout(() => {
            navigate(getUrlFromRoute(RoutePaths.SignIn))
        }, 0)
        return null
    }

    return children
}