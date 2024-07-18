type AnyFn = (...params: any[]) => void;

export default function debounce<T extends AnyFn>(fn: T, ms = 500) {
    let timeout: ReturnType<typeof setTimeout> | undefined;

    return (...params: Parameters<T>) => {
        if (timeout) clearTimeout(timeout);

        timeout = setTimeout(() => {
            fn(...params);
        }, ms);
    };
}
