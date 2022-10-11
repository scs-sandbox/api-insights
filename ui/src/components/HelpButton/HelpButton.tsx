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

import { useState, useEffect, ReactNode } from 'react';
import IconButton from '../IconButton/IconButton';
import TalkIcon from './images/talk.png';
import './HelpButton.scss';

type Props = {
  show?: boolean;
  title: ReactNode;
  message: ReactNode;
};

export default function HelpButton(props: Props) {
  const [openDialog, setOpenDialog] = useState(props.show);

  useEffect(() => {
    setOpenDialog(props.show);
  }, [props.show]);

  function renderButton() {
    return (
      <IconButton onClick={() => setOpenDialog(!openDialog)}>
        ?
      </IconButton>
    );
  }

  function renderBar() {
    if (!openDialog) return null;

    return (
      <div className="help-bar">
        <img className="talk-icon" alt="help" src={TalkIcon} />
        <div className="title">{props.title}</div>
        <div className="description">{props.message}</div>
      </div>
    );
  }

  const button = renderButton();
  const bar = renderBar();

  return (
    <div className={`help-container ${openDialog ? 'show' : ''}`}>
      <div className="help-button">
        {button}
      </div>
      {bar}
    </div>
  );
}
