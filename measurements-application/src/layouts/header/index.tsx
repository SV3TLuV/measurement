import {Avatar, Dropdown, Layout, MenuProps} from "antd";
import classNames from "classnames";
import {RoutePaths} from "../../lib/router/enums/routePaths.ts";
import {getUrlFromRoute} from "../../utils/getUrlFromRoute.ts";
import {useNavigate} from "react-router-dom";
import {UserOutlined} from "@ant-design/icons";
import './style.scss';
import {useShowResult} from "../../pages/result/hooks/useShowResult.tsx";
import {Roles} from "../../shared/enums/roles.ts";
import {useLogoutMutation} from "../../features/auth/api/authApi.ts";
import {useGetMeQuery} from "../../features/users/api/userApi.ts";

const {Header} = Layout;

export const AppHeader = () => {
    const { data: user } = useGetMeQuery()
    const navigate = useNavigate()
    const showResult = useShowResult()

    const [logout] = useLogoutMutation()

    const handleLogout = async () => {
        const response = await logout()
        await showResult({
            response: response,
            showError: false,
            getSuccessMessage: () => ({
                type: 'message',
                title: 'Вы успешно вышли'
            })
        })
    }

    const menuItems: MenuProps['items'] = [
        {
            key: 'data',
            label: 'Измерения',
            onClick: () => navigate(getUrlFromRoute(RoutePaths.Measurements)),
        },
        ...(user?.role.id === Roles.Admin
            ? [
                {
                    key: RoutePaths.Users,
                    label: 'Пользователи',
                    onClick: () => navigate(getUrlFromRoute(RoutePaths.Users)),
                },
                {
                    key: RoutePaths.Posts,
                    label: 'Станции',
                    onClick: () => navigate(getUrlFromRoute(RoutePaths.Posts)),
                },
                {
                    key: RoutePaths.Configuration,
                    label: 'Конфигурация',
                    onClick: () => navigate(getUrlFromRoute(RoutePaths.Configuration)),
                },
                {
                    key: RoutePaths.SchedulerInformation,
                    label: 'Статистика',
                    onClick: () => navigate(getUrlFromRoute(RoutePaths.SchedulerInformation)),
                },
            ]
            : []),
        {
            key: 'logout',
            label: 'Выйти',
            onClick: handleLogout,
        },
    ];

    return (
        <Header className={classNames('app-header')}>
            {user && (
                <Dropdown
                    menu={{ items: menuItems }}
                    trigger={['hover']}
                >
                    <div className={classNames('app-header-user-info')}>
                        <Avatar icon={<UserOutlined />} />
                        <span className={classNames('app-header-user-info-login')}>
                        {user.login}
                    </span>
                    </div>
                </Dropdown>
            )}
        </Header>
    )
}