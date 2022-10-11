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

import './DownloadIcon.scss';

export default function DownloadIcon() {
  return (
    <i className="download-icon">
      <svg className="download-icon-svg" viewBox="0 0 8 10">
        <path d="M7.06641 5.40625C7.10547 5.36719 7.125 5.30859 7.125 5.25C7.125 5.19141 7.10547 5.13281 7.06641 5.07422L6.91016 4.9375C6.87109 4.89844 6.8125 4.85938 6.75391 4.85938C6.67578 4.85938 6.61719 4.89844 6.57812 4.9375L4.33203 7.20312V1.10938C4.33203 1.05078 4.29297 0.992188 4.25391 0.953125C4.21484 0.914062 4.15625 0.875 4.09766 0.875H3.90234C3.82422 0.875 3.76562 0.914062 3.72656 0.953125C3.6875 0.992188 3.66797 1.05078 3.66797 1.10938V7.20312L1.42188 4.9375C1.36328 4.89844 1.30469 4.85938 1.24609 4.85938C1.16797 4.85938 1.12891 4.89844 1.08984 4.9375L0.933594 5.07422C0.894531 5.13281 0.875 5.19141 0.875 5.25C0.875 5.30859 0.894531 5.36719 0.953125 5.40625L3.82422 8.31641C3.86328 8.35547 3.92188 8.375 4 8.375C4.05859 8.375 4.11719 8.35547 4.17578 8.29688L7.06641 5.40625ZM7.75 9.39062C7.75 9.46875 7.71094 9.52734 7.67188 9.56641C7.63281 9.60547 7.57422 9.625 7.51562 9.625H0.484375C0.40625 9.625 0.347656 9.60547 0.308594 9.56641C0.269531 9.52734 0.25 9.46875 0.25 9.39062V9.23438C0.25 9.17578 0.269531 9.11719 0.308594 9.07812C0.347656 9.03906 0.40625 9 0.484375 9H7.51562C7.57422 9 7.63281 9.03906 7.67188 9.07812C7.71094 9.11719 7.75 9.17578 7.75 9.23438V9.39062Z" />
      </svg>
    </i>
  );
}
