/**
 * This file was auto-generated by Fern from our API Definition.
 */

import * as serializers from "../../..";
import * as SeedTrace from "../../../../api";
import * as core from "../../../../core";

export const ReservedKeywordEnum: core.serialization.Schema<
    serializers.ReservedKeywordEnum.Raw,
    SeedTrace.ReservedKeywordEnum
> = core.serialization.enum_(["is", "as"]);

export declare namespace ReservedKeywordEnum {
    type Raw = "is" | "as";
}
