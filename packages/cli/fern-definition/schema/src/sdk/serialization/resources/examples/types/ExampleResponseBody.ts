/**
 * This file was auto-generated by Fern from our API Definition.
 */

import * as serializers from "../../../index";
import * as FernDefinition from "../../../../api/index";
import * as core from "../../../../core";
import { ExampleTypeReference } from "./ExampleTypeReference";

export const ExampleResponseBody: core.serialization.ObjectSchema<
    serializers.ExampleResponseBody.Raw,
    FernDefinition.ExampleResponseBody
> = core.serialization.object({
    error: core.serialization.string().optional(),
    body: ExampleTypeReference.optional(),
});

export declare namespace ExampleResponseBody {
    interface Raw {
        error?: string | null;
        body?: (ExampleTypeReference.Raw | undefined) | null;
    }
}
