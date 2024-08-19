import { Modal, Form, Input, Button } from 'antd';
import classNames from "classnames";
import {useChangePasswordMutation} from "../api/userApi.ts";
import {User} from "../types/user.ts";
import {ChangePasswordCommand} from "../types/changePasswordCommand.ts";
import {useShowResult} from "../../../pages/result/hooks/useShowResult.tsx";
import './ChangePasswordForm.scss';

type ChangePasswordFormProps = {
    user: User
    open: boolean
    close: () => void
}

export const ChangePasswordForm = ({user, open, close}: ChangePasswordFormProps) => {
    const [changePassword] = useChangePasswordMutation()
    const [form] = Form.useForm<ChangePasswordCommand>();

    const showResult = useShowResult()

    const handleSubmit = async () => {
        try {
            const values = await form.validateFields();
            values.userId = user.id;

            handleClose();

            const response = await changePassword(values)
            await showResult({
                response: response,
                getSuccessMessage: () => ({
                    type: 'message',
                    title: `Пароль пользователя: ${user.login} изменен`
                })
            })
        } catch { /* empty */ }
    };

    const handleClose = () => {
        form.resetFields()
        close()
    }

    return (
        <Modal
            title="Смена пароля"
            open={open}
            onCancel={handleClose}
            className={classNames('change-password-modal')}
            footer={[
                <Button
                    className={classNames('change-button')}
                    key="submit"
                    type="primary"
                    onClick={handleSubmit}
                >
                    Сменить пароль
                </Button>,
            ]}
        >
            <Form className={classNames('form-content')} form={form} layout="horizontal">
                <Form.Item
                    name="password"
                    label="Новый пароль"
                    rules={[
                        {
                            required: true,
                            message: 'Пожалуйста, введите новый пароль!'
                        },
                        {
                            max: 64,
                            message: 'Пароль не должен превышать 64 символов!'
                        },
                        {
                            min: 12,
                            message: 'Пароль не должен быть меньше 12 символов!'
                        },
                        {
                            pattern: /^[a-zA-Z0-9!@#$%^&*()-]+$/,
                            message: 'Разрешены только цифры, специальные и английские символы!'
                        }
                    ]}
                >
                    <Input.Password
                        count={{
                            show: true,
                            max: 64
                        }}
                    />
                </Form.Item>
                <Form.Item
                    name="confirm"
                    label="Подтвердите новый пароль"
                    dependencies={['newPassword']}
                    hasFeedback
                    rules={[
                        {
                            required: true,
                            message: 'Пожалуйста, подтвердите новый пароль!'
                        },
                        {
                            max: 64,
                            message: 'Подтверждение пароля не должно превышать 64 символов!'
                        },
                        {
                            min: 12,
                            message: 'Подтверждение пароля не должно быть меньше 12 символов!'
                        },
                        {
                            pattern: /^[a-zA-Z0-9!@#$%^&*()-]+$/,
                            message: 'Разрешены только цифры, специальные и английские символы!'
                        },
                        ({ getFieldValue }) => ({
                            validator(_, value) {
                                if (!value || getFieldValue('password') === value) {
                                    return Promise.resolve();
                                }
                                return Promise.reject(new Error('Введенные пароли не совпадают!'));
                            },
                        }),
                    ]}
                >
                    <Input.Password
                        count={{
                            show: true,
                            max: 64
                        }}
                    />
                </Form.Item>
            </Form>
        </Modal>
    )
}