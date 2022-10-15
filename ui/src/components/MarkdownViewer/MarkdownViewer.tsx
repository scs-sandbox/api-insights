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

import { HTMLAttributes } from 'react';
import DOMPurify from 'dompurify';
import { marked } from 'marked';
import classNames from '../../utils/className';
import './MarkdownViewer.scss';

type Props = HTMLAttributes<HTMLElement> & {
  text?: string;
};

export default function MarkdownViewer(props: Props) {
  const {
    text, className, ...other
  } = props;

  const fullClassName = classNames('markdown-viewer', className);

  if (!text) return (<div className={fullClassName} />);

  const dangerousHtml = marked(text, { gfm: true, breaks: true });
  const secureHtml = DOMPurify.sanitize(dangerousHtml);
  const html = { __html: secureHtml };

  const newProps = {
    ...other,
    className: fullClassName,
    dangerouslySetInnerHTML: html,
  };

  return <div {...newProps} />;
}
