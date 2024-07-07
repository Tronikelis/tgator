export default function readableNumber(num: number): string {
    let div = 1e3;

    if (num < div) {
        return num + "";
    }

    if ((num /= div) < div) {
        return num.toFixed(1) + "K";
    }

    if ((num /= div) < div) {
        return num.toFixed(1) + "M";
    }

    num /= div;

    return num.toFixed(1) + "B";
}
