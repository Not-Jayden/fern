/**
 * This file was auto-generated by Fern from our API Definition.
 */

import * as serializers from "../../../index";
import * as FernDefinition from "../../../../api/index";
import * as core from "../../../../core";
import { VersionHeaderDetailed } from "./VersionHeaderDetailed";

export const VersionHeader: core.serialization.Schema<serializers.VersionHeader.Raw, FernDefinition.VersionHeader> =
    core.serialization.undiscriminatedUnion([core.serialization.string(), VersionHeaderDetailed]);

export declare namespace VersionHeader {
    type Raw = string | VersionHeaderDetailed.Raw;
}
