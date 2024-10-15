/**
 * This file was auto-generated by Fern from our API Definition.
 */

import * as serializers from "../../../index";
import * as FernDefinition from "../../../../api/index";
import * as core from "../../../../core";
import { ExampleType } from "../../examples/types/ExampleType";
import { WithDocs } from "../../commons/types/WithDocs";

export const ErrorDeclaration: core.serialization.ObjectSchema<
    serializers.ErrorDeclaration.Raw,
    FernDefinition.ErrorDeclaration
> = core.serialization
    .object({
        statusCode: core.serialization.property("status-code", core.serialization.number()),
        type: core.serialization.string().optional(),
        examples: core.serialization.list(ExampleType).optional(),
    })
    .extend(WithDocs);

export declare namespace ErrorDeclaration {
    interface Raw extends WithDocs.Raw {
        "status-code": number;
        type?: string | null;
        examples?: ExampleType.Raw[] | null;
    }
}
