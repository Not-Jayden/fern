/**
 * This file was auto-generated by Fern from our API Definition.
 */

import * as core from "./core";
import * as SeedPackageYml from "./api/index";
import * as serializers from "./serialization/index";
import urlJoin from "url-join";
import * as errors from "./errors/index";
import { Service } from "./api/resources/service/client/Client";

export declare namespace SeedPackageYmlClient {
    interface Options {
        environment: core.Supplier<string>;
        id: string;
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

export class SeedPackageYmlClient {
    constructor(protected readonly _options: SeedPackageYmlClient.Options) {}

    /**
     * @param {SeedPackageYml.EchoRequest} request
     * @param {SeedPackageYmlClient.RequestOptions} requestOptions - Request-specific configuration.
     *
     * @example
     *     await client.echo({
     *         name: "Hello world!",
     *         size: 20
     *     })
     */
    public echo(
        request: SeedPackageYml.EchoRequest,
        requestOptions?: SeedPackageYmlClient.RequestOptions
    ): core.APIPromise<string> {
        return core.APIPromise.from(
            (async () => {
                const _response = await core.fetcher({
                    url: urlJoin(
                        await core.Supplier.get(this._options.environment),
                        `/${encodeURIComponent(this._options.id)}/`
                    ),
                    method: "POST",
                    headers: {
                        "X-Fern-Language": "JavaScript",
                        "X-Fern-SDK-Name": "@fern/package-yml",
                        "X-Fern-SDK-Version": "0.0.1",
                        "User-Agent": "@fern/package-yml/0.0.1",
                        "X-Fern-Runtime": core.RUNTIME.type,
                        "X-Fern-Runtime-Version": core.RUNTIME.version,
                        ...requestOptions?.headers,
                    },
                    contentType: "application/json",
                    requestType: "json",
                    body: serializers.EchoRequest.jsonOrThrow(request, { unrecognizedObjectKeys: "strip" }),
                    timeoutMs:
                        requestOptions?.timeoutInSeconds != null ? requestOptions.timeoutInSeconds * 1000 : 60000,
                    maxRetries: requestOptions?.maxRetries,
                    abortSignal: requestOptions?.abortSignal,
                });
                if (_response.ok) {
                    return {
                        ok: _response.ok,
                        body: serializers.echo.Response.parseOrThrow(_response.body, {
                            unrecognizedObjectKeys: "passthrough",
                            allowUnrecognizedUnionMembers: true,
                            allowUnrecognizedEnumValues: true,
                            breadcrumbsPrefix: ["response"],
                        }),
                        headers: _response.headers,
                    };
                }
                if (_response.error.reason === "status-code") {
                    throw new errors.SeedPackageYmlError({
                        statusCode: _response.error.statusCode,
                        body: _response.error.body,
                    });
                }
                switch (_response.error.reason) {
                    case "non-json":
                        throw new errors.SeedPackageYmlError({
                            statusCode: _response.error.statusCode,
                            body: _response.error.rawBody,
                        });
                    case "timeout":
                        throw new errors.SeedPackageYmlTimeoutError("Timeout exceeded when calling POST /{id}/.");
                    case "unknown":
                        throw new errors.SeedPackageYmlError({
                            message: _response.error.errorMessage,
                        });
                }
            })()
        );
    }

    protected _service: Service | undefined;

    public get service(): Service {
        return (this._service ??= new Service(this._options));
    }
}
