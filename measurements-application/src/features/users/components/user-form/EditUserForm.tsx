import {User} from "../../types/user.ts";
import {useEditUserMutation} from "../../api/userApi.ts";
import {useShowResult} from "../../../../pages/result/hooks/useShowResult.tsx";
import {UserFormObject} from "../../types/userFormObject.ts";
import {UpdateUserCommand} from "../../types/updateUserCommand.ts";
import {UserForm} from "./UserForm.tsx";

type UpdateUserFormProps = {
    user: User
    open: boolean
    close: () => void
}

export const EditUserForm = ({open, close, user}: UpdateUserFormProps) => {
    const [update] = useEditUserMutation()

    const showResult = useShowResult()

    const handleSubmit = async (user: UserFormObject) => {
        const response = await update(user as UpdateUserCommand)
        await showResult({
            response: response,
            getSuccessMessage: () => ({
                type: 'message',
                title: `Данные пользователя ${user.login} изменены`
            })
        })
    }

    return (
        <UserForm
            title='Обновление пользователя'
            open={open}
            close={close}
            ok={handleSubmit}
            isCreateAction={false}
            user={user}
        />
    )
}