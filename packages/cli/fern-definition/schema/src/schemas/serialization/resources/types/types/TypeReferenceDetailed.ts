/**
 * This file was auto-generated by Fern from our API Definition.
 */

import * as serializers from "../../../index";
import * as FernDefinition from "../../../../api/index";
import * as core from "../../../../core";
import { BaseTypeReferenceSchema } from "./BaseTypeReferenceSchema";

export const TypeReferenceDetailed: core.serialization.ObjectSchema<
    serializers.TypeReferenceDetailed.Raw,
    FernDefinition.TypeReferenceDetailed
> = core.serialization.object({}).extend(BaseTypeReferenceSchema);

export declare namespace TypeReferenceDetailed {
    interface Raw extends BaseTypeReferenceSchema.Raw {}
}
