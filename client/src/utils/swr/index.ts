import { createMemo } from "solid-js";
import urlbat, { Params } from "urlbat";

export function createSwrKey(base: string, arg: () => Params | undefined) {
    return createMemo(() => {
        const a = arg();
        if (!a) return;
        return urlbat(base, a);
    });
}
