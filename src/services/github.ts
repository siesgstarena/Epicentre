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
import request from 'request';
import { Request, Response, NextFunction } from 'express';
import { getJSON } from '../utils/responses';
import * as travis from './travis';

const API_URL = 'https://api.github.com';

/**
 * Creates a GitHub webhook for current server
 * @param {string} token
 * @param {string} owner
 * @param {string} repo
 */
export const createWebhook = (token: string, owner: string, repo: string) => {
  /**
   * API options
   */
  const options = {
    url: API_URL + '/repos/' + owner + '/' + repo + '/hooks',
    headers: {
      'User-Agent': 'Epicentre',
      'Authorization': 'token ' + token
    },
    json: {
      'name': 'web',
      'active': true,
      'events': [
        'push'
      ],
      'config': {
        'url': process.env.SERVER_URL + '/actions/github/payload',
        'content_type': 'json'
      }
    }
  };
  request.post(options, (error, response, body) => {
    if (error) {
      console.log('Error creating GitHub webhook');
      console.log(error);
    } else {
      console.log('Successfully created GitHub webhook');
    }
  });
};

/**
 * Receives payload from the GitHub webhook
 * Reference: https://developer.github.com/v3/activity/events/types/
 * @param {Request} req
 * @param {Response} res
 * @param {NextFunction} next
 */
export const receivePayload = (req: Request, res: Response, next: NextFunction) => {
  /**
   * Only process the github push events not hook created events
   */
  if (!req.body.hook) {
    /**
     * Extract information for notification builder
     */
    const branch = req.body.ref.substr(11); // For refs/heads/{branch}
    const countCommits = req.body.commits.length;
    const author = req.body.head_commit.author.username;
    const notifyMessage = author + ' pushed ' + countCommits + ((countCommits > 1) ? ' commits to ' : ' commit to ') + branch;

    /**
     * Send Android Notifications on receiving Payload
     */
    // notifications.sendNotification("GitHub", notifyMessage)

    /**
     * If Travis CI is enabled, start the build with the latest push
     */
    if (process.env.TRAVIS_TOKEN) {
      travis.createBuild(process.env.TRAVIS_TOKEN, process.env.GITHUB_OWNER, process.env.GITHUB_REPO, process.env.GITHUB_BRANCH);
    }
  }
  /**
   * Send back a dummy response which is to complete HTTP cycle
   */
  res.status(200);
  res.json(getJSON(200));
};