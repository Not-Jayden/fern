/**
 * This file was auto-generated by Fern from our API Definition.
 */

import * as serializers from "../../../index";
import * as SeedTrace from "../../../../api/index";
import * as core from "../../../../core";
import { ExceptionInfo } from "./ExceptionInfo";

export const InternalError: core.serialization.ObjectSchema<serializers.InternalError.Raw, SeedTrace.InternalError> = core.serialization.object({
        "exceptionInfo": ExceptionInfo
    });

export declare namespace InternalError {
    interface Raw {
        "exceptionInfo": ExceptionInfo.Raw;
    }
}
