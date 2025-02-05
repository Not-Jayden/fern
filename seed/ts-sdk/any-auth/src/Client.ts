/**
 * This file was auto-generated by Fern from our API Definition.
 */

import * as core from "./core";
import { Auth } from "./api/resources/auth/client/Client";
import { User } from "./api/resources/user/client/Client";

export declare namespace SeedAnyAuthClient {
    interface Options {
        environment: core.Supplier<string>;
        clientId?: core.Supplier<string>;
        clientSecret?: core.Supplier<string>;
        token?: core.Supplier<core.BearerToken | undefined>;
        apiKey?: core.Supplier<string | undefined>;
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

export class SeedAnyAuthClient {
    private readonly _oauthTokenProvider: core.OAuthTokenProvider;

    constructor(protected readonly _options: SeedAnyAuthClient.Options) {
        const clientId = this._options.clientId ?? process.env["MY_CLIENT_ID"];
        if (clientId == null) {
            throw new Error(
                "clientId is required; either pass it as an argument or set the MY_CLIENT_ID environment variable"
            );
        }

        const clientSecret = this._options.clientSecret ?? process.env["MY_CLIENT_SECRET"];
        if (clientSecret == null) {
            throw new Error(
                "clientSecret is required; either pass it as an argument or set the MY_CLIENT_SECRET environment variable"
            );
        }

        this._oauthTokenProvider = new core.OAuthTokenProvider({
            clientId,
            clientSecret,
            authClient: new Auth({
                environment: this._options.environment,
            }),
        });
    }

    protected _auth: Auth | undefined;

    public get auth(): Auth {
        return (this._auth ??= new Auth({
            ...this._options,
            token: async () => await this._oauthTokenProvider.getToken(),
        }));
    }

    protected _user: User | undefined;

    public get user(): User {
        return (this._user ??= new User({
            ...this._options,
            token: async () => await this._oauthTokenProvider.getToken(),
        }));
    }
}
