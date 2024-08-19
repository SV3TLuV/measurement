import classNames from "classnames"
import {Result} from "antd";
import {Location, Navigate, useLocation} from "react-router-dom";
import './style.scss';
import {ResultObject} from "./models.ts";
import {useResultDataContext} from "../../context/result-data-context/hooks.ts";
import {getUrlFromRoute} from "../../utils/getUrlFromRoute.ts";
import {RoutePaths} from "../../lib/router/enums/routePaths.ts";

export const ResultPage = () => {
    const location: Location<ResultObject> = useLocation()
    const context = useResultDataContext()

    return (
        <div className={classNames('result-page')}>
            {location.state === null ? (
                <Navigate to={getUrlFromRoute(RoutePaths.Measurements)}/>
            ) : (
                <Result
                    status={location.state.status}
                    title={location.state.title}
                    subTitle={location.state.subTitle}
                    extra={context.extra}
                />
            )}
        </div>
    )
}