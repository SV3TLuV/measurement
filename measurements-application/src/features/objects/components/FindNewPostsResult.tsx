import {Button, Table, TableColumnsType, Tag} from "antd";
import classNames from "classnames";
import {useNavigate} from "react-router-dom";
import {GetAndAddNewPostsQueryResult} from "../types/getAndAddNewPostsQueryResult.ts";
import {Facility} from "../types/facility.ts";
import {RoutePaths} from "../../../lib/router/enums/routePaths.ts";
import {getUrlFromRoute} from "../../../utils/getUrlFromRoute.ts";
import './FindNewPostsResult.scss';

type FindNewPostsResultProps = {
    result: GetAndAddNewPostsQueryResult
}

export const FindNewPostsResult = ({ result }: FindNewPostsResultProps) => {
    const columns: TableColumnsType<Facility> = [
        {
            title: '',
            key: 'tags',
            dataIndex: 'tags',
            width: 40,
            render: (_, record) => {
                if (record.status) {
                    return (
                        <Tag
                            className={classNames('tag-center')}
                            key={record.id}
                            color='blue'
                        >
                            Новый
                        </Tag>
                    )
                }

                // TODO: fix it.

                if (record.status) {
                    return (
                        <Tag
                            className={classNames('tag-center')}
                            key={record.id}
                            color='yellow'
                        >
                            Обновлен
                        </Tag>
                    )
                }
            }
        },
        {
            title: 'Лаборатория',
            dataIndex: 'laboratory',
            key: 'laboratory',
            fixed: 'left',
            showSorterTooltip: {
                target: 'full-header'
            },
            sorter: (a, b) => a.laboratory.localeCompare(b.laboratory),
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
            sorter: (a, b) => a.city.localeCompare(b.city),
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
    ]

    return (
        <Table
            bordered
            columns={columns}
            pagination={false}
            rowKey='id'
            dataSource={result.posts}
        />
    )
}

export const FindNewPostResultEmpty = () => {
    const navigate = useNavigate()

    return (
        <div className={classNames('result-container')}>
            <Button
                size='large'
                type='primary'
                onClick={() => navigate(getUrlFromRoute(RoutePaths.Posts))}
            >
                Ок
            </Button>
        </div>
    )
}

type FindNewPostResultWithDataProps = {
    result: GetAndAddNewPostsQueryResult
}

export const FindNewPostResultWithData = (props: FindNewPostResultWithDataProps) => {
    const navigate = useNavigate()

    return (
        <div className={classNames('result-container')}>
            <div className={classNames('table-container')}>
                <FindNewPostsResult {...props} />
            </div>
            <div className={classNames('actions')}>
                <Button
                    size='large'
                    type='primary'
                    onClick={() => navigate(`/${RoutePaths.Posts}`)}
                >
                    Ок
                </Button>
            </div>
        </div>
    )
}