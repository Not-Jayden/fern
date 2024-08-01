/**
 * This file was auto-generated by Fern from our API Definition.
 */

import * as serializers from "../../../index";
import * as SeedTrace from "../../../../api/index";
import * as core from "../../../../core";

export const PlaylistId: core.serialization.Schema<serializers.PlaylistId.Raw, SeedTrace.PlaylistId> = core.serialization.string();

export declare namespace PlaylistId {
    type Raw = string;
}
