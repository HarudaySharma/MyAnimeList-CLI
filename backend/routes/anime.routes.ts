import express, { Router } from 'express';
import { searchAnime } from '../controllers/anime.controller.js';

const router: Router = express.Router();

router.get('search/', searchAnime);

export default router;


