import {ReactNode} from "react";
import {useTypedSelector} from "../../../hooks/redux.ts";
import {Navigate} from "react-router-dom";
import {RoutePaths} from "../enums/routePaths.ts";
import {getUrlFromRoute} from "../../../utils/getUrlFromRoute.ts";

type WithoutAuthRouteProps = {
    children: ReactNode
}

export const WithoutAuthRoute = (props: WithoutAuthRouteProps) => {
    const { children } = props

    const { isAuthorized } = useTypedSelector(state => state.auth)

    if (isAuthorized) {
        return <Navigate to={getUrlFromRoute(RoutePaths.Empty)}/>
    }

    return children
}