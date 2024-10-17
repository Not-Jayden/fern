import { python } from "../../../src";
import { Writer } from "../core/Writer";
import { Reference } from "../Reference";
import { Class } from "../Class";
import { Method } from "../Method";
import { CodeBlock } from "../CodeBlock";
import { Type } from "../Type";

describe("PythonFile", () => {
    let writer: Writer;

    beforeEach(() => {
        writer = new Writer();
    });

    it("Add a class with no references", async () => {
        const file = python.file({
            moduleName: "test_module",
            path: ["test"],
            name: "test_file"
        });

        const testClass = python.class_({ name: "TestClass" });
        file.addStatement(testClass);

        file.write(writer);
        expect(await writer.toStringFormatted()).toMatchSnapshot();
    });

    it("Add a class with a reference that uses a python standard library reference", async () => {
        const file = python.file({
            moduleName: "test_module",
            path: ["test"],
            name: "test_file"
        });

        const testClass = python.class_({
            name: "TestClass"
        });
        testClass.addBody(python.codeBlock("flat_list = list(itertools.chain([[1, 2], [3, 4]]))"));
        writer.addReference(python.reference({ modulePath: ["itertools"], name: "chain" }));

        file.addStatement(testClass);

        file.write(writer);
        expect(await writer.toStringFormatted()).toMatchSnapshot();
    });

    it("Add a class with a reference that uses a relative import", async () => {
        const file = python.file({
            moduleName: "my_module",
            path: ["level_1"],
            name: "test_file"
        });

        const relativeRef = python.reference({
            modulePath: ["my_module", "level_1"],
            name: "OtherClass"
        });
        const testClass = python.class_({
            name: "TestClass",
            extends_: [relativeRef]
        });
        writer.addReference(relativeRef);
        file.addStatement(testClass);

        // Add a class with a deeply nested relative import
        const deeplyNestedRef = python.reference({
            modulePath: ["my_module", "level_1", "level_2"],
            name: "DeepClass"
        });
        const deeplyNestedClass = python.class_({
            name: "DeeplyNestedTestClass",
            extends_: [deeplyNestedRef]
        });
        writer.addReference(deeplyNestedRef);
        file.addStatement(deeplyNestedClass);

        file.write(writer);
        expect(await writer.toStringFormatted()).toMatchSnapshot();
    });

    it("Add a Method", async () => {
        const file = python.file({
            moduleName: "test_module",
            path: ["test"],
            name: "test_file"
        });

        const testMethod = new Method({
            name: "test_method",
            parameters: [],
            return_: python.Type.str()
        });
        file.addStatement(testMethod);

        file.write(writer);
        expect(await writer.toStringFormatted()).toMatchSnapshot();
    });

    it("Add a code block", async () => {
        const file = python.file({
            moduleName: "test_module",
            path: ["test"],
            name: "test_file"
        });

        const codeBlock = new CodeBlock("print('Hello, World!')");
        file.addStatement(codeBlock);

        file.write(writer);
        expect(await writer.toStringFormatted()).toMatchSnapshot();
    });

    it("Add a class with an absolute import and alias", async () => {
        const file = python.file({
            moduleName: "test_module",
            path: ["test"],
            name: "test_file"
        });

        const absoluteRef = python.reference({
            modulePath: ["external_module", "submodule"],
            name: "ExternalClass",
            alias: "AliasedClass"
        });
        const testClass = python.class_({
            name: "TestClassWithAlias",
            extends_: [absoluteRef]
        });
        writer.addReference(absoluteRef);
        file.addStatement(testClass);

        file.write(writer);
        expect(await writer.toStringFormatted()).toMatchSnapshot();
    });

    it("Add a class with a relative import and alias", async () => {
        const file = python.file({
            moduleName: "test_module",
            path: ["test", "subdir"],
            name: "test_file"
        });

        const relativeRef = python.reference({
            modulePath: ["test_module", "test", "sibling_dir"],
            name: "SiblingClass",
            alias: "AliasedSibling"
        });
        const testClass = python.class_({
            name: "TestClassWithRelativeAlias",
            extends_: [relativeRef]
        });
        writer.addReference(relativeRef);
        file.addStatement(testClass);

        file.write(writer);
        expect(await writer.toStringFormatted()).toMatchSnapshot();
    });
});
