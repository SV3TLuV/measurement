import React, {useState} from "react";
import {Input, Table, TableColumnsType, TableProps} from "antd";
import DOMPurify from "dompurify";
import {UserFormObject} from "../../../types/userFormObject.ts";
import {useGetColumnsQuery} from "../../../../columns/api/columnApi.ts";
import {Column} from "../../../../columns/types/column.ts";

type TableRowSelection<T> = TableProps<T>['rowSelection']

const { Search } = Input;

type UserColumnsFormProps = {
    formObject: UserFormObject
    onChange: (user: UserFormObject) => void
}

export const UserColumnsForm = ({onChange, formObject}: UserColumnsFormProps) => {
    const [search, setSearch] = useState<string>('')

    const {columnsData, isFetching} = useGetColumnsQuery(undefined, {
        selectFromResult: ({ data, isFetching }) => ({
            columnsData: data?.filter(column =>
                column.title.toLowerCase().startsWith(search.toLowerCase())) ?? [],
            isFetching,
        })
    })

    const [selectedRowKeys, setSelectedRowKeys] =
        useState<React.Key[]>(formObject.columnIds);

    const onSelectChange = (newSelectedRowKeys: React.Key[]) => {
        setSelectedRowKeys(newSelectedRowKeys);
        onChange({ columnIds: newSelectedRowKeys } as UserFormObject)
    };

    const handleSearch = (text: string) => setSearch(text)

    const rowSelection: TableRowSelection<Column> = {
        selectedRowKeys,
        onChange: onSelectChange,
        selections: [
            Table.SELECTION_ALL,
            Table.SELECTION_INVERT,
            Table.SELECTION_NONE,
        ],
    };

    const columns: TableColumnsType<Column> = [
        {
            title: 'Название',
            dataIndex: 'title',
            key: 'title',
            fixed: 'left',
            showSorterTooltip: {
                target: 'full-header'
            },
            sorter: (a, b) => a.title.localeCompare(b.title),
            sortDirections: ['ascend', 'descend']
        },
        {
            title: 'Формула',
            dataIndex: 'formula',
            key: 'formula',
            fixed: 'left',
            render: (text) => (
                <div dangerouslySetInnerHTML={{__html: DOMPurify.sanitize(text)}}/>
            )
        },
    ]

    return (
        <>
            <Search
                title='Поиск по названию'
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
                dataSource={columnsData}
            />
        </>
    )
}