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

import { ReactNode } from 'react';
import IconButton from '@mui/material/IconButton';
import CloseIcon from '@mui/icons-material/Close';
import SettingsIcon from '@mui/icons-material/Settings';
import { DialogTitle as MuiDialogTitle, SxProps, Theme } from '@mui/material';
import './DialogTitle.scss';

type Props = {
  icon?: ReactNode;
  children?: ReactNode;
  onClose?: () => void;
};

export default function DialogTitle(props: Props) {
  const {
    icon, children, onClose, ...other
  } = props;

  const sx : SxProps<Theme> = {
    position: 'absolute',
    right: 8,
    top: 8,
    color: (theme) => theme.palette.grey[500],
  };

  return (
    <MuiDialogTitle className="dialog-title" sx={{ m: 0, p: 2 }} {...other}>
      {icon || <SettingsIcon className="title-icon" />}
      {children}
      <IconButton aria-label="close" onClick={onClose} sx={sx}>
        <CloseIcon />
      </IconButton>
    </MuiDialogTitle>
  );
}
