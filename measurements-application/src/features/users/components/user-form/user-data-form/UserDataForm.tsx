import {Form, Input, Select} from 'antd';
import {useCallback, useEffect} from "react";
import {debounce} from "lodash";
import {UserFormObject} from "../../../types/userFormObject.ts";
import {useGetRolesQuery} from "../../../../roles/api/roleApi.ts";

const { Option } = Select;

type UserDataFormProps = {
    isCreateAction: boolean
    formObject: UserFormObject
    onChange: (value: UserFormObject) => void
    enableNextButton: () => void
    disableNextButton: () => void
}

export const UserDataForm = (props: UserDataFormProps) => {
    const {
        isCreateAction,
        onChange,
        enableNextButton,
        disableNextButton,
        formObject
    } = props

    const [form] = Form.useForm<UserFormObject>()
    const values = Form.useWatch([], form);

    const {data: roles = []} = useGetRolesQuery()

    const debounceValidation = useCallback(debounce(() => {
        form.validateFields({ validateOnly: true })
            .then((value) => {
                enableNextButton()
                onChange(value)
            })
            .catch(() => disableNextButton());
    }, 200), [form]);

    useEffect(() => {
        debounceValidation();

        return () => debounceValidation.cancel();
    }, [debounceValidation, form, values]);

    useEffect(() => {
        form.setFieldsValue(formObject);
    }, [formObject, form]);

    return (
        <Form
            form={form}
            scrollToFirstError
        >
            <Form.Item
                name="login"
                label="Логин"
                rules={[
                    {
                        required: true,
                        message: 'Пожалуйста, введите логин!',
                    },
                    {
                        max: 100,
                        message: 'Логин не должен превышать 20 символов!'
                    },
                    {
                        min: 5,
                        message: 'Логин не должен быть меньше 8 символов!'
                    },
                    {
                        pattern: /^[a-zA-Z0-9]+$/,
                        message: 'Разрешены только цифры и английские символы!'
                    },
                ]}
            >
                <Input
                    count={{
                        show: true,
                        max: 100
                    }}
                />
            </Form.Item>

            {isCreateAction &&
                <Form.Item
                    name="password"
                    label="Пароль"
                    rules={[
                        {
                            required: true,
                            message: 'Пожалуйста, введите пароль!',
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
                    hasFeedback
                >
                    <Input.Password
                        count={{
                            show: true,
                            max: 64
                        }}
                    />
                </Form.Item>
            }

            {isCreateAction &&
                <Form.Item
                    name="confirm"
                    label="Подтвердите пароль"
                    dependencies={['password']}
                    hasFeedback
                    rules={[
                        {
                            required: true,
                            message: 'Пожалуйста, подтвердите пароль!',
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
            }

            <Form.Item
                name="roleId"
                label="Роль"
                rules={[
                    {
                        required: true,
                        message: 'Пожалуйста, выберите роль!'
                    }
                ]}
            >
                <Select
                    placeholder="Выберите роль"
                    showSearch
                >
                    {roles && roles.map(role => (
                        <Option key={role.id} value={role.id}>{role.title}</Option>
                    ))}
                </Select>
            </Form.Item>
        </Form>
    )
}