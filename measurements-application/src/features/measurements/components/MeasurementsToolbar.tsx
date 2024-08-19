import {Button, Dropdown, MenuProps, Space} from "antd";
import classNames from "classnames";
import './MeasurementsToolbar.scss';
import {MeasurementTabKeys} from "../enums/MeasurementTabKeys.ts";
import {DownOutlined} from "@ant-design/icons";
import {RefObject} from "react";
import html2canvas from "html2canvas";
import {useTypedSelector} from "../../../hooks/redux.ts";
import {useDownloadMeasurementsMutation} from "../api/measurementApi.ts";
import {ExportMeasurementsQuery} from "../types/exportMeasurementsQuery.ts";

type MeasurementsToolbarProps = {
    activeTab: MeasurementTabKeys
    mapRef: RefObject<HTMLDivElement>
}

export const MeasurementsToolbar = (props: MeasurementsToolbarProps) => {
    const { activeTab, mapRef } = props;

    const { selected } = useTypedSelector(state => state.menu)
    const { startDate, endDate, period } = useTypedSelector(state => state.measurementFilters)

    const [ download ] = useDownloadMeasurementsMutation()

    const handleSaveMap: MenuProps['onClick'] = async (e) => {
        if (mapRef.current) {
            const canvas = await html2canvas(mapRef.current);
            const link = document.createElement('a');
            link.href = canvas.toDataURL(`image/${e.key}`);
            link.download = `map.${e.key}`;
            link.click();
        }
    }

    const handleExportMeasurements: MenuProps['onClick'] = async (e) => {
        const query = {
            objectId: selected?.id,
            pageSize: 0,
            page: 0,
            period: period,
            startDate: startDate?.format("YYYY-MM-DD"),
            endDate: endDate?.format("YYYY-MM-DD"),
            format: e.key
        } as ExportMeasurementsQuery

        switch (e.key) {
            case 'xlsx':
                await download(query);
                break
            case 'csv':
                await download(query);
                break
        }
    }

    const saveMapMenu: MenuProps = {
        items: [
            {
                key: 'png',
                label: 'Экспортировать в PNG',
                onClick: handleSaveMap
            },
            {
                key: 'jpg',
                label: 'Экспортировать в JPG',
                onClick: handleSaveMap
            }
        ]
    }

    const exportMeasurementsMenu: MenuProps = {
        items: [
            {
                key: 'xlsx',
                label: 'Экспортировать в XLSX',
                onClick: handleExportMeasurements
            },
            {
                key: 'csv',
                label: 'Экспортировать в CSV',
                onClick: handleExportMeasurements
            }
        ]
    }

    const getContent = () => {
        switch (activeTab) {
            case MeasurementTabKeys.Map:
                return (
                    <Dropdown menu={saveMapMenu}>
                        <Button>
                            Экспортировать <DownOutlined />
                        </Button>
                    </Dropdown>
                )
            case MeasurementTabKeys.Table:
                return (
                    <Dropdown menu={exportMeasurementsMenu}>
                        <Button>
                            Экспортировать <DownOutlined />
                        </Button>
                    </Dropdown>
                )
            default:
                return null
        }
    }

    return (
        <div className={classNames('measurements-toolbar-container')}>
            <Space>
                {getContent()}
            </Space>
        </div>
    )
}