/**
 * This file was auto-generated by Fern from our API Definition.
 */

import * as serializers from "../../../index";
import * as FernOpenapiIr from "../../../../api/index";
import * as core from "../../../../core";

export const WithContentType: core.serialization.ObjectSchema<
    serializers.WithContentType.Raw,
    FernOpenapiIr.WithContentType
> = core.serialization.objectWithoutOptionalProperties({
    contentType: core.serialization.string().optional(),
});

export declare namespace WithContentType {
    interface Raw {
        contentType?: string | null;
    }
}
