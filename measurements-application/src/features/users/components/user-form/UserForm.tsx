import {Button, Modal, Steps} from "antd";
import {useEffect, useState} from "react";
import classNames from "classnames";
import {UserPostsForm} from "./user-posts-form/UserPostsForm.tsx";
import {UserColumnsForm} from "./user-columns-form/UserColumnsForm.tsx";
import {FaBroadcastTower, FaColumns, FaRegUser, FaUserShield} from "react-icons/fa";
import {useDisabled} from "../../../../hooks/useDisabled.ts";
import {useSteps} from "../../../../hooks/useSteps.ts";
import {UserFormObject} from "../../types/userFormObject.ts";
import {User} from "../../types/user.ts";
import {userToUserFormObject} from "../../utils/userToUserFormObject.ts";
import './UserForm.scss';
import {UserPermissionsForm} from "./user-permissions-form/UserPermissionsForm.tsx";
import {UserDataForm} from "./user-data-form/UserDataForm.tsx";

const {Step} = Steps;

type UserFormProps  = {
    title: string
    open: boolean
    close: () => void
    ok: (user: UserFormObject) => void
    isCreateAction: boolean
    user: User
}

export const UserForm = ({title, open, close, ok, isCreateAction, user}: UserFormProps) => {
    const [formObject, setFormObject] =
        useState<UserFormObject>(userToUserFormObject(user))
    const {enabled, enable, disable} = useDisabled(!isCreateAction)
    const {step, next, previous, reset: resetStep} = useSteps()

    const handleOk = () => {
        ok(formObject)
        handleClose()
    }
    const handleClose = () => {
        setFormObject(userToUserFormObject(user))
        resetStep()
        close()
    }
    const handleChange = (user: UserFormObject) => setFormObject(prev => ({...prev, ...user}))

    const renderContent = () => {
        switch (step) {
            case 0:
                return (
                    <UserDataForm
                        formObject={formObject}
                        onChange={handleChange}
                        isCreateAction={isCreateAction}
                        enableNextButton={enable}
                        disableNextButton={disable}
                    />
                )
            case 1:
                return (
                    <UserPermissionsForm
                        formObject={formObject}
                        onChange={handleChange}
                    />
                )
            case 2:
                return (
                    <UserPostsForm
                        formObject={formObject}
                        onChange={handleChange}
                    />
                )
            case 3:
                return (
                    <UserColumnsForm
                        formObject={formObject}
                        onChange={handleChange}
                    />
                )
            default:
                return null
        }
    }

    useEffect(() => {
        setFormObject(userToUserFormObject(user));
    }, [user]);

    return (
        <Modal
            title={title}
            onOk={handleOk}
            onCancel={handleClose}
            open={open}
            footer={null}
            width={800}
            className={classNames('user-modal')}
        >
            <div className={classNames('steps-content')}>
                {renderContent()}
            </div>
            <Steps current={step}>
                <Step title="Данные" icon={<FaRegUser/>} />
                <Step title="Права" icon={<FaUserShield/>} />
                <Step title="Станции" icon={<FaBroadcastTower/>} />
                <Step title="Колонки" icon={<FaColumns/>} />
            </Steps>
            <div className={classNames('steps-action')}>
                {step > 0 && (
                    <Button type="link" onClick={previous} className={classNames('back-button')}>
                        Назад
                    </Button>
                )}
                {step < 3 && (
                    <Button disabled={!enabled} type="link" onClick={next} className={classNames('next-button')}>
                        Далее
                    </Button>
                )}
                {step === 3 && (
                    <Button disabled={!enabled} type="link" onClick={handleOk} className={classNames('done-button')}>
                        Готово
                    </Button>
                )}
            </div>
        </Modal>
    )
}