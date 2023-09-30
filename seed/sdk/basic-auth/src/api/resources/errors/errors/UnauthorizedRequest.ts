/**
 * This file was auto-generated by Fern from our API Definition.
 */

import * as errors from "../../../../errors";
import * as SeedBasicAuth from "../../..";

export class UnauthorizedRequest extends errors.SeedBasicAuthError {
    constructor(body: SeedBasicAuth.UnauthorizedRequestErrorBody) {
        super({
            message: "UnauthorizedRequest",
            statusCode: 401,
            body: body,
        });
        Object.setPrototypeOf(this, UnauthorizedRequest.prototype);
    }
}
