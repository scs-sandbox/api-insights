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

import { ButtonHTMLAttributes, useState } from 'react';
import { Alert, Snackbar } from '@mui/material';
import IconButton from '../../IconButton/IconButton';
import BusyIcon, { BusyIconType } from '../../BusyIcon/BusyIcon';
import EnvIcon from '../SpecStateIcon/SpecStateIcon';
import { SpecData, useReconstructSnapshot } from '../../../query/spec';
import classNames from '../../../utils/className';
import { ApiError } from '../../../query/api';
import './ReconstructSnapshotButton.scss';

export type ReconstructedEvent = {
  id: string;
};

type Props = ButtonHTMLAttributes<HTMLButtonElement> & {
  serviceId: string;
  onReconstructed?: (data: ReconstructedEvent) => void;
};

export default function ReconstructSnapshotButton(props: Props) {
  const [message, setMessage] = useState('messsage unset');
  const {
    isLoading,
    isError,
    isSuccess,
    mutate: reconstructSnapshot,
    reset: resetReconstructSnapshot,
  } = useReconstructSnapshot(props.serviceId);

  const onClick = () => {
    reconstructSnapshot(undefined, {
      onSuccess: (data) => {
        if (props.onReconstructed) {
          props.onReconstructed(data);
        }
        setMessage('Succeed in reconstructing snapshot');
      },
      onError: (data: unknown) => {
        if ((data as ApiError).status === 400) {
          setMessage('apiclarity: no API traffic found');
        } else {
          setMessage('apiclarity not integrated for given service');
        }
      },
    });
  };

  const onCloseMessage = () => {
    resetReconstructSnapshot();
  };

  const renderMessage = () => {
    if (!isError && !isSuccess) return null;

    const severity = isSuccess ? 'success' : 'error';

    return (
      <Snackbar
        open
        anchorOrigin={{ vertical: 'top', horizontal: 'center' }}
        autoHideDuration={3000}
        onClose={onCloseMessage}
      >
        <Alert
          onClose={onCloseMessage}
          severity={severity}
          sx={{ width: '100%' }}
        >
          {message}
        </Alert>
      </Snackbar>
    );
  };

  const renderIcon = () => (
    <i className="recons-icon">
      <EnvIcon value={SpecData.SpecState.Release} />
      <BusyIcon type={BusyIconType.ArrowCircle} busy={isLoading} />
    </i>
  );

  const className = classNames('reconstruct-snapshot-button', props.className);
  const responseMessage = renderMessage();

  return (
    <div className={className}>
      <IconButton
        icon={renderIcon()}
        disabled={props.disabled || isLoading}
        onClick={onClick}
      >
        {props.children || 'Reconstruct Snapshot'}
      </IconButton>
      {responseMessage}
    </div>
  );
}
