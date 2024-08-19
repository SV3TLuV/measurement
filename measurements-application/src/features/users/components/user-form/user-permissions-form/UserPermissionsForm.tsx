import {Table, TableColumnsType, TableProps} from "antd";
import React, {useState} from "react";
import {UserFormObject} from "../../../types/userFormObject.ts";
import {useGetPermissionsQuery} from "../../../../permissions/api/permissionApi.ts";
import {Permission} from "../../../../permissions/types/permission.ts";

type TableRowSelection<T> = TableProps<T>['rowSelection']

type UserPermissionsFormProps = {
    formObject: UserFormObject
    onChange: (user: UserFormObject) => void
}

export const UserPermissionsForm = ({onChange, formObject}: UserPermissionsFormProps) => {
    const {data: permissionsData = [], isFetching} = useGetPermissionsQuery()

    const [selectedRowKeys, setSelectedRowKeys] =
        useState<React.Key[]>(formObject.permissionIds);

    const onSelectChange = (newSelectedRowKeys: React.Key[]) => {
        setSelectedRowKeys(newSelectedRowKeys);
        onChange({ permissionIds: newSelectedRowKeys } as UserFormObject)
    };

    const rowSelection: TableRowSelection<Permission> = {
        selectedRowKeys,
        onChange: onSelectChange,
        selections: [
            Table.SELECTION_ALL,
            Table.SELECTION_INVERT,
            Table.SELECTION_NONE,
        ],
    };

    const columns: TableColumnsType<Permission> = [
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
    ]

    return (
        <Table
            bordered
            loading={isFetching}
            columns={columns}
            pagination={false}
            rowSelection={rowSelection}
            rowKey='id'
            dataSource={permissionsData}
        />
    )
}