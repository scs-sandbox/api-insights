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

/*
example:
    const className = classNames("button", props.className, props.active ? "active" : "");

result:
    className will be "button start active" if the props.className
    is "start" and the props.active is true
*/
export default function classNames(...args: (string | undefined)[]): string {
  return args
    .filter((i) => (i || '').trim())
    .map((i) => (i || '').trim())
    .join(' ');
}
