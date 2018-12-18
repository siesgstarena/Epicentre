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
import httpStatusCode from 'http-status-code';

/**
 * Used to create a base response object for sending action responses
 * @param {Number} code
 */
export const getJSON = (code: number) => {
  return {
    status: httpStatusCode.getMessage(code),
    message: 'Epicentre Service Platform',
    serverInformation: {
      serverName: process.env.SERVER_NAME,
      apiVersion: process.env.SERVER_VERSION,
      currentTime: new Date().getTime()
    }
  };
};

/**
 * Used to create a base response object for sending service responses
 * @param {Number} code
 * @param {string} message
 */
export const getMessageJSON = (code: number, message: string) => {
  return {
    status: httpStatusCode.getMessage(code),
    message: message,
    serverInformation: {
      serverName: process.env.SERVER_NAME,
      apiVersion: process.env.SERVER_VERSION,
      currentTime: new Date().getTime()
    }
  };
};

/**
 * Used to create a base response object for sending service responses
 * @param {number} code
 * @param {string} message
 * @param {object} data
 */
export const getDataJSON = (code: number, message: string, data: object) => {
  return {
    status: httpStatusCode.getMessage(code),
    message: message,
    data: data,
    serverInformation: {
      serverName: process.env.SERVER_NAME,
      apiVersion: process.env.SERVER_VERSION,
      currentTime: new Date().getTime()
    }
  };
};