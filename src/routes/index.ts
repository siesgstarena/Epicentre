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
import { Request, Response, NextFunction } from 'express';
import { getJSON, getMessageJSON } from '../utils/responses';
import * as github from '../services/github';
import * as heroku from '../services/heroku';
import * as travis from '../services/travis';
import * as auth from '../services/auth';

export const initRoutes = (app: any) => {
  /**
   * Payload level routes
   */
  app.get('/', (req: Request, res: Response, next: NextFunction) => {
    res.status(200);
    res.json(getJSON(200));
  });
  /**
   * Github Payload
   */
  app.post('/actions/github/payload', github.receivePayload);

  /**
   * Travis CI Payload
   */
  app.post('/actions/travis/payload', travis.receivePayload);

  /**
   * Heroku Payload
   */
  app.post('/actions/heroku/payload', heroku.receivePayload);

  /**
   * Application endpoint routes
   */
  app.get('/api', (req: Request, res: Response, next: NextFunction) => {
    res.status(200);
    res.json(getMessageJSON(200, 'Successfully loaded Epicentre Services'));
  });

  /**
   * Authentication
   * Creates a JWT Token for valid user
   */
  app.post('/api/auth', auth.createToken);

  /**
   * Travis CI
   * Shows health and builds
   */
  app.get('/api/travis', auth.verifyToken, travis.index);
  app.get('/api/travis/builds', auth.verifyToken, travis.builds);

  /**
   * Heroku
   * Shows health, app info, dyno info, updates app settings and deploys
   */
  app.get('/api/heroku/', auth.verifyToken, heroku.index);
  app.get('/api/heroku/info', auth.verifyToken, heroku.info);
  app.get('/api/heroku/dynos', auth.verifyToken, heroku.dynos);
  app.get('/api/heroku/config', auth.verifyToken, heroku.getConfig);
  app.patch('/api/heroku/config', auth.verifyToken, heroku.updateConfig);
  app.patch('/api/heroku/maintenance', auth.verifyToken, heroku.setMaintenance);
  app.get('/api/heroku/deploy', auth.verifyToken, heroku.deploy);
};