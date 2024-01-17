import { Response, Request } from "express";

export class errorHandler extends Error{
    message: string;
    statusCode: number;
    constructor(statusCode: number, message: string) {
        super();
        this.statusCode = statusCode;
        this.message = message;
    }
}

export const errorMiddleware = (err: errorHandler, res: Response, req: Request) => {
    const message: string = err.message || "Internal Server Error";
    const statusCode: number = err.statusCode || 500;
    res.status(statusCode)
    .json({message: message});
}
