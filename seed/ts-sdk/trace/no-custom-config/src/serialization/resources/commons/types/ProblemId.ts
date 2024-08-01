/**
 * This file was auto-generated by Fern from our API Definition.
 */

import * as serializers from "../../../index";
import * as SeedTrace from "../../../../api/index";
import * as core from "../../../../core";

export const ProblemId: core.serialization.Schema<serializers.ProblemId.Raw, SeedTrace.ProblemId> = core.serialization.string();

export declare namespace ProblemId {
    type Raw = string;
}
