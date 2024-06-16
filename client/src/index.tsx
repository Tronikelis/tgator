import { render } from "solid-js/web";
import { Router, Route } from "@solidjs/router";
import Gauk from "gauk";
import { SWROptionsProvider } from "solid-swr";

import Idx from "routes/idx";
import SourceId from "routes/sources_[id]";

import "./main.css";

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
        <SWROptionsProvider value={{ fetcher: key => fetcher.get(key).then(x => x.data) }}>
            <Router>
                <Route path="/" component={Idx} />
                <Route path="/sources/:id" component={SourceId} />
            </Router>
        </SWROptionsProvider>
    );
}

render(() => <App />, document.getElementById("solid-root")!);
