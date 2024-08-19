import {Button, Form, Input} from "antd";
import {LockOutlined, UserOutlined} from "@ant-design/icons";
import classNames from "classnames";
import {useShowResult} from "../../../pages/result/hooks/useShowResult.tsx";
import {useLoginMutation} from "../api/authApi.ts";
import {LoginCommand} from "../types/loginCommand.ts";
import './SignInForm.scss';
import {getWelcomeMsg} from "../utils/getWelcomeMsg.ts";

export const SignInForm = () => {
    const [login] = useLoginMutation()
    const [form] = Form.useForm<LoginCommand>()
    const showResult = useShowResult()

    const onFinish = async () => {
        try {
            const values = await form.validateFields();
            const response = await login(values)
            await showResult({
                response: response,
                getSuccessMessage: (response) => ({
                    type: 'message',
                    title: getWelcomeMsg(response.accessToken),
                })
            })
            close();
        } catch { /* empty */ }
    }

    return (
        <Form
            form={form}
            className={classNames('sign-in-form')}
            onFinish={onFinish}
        >
            <h2 className={classNames('sign-in-form-header', 'form-item-center')}>
                Вход
            </h2>
            <Form.Item
                name="login"
                rules={[
                    { required: true, message: 'Пожалуйста, введите свой логин!' },
                    {
                        pattern: /^[a-zA-Z0-9]+$/,
                        message: 'Разрешены только цифры и английские символы!'
                    }
                ]}
            >
                <Input
                    prefix={<UserOutlined className={classNames('site-form-item-icon')} />}
                    placeholder="Логин"
                />
            </Form.Item>
            <Form.Item
                name="password"
                rules={[
                    { required: true, message: `Пожалуйста, введите свой пароль!` },
                    {
                        pattern: /^[a-zA-Z0-9!@#$%^&*()-]+$/,
                        message: 'Разрешены только цифры, специальные и английские символы!'
                    }
                ]}
            >
                <Input.Password
                    prefix={<LockOutlined className={classNames('site-form-item-icon')} />}
                    type="password"
                    placeholder="Пароль"
                />
            </Form.Item>

            <Form.Item className={classNames('form-item-center')}>
                <Button
                    type="primary"
                    htmlType="submit"
                    className={classNames('sign-in-form-button')}
                >
                    Войти
                </Button>
            </Form.Item>
        </Form>
    )
}