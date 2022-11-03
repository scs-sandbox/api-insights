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

import { useState, ButtonHTMLAttributes } from 'react';
import { useAddSpec } from '../../../query/spec';
import useSpecFile, { ReadSpecFunciton, SpecFileData } from './hooks/useSpecFile';
import UploadSpecDialog, {
  UploadingEvent,
} from '../UploadSpecDialog/UploadSpecDialog';
import IconButton from '../../IconButton/IconButton';
import UploadIcon from '../../UploadIcon/UploadIcon';
import classNames from '../../../utils/className';
import SnackAlert from '../../SnackAlert/SnackAlert';
import { ApiError } from '../../../query/api';
import './UploadSpecButton.scss';

export type UploadSpecParam = {
  doc: string;
  revision: string;
  service_id: string;
  disabled: boolean;
};

type Props = ButtonHTMLAttributes<HTMLButtonElement> & {
  serviceId: string;
  onUploaded: (data: UploadSpecParam) => void;
};

export default function UploadSpecButton(props: Props) {
  const [openDialog, setOpenDialog] = useState(false);
  const [readSpec, specReading, specReadingError, setSpecReadingError] = useSpecFile();
  const {
    isLoading: isSpecAdding,
    error: specAddingError,
    isSuccess: isSpecAddingSuccess,
    mutate: addSpec,
    reset: resetAddSpecResult,
  } = useAddSpec();

  const isUploading = !!(specReading || isSpecAdding);
  const setSafeSpecReadingError = setSpecReadingError as
    React.Dispatch<React.SetStateAction<Error | undefined>>;

  const onOpenDialog = () => {
    setOpenDialog(true);
  };

  const onCloseDialog = () => {
    setOpenDialog(false);
  };

  const onUploading = (uploadingEvent: UploadingEvent) => {
    const { file, revision } = uploadingEvent;
    (readSpec as ReadSpecFunciton)(file).then((specFileDataEvent: SpecFileData) => {
      const valid = specFileDataEvent.parsedSpec
        && (specFileDataEvent.parsedSpec.openapi || specFileDataEvent.parsedSpec.swagger);

      if (!valid) {
        if (setSpecReadingError) {
          setSafeSpecReadingError(new Error('Wrong Format'));
        }
        return;
      }

      const data = {
        doc: specFileDataEvent.text,
        revision,
        service_id: props.serviceId,
      };

      addSpec(data, {
        onSuccess: (successData) => {
          setOpenDialog(false);
          if (props.onUploaded) {
            props.onUploaded(successData);
          }
        },
      });
    });
  };

  const onCloseMessage = () => {
    setSafeSpecReadingError(undefined);
    resetAddSpecResult();
  };

  const renderUploadButton = () => (
    <IconButton
      icon={<UploadIcon busy={isUploading} />}
      disabled={props.disabled || isUploading}
      className={`${props.disabled ? 'disabled' : ''}`}
      onClick={onOpenDialog}
    >
      {props.children || 'Upload New Spec'}
    </IconButton>
  );

  const renderDialog = () => {
    if (!openDialog) return null;

    return (
      <UploadSpecDialog
        busy={isUploading}
        open
        handleClose={onCloseDialog}
        onUploading={onUploading}
      />
    );
  };

  const renderErrorMessageText = (specFormatError: unknown, specAddingApiError: ApiError) => {
    if (specFormatError) return 'Wrong Format!';

    return specAddingApiError.status === 409 ? specAddingApiError.message : 'Failed to upload!';
  };

  const renderErrorMessage = () => {
    if (!specReadingError && !specAddingError) return null;

    const specAddingApiError = specAddingError as ApiError;

    const messageText = renderErrorMessageText(specReadingError, specAddingApiError);

    return (
      <SnackAlert severity="error" message={messageText} onClose={onCloseMessage} />
    );
  };

  const renderSuccessMessage = () => {
    if (!isSpecAddingSuccess) return null;

    return (
      <SnackAlert
        severity="success"
        message="The spec has been uploaded"
        onClose={onCloseMessage}
      />
    );
  };

  const className = classNames('upload-spec-button', props.className);
  const button = renderUploadButton();
  const dialog = renderDialog();
  const errorMessage = renderErrorMessage();
  const successMessage = renderSuccessMessage();

  return (
    <div className={className}>
      {button}
      {dialog}
      {errorMessage}
      {successMessage}
    </div>
  );
}
