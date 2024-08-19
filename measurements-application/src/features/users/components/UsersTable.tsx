
import {Button, Table, TableColumnsType, Pagination, Input, TableProps, Popconfirm, Switch} from "antd";
import {AiOutlineDelete, AiOutlineEdit} from "react-icons/ai";
import classNames from "classnames";
import {ChangePasswordForm} from "./ChangePasswordForm.tsx";
import {useDialog} from "../../../hooks/useDialog.ts";
import {useState} from "react";
import {PlusOutlined} from "@ant-design/icons";
import {
    useDeleteUserMutation,
    useBanUserMutation,
    useUnbanUserMutation, useGetMeQuery,
    useGetUsersQuery
} from "../api/userApi.ts";
import {useShowResult} from "../../../pages/result/hooks/useShowResult.tsx";
import {useGetRolesQuery} from "../../roles/api/roleApi.ts";
import {User} from "../types/user.ts";
import {usePagination} from "../../../hooks/usePagination.ts";
import './UsersTable.scss';
import {useMediaQuery} from "react-responsive";
import {CreateUserForm} from "./user-form/CreateUserForm.tsx";
import {EditUserForm} from "./user-form/EditUserForm.tsx";

const { Search } = Input;

export const UsersTable = () => {
    const {data: user} = useGetMeQuery()

    const [search, setSearch] = useState<string>('')
    const [roleIds, setRoleIds] = useState<number[] | null>(null)
    const {page, pageSize, setPage, setPageSize} = usePagination()

    const {data} = useGetUsersQuery({
        search: search,
        pageSize: pageSize,
        page: page,
        roleIds: roleIds
    })

    const {data: roles = []} = useGetRolesQuery()

    const isMobile = useMediaQuery({ maxWidth: 767 });

    const showResult = useShowResult()

    const [deleteUser] = useDeleteUserMutation()
    const [unbanUser] = useUnbanUserMutation()
    const [banUser] = useBanUserMutation()

    const [currentUser, setCurrentUser] = useState<User | null>(null)

    const changePasswordDialog = useDialog()
    const createUserDialog = useDialog()
    const editUserDialog = useDialog()

    const handleDelete = async (user: User) => {
        const response = await deleteUser(user.id)
        await showResult({
            response: response,
            getSuccessMessage: () => ({
                type: 'message',
                title: `Пользователь ${user.login} удален`
            })
        })
    }

    const handleDisableOrEnable = async (user: User) => {
        if (user.isBlocked) {
            const response = await unbanUser(user.id)
            await showResult({
                response: response,
                getSuccessMessage: () => ({
                    type: 'message',
                    title: `Пользователь ${user.login} разблокирован`
                })
            })
        } else {
            const response = await banUser(user.id)
            await showResult({
                response: response,
                getSuccessMessage: () => ({
                    type: 'message',
                    title: `Пользователь ${user.login} заблокирован`
                })
            })
        }
    }

    const handleChangePassword = (user: User) => {
        setCurrentUser(user)
        changePasswordDialog.show()
    }

    const handleEdit = (user: User) => {
        setCurrentUser(user)
        editUserDialog.show()
    }

    const handleSearch = (text: string) => setSearch(text)

    const handleChangeFilter: TableProps['onChange'] =
        (_, filters)=> {
            setRoleIds(filters['role'] as number[])
        }

    const columns: TableColumnsType<User> = [
        {
            title: 'Логин',
            dataIndex: 'login',
            key: 'login',
            fixed: 'left',
            showSorterTooltip: {
                target: 'full-header'
            },
            sorter: (a, b) => a.login.localeCompare(b.login),
            sortDirections: ['ascend', 'descend']
        },
        {
            title: 'Роль',
            dataIndex: 'role',
            key: 'role',
            fixed: 'left',
            filters: roles.map(role => ({
                text: role.title,
                value: role.id
            })),
            showSorterTooltip: {
                target: 'full-header'
            },
            sortDirections: ['ascend', 'descend'],
            sorter: (a, b) => a.role.title.localeCompare(b.role.title),
            render: (_, record) => (
                <div>{record.role.title}</div>
            )
        },
        {
            title: 'Заблокирован',
            key: 'is_disabled',
            fixed: 'left',
            width: 80,
            render: (_, record) => (
                <div className={classNames('button-center')}>
                    <Switch
                        checked={record.isBlocked}
                        disabled={user !== null && user?.id === record.id}
                        onClick={async () => await handleDisableOrEnable(record)}
                    />
                </div>
            ),
        },
        {
            title: 'Изменить пароль',
            key: 'operation',
            fixed: 'right',
            width: 160,
            render: (_, record) => (
                <div className={classNames('button-center')}>
                    <Button onClick={() => handleChangePassword(record)}>
                        Изменить пароль
                    </Button>
                </div>
            ),
        },
        {
            title: 'Редактировать',
            key: 'operation',
            fixed: 'right',
            width: 80,
            render: (_, record) => (
                <div className={classNames('button-center')}>
                    <Button
                        onClick={() => handleEdit(record)}
                        icon={<AiOutlineEdit/>}
                    />
                </div>
            ),
        },
        {
            title: 'Удалить',
            key: 'operation',
            fixed: 'right',
            width: 80,
            render: (_, record) => (
                <div className={classNames('button-center')}>
                    <Popconfirm
                        title='Удаление пользователя'
                        description={`Вы уверены что хотите удалить пользователя: ${record.login}?`}
                        okText='Да'
                        cancelText='Отмена'
                        onConfirm={async () => await handleDelete(record)}
                    >
                        <Button
                            danger
                            disabled={user !== null && user?.id === record.id}
                            icon={<AiOutlineDelete/>}
                        />
                    </Popconfirm>
                </div>
            ),
        },
    ]

    return (
        <div className={classNames('users-table')}>
            <div className={classNames('table-header')}>
                <Search
                    title='Поиск по логину и роли'
                    size='large'
                    allowClear onSearch={handleSearch}
                />
                <Button
                    title='Добавить пользователя'
                    onClick={() => createUserDialog.show()}
                    type='primary'
                    shape='circle'
                    icon={<PlusOutlined/>}
                />
            </div>
            <div className={classNames('users-table-container', 'ant-table-wrapper')}>
                <Table
                    bordered
                    columns={columns}
                    pagination={false}
                    rowKey='id'
                    dataSource={data?.items ?? []}
                    onChange={handleChangeFilter}
                />
                {currentUser && (
                    <>
                        <ChangePasswordForm user={currentUser} {...changePasswordDialog} />
                        <EditUserForm user={currentUser} {...editUserDialog} />
                    </>
                )}
                <CreateUserForm {...createUserDialog} />
            </div>
            <div className={classNames('pagination-wrapper')}>
                <Pagination
                    showSizeChanger
                    showTotal={(total, range) => `Показано ${range[1]} из ${total}`}
                    defaultPageSize={pageSize}
                    current={page}
                    pageSize={pageSize}
                    total={data?.total}
                    simple={isMobile}
                    onChange={(page, size) => {
                        setPage(page)
                        setPageSize(size)
                    }}
                />
            </div>
        </div>
    )
}