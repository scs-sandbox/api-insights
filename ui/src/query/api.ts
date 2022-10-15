/*
 * Copyright 2022 Cisco Systems, Inc. and its affiliates.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

// eslint-disable-next-line no-use-before-define
type JSONValue = string | number | boolean | Array<JSONValue> | JSONObject;

interface JSONObject {
  [x: string]: JSONValue;
}

export class ApiError extends Error {
  status: number;

  constructor(status, message) {
    super(message);
    this.status = status;
    this.name = 'ApiError';
  }
}

export function buildApiAbsoluteUrl(relativeUrl: string) {
  return (
    process.env.REACT_APP_API_ENDPOINT_URL
    + process.env.REACT_APP_API_BASE_PATH
    + relativeUrl
  );
}

function buildMethodOptions(
  method: string,
  payload?: JSONObject,
  options?: RequestInit,
) {
  const types = method === 'GET'
    ? {}
    : {
      'Content-Type': 'application/json',
    };
  const headers = options && options.headers ? { ...options.headers, ...types } : types;
  const credentials = (options && options.credentials) ? options.credentials : 'include';
  const body = payload
    ? {
      body: JSON.stringify(payload),
    }
    : {};

  return {
    ...options,
    method,
    credentials,
    headers,
    ...body,
  };
}

function fetchJson(
  method: string,
  url: string,
  payload?: JSONObject,
  options?: RequestInit,
) {
  const absoluteUrl = buildApiAbsoluteUrl(url);

  const finalOptions = buildMethodOptions(method, payload, options);

  return fetch(absoluteUrl, finalOptions).then(async (resp) => {
    if (resp.ok) return resp.json();
    const errMsg = await resp.text();
    throw new ApiError(resp.status, errMsg);
  });
}

const API = {
  get(url: string, options?: RequestInit) {
    return fetchJson('GET', url, null, options);
  },
  post(url: string, payload?: JSONObject, options?: RequestInit) {
    return fetchJson('POST', url, payload, options);
  },
  put(url: string, payload?: JSONObject, options?: RequestInit) {
    return fetchJson('PUT', url, payload, options);
  },
  patch(url: string, payload?: JSONObject, options?: RequestInit) {
    return fetchJson('PATCH', url, payload, options);
  },
  delete(url: string, payload?: JSONObject, options?: RequestInit) {
    return fetchJson('DELETE', url, payload, options);
  },
};

export default API;
