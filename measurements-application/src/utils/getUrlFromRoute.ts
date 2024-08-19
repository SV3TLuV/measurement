import {RoutePaths} from "../lib/router/enums/routePaths.ts";

export const getUrlFromRoute = (route: RoutePaths) => {
    return `/${route}`
}