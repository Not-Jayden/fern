import { AbsoluteFilePath } from "@fern-api/fs-utils";
import { BaseTypescriptCustomConfigSchema, BaseTypescriptGeneratorContext } from "@fern-api/typescript-codegen";
import { HttpService, IntermediateRepresentation } from "@fern-fern/ir-sdk/api";
import {
    convertExportedFilePathToFilePath,
    CoreUtilitiesManager,
    DependencyManager,
    ExportedDirectory,
    ExportedFilePath,
    ExportsManager,
    ImportsManager,
    NpmPackage,
    PackageId,
    SimpleTypescriptProject,
    TypescriptProject
} from "@fern-typescript/commons";
import { GeneratorContext } from "@fern-typescript/contexts";
import { ExpressEndpointTypeSchemasGenerator } from "@fern-typescript/express-endpoint-type-schemas-generator";
import { ExpressErrorGenerator } from "@fern-typescript/express-error-generator";
import { ExpressErrorSchemaGenerator } from "@fern-typescript/express-error-schema-generator";
import { ExpressInlinedRequestBodyGenerator } from "@fern-typescript/express-inlined-request-body-generator";
import { ExpressInlinedRequestBodySchemaGenerator } from "@fern-typescript/express-inlined-request-schema-generator";
import { ExpressRegisterGenerator } from "@fern-typescript/express-register-generator";
import { ExpressServiceGenerator } from "@fern-typescript/express-service-generator";
import { GenericAPIExpressErrorGenerator } from "@fern-typescript/generic-express-error-generators";
import { ErrorResolver, PackageResolver, TypeResolver } from "@fern-typescript/resolvers";
import { TypeGenerator } from "@fern-typescript/type-generator";
import { TypeReferenceExampleGenerator } from "@fern-typescript/type-reference-example-generator";
import { TypeSchemaGenerator } from "@fern-typescript/type-schema-generator";
import { Directory, Project, SourceFile } from "ts-morph";
import { ExpressContextImpl } from "./contexts/ExpressContextImpl";
import { EndpointDeclarationReferencer } from "./declaration-referencers/EndpointDeclarationReferencer";
import { ExpressErrorDeclarationReferencer } from "./declaration-referencers/ExpressErrorDeclarationReferencer";
import { ExpressInlinedRequestBodyDeclarationReferencer } from "./declaration-referencers/ExpressInlinedRequestBodyDeclarationReferencer";
import { ExpressRegisterDeclarationReferencer } from "./declaration-referencers/ExpressRegisterDeclarationReferencer";
import { ExpressServiceDeclarationReferencer } from "./declaration-referencers/ExpressServiceDeclarationReferencer";
import { GenericAPIExpressErrorDeclarationReferencer } from "./declaration-referencers/GenericAPIExpressErrorDeclarationReferencer";
import { TypeDeclarationReferencer } from "./declaration-referencers/TypeDeclarationReferencer";
import { FernGeneratorExec } from "@fern-fern/generator-exec-sdk";
import { GeneratorNotificationService } from "@fern-api/generator-commons";

const FILE_HEADER = `/**
 * This file was auto-generated by Fern from our API Definition.
 */
`;

export declare namespace ExpressGenerator {
    export interface Init {
        generatorExecConfig: FernGeneratorExec.GeneratorConfig;
        namespaceExport: string;
        intermediateRepresentation: IntermediateRepresentation;
        context: GeneratorContext;
        npmPackage: NpmPackage;
        config: Config;
    }

    export interface Config {
        shouldUseBrandedStringAliases: boolean;
        areImplementationsOptional: boolean;
        doNotHandleUnrecognizedErrors: boolean;
        includeUtilsOnUnionMembers: boolean;
        includeOtherInUnionTypes: boolean;
        treatUnknownAsAny: boolean;
        includeSerdeLayer: boolean;
        outputEsm: boolean;
        retainOriginalCasing: boolean;
        allowExtraFields: boolean;
        skipRequestValidation: boolean;
        skipResponseValidation: boolean;
        requestValidationStatusCode: number;
        useBigInt: boolean;
        noOptionalProperties: boolean;
    }
}

export class ExpressGenerator {
    private context: GeneratorContext;
    private intermediateRepresentation: IntermediateRepresentation;
    private npmPackage: NpmPackage;
    private config: ExpressGenerator.Config;

    private project: Project;
    private rootDirectory: Directory;
    private exportsManager: ExportsManager;
    private dependencyManager = new DependencyManager();
    private coreUtilitiesManager: CoreUtilitiesManager;
    private typeResolver: TypeResolver;
    private errorResolver: ErrorResolver;
    private packageResolver: PackageResolver;

    private typeDeclarationReferencer: TypeDeclarationReferencer;
    private typeSchemaDeclarationReferencer: TypeDeclarationReferencer;
    private expressInlinedRequestBodyDeclarationReferencer: ExpressInlinedRequestBodyDeclarationReferencer;
    private expressInlinedRequestBodySchemaDeclarationReferencer: ExpressInlinedRequestBodyDeclarationReferencer;
    private expressEndpointSchemaDeclarationReferencer: EndpointDeclarationReferencer;
    private expressServiceDeclarationReferencer: ExpressServiceDeclarationReferencer;
    private expressRegisterDeclarationReferencer: ExpressRegisterDeclarationReferencer;
    private genericApiExpressErrorDeclarationReferencer: GenericAPIExpressErrorDeclarationReferencer;
    private expressErrorDeclarationReferencer: ExpressErrorDeclarationReferencer;
    private expressErrorSchemaDeclarationReferencer: ExpressErrorDeclarationReferencer;

    private typeGenerator: TypeGenerator;
    private typeSchemaGenerator: TypeSchemaGenerator;
    private typeReferenceExampleGenerator: TypeReferenceExampleGenerator;
    private expressInlinedRequestBodyGenerator: ExpressInlinedRequestBodyGenerator;
    private expressInlinedRequestBodySchemaGenerator: ExpressInlinedRequestBodySchemaGenerator;
    private expressEndpointTypeSchemasGenerator: ExpressEndpointTypeSchemasGenerator;
    private expressServiceGenerator: ExpressServiceGenerator;
    private expressRegisterGenerator: ExpressRegisterGenerator;
    private genericApiExpressErrorGenerator: GenericAPIExpressErrorGenerator;
    private expressErrorGenerator: ExpressErrorGenerator;
    private expressErrorSchemaGenerator: ExpressErrorSchemaGenerator;
    private v2Context: BaseTypescriptGeneratorContext<BaseTypescriptCustomConfigSchema>;

    constructor({
        generatorExecConfig,
        namespaceExport,
        intermediateRepresentation,
        context,
        npmPackage,
        config
    }: ExpressGenerator.Init) {
        this.context = context;
        this.intermediateRepresentation = intermediateRepresentation;
        this.npmPackage = npmPackage;
        this.config = config;

        this.v2Context = new BaseTypescriptGeneratorContext(
            this.intermediateRepresentation,
            generatorExecConfig,
            { namespaceExport, noSerdeLayer: !this.config.includeSerdeLayer },
            new GeneratorNotificationService(generatorExecConfig.environment)
        );

        this.exportsManager = new ExportsManager();
        this.coreUtilitiesManager = new CoreUtilitiesManager();

        this.project = new Project({
            useInMemoryFileSystem: true
        });
        this.rootDirectory = this.project.createDirectory("/");
        this.typeResolver = new TypeResolver(intermediateRepresentation);
        this.errorResolver = new ErrorResolver(intermediateRepresentation);
        this.packageResolver = new PackageResolver(intermediateRepresentation);

        const apiDirectory: ExportedDirectory[] = [
            {
                nameOnDisk: "api",
                exportDeclaration: { namespaceExport }
            }
        ];

        const schemaDirectory: ExportedDirectory[] = [
            {
                nameOnDisk: "serialization"
            }
        ];

        this.typeDeclarationReferencer = new TypeDeclarationReferencer({
            containingDirectory: apiDirectory,
            namespaceExport
        });
        this.typeSchemaDeclarationReferencer = new TypeDeclarationReferencer({
            containingDirectory: schemaDirectory,
            namespaceExport
        });
        this.expressInlinedRequestBodyDeclarationReferencer = new ExpressInlinedRequestBodyDeclarationReferencer({
            packageResolver: this.packageResolver,
            containingDirectory: apiDirectory,
            namespaceExport
        });
        this.expressInlinedRequestBodySchemaDeclarationReferencer = new ExpressInlinedRequestBodyDeclarationReferencer({
            packageResolver: this.packageResolver,
            containingDirectory: schemaDirectory,
            namespaceExport
        });
        this.expressEndpointSchemaDeclarationReferencer = new EndpointDeclarationReferencer({
            packageResolver: this.packageResolver,
            containingDirectory: schemaDirectory,
            namespaceExport
        });
        this.expressServiceDeclarationReferencer = new ExpressServiceDeclarationReferencer({
            packageResolver: this.packageResolver,
            containingDirectory: apiDirectory,
            namespaceExport
        });
        this.expressRegisterDeclarationReferencer = new ExpressRegisterDeclarationReferencer({
            containingDirectory: [],
            namespaceExport
        });
        this.genericApiExpressErrorDeclarationReferencer = new GenericAPIExpressErrorDeclarationReferencer({
            containingDirectory: [],
            namespaceExport
        });
        this.expressErrorDeclarationReferencer = new ExpressErrorDeclarationReferencer({
            containingDirectory: apiDirectory,
            namespaceExport
        });
        this.expressErrorSchemaDeclarationReferencer = new ExpressErrorDeclarationReferencer({
            containingDirectory: schemaDirectory,
            namespaceExport
        });

        this.typeGenerator = new TypeGenerator({
            useBrandedStringAliases: config.shouldUseBrandedStringAliases,
            includeUtilsOnUnionMembers: config.includeUtilsOnUnionMembers,
            includeOtherInUnionTypes: config.includeOtherInUnionTypes,
            includeSerdeLayer: config.includeSerdeLayer,
            retainOriginalCasing: config.retainOriginalCasing,
            noOptionalProperties: config.noOptionalProperties,
            respectInlinedTypes: true
        });
        this.typeSchemaGenerator = new TypeSchemaGenerator({
            includeUtilsOnUnionMembers: config.includeUtilsOnUnionMembers,
            noOptionalProperties: config.noOptionalProperties
        });
        this.typeReferenceExampleGenerator = new TypeReferenceExampleGenerator();
        this.expressInlinedRequestBodyGenerator = new ExpressInlinedRequestBodyGenerator();
        this.expressInlinedRequestBodySchemaGenerator = new ExpressInlinedRequestBodySchemaGenerator({
            includeSerdeLayer: config.includeSerdeLayer,
            skipRequestValidation: config.skipRequestValidation
        });
        this.expressEndpointTypeSchemasGenerator = new ExpressEndpointTypeSchemasGenerator({
            includeSerdeLayer: config.includeSerdeLayer,
            allowExtraFields: config.allowExtraFields,
            skipRequestValidation: config.skipRequestValidation,
            skipResponseValidation: config.skipResponseValidation
        });
        this.expressServiceGenerator = new ExpressServiceGenerator({
            packageResolver: this.packageResolver,
            doNotHandleUnrecognizedErrors: config.doNotHandleUnrecognizedErrors,
            includeSerdeLayer: config.includeSerdeLayer,
            skipRequestValidation: config.skipRequestValidation,
            skipResponseValidation: config.skipResponseValidation,
            requestValidationStatusCode: config.requestValidationStatusCode
        });
        this.expressRegisterGenerator = new ExpressRegisterGenerator({
            packageResolver: this.packageResolver,
            intermediateRepresentation: this.intermediateRepresentation,
            registerFunctionName: this.expressRegisterDeclarationReferencer.getRegisterFunctionName(),
            areImplementationsOptional: config.areImplementationsOptional
        });
        this.genericApiExpressErrorGenerator = new GenericAPIExpressErrorGenerator();
        this.expressErrorGenerator = new ExpressErrorGenerator();
        this.expressErrorSchemaGenerator = new ExpressErrorSchemaGenerator({
            includeSerdeLayer: config.includeSerdeLayer,
            allowExtraFields: config.allowExtraFields
        });
    }

    public async generate(): Promise<TypescriptProject> {
        this.generateTypeDeclarations();
        this.generateInlinedRequestBodies();
        this.generateExpressServices();
        this.generateExpressRegister();
        this.generateGenericApiExpressGenerator();
        this.generateErrorDeclarations();

        if (this.config.includeSerdeLayer) {
            this.generateTypeSchemas();
            this.generateInlinedRequestBodySchemas();
            this.generateEndpointTypeSchemas();
            this.generateErrorSchemas();
        }

        this.coreUtilitiesManager.finalize(this.exportsManager, this.dependencyManager);
        this.exportsManager.writeExportsToProject(this.rootDirectory);

        return new SimpleTypescriptProject({
            runScripts: true,
            npmPackage: this.npmPackage,
            dependencies: this.dependencyManager.getDependencies(),
            tsMorphProject: this.project,
            outputEsm: this.config.outputEsm,
            extraDependencies: {},
            extraDevDependencies: {},
            extraFiles: {},
            extraScripts: {},
            extraPeerDependencies: {},
            extraPeerDependenciesMeta: {},
            resolutions: {
                "@types/mime": "3.0.4"
            },
            extraConfigs: undefined,
            outputJsr: false
        });
    }

    public async copyCoreUtilities({
        pathToSrc,
        pathToRoot
    }: {
        pathToSrc: AbsoluteFilePath;
        pathToRoot: AbsoluteFilePath;
    }): Promise<void> {
        await this.coreUtilitiesManager.copyCoreUtilities({ pathToSrc, pathToRoot });
    }

    private generateTypeDeclarations() {
        for (const typeDeclaration of Object.values(this.intermediateRepresentation.types)) {
            this.withSourceFile({
                filepath: this.typeDeclarationReferencer.getExportedFilepath(typeDeclaration.name),
                run: ({ sourceFile, importsManager }) => {
                    const context = this.generateExpressContext({ sourceFile, importsManager });
                    context.type.getGeneratedType(typeDeclaration.name).writeToFile(context);
                }
            });
        }
    }

    private generateTypeSchemas() {
        for (const typeDeclaration of Object.values(this.intermediateRepresentation.types)) {
            this.withSourceFile({
                filepath: this.typeSchemaDeclarationReferencer.getExportedFilepath(typeDeclaration.name),
                run: ({ sourceFile, importsManager }) => {
                    const context = this.generateExpressContext({ sourceFile, importsManager });
                    context.typeSchema.getGeneratedTypeSchema(typeDeclaration.name).writeToFile(context);
                }
            });
        }
    }

    private generateErrorDeclarations() {
        for (const errorDeclaration of Object.values(this.intermediateRepresentation.errors)) {
            this.withSourceFile({
                filepath: this.expressErrorDeclarationReferencer.getExportedFilepath(errorDeclaration.name),
                run: ({ sourceFile, importsManager }) => {
                    const context = this.generateExpressContext({ sourceFile, importsManager });
                    context.expressError.getGeneratedExpressError(errorDeclaration.name).writeToFile(context);
                }
            });
        }
    }

    private generateErrorSchemas() {
        for (const errorDeclaration of Object.values(this.intermediateRepresentation.errors)) {
            this.withSourceFile({
                filepath: this.expressErrorSchemaDeclarationReferencer.getExportedFilepath(errorDeclaration.name),
                run: ({ sourceFile, importsManager }) => {
                    const context = this.generateExpressContext({ sourceFile, importsManager });
                    context.expressErrorSchema
                        .getGeneratedExpressErrorSchema(errorDeclaration.name)
                        ?.writeToFile(context);
                }
            });
        }
    }

    private generateInlinedRequestBodies() {
        this.forEachService((service, packageId) => {
            for (const endpoint of service.endpoints) {
                if (endpoint.requestBody?.type === "inlinedRequestBody") {
                    this.withSourceFile({
                        filepath: this.expressInlinedRequestBodyDeclarationReferencer.getExportedFilepath({
                            packageId,
                            endpoint
                        }),
                        run: ({ sourceFile, importsManager }) => {
                            const context = this.generateExpressContext({ sourceFile, importsManager });
                            context.expressInlinedRequestBody
                                .getGeneratedInlinedRequestBody(packageId, endpoint.name)
                                .writeToFile(context);
                        }
                    });
                }
            }
        });
    }

    private generateInlinedRequestBodySchemas() {
        this.forEachService((service, packageId) => {
            for (const endpoint of service.endpoints) {
                if (endpoint.requestBody?.type === "inlinedRequestBody") {
                    this.withSourceFile({
                        filepath: this.expressInlinedRequestBodySchemaDeclarationReferencer.getExportedFilepath({
                            packageId,
                            endpoint
                        }),
                        run: ({ sourceFile, importsManager }) => {
                            const context = this.generateExpressContext({ sourceFile, importsManager });
                            context.expressInlinedRequestBodySchema
                                .getGeneratedInlinedRequestBodySchema(packageId, endpoint.name)
                                .writeToFile(context);
                        }
                    });
                }
            }
        });
    }

    private generateEndpointTypeSchemas() {
        this.forEachService((service, packageId) => {
            for (const endpoint of service.endpoints) {
                this.withSourceFile({
                    filepath: this.expressEndpointSchemaDeclarationReferencer.getExportedFilepath({
                        packageId,
                        endpoint
                    }),
                    run: ({ sourceFile, importsManager }) => {
                        const context = this.generateExpressContext({ sourceFile, importsManager });
                        context.expressEndpointTypeSchemas
                            .getGeneratedEndpointTypeSchemas(packageId, endpoint.name)
                            .writeToFile(context);
                    }
                });
            }
        });
    }

    private generateExpressServices() {
        this.forEachService((_service, packageId) => {
            this.withSourceFile({
                filepath: this.expressServiceDeclarationReferencer.getExportedFilepath(packageId),
                run: ({ sourceFile, importsManager }) => {
                    const context = this.generateExpressContext({ sourceFile, importsManager });
                    context.expressService.getGeneratedExpressService(packageId).writeToFile(context);
                }
            });
        });
    }

    private generateExpressRegister() {
        this.withSourceFile({
            filepath: this.expressRegisterDeclarationReferencer.getExportedFilepath(),
            run: ({ sourceFile, importsManager }) => {
                const context = this.generateExpressContext({ sourceFile, importsManager });
                context.expressRegister.getGeneratedExpressRegister()?.writeToFile(context);
            }
        });
    }

    private generateGenericApiExpressGenerator() {
        this.withSourceFile({
            filepath: this.genericApiExpressErrorDeclarationReferencer.getExportedFilepath(),
            run: ({ sourceFile, importsManager }) => {
                const context = this.generateExpressContext({ sourceFile, importsManager });
                context.genericAPIExpressError.getGeneratedGenericAPIExpressError().writeToFile(context);
            }
        });
    }

    private withSourceFile({
        run,
        filepath
    }: {
        run: (args: { sourceFile: SourceFile; importsManager: ImportsManager }) => void;
        filepath: ExportedFilePath;
    }) {
        const filepathStr = convertExportedFilePathToFilePath(filepath);
        this.context.logger.debug(`Generating ${filepathStr}`);

        const sourceFile = this.rootDirectory.createSourceFile(filepathStr);
        const importsManager = new ImportsManager();

        run({ sourceFile, importsManager });

        if (sourceFile.getStatements().length === 0) {
            sourceFile.delete();
            this.context.logger.debug(`Skipping ${filepathStr} (no content)`);
        } else {
            importsManager.writeImportsToSourceFile(sourceFile);
            this.exportsManager.addExportsForFilepath(filepath);

            // this needs to be last.
            // https://github.com/dsherret/ts-morph/issues/189#issuecomment-414174283
            sourceFile.insertText(0, (writer) => {
                writer.writeLine(FILE_HEADER);
            });

            this.context.logger.debug(`Generated ${filepathStr}`);
        }
    }

    private getAllPackageIds(): PackageId[] {
        return [
            { isRoot: true },
            ...Object.keys(this.intermediateRepresentation.subpackages).map(
                (subpackageId): PackageId => ({ isRoot: false, subpackageId })
            )
        ];
    }

    private forEachService(run: (service: HttpService, packageId: PackageId) => void): void {
        for (const packageId of this.getAllPackageIds()) {
            const service = this.packageResolver.getServiceDeclaration(packageId);
            if (service != null) {
                run(service, packageId);
            }
        }
    }

    private generateExpressContext({
        sourceFile,
        importsManager
    }: {
        sourceFile: SourceFile;
        importsManager: ImportsManager;
    }): ExpressContextImpl {
        return new ExpressContextImpl({
            v2Context: this.v2Context,
            sourceFile,
            coreUtilitiesManager: this.coreUtilitiesManager,
            dependencyManager: this.dependencyManager,
            fernConstants: this.intermediateRepresentation.constants,
            importsManager,
            typeResolver: this.typeResolver,
            typeDeclarationReferencer: this.typeDeclarationReferencer,
            typeSchemaDeclarationReferencer: this.typeSchemaDeclarationReferencer,
            typeReferenceExampleGenerator: this.typeReferenceExampleGenerator,
            expressEndpointSchemaDeclarationReferencer: this.expressEndpointSchemaDeclarationReferencer,
            typeGenerator: this.typeGenerator,
            packageResolver: this.packageResolver,
            expressEndpointTypeSchemasGenerator: this.expressEndpointTypeSchemasGenerator,
            typeSchemaGenerator: this.typeSchemaGenerator,
            expressInlinedRequestBodyDeclarationReferencer: this.expressInlinedRequestBodyDeclarationReferencer,
            expressInlinedRequestBodyGenerator: this.expressInlinedRequestBodyGenerator,
            expressInlinedRequestBodySchemaDeclarationReferencer:
                this.expressInlinedRequestBodySchemaDeclarationReferencer,
            expressInlinedRequestBodySchemaGenerator: this.expressInlinedRequestBodySchemaGenerator,
            expressServiceDeclarationReferencer: this.expressServiceDeclarationReferencer,
            expressServiceGenerator: this.expressServiceGenerator,
            expressErrorGenerator: this.expressErrorGenerator,
            errorDeclarationReferencer: this.expressErrorDeclarationReferencer,
            errorResolver: this.errorResolver,
            genericAPIExpressErrorDeclarationReferencer: this.genericApiExpressErrorDeclarationReferencer,
            genericAPIExpressErrorGenerator: this.genericApiExpressErrorGenerator,
            treatUnknownAsAny: this.config.treatUnknownAsAny,
            expressRegisterGenerator: this.expressRegisterGenerator,
            expressErrorSchemaDeclarationReferencer: this.expressErrorSchemaDeclarationReferencer,
            expressErrorSchemaGenerator: this.expressErrorSchemaGenerator,
            includeSerdeLayer: this.config.includeSerdeLayer,
            retainOriginalCasing: this.config.retainOriginalCasing,
            useBigInt: this.config.useBigInt
        });
    }
}
