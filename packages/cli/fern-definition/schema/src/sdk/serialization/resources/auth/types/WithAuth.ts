/**
 * This file was auto-generated by Fern from our API Definition.
 */

import * as serializers from "../../../index";
import * as FernDefinition from "../../../../api/index";
import * as core from "../../../../core";
import { Auth } from "./Auth";
import { AuthScheme } from "./AuthScheme";

export const WithAuth: core.serialization.ObjectSchema<serializers.WithAuth.Raw, FernDefinition.WithAuth> =
    core.serialization.object({
        auth: Auth,
        authSchemes: core.serialization.property(
            "auth-schemes",
            core.serialization.record(core.serialization.string(), AuthScheme).optional()
        ),
    });

export declare namespace WithAuth {
    interface Raw {
        auth: Auth.Raw;
        "auth-schemes"?: Record<string, AuthScheme.Raw> | null;
    }
}
