/**
 * This file was auto-generated by Fern from our API Definition.
 */

import * as core from "../../../../core";
import * as SeedNurseryApi from "../../../index";
import * as errors from "../../../../errors/index";

export declare namespace Package {
    interface Options {
        environment: core.Supplier<string>;
    }

    interface RequestOptions {
        /** The maximum time to wait for a response in seconds. */
        timeoutInSeconds?: number;
        /** The number of times to retry the request. Defaults to 2. */
        maxRetries?: number;
        /** A hook to abort the request. */
        abortSignal?: AbortSignal;
    }
}

export class Package {
    constructor(protected readonly _options: Package.Options) {
    }

    /**
     * @param {SeedNurseryApi.TestRequest} request
     * @param {Package.RequestOptions} requestOptions - Request-specific configuration.
     *
     * @example
     *     await client.package.test({
     *         for: "string"
     *     })
     */
    public async test(request: SeedNurseryApi.TestRequest, requestOptions?: Package.RequestOptions): Promise<void> {
        const { "for": for_ } = request;
        const _queryParams: Record<string, string | string[] | object | object[]> = {};
        _queryParams["for"] = for_;
        const _response = await core.fetcher({
            url: await core.Supplier.get(this._options.environment),
            method: "POST",
            headers: {
                "X-Fern-Language": "JavaScript",
                "X-Fern-SDK-Name": "@fern/reserved-keywords",
                "X-Fern-SDK-Version": "0.0.1",
                "X-Fern-Runtime": core.RUNTIME.type,
                "X-Fern-Runtime-Version": core.RUNTIME.version
            },
            contentType: "application/json",
            queryParameters: _queryParams,
            requestType: "json",
            timeoutMs: requestOptions?.timeoutInSeconds != null ? (requestOptions.timeoutInSeconds * 1000) : 60000,
            maxRetries: requestOptions?.maxRetries,
            abortSignal: requestOptions?.abortSignal
        });
        if (_response.ok) {
            return;
        }

        if (_response.error.reason === "status-code") {
            throw new errors.SeedNurseryApiError({
                statusCode: _response.error.statusCode,
                body: _response.error.body
            });
        }

        switch (_response.error.reason) {
            case "non-json": throw new errors.SeedNurseryApiError({
                statusCode: _response.error.statusCode,
                body: _response.error.rawBody
            });
            case "timeout": throw new errors.SeedNurseryApiTimeoutError;
            case "unknown": throw new errors.SeedNurseryApiError({
                message: _response.error.errorMessage
            });
        }
    }
}
