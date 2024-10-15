/**
 * This file was auto-generated by Fern from our API Definition.
 */

import * as serializers from "../../../index";
import * as FernDefinition from "../../../../api/index";
import * as core from "../../../../core";
import { TypeReferenceDetailed } from "../../types/types/TypeReferenceDetailed";

export const HttpInlineRequestBodyProperty: core.serialization.ObjectSchema<
    serializers.HttpInlineRequestBodyProperty.Raw,
    FernDefinition.HttpInlineRequestBodyProperty
> = core.serialization
    .object({
        contentType: core.serialization.property("content-type", core.serialization.string().optional()),
    })
    .extend(TypeReferenceDetailed);

export declare namespace HttpInlineRequestBodyProperty {
    interface Raw extends TypeReferenceDetailed.Raw {
        "content-type"?: string | null;
    }
}
