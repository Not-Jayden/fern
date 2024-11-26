/**
 * This file was auto-generated by Fern from our API Definition.
 */

import * as core from "./core";
import { Imdb } from "./api/resources/imdb/client/Client";

export declare namespace SeedApiClient {
    interface Options {
        environment: core.Supplier<string>;
        token?: core.Supplier<core.BearerToken | undefined>;
    }

    interface RequestOptions {
        /** The maximum time to wait for a response in seconds. */
        timeoutInSeconds?: number;
        /** The number of times to retry the request. Defaults to 2. */
        maxRetries?: number;
        /** A hook to abort the request. */
        abortSignal?: AbortSignal;
        /** Additional headers to include in the request. */
        headers?: Record<string, string>;
    }
}

export class SeedApiClient {
    constructor(protected readonly _options: SeedApiClient.Options) {
    }

    protected _imdb: Imdb | undefined;

    public get imdb(): Imdb {
        return (this._imdb ??= new Imdb(this._options));
    }
}
