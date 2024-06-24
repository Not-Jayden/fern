import { GeneratorName } from "@fern-api/configuration";
import { IrSerialization } from "../../ir-serialization";
import { IrVersions } from "../../ir-versions";
import { GeneratorWasNeverUpdatedToConsumeNewIR, IrMigration } from "../../types/IrMigration";

export const V49_TO_V48_MIGRATION: IrMigration<
    IrVersions.V49.ir.IntermediateRepresentation,
    IrVersions.V48.ir.IntermediateRepresentation
> = {
    laterVersion: "v49",
    earlierVersion: "v48",
    firstGeneratorVersionToConsumeNewIR: {
        [GeneratorName.TYPESCRIPT_NODE_SDK]: GeneratorWasNeverUpdatedToConsumeNewIR,
        [GeneratorName.TYPESCRIPT_BROWSER_SDK]: GeneratorWasNeverUpdatedToConsumeNewIR,
        [GeneratorName.TYPESCRIPT]: GeneratorWasNeverUpdatedToConsumeNewIR,
        [GeneratorName.TYPESCRIPT_SDK]: GeneratorWasNeverUpdatedToConsumeNewIR,
        [GeneratorName.TYPESCRIPT_EXPRESS]: GeneratorWasNeverUpdatedToConsumeNewIR,
        [GeneratorName.JAVA]: GeneratorWasNeverUpdatedToConsumeNewIR,
        [GeneratorName.JAVA_MODEL]: GeneratorWasNeverUpdatedToConsumeNewIR,
        [GeneratorName.JAVA_SDK]: GeneratorWasNeverUpdatedToConsumeNewIR,
        [GeneratorName.JAVA_SPRING]: GeneratorWasNeverUpdatedToConsumeNewIR,
        [GeneratorName.PYTHON_FASTAPI]: GeneratorWasNeverUpdatedToConsumeNewIR,
        [GeneratorName.PYTHON_PYDANTIC]: GeneratorWasNeverUpdatedToConsumeNewIR,
        [GeneratorName.OPENAPI_PYTHON_CLIENT]: GeneratorWasNeverUpdatedToConsumeNewIR,
        [GeneratorName.OPENAPI]: GeneratorWasNeverUpdatedToConsumeNewIR,
        [GeneratorName.STOPLIGHT]: GeneratorWasNeverUpdatedToConsumeNewIR,
        [GeneratorName.POSTMAN]: GeneratorWasNeverUpdatedToConsumeNewIR,
        [GeneratorName.PYTHON_SDK]: GeneratorWasNeverUpdatedToConsumeNewIR,
        [GeneratorName.GO_FIBER]: GeneratorWasNeverUpdatedToConsumeNewIR,
        [GeneratorName.GO_MODEL]: GeneratorWasNeverUpdatedToConsumeNewIR,
        [GeneratorName.GO_SDK]: GeneratorWasNeverUpdatedToConsumeNewIR,
        [GeneratorName.RUBY_MODEL]: GeneratorWasNeverUpdatedToConsumeNewIR,
        [GeneratorName.RUBY_SDK]: GeneratorWasNeverUpdatedToConsumeNewIR,
        [GeneratorName.CSHARP_MODEL]: GeneratorWasNeverUpdatedToConsumeNewIR,
        [GeneratorName.CSHARP_SDK]: GeneratorWasNeverUpdatedToConsumeNewIR
    },
    jsonifyEarlierVersion: (ir) =>
        IrSerialization.V48.IntermediateRepresentation.jsonOrThrow(ir, {
            unrecognizedObjectKeys: "strip",
            skipValidation: true
        }),
    migrateBackwards: (v49): IrVersions.V48.ir.IntermediateRepresentation => {
        const services = Object.fromEntries(
            Object.entries(v49.services).map(([serviceId, service]) => {
                const convertedEndpoints: IrVersions.V48.HttpEndpoint[] = service.endpoints.flatMap((endpoint) => {
                    let response: IrVersions.V48.HttpResponseBody | undefined = undefined;
                    if (endpoint.response?.body != null) {
                        response = endpoint.response.body?._visit<IrVersions.V48.HttpResponseBody | undefined>({
                            json: (val) => IrVersions.V48.HttpResponseBody.json({ ...val }),
                            fileDownload: (val) => IrVersions.V48.HttpResponseBody.fileDownload({ ...val }),
                            text: (val) => IrVersions.V48.HttpResponseBody.text({ ...val }),
                            streaming: (val) => IrVersions.V48.HttpResponseBody.streaming({ ...val }),
                            streamParameter: () => undefined,
                            _other: () => undefined
                        });
                        if (response == null) {
                            return [];
                        }
                    }
                    const convertedEndpoint: IrVersions.V48.HttpEndpoint = {
                        ...endpoint,
                        response:
                            response == null
                                ? undefined
                                : {
                                      statusCode: endpoint.response?.statusCode,
                                      body: response
                                  }
                    };
                    return convertedEndpoint;
                });
                return [
                    serviceId,
                    {
                        ...service,
                        endpoints: convertedEndpoints
                    }
                ];
            })
        );
        return {
            ...v49,
            services
        };
    }
};
