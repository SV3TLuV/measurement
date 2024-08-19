import {MeasurementTabKeys} from "../enums/MeasurementTabKeys.ts";
import {ObjectMap} from "./ObjectMap.tsx";
import {MeasurementsTable} from "./MeasurementsTable.tsx";
import {Tabs, TabsProps} from "antd";
import {MeasurementsToolbar} from "./MeasurementsToolbar.tsx";
import {useRef, useState} from "react";
import {useUserPermissions} from "../../../hooks/useUserPermissions.ts";

export const MeasurementsTabs = () => {
    const [activeTab, setActiveTab] = useState(MeasurementTabKeys.Map)
    const mapRef = useRef<HTMLDivElement>(null)

    const permissions = useUserPermissions()

    const handleChangeTab: TabsProps['onChange'] = (tab) => {
        setActiveTab(tab as MeasurementTabKeys)
    }

    return (
        <Tabs
            type='card'
            onChange={handleChangeTab}
            tabBarExtraContent={permissions.canExport && (
                <MeasurementsToolbar
                    activeTab={activeTab}
                    mapRef={mapRef}
                />
            )}
            items={[
                {
                    label: 'Карта',
                    key: MeasurementTabKeys.Map,
                    children: (
                        <ObjectMap
                            mapRef={mapRef}
                        />
                    )
                },
                {
                    label: 'Данные',
                    key: MeasurementTabKeys.Table,
                    children: (
                        <MeasurementsTable/>
                    ),
                }
            ]}
        />
    )
}