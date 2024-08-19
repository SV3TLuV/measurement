import {useAppDispatch, useTypedSelector} from "../../../hooks/redux.ts";
import {
    DatePicker,
    DatePickerProps,
    Pagination,
    Select,
    Space,
    Table,
    TableColumnsType
} from "antd";
import classNames from "classnames";
import {useEffect, useRef, useState} from "react";
import dayjs from "dayjs";
import {Loading} from "../../../components/ui/loading/Loading.tsx";
import {useGetMeQuery, useGetUserColumnsQuery} from "../../users/api/userApi.ts";
import {useGetQualitiesQuery} from "../../quality/api/qualityApi.ts";
import {Period} from "../../../shared/enums/period.ts";
import {Measurement} from "../types/measurement.ts";
import {getClassNameForColumn} from "../utils/getClassNameForColumn.ts";
import {useGetMeasurementsQuery} from "../api/measurementApi.ts";
import './MeasurementsTable.scss';
import {useMediaQuery} from "react-responsive";
import {setEndDate, setPeriod, setStartDate} from "../stores/measurementFiltersSlice.ts";
import {useUserPermissions} from "../../../hooks/useUserPermissions.ts";
import {convertFieldName} from "../utils/convertFieldName.ts";

const { Option } = Select;

const defaultPageSize = 50;
const defaultPage = 1;
const defaultPeriod = Period.All

const periods = [
    { period: Period.All, title: "за весь период" },
    { period: Period.Year, title: "за год" },
    { period: Period.Month, title: "за месяц" },
    { period: Period.Week, title: "за неделю" },
    { period: Period.Day, title: "за сутки (24 часа)" },
]

export const MeasurementsTable = () => {
    const {data: user } = useGetMeQuery()
    const { selected } = useTypedSelector(state => state.menu)
    const { startDate, endDate, period } = useTypedSelector(state => state.measurementFilters)

    const dispatch = useAppDispatch()

    const tableRef = useRef<HTMLDivElement>(null);

    const [pageSize, setPageSize] = useState<number>(defaultPageSize)
    const [page, setPage] = useState<number>(defaultPage)

    const isMobile = useMediaQuery({ maxWidth: 767 });

    const {data: availableColumns = []} = useGetUserColumnsQuery(user?.id ?? 0, {
        skip: !user
    })
    const {data: qualities = []} = useGetQualitiesQuery()

    const {data: measurementData, isFetching} = useGetMeasurementsQuery({
        objectId: selected?.id,
        pageSize: pageSize,
        page: page,
        period: period,
        start: startDate ? startDate.format("YYYY-MM-DD") + "T00:00:00Z" : null,
        end: endDate ? endDate.format("YYYY-MM-DD") + "T23:59:59Z" : null,
    })

    const permissions = useUserPermissions()

    const handleChangePeriod = (value: string) => dispatch(setPeriod(Number(value) as Period))

    const handleStartDateChange: DatePickerProps['onChange'] = (date) => {
        dispatch(setStartDate(date ? date : null));
    };

    const handleEndDateChange: DatePickerProps['onChange'] = (date) => {
        dispatch(setEndDate(date ? date : null));
    };

    const disabledDate: DatePickerProps['disabledDate'] = (current) => {
        const today = dayjs().startOf('day');
        if (startDate) {
            return current.isAfter(today) || current.isBefore(startDate);
        }
        return current.isAfter(today);
    }

    const columns: TableColumnsType<Measurement> = availableColumns.map(column => ({
        title: (
            <>
                {column.code && (
                    <>
                        <>({column.code})</>
                        <br/>
                    </>
                )}
                {column.title}
            </>
        ),
        dataIndex: convertFieldName(column.objField),
        key: convertFieldName(column.objField),
        align: 'center',
        fixed: 'left',
        render: (text, value) => {
            const withRounding = [
                'temp',
                'pressure'
            ]

            return (
                <div
                    className={classNames(
                        'measurements-table-cell',
                        getClassNameForColumn(column.objField, text, value, qualities)
                    )}
                >
                    {text && withRounding.includes(column.objField)
                        ? parseFloat(text).toFixed(1)
                        : text}
                </div>
            )
        }
    }))

    useEffect(() => {
        setPageSize(defaultPageSize)
        setPage(defaultPage)
        dispatch(setPeriod(defaultPeriod))
        dispatch(setStartDate(null))
        dispatch(setEndDate(null))
    }, [selected])

    return (
        <div className={classNames('measurements-table')}>
            <div className={classNames('ant-table-wrapper')}>
                {permissions.canFilter && (
                    <Space className={classNames('measurements-table-toolbar')}>
                        <Select
                            className={classNames('measurements-table-toolbar-select')}
                            defaultValue={`${defaultPeriod}`}
                            onChange={handleChangePeriod}
                            showSearch
                        >
                            {periods.map(period => (
                                <Option key={period.period} value={`${period.period}`}>
                                    {period.title}
                                </Option>
                            ))}
                        </Select>
                        <DatePicker
                            placeholder='С даты'
                            onChange={handleStartDateChange}
                            value={startDate}
                            disabledDate={disabledDate}
                        />
                        <DatePicker
                            placeholder='До даты'
                            onChange={handleEndDateChange}
                            value={endDate}
                            disabledDate={disabledDate}
                        />
                    </Space>
                )}
            </div>
            <div className={classNames('table-container')} ref={tableRef}>
                <Table
                    size='small'
                    bordered
                    pagination={false}
                    columns={columns}
                    rowKey='id'
                    dataSource={measurementData?.items ?? []}
                />
            </div>
            <div className={classNames('pagination-wrapper')}>
                <Pagination
                    showSizeChanger
                    showTotal={(total, range) => `Показано ${range[1]} из ${total}`}
                    defaultPageSize={defaultPageSize}
                    current={page}
                    pageSize={pageSize}
                    simple={isMobile}
                    total={measurementData?.total}
                    onChange={(page, size) => {
                        setPage(page)
                        setPageSize(size)
                    }}
                />
            </div>
            <Loading loading={isFetching}/>
        </div>
    )
}