/**
 * This file was auto-generated by Fern from our API Definition.
 */

import * as serializers from "../../../index";
import * as FernDefinition from "../../../../api/index";
import * as core from "../../../../core";
import { OAuthRefreshTokenRequestProperties } from "./OAuthRefreshTokenRequestProperties";
import { OAuthRefreshTokenResponseProperties } from "./OAuthRefreshTokenResponseProperties";

export const OAuthRefreshTokenEndpoint: core.serialization.ObjectSchema<
    serializers.OAuthRefreshTokenEndpoint.Raw,
    FernDefinition.OAuthRefreshTokenEndpoint
> = core.serialization.object({
    endpoint: core.serialization.string(),
    requestProperties: core.serialization.property("request-properties", OAuthRefreshTokenRequestProperties.optional()),
    responseProperties: core.serialization.property(
        "response-properties",
        OAuthRefreshTokenResponseProperties.optional()
    ),
});

export declare namespace OAuthRefreshTokenEndpoint {
    interface Raw {
        endpoint: string;
        "request-properties"?: OAuthRefreshTokenRequestProperties.Raw | null;
        "response-properties"?: OAuthRefreshTokenResponseProperties.Raw | null;
    }
}
