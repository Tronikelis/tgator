import { render } from "solid-js/web";
import { Router, Route } from "@solidjs/router";
import Gauk from "gauk";
import { SWROptionsProvider } from "solid-swr";

import Idx from "routes/idx";

function App() {
    const fetcher = new Gauk({
        options: {
            baseUrl: "/api/v1",
            headers: {
                "Content-Type": "application/json",
                Accept: "application/json",
            },
        },
    });

    return (
        <SWROptionsProvider value={{ fetcher: key => fetcher.get(key) }}>
            <Router>
                <Route path="/" component={Idx} />
            </Router>
        </SWROptionsProvider>
    );
}

render(() => <App />, document.getElementById("solid-root")!);
