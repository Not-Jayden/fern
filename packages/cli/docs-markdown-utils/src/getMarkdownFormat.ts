import { AbsoluteFilePath, getFilename, RelativeFilePath } from "@fern-api/fs-utils";

export function getMarkdownFormat(absoluteFilepath: AbsoluteFilePath | RelativeFilePath): "mdx" | "md" {
    const filename = getFilename(absoluteFilepath);
    if (filename == null) {
        throw new Error(`Filepath ${absoluteFilepath} does not have a filename`);
    }
    const format = filename.endsWith(".mdx") ? "mdx" : filename.endsWith(".md") ? "md" : undefined;
    if (format == null) {
        throw new Error(`Filepath ${absoluteFilepath} does not have a markdown extension`);
    }
    return format;
}
