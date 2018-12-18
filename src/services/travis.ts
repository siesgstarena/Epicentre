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
import request from  'request';
import { Request, Response, NextFunction } from 'express';
import { getJSON, getDataJSON, getMessageJSON } from '../utils/responses';

/**
 * Travis Constants
 */
const API_URL = 'https://api.travis-ci.com';
const TOKEN = process.env.TRAVIS_TOKEN;
const GITHUB_OWNER = process.env.GITHUB_OWNER;
const GITHUB_REPO = process.env.GITHUB_REPO;

/**
 * Common headers used in all requests
 */
const REQUEST_HEADERS = {
  'Content-Type': 'application/json',
  'Travis-API-Version': '3',
  'User-Agent': 'Epicentre',
  'Authorization': 'token ' + TOKEN
};

/**
 * Disables direct push builds for Travis CI so that manual build via GitHub can be done for webhook to this server
 * Note: If Direct push builds are enabled, it will not allow to set up webhook
 * Other Way: Set the notifications webhook in .travis.yml file for sending payload to this server hosted URL
 * @param {string} token
 * @param {string} owner
 * @param {string} repo
 */
export const disablePushBuilds = (token: string, owner: string, repo: string) => {
  /**
   * API options
   */
  const options = {
    url: API_URL + '/repo/' + owner + '%2F' + repo + '/setting/build_pushes',
    headers: REQUEST_HEADERS,
    json: {
      'setting.value': false // Set false to avoid push builds to this repository
    }
  };
  request.patch(options, (error, response, body) => {
    if (error) {
      console.log('Error in updating Travis CI Setting');
      console.log(error);
    } else {
      console.log('Successfully updated Travis CI Setting');
    }
  });
};

/**
 * Creates a Travis CI Build with custom config for receiving back notification via webhook
 * Note: This will update configuration with deep_mege
 * Note: This will add webhook URL as current server URL to receive payload after build finishes
 * @param {string} token
 * @param {string} owner
 * @param {string} repo
 * @param {string} branch
 */
export const createBuild = (token: string, owner: string, repo: string, branch: string) => {
  /**
   * API options
   */
  const options = {
    url: API_URL + '/repo/' + owner + '%2F' + repo + '/requests',
    headers: REQUEST_HEADERS,
    json: {
      'request': {
        'branch': branch,
        'config': {
          'merge_mode': 'deep_merge',
          'notifications': {
            'webhooks': process.env.SERVER_URL + '/actions/travis/payload',
          }
        }
      }
    }
  };
  request.post(options, (error, response, body) => {
    if (error) {
      console.log('Error in creating Travis CI Build Request');
      console.log(error);
    } else {
      console.log('Successfully created Travis CI Build Request');
    }
  });
};

/**
 * Receives payload from the GitHub webhook
 * Reference: https://docs.travis-ci.com/user/notifications#configuring-webhook-notifications
 * @param {Object} payload
 */
export const receivePayload = (req: Request, res: Response, next: NextFunction) => {
  /**
   * Extract information for notification builder
   */
  const bodyJSON = JSON.parse(req.body.payload);
  const state = bodyJSON.state;
  const branch = bodyJSON.branch;
  const notifyMessage = 'Build on ' + branch + ' ' + state;

  /**
   * Send Android notification for Travis CI build updates
   */
  // notifications.sendNotification("Travis CI", notifyMessage)

  /**
   * Send back a dummy response which is to complete HTTP cycle
   */
  res.status(200);
  res.json(getJSON(200));
};

/**
 * Shows health and normal status
 * Reference: https://developer.travis-ci.com/resource/builds#find
 * @param {string} GITHUB_OWNER
 * @param {string} GITHUB_REPO
 */
export const index = (req: Request, res: Response, next: NextFunction) => {
  const options = {
    url: API_URL + '/repo/' + GITHUB_OWNER + '%2F' + GITHUB_REPO + '/builds?sort_by=updated_at:desc&limit=1',
    headers: REQUEST_HEADERS
  };
  request.get(options, (error, response, body) => {
    if (error) {
      console.log(error);
      res.status(500);
      res.json(getMessageJSON(500, 'Some error. Try again'));
    } else {
      const element = JSON.parse(body).builds[0] ? JSON.parse(body).builds[0] : {};
      res.status(200);
      res.json(getDataJSON(200, 'Travis Build Status', {
        branch: element.branch && element.branch.name ? element.branch.name : 'not defined',
        message: element.commit.message,
        number: element.number,
        state: element.state ? element.state : 'unknown',
        started_at: element.started_at ? element.started_at : 'NA',
        finished_at: element.finished_at ? element.finished_at : 'NA'
      }));
    }
  });
};

/**
 * Shows latest 10 builds for the given project
 * Reference: https://developer.travis-ci.com/resource/builds#find
 * @param {string} GITHUB_OWNER
 * @param {string} GITHUB_REPO
 */
export const builds = (req: Request, res: Response, next: NextFunction) => {
  const options = {
    url: API_URL + '/repo/' + GITHUB_OWNER + '%2F' + GITHUB_REPO + '/builds?sort_by=updated_at:desc',
    headers: REQUEST_HEADERS
  };
  request.get(options, (error, response, body) => {
    if (error) {
      console.log(error);
      res.status(500);
      res.json(getMessageJSON(500, 'Some error. Try again'));
    } else {
      res.status(200);
      res.json(getDataJSON(200, 'Successfully loaded Travis CI builds', JSON.parse(body).builds));
    }
  });
};