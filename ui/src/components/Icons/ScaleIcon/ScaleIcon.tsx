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

type Props = {
  className?: string;
}

export default function ScaleIcon(props: Props) {
  return (
    <svg
      className={props.className}
      fill="transparent"
      stroke="currentColor"
      strokeWidth="4px"
      viewBox="0 0 139 139"
    >
      <path className="st0" d="M34.8,44.4L12.5,78.5c0,0,5.4,8,22.2,8" />
      <path className="st1" d="M34.8,86.5c16.9,0,22.2-8,22.2-8L34.8,44.4v-4.9h34.7h34.7l0,4.9l22.2,34.1c0,0-5.4,8-22.2,8" />
      <path className="st2" d="M104.2,44.4L82,78.5c0,0,5.4,8,22.2,8" />
      <line className="st3" x1="69.5" x2="69.5" y1="106.6" y2="31" />
      <line className="st4" x1="94.7" x2="44.2" y1="106.6" y2="106.6" />
      <path className="st5" d="M12.5,78.5H57C57,78.5,34.7,95,12.5,78.5z" />
      <path className="st6" d="M82,78.5h44.5C126.5,78.5,104.2,95,82,78.5z" />
    </svg>
  );
}
