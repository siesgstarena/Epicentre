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
import firebaseAdmin from 'firebase-admin';
import { spawn } from 'child_process';
import fs from 'fs';
import path from 'path';
import dotenv from 'dotenv';
dotenv.config({ path: '.env' });

/**
 * Initialize Firebase Admin for current application
 */
firebaseAdmin.initializeApp({
  credential: firebaseAdmin.credential.cert(path.join(__dirname, '../config/serviceAccount.json')),
  storageBucket: process.env.FIREBASE_BUCKET
});
const bucket = firebaseAdmin.storage().bucket();

/**
 * Constants for storing daily backups with date
 * Note: Replacing this will ensure you are uploading data dumps in same folder
 */
const date = new Date();

/**
 * Individually upload all the files to Firebase Storage
 * @param {string} fileName
 */
async function uploadFile (fileName: string) {
  /**
   * Upload individual files to storage and wait for the process to finish
   */
  await bucket.upload('dump/test/' + fileName, {
    destination: date + '/' + fileName,
    public: true
  }).then((file) => {
    console.log(fileName + ' uploaded successfully');
  }).catch((err) => {
    console.log(err);
  });
}

/**
 * Start mongodump
 * Reference: https://docs.mongodb.com/manual/reference/program/mongodump/
 */
console.log('Starting MongoDB Dump...');

/**
 * This will run a process of mongodump with details provided for MongoDB
 * Note: Make sure mongo commands are installed or MongoDB is installed on this server
 * Note: If this process exits with error on Heroku, make sure you add this buildpack in your application
 * Reference Buildpack: https://github.com/siesgstarena/heroku-buildpack-mongo
 */
const backup = spawn('mongodump',
    ['--ssl', '--host', process.env.MONGO_REPLSET + '/' + process.env.MONGO_NODES,
    '--authenticationDatabase', 'admin', '-u', process.env.MONGO_USERNAME, '-p', process.env.MONGO_PASSWORD]);
backup.stderr.on('data', (data) => {
    console.log(data.toString());
});
backup.on('exit', () => {
  console.log('Finished MongoDB Dump...');
  console.log('Starting Remote Backup...');
  /**
   * Start uploading dumped files to remote storage
   */
  fs.readdirSync('dump/test').forEach(file => {
    uploadFile(file);
  });
});
