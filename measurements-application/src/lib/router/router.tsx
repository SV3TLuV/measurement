import {createBrowserRouter, Navigate, Outlet, useLocation, useNavigate} from "react-router-dom";
import {ResultDataContext} from "../../context/result-data-context/ResultDataContext.tsx";
import {RoutePaths} from "./enums/routePaths.ts";
import {SignInPage} from "../../pages/SignInPage.tsx";
import {WithoutAuthRoute} from "./components/WithoutAuthRoute.tsx";
import {WithAuthRoute} from "./components/WithAuthRoute.tsx";
import {getUrlFromRoute} from "../../utils/getUrlFromRoute.ts";
import {UnavailablePage} from "../../pages/UnavailablePage.tsx";
import {AdminRoute} from "./components/AdminRoute.tsx";
import {UnavailableChecker} from "./components/UnavailableChecker.tsx";
import {AppHeader} from "../../layouts/header";
import {Layout} from "antd";
import {ConfigurationPage} from "../../pages/ConfigurationPage.tsx";
import {SchedulerInformationPage} from "../../pages/SchedulerInformationPage.tsx";
import {PostsPage} from "../../pages/PostsPage.tsx";
import {UsersPage} from "../../pages/UsersPage.tsx";
import {MeasurementsPage} from "../../pages/MeasurementsPage.tsx";
import {ResultPage} from "../../pages/result/page.tsx";
import {startSignalR} from "../signalr/client.ts";
import {useEffect, useState} from "react";
import {useAppDispatch} from "../../hooks/redux.ts";
import {useGetMeQuery} from "../../features/users/api/userApi.ts";

const Root = () => {
    const { pathname } = useLocation()
    const navigate = useNavigate()
    // const dispatch = useAppDispatch()
    // const { data: user } = useGetMeQuery()

    // eslint-disable-next-line
    // @ts-ignore
    // eslint-disable-next-line @typescript-eslint/no-unused-vars
    // const [connection, setConnection]
    //     = useState<signalR.HubConnection | null>(startSignalR(dispatch, user))

    useEffect(() => {
        if (pathname === "/") {
            navigate(getUrlFromRoute(RoutePaths.Measurements))
        }
    }, [pathname, navigate]);

    return (
        <ResultDataContext>
            <UnavailableChecker/>
            <Outlet/>
        </ResultDataContext>
    )
}

export const router = createBrowserRouter([
    {
        element: <Root/>,
        children: [
            {
                path: '/',
                element: (
                    <WithAuthRoute>
                        <Layout>
                            <AppHeader/>
                            <Outlet/>
                        </Layout>
                    </WithAuthRoute>
                ),
                children: [
                    {
                        path: RoutePaths.Measurements,
                        element: (
                            <MeasurementsPage/>
                        )
                    },
                    {
                        path: RoutePaths.Posts,
                        element: (
                            <AdminRoute>
                                <PostsPage/>
                            </AdminRoute>
                        )
                    },
                    {
                        path: RoutePaths.Users,
                        element: (
                            <AdminRoute>
                                <UsersPage/>
                            </AdminRoute>
                        )
                    },
                    {
                        path: RoutePaths.Configuration,
                        element: (
                            <AdminRoute>
                                <ConfigurationPage/>
                            </AdminRoute>
                        )
                    },
                    {
                        path: RoutePaths.SchedulerInformation,
                        element: (
                            <AdminRoute>
                                <SchedulerInformationPage/>
                            </AdminRoute>
                        )
                    },
                ]
            },
            {
                path: RoutePaths.SignIn,
                element: (
                    <WithoutAuthRoute>
                        <SignInPage/>
                    </WithoutAuthRoute>
                )
            },
            {
                path: RoutePaths.Result,
                element: <ResultPage/>
            },
            {
                path: RoutePaths.Unavailable,
                element: <UnavailablePage/>
            },
        ]
    },
    {
        path: '*',
        element: <Navigate to={getUrlFromRoute(RoutePaths.Measurements)}/>
    }
])