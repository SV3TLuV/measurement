import classNames from "classnames";
import './SchedulerInformationPage.scss';
import {SchedulerInformationCard} from "../features/collector/components/SchedulerInformationCard.tsx";

export const SchedulerInformationPage = () => {
    return (
        <div className={classNames('scheduler-information-page')}>
            <SchedulerInformationCard/>
        </div>
    );
};