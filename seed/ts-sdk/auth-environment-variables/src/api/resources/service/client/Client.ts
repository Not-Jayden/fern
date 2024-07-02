/**
 * This file was auto-generated by Fern from our API Definition.
 */

import * as core from "../../../../core";
import urlJoin from "url-join";
import * as serializers from "../../../../serialization/index";
import * as errors from "../../../../errors/index";
import * as SeedAuthEnvironmentVariables from "../../../index";

export declare namespace Service {
    interface Options {
        environment: core.Supplier<string>;
        apiKey?: core.Supplier<string | undefined>;
        xAnotherHeader: core.Supplier<string>;
    }

    interface RequestOptions {
        /** The maximum time to wait for a response in seconds. */
        timeoutInSeconds?: number;
        /** The number of times to retry the request. Defaults to 2. */
        maxRetries?: number;
        /** A hook to abort the request. */
        abortSignal?: AbortSignal;
        /** Override the X-Another-Header header */
        xAnotherHeader?: string;
    }
}

export class Service {
    constructor(protected readonly _options: Service.Options) {}

    /**
     * GET request with custom api key
     *
     * @param {Service.RequestOptions} requestOptions - Request-specific configuration.
     *
     * @example
     *     await client.service.getWithApiKey()
     */
    public async getWithApiKey(requestOptions?: Service.RequestOptions): Promise<string> {
        const _response = await core.fetcher({
            url: urlJoin(await core.Supplier.get(this._options.environment), "apiKey"),
            method: "GET",
            headers: {
                "X-Another-Header": await core.Supplier.get(this._options.xAnotherHeader),
                "X-Fern-Language": "JavaScript",
                "X-Fern-SDK-Name": "@fern/auth-environment-variables",
                "X-Fern-SDK-Version": "0.0.1",
                "X-Fern-Runtime": core.RUNTIME.type,
                "X-Fern-Runtime-Version": core.RUNTIME.version,
                ...(await this._getCustomAuthorizationHeaders()),
            },
            contentType: "application/json",
            timeoutMs: requestOptions?.timeoutInSeconds != null ? requestOptions.timeoutInSeconds * 1000 : 60000,
            maxRetries: requestOptions?.maxRetries,
            abortSignal: requestOptions?.abortSignal,
        });
        if (_response.ok) {
            return await serializers.service.getWithApiKey.Response.parseOrThrow(_response.body, {
                unrecognizedObjectKeys: "passthrough",
                allowUnrecognizedUnionMembers: true,
                allowUnrecognizedEnumValues: true,
                breadcrumbsPrefix: ["response"],
            });
        }

        if (_response.error.reason === "status-code") {
            throw new errors.SeedAuthEnvironmentVariablesError({
                statusCode: _response.error.statusCode,
                body: _response.error.body,
            });
        }

        switch (_response.error.reason) {
            case "non-json":
                throw new errors.SeedAuthEnvironmentVariablesError({
                    statusCode: _response.error.statusCode,
                    body: _response.error.rawBody,
                });
            case "timeout":
                throw new errors.SeedAuthEnvironmentVariablesTimeoutError();
            case "unknown":
                throw new errors.SeedAuthEnvironmentVariablesError({
                    message: _response.error.errorMessage,
                });
        }
    }

    /**
     * GET request with custom api key
     *
     * @param {SeedAuthEnvironmentVariables.HeaderAuthRequest} request
     * @param {Service.RequestOptions} requestOptions - Request-specific configuration.
     *
     * @example
     *     await client.service.getWithHeader({
     *         xEndpointHeader: "string"
     *     })
     */
    public async getWithHeader(
        request: SeedAuthEnvironmentVariables.HeaderAuthRequest,
        requestOptions?: Service.RequestOptions
    ): Promise<string> {
        const { xEndpointHeader } = request;
        const _response = await core.fetcher({
            url: urlJoin(await core.Supplier.get(this._options.environment), "apiKeyInHeader"),
            method: "GET",
            headers: {
                "X-Another-Header": await core.Supplier.get(this._options.xAnotherHeader),
                "X-Fern-Language": "JavaScript",
                "X-Fern-SDK-Name": "@fern/auth-environment-variables",
                "X-Fern-SDK-Version": "0.0.1",
                "X-Fern-Runtime": core.RUNTIME.type,
                "X-Fern-Runtime-Version": core.RUNTIME.version,
                "X-Endpoint-Header": xEndpointHeader,
                ...(await this._getCustomAuthorizationHeaders()),
            },
            contentType: "application/json",
            timeoutMs: requestOptions?.timeoutInSeconds != null ? requestOptions.timeoutInSeconds * 1000 : 60000,
            maxRetries: requestOptions?.maxRetries,
            abortSignal: requestOptions?.abortSignal,
        });
        if (_response.ok) {
            return await serializers.service.getWithHeader.Response.parseOrThrow(_response.body, {
                unrecognizedObjectKeys: "passthrough",
                allowUnrecognizedUnionMembers: true,
                allowUnrecognizedEnumValues: true,
                breadcrumbsPrefix: ["response"],
            });
        }

        if (_response.error.reason === "status-code") {
            throw new errors.SeedAuthEnvironmentVariablesError({
                statusCode: _response.error.statusCode,
                body: _response.error.body,
            });
        }

        switch (_response.error.reason) {
            case "non-json":
                throw new errors.SeedAuthEnvironmentVariablesError({
                    statusCode: _response.error.statusCode,
                    body: _response.error.rawBody,
                });
            case "timeout":
                throw new errors.SeedAuthEnvironmentVariablesTimeoutError();
            case "unknown":
                throw new errors.SeedAuthEnvironmentVariablesError({
                    message: _response.error.errorMessage,
                });
        }
    }

    protected async _getCustomAuthorizationHeaders() {
        const apiKeyValue = (await core.Supplier.get(this._options.apiKey)) ?? process?.env["FERN_API_KEY"];
        return { "X-FERN-API-KEY": apiKeyValue };
    }
}
