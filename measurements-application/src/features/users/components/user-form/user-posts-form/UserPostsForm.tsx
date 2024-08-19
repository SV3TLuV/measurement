import {Input, Table, TableColumnsType, TableProps} from "antd";
import React, {useState} from "react";
import {UserFormObject} from "../../../types/userFormObject.ts";
import {Facility} from "../../../../objects/types/facility.ts";
import {useGetObjectsQuery} from "../../../../objects/api/objectApi.ts";
import {ObjectTypeKeys} from "../../../../objects/types/objectType.ts";

type TableRowSelection<T> = TableProps<T>['rowSelection']

const { Search } = Input;

type UserPostsFormProps = {
    formObject: UserFormObject
    onChange: (user: UserFormObject) => void
}

export const UserPostsForm = ({onChange, formObject}: UserPostsFormProps) => {
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

    const [selectedRowKeys, setSelectedRowKeys] =
        useState<React.Key[]>(formObject.postIds);

    const onSelectChange = (newSelectedRowKeys: React.Key[]) => {
        setSelectedRowKeys(newSelectedRowKeys);
        onChange({ postIds: newSelectedRowKeys } as UserFormObject)
    };

    const handleSearch = (text: string) => setSearch(text)

    const handleChangeFilter: TableProps['onChange'] =
        (_, filters) => {
            setLaboratoryIds(filters['laboratory'] as number[])
            setCityIds(filters['city'] as number[])
        }

    const rowSelection: TableRowSelection<Facility> = {
        selectedRowKeys,
        onChange: onSelectChange,
        selections: [
            Table.SELECTION_ALL,
            Table.SELECTION_INVERT,
            Table.SELECTION_NONE,
        ],
    };

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
            sorter: (a, b) => a.laboratory!.localeCompare(b.laboratory!),
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
            sorter: (a, b) => a.city!.localeCompare(b.city!),
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
            render: (_, record) => (
                <div>
                    Пост №{record.title}
                </div>
            )
        },
        {
            title: 'Адрес',
            dataIndex: 'address',
            key: 'address',
            fixed: 'left',
        },
    ]

    return (
        <>
            <Search
                title='Поиск по лаборатории, горороду, названию и адрессу'
                allowClear
                onSearch={handleSearch}
            />
            <Table
                bordered
                loading={isFetching}
                columns={columns}
                pagination={false}
                rowSelection={rowSelection}
                rowKey='id'
                onChange={handleChangeFilter}
                dataSource={posts}
            />
        </>
    )
}