/**
 * This file was auto-generated by Fern from our API Definition.
 */

import * as serializers from "../../../index";
import * as FernDefinition from "../../../../api/index";
import * as core from "../../../../core";
import { ExampleTypeSchema } from "../../examples/types/ExampleTypeSchema";
import { EncodingSchema } from "../../encoding/types/EncodingSchema";
import { SourceSchema } from "../../source/types/SourceSchema";
import { WithDocs } from "../../commons/types/WithDocs";
import { WithAvailability } from "../../commons/types/WithAvailability";
import { WithDisplayName } from "../../commons/types/WithDisplayName";
import { WithAudiences } from "../../commons/types/WithAudiences";

export const BaseTypeDeclarationSchema: core.serialization.ObjectSchema<
    serializers.BaseTypeDeclarationSchema.Raw,
    FernDefinition.BaseTypeDeclarationSchema
> = core.serialization
    .object({
        examples: ExampleTypeSchema.optional(),
        encoding: EncodingSchema.optional(),
        source: SourceSchema.optional(),
    })
    .extend(WithDocs)
    .extend(WithAvailability)
    .extend(WithDisplayName)
    .extend(WithAudiences);

export declare namespace BaseTypeDeclarationSchema {
    interface Raw extends WithDocs.Raw, WithAvailability.Raw, WithDisplayName.Raw, WithAudiences.Raw {
        examples?: ExampleTypeSchema.Raw | null;
        encoding?: EncodingSchema.Raw | null;
        source?: SourceSchema.Raw | null;
    }
}
