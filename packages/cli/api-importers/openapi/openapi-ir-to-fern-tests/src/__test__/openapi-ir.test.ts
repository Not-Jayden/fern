import { AbsoluteFilePath, doesPathExist, join, RelativeFilePath } from "@fern-api/fs-utils";
import { createMockTaskContext } from "@fern-api/task-context";
import { loadAPIWorkspace } from "@fern-api/workspace-loader";
import { readdir } from "fs/promises";
import { OSSWorkspace } from "@fern-api/lazy-fern-workspace";
import { InMemoryOpenAPILoader } from "@fern-api/browser-compatible-fern-workspace";
import fs from "fs/promises";
import { OpenAPI } from "openapi-types";
import yaml from "js-yaml";

const FIXTURES_DIR = join(AbsoluteFilePath.of(__dirname), RelativeFilePath.of("fixtures"));

// eslint-disable-next-line @typescript-eslint/no-misused-promises
describe("openapi-ir", async () => {
    for (const fixture of await readdir(FIXTURES_DIR, { withFileTypes: true })) {
        if (!fixture.isDirectory()) {
            continue;
        }

        it(
            fixture.name,
            async () => {
                const fixturePath = join(FIXTURES_DIR, RelativeFilePath.of(fixture.name), RelativeFilePath.of("fern"));
                const context = createMockTaskContext();
                const workspace = await loadAPIWorkspace({
                    absolutePathToWorkspace: fixturePath,
                    context,
                    cliVersion: "0.0.0",
                    workspaceName: fixture.name
                });
                if (!workspace.didSucceed) {
                    throw new Error(
                        `Failed to load OpenAPI fixture ${fixture.name}\n${JSON.stringify(workspace.failures)}`
                    );
                }

                if (workspace.workspace instanceof OSSWorkspace) {
                    const openApiIr = await (workspace.workspace as OSSWorkspace).getOpenAPIIr({ context });
                    // eslint-disable-next-line jest/no-standalone-expect
                    expect(JSON.stringify(openApiIr, undefined, 2)).toMatchFileSnapshot(
                        `./__snapshots__/openapi-ir/${fixture.name}.json`
                    );
                }
            },
            90_000
        );
    }
});

// eslint-disable-next-line @typescript-eslint/no-misused-promises
describe("openapi-ir-in-memory", async () => {
    const loader = new InMemoryOpenAPILoader();
    for (const fixture of await readdir(FIXTURES_DIR, { withFileTypes: true })) {
        if (!fixture.isDirectory()) {
            continue;
        }
        it(
            fixture.name,
            async () => {
                const fixtureFilePath = await getTestFixturePath(join(FIXTURES_DIR, RelativeFilePath.of(fixture.name)));
                const document = await loader.loadDocument({
                    parsed: await readAndParseOpenAPI(fixtureFilePath)
                });
                expect(JSON.stringify(document, undefined, 2)).toMatchFileSnapshot(
                    `./__snapshots__/openapi-ir-in-memory/${fixture.name}.json`
                );
            },
            90_000
        );
    }
});

async function readAndParseOpenAPI(absolutePathToOpenAPI: AbsoluteFilePath): Promise<OpenAPI.Document> {
    const content = await fs.readFile(absolutePathToOpenAPI, "utf-8");
    return yaml.load(content) as OpenAPI.Document;
}

async function getTestFixturePath(fixtureFilePath: AbsoluteFilePath): Promise<AbsoluteFilePath> {
    const yamlFixturePath = join(fixtureFilePath, RelativeFilePath.of("openapi.yaml"));
    const ymlFixturePath = join(fixtureFilePath, RelativeFilePath.of("openapi.yml"));
    const jsonFixturePath = join(fixtureFilePath, RelativeFilePath.of("openapi.json"));
    const swaggerFixturePath = join(fixtureFilePath, RelativeFilePath.of("swagger.json"));
    return (await doesPathExist(yamlFixturePath))
        ? yamlFixturePath
        : (await doesPathExist(ymlFixturePath))
        ? ymlFixturePath
        : (await doesPathExist(jsonFixturePath))
        ? jsonFixturePath
        : swaggerFixturePath;
}
