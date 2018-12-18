/**
 * @license
 * Copyright 2019 SIESGSTarena
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 * =============================================================================
 */
import express from 'express';
import logger from 'morgan';
import compression from 'compression';
import helmet from 'helmet';
import bodyParser from 'body-parser';
import dotenv from 'dotenv';
dotenv.config({ path: '.env' });
import { initRoutes } from './routes/index';
import * as epicentre from './services/epicentre';
import { Request, Response, NextFunction } from 'express';
import { getJSON } from './utils/responses';
class App {

  public app: express.Application;

  constructor() {
    this.app = express();
    /**
     * Initialize all middlewares
     * @param app
     */
    this.initMiddleware(this.app);

    /**
     * Initialize epicentre with options
     * @param {object}
     */
    epicentre.init({
      actions: false // Avoid generating webhooks, only for testing services
    });

    /**
     * Initialize all the routes
     * @param app
     */
    initRoutes(this.app);

    /**
     * Error generator for the app
     */
    this.app.use((req: Request, res: Response, next: NextFunction) => {
      const err = new Error('Not Found');
      next(err);
    });

    /**
     * Error handler for the app
     * @param {Error} err
     */
    this.app.use((err: object, req: Request, res: Response, next: NextFunction) => {
      res.status(404);
      res.json(getJSON(404));
    });
  }

  private initMiddleware(app: express.Application): void {
    app.use(logger('dev'));
    app.use(helmet());
    app.disable('x-powered-by');
    app.use(compression());
    app.use(bodyParser.json());
    app.use(bodyParser.urlencoded({ extended: false }));
    app.set('json spaces', 2);
  }
}

export default new App().app;
