export default function safeJsonPretty(str: string) {
    try {
        str = JSON.stringify(JSON.parse(str), null, 2);
        str = str.replaceAll("\\n", "\n");
        return str;
    } catch {
        return str;
    }
}
