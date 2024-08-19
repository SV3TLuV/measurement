import classNames from "classnames";
import {SignInForm} from "../features/auth/components/SignInForm.tsx";
import './SignInPage.scss';

export const SignInPage = () => {
    return (
        <div className={classNames('sign-in-page')}>
            <div className={classNames('sign-in-form-container')}>
                <SignInForm />
            </div>
        </div>
    )
}