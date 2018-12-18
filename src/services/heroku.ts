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
import request from 'request';
import syncRequest from 'sync-request';
import { getJSON, getMessageJSON, getDataJSON } from '../utils/responses';

/**
 * Heroku Constants
 */
const API_URL = 'https://api.heroku.com';
const APP_NAME = process.env.HEROKU_APP;
const BRANCH = process.env.HEROKU_BRANCH;
const TOKEN = process.env.HEROKU_TOKEN;

/**
 * Github Constants
 */
const GITHUB_TOKEN = process.env.GITHUB_TOKEN;
const GITHUB_API_URL = 'https://api.github.com';
const GITHUB_OWNER = process.env.GITHUB_OWNER;
const GITHUB_REPO = process.env.GITHUB_REPO;

/**
 * Common headers used in all requests
 */
const REQUEST_HEADERS = {
  'Content-Type': 'application/json',
  'Accept': 'application/vnd.heroku+json; version=3',
  'Authorization': 'Bearer ' + TOKEN
};

/**
 * Deletes all previous Heroku webhooks
 * @param {string} token
 * @param {string} appName
 */
const deleteWebhooks = (token: string, appName: string) => {
  return new Promise((resolve, reject) => {
    /**
     * API options
     */
    const options = {
      url: API_URL + '/apps/' + appName + '/webhooks',
      headers: {
        'Content-Type': 'application/json',
        'Accept': 'application/vnd.heroku+json; version=3.webhooks',
        'Authorization': 'Bearer ' + token
      }
    };
    request.get(options, (error, response, body) => {
      if (error) {
        console.log('Error getting previous Heroku webhooks');
        console.log(error);
      } else {
        /**
         * API options
         */
        const optionsEach = {
          headers: {
            'Content-Type': 'application/json',
            'Accept': 'application/vnd.heroku+json; version=3.webhooks',
            'Authorization': 'Bearer ' + token
          }
        };
        const bodyJSON = JSON.parse(body);
        let error = false;
        if (bodyJSON.length === 0) {
          console.log('There are no previous Heroku webhooks');
          resolve(true);
        }
        /**
         * Delete individually all previous webhooks
         */
        bodyJSON.forEach((webhook: any) => {
          const url = API_URL + '/apps/' + appName + '/webhooks/' + webhook.id;
          const res = syncRequest('DELETE', url, optionsEach);
          if (res.statusCode !== 200) {
            error = true;
            console.log('Error deleting a previous Heroku webhook');
            resolve(false);
          }
        });
        if (!error) {
          console.log('Successfully deleted all previous Heroku webhooks');
          resolve(true);
        }
      }
    });
  });
};

/**
 * Creates a Heroku webhook for current server
 * @param {string} token
 * @param {string} appName
 */
export const createWebhook = (token: string, appName: string) => {
  /**
   * Heroku creates new webhook with each request
   * So, delete all previous webhooks and create a new one for system on startup
   */
  deleteWebhooks(token, appName).then((result) => {
    if (result) {
      /**
       * API options
       */
      const options = {
        url: API_URL + '/apps/' + appName + '/webhooks',
        headers: {
          'Content-Type': 'application/json',
          'Accept': 'application/vnd.heroku+json; version=3',
          'Authorization': 'Bearer ' + token
        },
        json: {
          'include': [
            'api:release',
            'api:dyno',
            'api:formation',
            'dyno'
          ],
          'level': 'notify',
          'url': process.env.SERVER_URL + '/actions/heroku/payload'
        }
      };
      request.post(options, (error, response, body) => {
        if (error) {
          console.log('Error creating Heroku webhook');
          console.log(error);
        } else {
          console.log('Successfully created Heroku webhook');
        }
      });
    } else {
      console.log('Cannot create Heroku webhook due to previous deletion error');
    }
  });
};

/**
 * Receives payload from the Heroku webhook
 * Reference: https://devcenter.heroku.com/articles/app-webhooks-schema#delivery
 * @param {Object} req.body
 */
export const receivePayload = (req: Request, res: Response, next: NextFunction) => {
  /**
   * Extract information for notification builder
   */
  const action = req.body.action;
  const resource = req.body.resource;
  const state = req.body.data.state;

  /**
   * Only notify for action "create" for release type of payload
   * Only notify for dyno state "up", "down" or "crashed"
   */
  if (resource === 'release' && action === 'create') {
    /**
     * Send Android Notification for release updates
     */
    // notifications.sendNotification("Heroku", "Released a new version")
  } else if (state === 'up' || state === 'down' || state === 'crashed') {
    /**
     * Send Android Notification for dyno updates
     */
    // notifications.sendNotification("Heroku", "Dyno is " + state)
  }
  /**
   * Send back a dummy response which is to complete HTTP cycle
   */
  res.status(200);
  res.json(getJSON(200));
};

/**
 * Shows health and normal status
 * @param none
 */
export const index = (req: Request, res: Response, next: NextFunction) => {
  /**
   * API options
   */
  const options = {
    url: API_URL + '/apps/' + APP_NAME + '/dynos',
    headers: REQUEST_HEADERS
  };
  /**
   * Send back response
   */
  request.get(options, (error, response, body) => {
    if (error) {
      console.log(error);
      res.status(500);
      res.json(getMessageJSON(500, 'Some error. Try again'));
    } else {
      const state = JSON.parse(body)[0].state;
      const dyno = JSON.parse(body)[0];
      /**
       * API options
       */
      const options = {
        url: API_URL + '/apps/' + APP_NAME,
        headers: REQUEST_HEADERS
      };
      /**
       * Send back response
       */
      request.get(options, (error, response, body) => {
        if (error) {
          console.log(error);
          res.status(500);
          res.json(getMessageJSON(500, 'Some error. Try again'));
        } else {
          res.status(200);
          res.json(getDataJSON(200, 'Application is ' + state, {
            info: JSON.parse(body),
            dyno: dyno
          }));
        }
      });
    }
  });
};

/**
 * Shows information about the current heroku app
 * Reference: https://devcenter.heroku.com/articles/platform-api-reference#app-info
 * @param none
 */
export const info = (req: Request, res: Response, next: NextFunction) => {
  const options = {
    url: API_URL + '/apps/' + APP_NAME,
    headers: REQUEST_HEADERS
  };
  request.get(options, (error, response, body) => {
    if (error) {
      console.log(error);
      res.status(500);
      res.json(getMessageJSON(500, 'Some error. Try again'));
    } else {
      res.status(200);
      res.json(getDataJSON(200, 'Successfully loaded Heroku app information', JSON.parse(body)));
    }
  });
};

/**
 * Shows dyno information associated with the current heroku app
 * Reference: https://devcenter.heroku.com/articles/platform-api-reference#dyno-info
 * @param none
 */
export const dynos = (req: Request, res: Response, next: NextFunction) => {
  const options = {
    url: API_URL + '/apps/' + APP_NAME + '/dynos',
    headers: REQUEST_HEADERS
  };
  request.get(options, (error, response, body) => {
    if (error) {
      console.log(error);
      res.status(500);
      res.json(getMessageJSON(500, 'Some error. Try again'));
    } else {
      res.status(200);
      res.json(getDataJSON(200, 'Successfully loaded Heroku dynos information', JSON.parse(body)[0]));
    }
  });
};

/**
 * Lists all config vars for current heroku app
 * Reference: https://devcenter.heroku.com/articles/platform-api-reference#config-vars-info-for-app
 * @param none
 */
export const getConfig = (req: Request, res: Response, next: NextFunction) => {
  const options = {
    url: API_URL + '/apps/' + APP_NAME + '/config-vars',
    headers: REQUEST_HEADERS
  };
  request.get(options, (error, response, body) => {
    if (error) {
      console.log(error);
      res.status(500);
      res.json(getMessageJSON(500, 'Some error. Try again'));
    } else {
      res.status(200);
      res.json(getDataJSON(200, 'Successfully loaded Heroku config vars', JSON.parse(body)));
    }
  });
};

/**
 * Adds or Updates a config vars for current heroku app
 * Reference: https://devcenter.heroku.com/articles/platform-api-reference#config-vars-update
 * @param {object} config
 */
export const updateConfig = (req: Request, res: Response, next: NextFunction) => {
  const config = req.body.config;
  if (config) {
    const options = {
      url: API_URL + '/apps/' + APP_NAME + '/config-vars',
      headers: REQUEST_HEADERS,
      json: config
    };
    request.patch(options, (error, response, body) => {
      if (error) {
        console.log(error);
        res.status(500);
        res.json(getMessageJSON(500, 'Some error. Try again'));
      } else {
        res.status(200);
        res.json(getDataJSON(200, 'Successfully udpated Heroku config vars', body));
      }
    });
    // TODO:
  } else {
    res.status(422);
    res.json(getMessageJSON(422, 'Some fields were missing or were not set correctly'));
  }
};

/**
 * Enable or disable maintenance mode for the app
 * Reference: https://devcenter.heroku.com/articles/platform-api-reference#app-update
 * @param {boolean} maintenance
 */
export const setMaintenance = (req: Request, res: Response, next: NextFunction) => {
  const options = {
    url: API_URL + '/apps/' + APP_NAME,
    headers: REQUEST_HEADERS,
    json: {
      'maintenance': req.body.maintenance
    }
  };
  request.patch(options, (error, response, body) => {
    if (error) {
      console.log(error);
      res.status(500);
      res.json(getMessageJSON(500, 'Some error. Try again'));
    } else {
      res.status(200);
      res.json(getMessageJSON(200, 'Successfully updated Heroku maintenance mode'));
    }
  });
};

/**
 * Extracts the URL for latest code (in tarball format) from GitHub branch and starts Heroku build
 * Note: Builds on Heroku require source to be in .tar format
 * The created slug is deployed as a new release for the given app
 * Reference: https://devcenter.heroku.com/articles/platform-api-reference#build-create
 * @param {string} branch
 */
export const deploy = (req: Request, res: Response, next: NextFunction) => {
  /**
   * Extract GitHub code in tarball with latest commit from branch
   * @param {string} GITHUB_OWNER
   * @param {string} GITHUB_REPO
   * @param {string} BRANCH
   * @param {string} GITHUB_TOKEN
   */
  const tarballUrl = GITHUB_API_URL + '/repos/' + GITHUB_OWNER + '/' + GITHUB_REPO +
            '/tarball/' + BRANCH + '?access_token=' + GITHUB_TOKEN;

  /**
   * API options
   */
  const options = {
    url: API_URL + '/apps/' + APP_NAME + '/builds',
    headers: REQUEST_HEADERS,
    json: {
      source_blob: {
        url: tarballUrl,
        version: '',
        checksum: ''
      }
    }
  };

  /**
   * Start the release after extracting code from GitHub
   */
  request.post(options, (error, response, body) => {
    if (error) {
      res.status(500);
      res.json(getMessageJSON(500, 'Error while building the source code'));
    } else {
      res.status(200);
      res.json(getMessageJSON(200, 'Successfully started building on Heroku'));
    }
  });
};