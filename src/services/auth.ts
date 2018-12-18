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
import fs from 'fs';
import path from 'path';
import { getDataJSON, getMessageJSON } from '../utils/responses';
import * as jwt from 'jsonwebtoken';

/**
 * Note: Token is generated using JSON Web Tokens
 * Package Reference: https://github.com/auth0/node-jsonwebtoken
 * We use TOKEN_OPTIONS to sign the token and verify the same
 *
 * Note: Adjust "algorithm" based on Private-Public Key pair used
 * e.g. default key size of 512 bit size uses RS256
 *
 * Generate Keys: http://travistidwell.com/jsencrypt/demo/
 */
const TOKEN_OPTIONS = {
  issuer: process.env.SERVER_NAME,
  subject: 'codechef@siesgst.ac.in',
  audience: process.env.SERVER_URL,
  expiresIn: 60 * 60 * 24,
  algorithm: 'RS256'
};

/**
 * Used for authenticating app users or other application end users
 * Creates a token and signs the payload
 * Note: the token generated should be used for all other routes as authorization
 * @param {object} payload (process.env.SERVER_CODE)
 * @param {file} keys
 * @param {object} signOptions
 */
export const createToken = (req: Request, res: Response, next: NextFunction) => {
  if (req.body && req.body.code) {
    /**
     * Server Configuration is used as payload data for signing the token
     */
    if (process.env.SERVER_CODE === req.body.code) {
      /**
       * Read the contents for extracting private key information
       */
      fs.readFile(path.join(__dirname, '../../config/keys/private.key'), 'utf8', (err, privateKey) => {
        if (err) {
          console.log(err);
          res.status(500);
          res.json(getMessageJSON(500, 'Some error. Try again'));
        } else {
          /**
           * Generate a new token and send back to requested user
           */
          jwt.sign({
            code: process.env.SERVER_CODE
          },
          privateKey, TOKEN_OPTIONS, (err, token) => {
            if (err) {
              console.log(err);
              res.status(500);
              res.json(getMessageJSON(500, 'Some error generating token. Try again'));
            } else {
              res.status(200);
              res.json(getDataJSON(200, 'Signing in...', {
                token: token
              }));
            }
          });
        }
      });
    } else {
      res.status(400);
      res.json(getMessageJSON(400, 'Incorrect secret code'));
    }
  } else {
    res.status(422);
    res.json(getMessageJSON(422, 'Some fields were missing'));
  }
};

/**
 * Errors generated due to token by used NPM module
 * Generates a human readable error message for user
 * @param {Object} error
 */
const getErrorMessage = (error: any) => {
  switch (error.name) {
  case 'TokenExpiredError':
    return 'Token has expired';
  default:
    return 'Token is invalid or malformed';
  }
};

/**
 * This is used as middlware for securing specific routes by adding authenticating layer
 * Verifies the given token
 * Note: Headers should consist of authorization to extract the token
 * @param {Object} Headers
 * @param {Key} authorization
 */
export const verifyToken = (req: Request, res: Response, next: NextFunction) => {
  /**
   * Access Headers of incoming request
   */
  const headers = req.headers;
  if (headers['authorization']) {
    const token = headers['authorization'];

    /**
     * Read the contents for extracting public key information
     */
    fs.readFile(path.join(__dirname, '../../config/keys/public.key'), 'utf8', (err, publicKey) => {
      if (err) {
        console.log(err);
        res.status(500);
        res.json(getMessageJSON(500, 'Some error verifying the token. Try again'));
      } else {
        /**
         * Verifies the token with payload used from project's configuration and public key signing
         */
        jwt.verify(token, publicKey, TOKEN_OPTIONS, (err, result) => {
          if (err) {
            res.status(401);
            res.json(getMessageJSON(401, getErrorMessage(err)));
          } else {
            next(err);
          }
        });
      }
    });
  } else {
    res.status(401);
    res.json(getMessageJSON(401, 'Cannot find the authorization token'));
  }
};