/**
 * This file was auto-generated by Fern from our API Definition.
 */

import * as serializers from "../../../index";
import * as FernDefinition from "../../../../api/index";
import * as core from "../../../../core";
import { NonInlinedTypeReferenceSchema } from "./NonInlinedTypeReferenceSchema";

export const InlineableTypeReferenceSchemaDetailed: core.serialization.Schema<
    serializers.InlineableTypeReferenceSchemaDetailed.Raw,
    FernDefinition.InlineableTypeReferenceSchemaDetailed
> = core.serialization.undiscriminatedUnion([
    NonInlinedTypeReferenceSchema,
    core.serialization.lazyObject(() => serializers.InlinedTypeReferenceSchema),
    core.serialization.lazyObject(() => serializers.InlinedListTypeReferenceSchema),
]);

export declare namespace InlineableTypeReferenceSchemaDetailed {
    type Raw =
        | NonInlinedTypeReferenceSchema.Raw
        | serializers.InlinedTypeReferenceSchema.Raw
        | serializers.InlinedListTypeReferenceSchema.Raw;
}
