import {ConfigurationForm} from "../features/app/components/ConfigurationForm.tsx";
import classNames from "classnames";
import './ConfigurationPage.scss';

export const ConfigurationPage = () => {
    return (
        <div className={classNames('configuration-page')}>
            <ConfigurationForm/>
        </div>
    );
}