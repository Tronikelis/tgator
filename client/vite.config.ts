import { defineConfig } from "vite";
import tsconfigPaths from "vite-tsconfig-paths";
import solid from "vite-plugin-solid";

export default defineConfig({
    plugins: [solid(), tsconfigPaths()],
    server: {
        proxy: {
            "/api/v1": "http://localhost:3000",
        },
    },
});
