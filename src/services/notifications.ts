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
import * as firebaseAdmin from 'firebase-admin';
import path from 'path';

/**
 * Initialize Firebase Admin for current application
 */
firebaseAdmin.initializeApp({
  credential: firebaseAdmin.credential.cert(path.join(__dirname, '../../config/serviceAccount.json')),
  databaseURL: process.env.FIREBASE_DATABASE_URL
});

/**
 * Send notifications to all registered app users
 * @param {string} title
 * @param {string} body
 */
module.exports.sendNotification = (title: string, body: string) => {
  /**
   * Get all tokens and notify them sequentially
   * Tokens are stored in Realtime Database of the project
   * Tokens are FCM Device Tokens generated during Sign In in Android App
   * TODO: Switch to Topcis for better performance
   */
  const db = firebaseAdmin.database();
  const ref = db.ref('tokens');
  ref.once('value')
    .then((snapshot) => {
      snapshot.forEach((data) => {
        /**
         * Notification Options
         * @param {string} title Service name like Heroku, Travis CI or GitHub
         * @param {string} body Readable description for notification
         */
        const message = {
          notification: {
            title: title,
            body: body
          },
          token: data.val()
        };

        /**
         * Send Notifications for each token in Realtime Database
         */
        firebaseAdmin.messaging().send(message)
          .then((response) => {
            console.log(response);
          })
          .catch((error) => {
            console.log('Notification Error:' + error);
          });
      });
    });
};