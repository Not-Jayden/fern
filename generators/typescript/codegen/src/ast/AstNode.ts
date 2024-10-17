import { AbstractAstNode } from "@fern-api/generator-commons";
import { Writer } from "./Writer";
import * as prettier from "prettier";

export abstract class AstNode extends AbstractAstNode {
    /**
     * Writes the node to a string.
     */
    public toString(): string {
        const writer = new Writer();
        this.write(writer);
        const response = writer.toString();
        console.log(response); // TODO: comment
        return response;
    }

    public toStringFormatted(): string {
        return prettier.format(this.toString(), { parser: "typescript", tabWidth: 4, printWidth: 120 });
    }
}
