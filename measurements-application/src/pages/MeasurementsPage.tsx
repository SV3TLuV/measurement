import classNames from "classnames";
import {Split} from "@geoffcox/react-splitter";
import {TreeMenu} from "../features/measurements/components/TreeMenu.tsx";
import {useState} from "react";
import {Drawer, FloatButton} from "antd";
import {Loading} from "../components/ui/loading/Loading.tsx";
import {useMediaQuery} from "react-responsive";
import './MeasurementsPage.scss';
import {MenuOutlined} from "@ant-design/icons";
import {useDialog} from "../hooks/useDialog.ts";
import {MeasurementsTabs} from "../features/measurements/components/MeasurementsTabs.tsx";

export const MeasurementsPage = () => {
    const [loading, setLoading] = useState(true)
    const {open, show, close} = useDialog(true)

    const handleLoading = (loading: boolean) => setLoading(loading)

    const isMobile = useMediaQuery({ maxWidth: 767 });
    const isTablet = useMediaQuery({ minWidth: 768, maxWidth: 1024 });

    return (
        <div className={classNames('measurements-page')}>
            {isMobile || isTablet ? (
                <>
                    <Drawer
                        title='Лаборатории'
                        placement="left"
                        onClose={close}
                        open={open}
                        className={classNames('tree-menu-drawer')}
                    >
                        <div className={classNames('tree-menu-container')}>
                            <TreeMenu
                                onLoading={handleLoading}
                            />
                        </div>
                    </Drawer>
                    <FloatButton
                        onClick={show}
                        icon={<MenuOutlined/>}
                    />
                    <div className={classNames('map-table-container')}>
                        <MeasurementsTabs/>
                    </div>
                </>
            ) : (
                <Split initialPrimarySize='360px'>
                    <div className={classNames('tree-menu-container')}>
                        <TreeMenu
                            onLoading={handleLoading}
                        />
                    </div>
                    <div className={classNames('map-table-container')}>
                        <MeasurementsTabs/>
                    </div>
                </Split>
            )}
            <Loading loading={loading}/>
        </div>
    )
}