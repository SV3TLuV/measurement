import {Button, Input, Switch, Table, TableColumnsType, TableProps} from "antd";
import {useState} from "react";
import classNames from "classnames";
import {SyncOutlined} from "@ant-design/icons";
import {useShowResult} from "../../../pages/result/hooks/useShowResult.tsx";
import {FindNewPostResultEmpty, FindNewPostResultWithData} from "./FindNewPostsResult.tsx";
import {
    useDisablePostListenedMutation,
    useEnablePostListenedMutation, useGetObjectsQuery, useSearchNewObjectsMutation,
} from "../api/objectApi.ts";
import {Facility} from "../types/facility.ts";
import './PostsTable.scss';
import {getObjectTitle} from "../../measurements/utils/getObjectTitle.ts";
import {ObjectTypeKeys} from "../types/objectType.ts";

const { Search } = Input;

export const PostsTable = () => {
    const [search, setSearch] = useState<string | undefined>(undefined)
    const [cityIds, setCityIds] = useState<number[] | undefined>(undefined)
    const [laboratoryIds, setLaboratoryIds] = useState<number[] | undefined>(undefined)

    const {data: posts = [], isFetching} = useGetObjectsQuery({
        search: search,
        typeId: ObjectTypeKeys.Post,
        parentIds: cityIds || laboratoryIds,
    })
    const {data: cities = []} = useGetObjectsQuery({
        typeId: ObjectTypeKeys.City,
        parentIds: laboratoryIds
    })
    const {data: laboratories = []} = useGetObjectsQuery({
        typeId: ObjectTypeKeys.Laboratory,
    })

    const [enablePost] = useEnablePostListenedMutation()
    const [disablePost] = useDisablePostListenedMutation()
    const [searchNewObjects, { isLoading }] = useSearchNewObjectsMutation()

    const showResult = useShowResult()

    const handleDisableOrEnableListened = async (post: Facility) => {
        if (post.isListened) {
            const response = await disablePost(post.id)
            await showResult({
                response: response,
                getSuccessMessage: () => ({
                    type: 'message',
                    title: `Опрос станции ${getObjectTitle(post, true)} отключен`
                })
            })
        } else {
            const response = await enablePost(post.id)
            await showResult({
                response: response,
                getSuccessMessage: () => ({
                    type: 'message',
                    title: `Опрос станции ${getObjectTitle(post, true)} включен`
                })
            })
        }
    }

    const handleSynchronize = async () => {
        const response = await searchNewObjects()
        await showResult({
            response: response,
            getSuccessMessage: (result) => ({
                type: 'result',
                status: (
                    result.posts.length === 0 ? '404' : 'success'
                ),
                title: (
                    result.posts.length === 0
                        ? 'Ничего не найдено. Данные актуальны.'
                        : `Найдено новых: ${result.newPostsCount}. Обновлено: ${result.updatedPostsCount}`
                ),
                extra: (
                    result.posts.length === 0
                        ? (
                            <FindNewPostResultEmpty/>
                        ) : (
                            <FindNewPostResultWithData result={result} />
                        )
                )
            })
        })
    }

    const handleSearch = (text: string) => setSearch(text)

    const handleChangeFilter: TableProps['onChange'] =
        (_, filters) => {
            setLaboratoryIds(filters['laboratory'] as number[])
            setCityIds(filters['city'] as number[])
        }

    const columns: TableColumnsType<Facility> = [
        {
            title: 'Лаборатория',
            dataIndex: 'laboratory',
            key: 'laboratory',
            fixed: 'left',
            showSorterTooltip: {
                target: 'full-header'
            },
            filters: laboratories.map(laboratory => ({
                text: laboratory.title,
                value: laboratory.id
            })),
            sorter: (a, b) => a.title.localeCompare(b.title),
            sortDirections: ['ascend', 'descend']
        },
        {
            title: 'Город',
            dataIndex: 'city',
            key: 'city',
            fixed: 'left',
            showSorterTooltip: {
                target: 'full-header'
            },
            filters: cities.map(city => ({
                text: city.title,
                value: city.id
            })),
            sorter: (a, b) => a.title.localeCompare(b.title),
            sortDirections: ['ascend', 'descend']
        },
        {
            title: 'Название',
            dataIndex: 'title',
            key: 'title',
            fixed: 'left',
            showSorterTooltip: {
                target: 'full-header'
            },
            sorter: (a, b) => parseInt(a.title) - parseInt(b.title),
            sortDirections: ['ascend', 'descend'],
            render: (_, record) => <div>Пост №{record.title}</div>
        },
        {
            title: 'Адрес',
            dataIndex: 'address',
            key: 'address',
            fixed: 'left',
        },
        {
            title: 'Опрашивается',
            key: 'is_listened',
            fixed: 'left',
            width: 80,
            render: (_, record) => (
                <div className={classNames('button-center')}>
                    <Switch
                        checked={!!record.isListened}
                        onClick={async () => await handleDisableOrEnableListened(record)}
                    />
                </div>
            ),
        },
    ]

    return (
        <div className={classNames('posts-table')}>
            <div className={classNames('posts-table-toolbar')}>
                <Search
                    title='Поиск по лаборатории, горороду, названию и адрессу'
                    size='large'
                    allowClear
                    onSearch={handleSearch}
                />
                <Button
                    title='Синхронизирует станции с АСОИЗА+'
                    icon={<SyncOutlined />}
                    size='large'
                    loading={isLoading}
                    onClick={handleSynchronize}
                >
                    Синхронизировать
                </Button>
            </div>
            <div>

            </div>
            <div className={classNames('posts-table-container')}>
                <Table
                    bordered
                    loading={isFetching}
                    columns={columns}
                    pagination={false}
                    rowKey='id'
                    onChange={handleChangeFilter}
                    dataSource={posts}
                />
            </div>
        </div>
    )
}