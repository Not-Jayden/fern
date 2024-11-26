/**
 * This file was auto-generated by Fern from our API Definition.
 */

import * as FernDocsConfig from "../../..";

export interface VersionConfig extends FernDocsConfig.WithPermissions {
    displayName: string;
    /** The relative path to the version's docs.yml file. */
    path: string;
    /** The "slug" is this version's basePath. If not set, the slug will be generated from the display-name. */
    slug?: string;
    /** If `availability` is set to `deprecated`, Fern will display a warning banner on the docs site. */
    availability?: FernDocsConfig.VersionAvailability;
}
