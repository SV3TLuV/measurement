import {useState} from "react";

export const useDisabled = (defaultValue: boolean = false) => {
    const [enabled, setEnabled] = useState(defaultValue);
    const enable = () => setEnabled(true);
    const disable = () => setEnabled(false);
    return {enabled, enable, disable}
}