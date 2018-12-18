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
import * as github from './github';
import * as travis from './travis';
import * as heroku from './heroku';

/**
 * Initializes actions which require on server startup
 * @param {Object} options
 */
export const init = (options: any) => {
  const opts = options || {};

  /**
   * Checks for available options and if true creates webhooks for given configs
   */
  if (opts.actions && opts.actions === true) {
    /**
     * Creates a GitHub webhook
     */
    if (process.env.GITHUB_OWNER && process.env.GITHUB_REPO && process.env.GITHUB_TOKEN) {
      github.createWebhook(process.env.GITHUB_TOKEN, process.env.GITHUB_OWNER, process.env.GITHUB_REPO);
    }

    /**
     * Creates a Travis CI webhook
     */
    if (process.env.TRAVIS_TOKEN) {
      // Disable direct push builds
      travis.disablePushBuilds(process.env.TRAVIS_TOKEN, process.env.GITHUB_OWNER, process.env.GITHUB_REPO);
    }

    /**
     * Creates a Heroku webhook
     */
    if (process.env.HEROKU_TOKEN && process.env.HEROKU_APP) {
      heroku.createWebhook(process.env.HEROKU_TOKEN, process.env.HEROKU_APP);
    }
  }
};