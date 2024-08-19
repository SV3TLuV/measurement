import classNames from "classnames";
import {UsersTable} from "../features/users/components/UsersTable.tsx";
import './UsersPage.scss';

export const UsersPage = () => {
    return (
        <div className={classNames('users-page')}>
            <UsersTable/>
        </div>
    )
}