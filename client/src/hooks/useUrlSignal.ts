import { createEffect, createSignal, on } from "solid-js";
import { useBeforeLeave, useSearchParams } from "@solidjs/router";

import debounce from "utils/debounce";

type Arg<T> = {
    key: string;
    fromQuery: (query: string) => T;
    def: T;
};

export default function useUrlSignal<T extends string | number | boolean>({
    key,
    def,
    fromQuery,
}: Arg<T>) {
    const [params, setParams] = useSearchParams();

    const getParam = () => params[key];
    const setParam = (value: string) => setParams({ [key]: value }, { replace: true });

    const [value, setValue] = createSignal<T>(def);

    const setValueToParam = () => {
        const param = getParam();
        if (param) setValue(() => fromQuery(param));
    };

    setValueToParam();

    useBeforeLeave(ev => {
        if (typeof ev.to !== "string") return;

        const [, query] = ev.to.split("?");

        if (!query) {
            setValue(() => def);
            return;
        }

        const latest = new URLSearchParams(query).get(key);

        if (!latest) {
            setValue(() => def);
            return;
        }

        setValue(() => fromQuery(latest));
    });

    createEffect(
        on(
            value,
            debounce(value => setParam(String(value)), 1e3),
            { defer: true }
        )
    );

    return [value, setValue] as const;
}
