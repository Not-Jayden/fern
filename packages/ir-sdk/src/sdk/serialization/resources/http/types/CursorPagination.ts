/**
 * This file was auto-generated by Fern from our API Definition.
 */

import * as serializers from "../../..";
import * as FernIr from "../../../../api";
import * as core from "../../../../core";

export const CursorPagination: core.serialization.ObjectSchema<
    serializers.CursorPagination.Raw,
    FernIr.CursorPagination
> = core.serialization.objectWithoutOptionalProperties({
    page: core.serialization.lazyObject(async () => (await import("../../..")).RequestProperty),
    next: core.serialization.lazyObject(async () => (await import("../../..")).ResponseProperty),
    results: core.serialization.lazyObject(async () => (await import("../../..")).ResponseProperty),
});

export declare namespace CursorPagination {
    interface Raw {
        page: serializers.RequestProperty.Raw;
        next: serializers.ResponseProperty.Raw;
        results: serializers.ResponseProperty.Raw;
    }
}
