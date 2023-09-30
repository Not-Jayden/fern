/**
 * This file was auto-generated by Fern from our API Definition.
 */

import * as SeedExamples from "../../../../..";
import express from "express";
import * as serializers from "../../../../../../serialization";
import * as errors from "../../../../../../errors";

export interface ServiceServiceMethods {
    getFile(
        req: express.Request<
            {
                filename: string;
            },
            SeedExamples.File_,
            never,
            never
        >,
        res: {
            send: (responseBody: SeedExamples.File_) => Promise<void>;
            cookie: (cookie: string, value: string, options?: express.CookieOptions) => void;
            locals: any;
        }
    ): void | Promise<void>;
}

export class ServiceService {
    private router;

    constructor(private readonly methods: ServiceServiceMethods, middleware: express.RequestHandler[] = []) {
        this.router = express.Router({ mergeParams: true }).use(
            express.json({
                strict: false,
            }),
            ...middleware
        );
    }

    public addMiddleware(handler: express.RequestHandler): this {
        this.router.use(handler);
        return this;
    }

    public toRouter(): express.Router {
        this.router.get("/:filename", async (req, res, next) => {
            try {
                await this.methods.getFile(req as any, {
                    send: async (responseBody) => {
                        res.json(
                            await serializers.File_.jsonOrThrow(responseBody, { unrecognizedObjectKeys: "strip" })
                        );
                    },
                    cookie: res.cookie.bind(res),
                    locals: res.locals,
                });
                next();
            } catch (error) {
                console.error(error);
                if (error instanceof errors.SeedExamplesError) {
                    switch (error.errorName) {
                        case "NotFoundError":
                            break;
                        default:
                            console.warn(
                                `Endpoint 'getFile' unexpectedly threw ${error.constructor.name}.` +
                                    ` If this was intentional, please add ${error.constructor.name} to` +
                                    " the endpoint's errors list in your Fern Definition."
                            );
                    }
                    await error.send(res);
                } else {
                    res.status(500).json("Internal Server Error");
                }
                next(error);
            }
        });
        return this.router;
    }
}
