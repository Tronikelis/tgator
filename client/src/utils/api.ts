import Gauk from "gauk";

export const api = new Gauk({
    options: {
        baseUrl: "/api/v1",
        headers: {
            "Content-Type": "application/json",
            Accept: "application/json",
        },
    },
});
