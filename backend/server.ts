import express ,{Application, Request, Response} from 'express'
import { config } from 'dotenv';
import mongoose from 'mongoose';

import {errorMiddleware} from './types/interfaces.js'
import animeRoutes from './routes/anime.routes.js';

config();

const MONGO: string = process.env.MONGO || "NULL";
mongoose.connect(MONGO)
.then(() => {
    console.log("connected to db");
})
.catch((err : Error) => {
    console.log("db not connected");
    console.error(err)
})

const app: Application = express();
const PORT: number | string = process.env.PORT || 3000;

app.listen(PORT, () => {
    console.log(`server running at port: ${process.env.PORT}`);
})


app.use('api/anime/', animeRoutes);

app.use(errorMiddleware);
