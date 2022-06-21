import { HttpService } from "@fern-api/api";
import {
    addFernServiceUtilsDependency,
    DependencyManager,
    ErrorResolver,
    getOrCreateDirectory,
    getOrCreateSourceFile,
    getTextOfTsKeyword,
    getTextOfTsNode,
    maybeAddDocs,
    TypeResolver,
} from "@fern-typescript/commons";
import { HelperManager } from "@fern-typescript/helper-manager";
import { ClassDeclaration, Directory, Scope, ts } from "ts-morph";
import { ClientConstants } from "../constants";
import { generateJoinPathsCall } from "../utils/generateJoinPathsCall";
import { addEndpointToService } from "./endpoints/addEndpointToService";

const SERVICE_INIT_TYPE = ts.factory.createTypeReferenceNode(
    ts.factory.createQualifiedName(
        ts.factory.createIdentifier(ClientConstants.HttpService.ServiceUtils.Imported.SERVICE_NAMESPACE),
        ts.factory.createIdentifier(ClientConstants.HttpService.ServiceUtils.ServiceInit.TYPE_NAME)
    )
);

export async function generateHttpService({
    servicesDirectory,
    modelDirectory,
    encodersDirectory,
    service,
    typeResolver,
    errorResolver,
    helperManager,
    dependencyManager,
}: {
    servicesDirectory: Directory;
    modelDirectory: Directory;
    encodersDirectory: Directory;
    service: HttpService;
    typeResolver: TypeResolver;
    errorResolver: ErrorResolver;
    helperManager: HelperManager;
    dependencyManager: DependencyManager;
}): Promise<void> {
    const serviceDirectory = getOrCreateDirectory(servicesDirectory, service.name.name, {
        exportOptions: {
            type: "namespace",
            namespace: service.name.name,
        },
    });
    await generateService({
        service,
        serviceDirectory,
        modelDirectory,
        encodersDirectory,
        servicesDirectory,
        typeResolver,
        errorResolver,
        helperManager,
        dependencyManager,
    });
}

async function generateService({
    service,
    serviceDirectory,
    modelDirectory,
    servicesDirectory,
    encodersDirectory,
    typeResolver,
    errorResolver,
    helperManager,
    dependencyManager,
}: {
    service: HttpService;
    serviceDirectory: Directory;
    modelDirectory: Directory;
    servicesDirectory: Directory;
    encodersDirectory: Directory;
    typeResolver: TypeResolver;
    errorResolver: ErrorResolver;
    helperManager: HelperManager;
    dependencyManager: DependencyManager;
}): Promise<void> {
    const serviceFile = getOrCreateSourceFile(serviceDirectory, `${service.name.name}.ts`);
    serviceFile.addImportDeclaration({
        namedImports: [
            ClientConstants.HttpService.ServiceUtils.Imported.FETCHER_TYPE_NAME,
            ClientConstants.HttpService.ServiceUtils.Imported.DEFAULT_FETCHER,
            ClientConstants.HttpService.ServiceUtils.Imported.SERVICE_NAMESPACE,
            ClientConstants.HttpService.ServiceUtils.Imported.IS_RESPONSE_OK_FUNCTION,
            ClientConstants.HttpService.ServiceUtils.Imported.TOKEN_TYPE_NAME,
        ],
        moduleSpecifier: "@fern-typescript/service-utils",
    });
    addFernServiceUtilsDependency(dependencyManager);

    const serviceInterface = serviceFile.addInterface({
        name: ClientConstants.HttpService.CLIENT_NAME,
        isExported: true,
    });

    const serviceClass = serviceFile.addClass({
        name: ClientConstants.HttpService.CLIENT_NAME,
        implements: [ClientConstants.HttpService.CLIENT_NAME],
        isExported: true,
    });
    maybeAddDocs(serviceClass, service.docs);

    serviceClass.addProperty({
        name: ClientConstants.HttpService.PrivateMembers.BASE_URL,
        scope: Scope.Private,
        type: getTextOfTsKeyword(ts.SyntaxKind.StringKeyword),
    });

    serviceClass.addProperty({
        name: ClientConstants.HttpService.PrivateMembers.FETCHER,
        scope: Scope.Private,
        type: getTextOfTsNode(
            ts.factory.createIdentifier(ClientConstants.HttpService.ServiceUtils.Imported.FETCHER_TYPE_NAME)
        ),
    });

    serviceClass.addProperty({
        name: ClientConstants.HttpService.PrivateMembers.TOKEN,
        scope: Scope.Private,
        type: getTextOfTsNode(
            ts.factory.createUnionTypeNode([
                ts.factory.createTypeReferenceNode(
                    ts.factory.createIdentifier(ClientConstants.HttpService.ServiceUtils.Imported.TOKEN_TYPE_NAME)
                ),
                ts.factory.createKeywordTypeNode(ts.SyntaxKind.UndefinedKeyword),
            ])
        ),
    });

    addConstructor({ serviceClass, serviceDefinition: service });

    const endpointsDirectory = getOrCreateDirectory(
        serviceDirectory,
        ClientConstants.HttpService.Files.ENDPOINTS_DIRECTORY_NAME
    );

    for (const endpoint of service.endpoints) {
        await addEndpointToService({
            endpoint,
            serviceInterface,
            serviceClass,
            serviceDefinition: service,
            modelDirectory,
            endpointsDirectory,
            servicesDirectory,
            encodersDirectory,
            typeResolver,
            errorResolver,
            helperManager,
            dependencyManager,
        });
    }
}
function addConstructor({
    serviceClass,
    serviceDefinition,
}: {
    serviceClass: ClassDeclaration;
    serviceDefinition: HttpService;
}) {
    const SERVICE_INIT_PARAMETER_NAME = "args";
    serviceClass.addConstructor({
        parameters: [
            {
                name: SERVICE_INIT_PARAMETER_NAME,
                type: getTextOfTsNode(SERVICE_INIT_TYPE),
            },
        ],
        statements: [
            getTextOfTsNode(
                createClassMemberAssignment({
                    member: ClientConstants.HttpService.PrivateMembers.FETCHER,
                    initialValue: ts.factory.createBinaryExpression(
                        ts.factory.createPropertyAccessExpression(
                            ts.factory.createIdentifier(SERVICE_INIT_PARAMETER_NAME),
                            ts.factory.createIdentifier(
                                ClientConstants.HttpService.ServiceUtils.ServiceInit.Properties.FETCHER
                            )
                        ),
                        ts.factory.createToken(ts.SyntaxKind.QuestionQuestionToken),
                        ts.factory.createIdentifier(ClientConstants.HttpService.ServiceUtils.Imported.DEFAULT_FETCHER)
                    ),
                })
            ),
            getTextOfTsNode(
                createClassMemberAssignment({
                    member: ClientConstants.HttpService.PrivateMembers.BASE_URL,
                    initialValue:
                        serviceDefinition.basePath != null
                            ? generateJoinPathsCall({
                                  file: serviceClass.getSourceFile(),
                                  paths: [
                                      ts.factory.createPropertyAccessExpression(
                                          ts.factory.createIdentifier(SERVICE_INIT_PARAMETER_NAME),
                                          ts.factory.createIdentifier(
                                              ClientConstants.HttpService.ServiceUtils.ServiceInit.Properties.ORIGIN
                                          )
                                      ),
                                      ts.factory.createStringLiteral(serviceDefinition.basePath),
                                  ],
                              })
                            : ts.factory.createPropertyAccessExpression(
                                  ts.factory.createIdentifier(SERVICE_INIT_PARAMETER_NAME),
                                  ts.factory.createIdentifier(
                                      ClientConstants.HttpService.ServiceUtils.ServiceInit.Properties.ORIGIN
                                  )
                              ),
                })
            ),
            getTextOfTsNode(
                createClassMemberAssignment({
                    member: ClientConstants.HttpService.PrivateMembers.TOKEN,
                    initialValue: ts.factory.createPropertyAccessExpression(
                        ts.factory.createIdentifier(SERVICE_INIT_PARAMETER_NAME),
                        ts.factory.createIdentifier(
                            ClientConstants.HttpService.ServiceUtils.ServiceInit.Properties.TOKEN
                        )
                    ),
                })
            ),
        ],
    });
}

function createClassMemberAssignment({
    member,
    initialValue,
}: {
    member: string;
    initialValue: ts.Expression;
}): ts.ExpressionStatement {
    return ts.factory.createExpressionStatement(
        ts.factory.createBinaryExpression(
            ts.factory.createPropertyAccessExpression(ts.factory.createThis(), ts.factory.createIdentifier(member)),
            ts.factory.createToken(ts.SyntaxKind.EqualsToken),
            initialValue
        )
    );
}
