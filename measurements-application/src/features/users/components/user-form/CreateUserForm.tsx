import {useCreateUserMutation} from "../../api/userApi.ts";
import {useShowResult} from "../../../../pages/result/hooks/useShowResult.tsx";
import {UserFormObject} from "../../types/userFormObject.ts";
import {CreateUserCommand} from "../../types/createUserCommand.ts";
import {UserForm} from "./UserForm.tsx";
import {User} from "../../types/user.ts";

type CreateUserFormProps = {
    open: boolean
    close: () => void
}

export const CreateUserForm = ({open, close}: CreateUserFormProps) => {
    const [create] = useCreateUserMutation()

    const showResult = useShowResult()

    const handleSubmit = async (user: UserFormObject) => {
        const response = await create(user as CreateUserCommand)
        await showResult({
            response: response,
            getSuccessMessage: () => ({
                type: 'message',
                title: `Пользователь: ${user.login} добавлен`
            })
        })
    }

    return (
        <UserForm
            title='Добавление пользователя'
            open={open}
            close={close}
            ok={handleSubmit}
            isCreateAction={true}
            user={{} as User}
        />
    )
}