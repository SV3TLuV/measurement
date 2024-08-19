import {useEffect, useState} from "react";
import {RoutePaths} from "../enums/routePaths.ts";
import {useTypedSelector} from "../../../hooks/redux.ts";
import {useLocation, useNavigate} from "react-router-dom";
import {getUrlFromRoute} from "../../../utils/getUrlFromRoute.ts";

export const UnavailableChecker = () => {
    const [redirectFrom, setRedirectFrom] = useState<string>(RoutePaths.Empty)
    const {available} = useTypedSelector(state => state.app)
    const {pathname} = useLocation()
    const navigate = useNavigate()

    useEffect(() => {
        const unavailableUrl = getUrlFromRoute(RoutePaths.Unavailable)

        if (available && pathname === unavailableUrl) {
            navigate(redirectFrom)
        } else if (!available && pathname !== unavailableUrl) {
            setRedirectFrom(pathname)
            setTimeout(() => navigate(unavailableUrl), 0)
        }
    }, [available, pathname]);

    return null
}