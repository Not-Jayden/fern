/**
 * This file was auto-generated by Fern from our API Definition.
 */

import * as SeedTrace from "../../../../..";

export type AssertCorrectnessCheck =
    | SeedTrace.v2.AssertCorrectnessCheck.DeepEquality
    | SeedTrace.v2.AssertCorrectnessCheck.Custom;

export declare namespace AssertCorrectnessCheck {
    interface DeepEquality extends SeedTrace.v2.DeepEqualityCorrectnessCheck {
        type: "deepEquality";
    }

    interface Custom extends SeedTrace.v2.VoidFunctionDefinitionThatTakesActualResult {
        type: "custom";
    }
}
