import { FernFilepath } from "@fern-fern/ir-model/commons";
import { HttpEndpoint } from "@fern-fern/ir-model/http";
import { Reference } from "@fern-typescript/contexts";
import { ExportedFilePath } from "../exports-manager/ExportedFilePath";
import { AbstractDeclarationReferencer } from "./AbstractDeclarationReferencer";
import { AbstractServiceDeclarationReferencer } from "./AbstractServiceDeclarationReferencer";
import { DeclarationReferencer } from "./DeclarationReferencer";
import { ServiceDeclarationReferencer } from "./ServiceDeclarationReferencer";

export declare namespace EndpointDeclarationReferencer {
    export interface Init extends AbstractDeclarationReferencer.Init {
        serviceDeclarationReferencer: ServiceDeclarationReferencer;
    }

    export interface Name {
        service: FernFilepath;
        endpoint: HttpEndpoint;
    }
}
export class EndpointDeclarationReferencer extends AbstractServiceDeclarationReferencer<EndpointDeclarationReferencer.Name> {
    public getExportedFilepath(name: EndpointDeclarationReferencer.Name): ExportedFilePath {
        return {
            directories: this.getExportedDirectory(name.service),
            file: {
                nameOnDisk: this.getFilename(name),
                exportDeclaration: {
                    namespaceExport: this.getNamespaceExport(name),
                },
            },
        };
    }

    public getFilename(name: EndpointDeclarationReferencer.Name): string {
        return `${this.getNamespaceExport(name)}.ts`;
    }

    private getNamespaceExport({ endpoint }: EndpointDeclarationReferencer.Name): string {
        return endpoint.name.camelCase.unsafeName;
    }

    public getReferenceToEndpointExport(
        args: DeclarationReferencer.getReferenceTo.Options<EndpointDeclarationReferencer.Name>
    ): Reference {
        return this.getReferenceTo(this.getNamespaceExport(args.name), args);
    }

    protected override getExportedFilepathForReference(name: EndpointDeclarationReferencer.Name): ExportedFilePath {
        return {
            directories: this.getExportedDirectory(name.service),
            file: undefined,
        };
    }
}
